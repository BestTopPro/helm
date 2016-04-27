package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/deis/tiller/pkg/client"
	"github.com/deis/tiller/pkg/kubectl"
	"github.com/spf13/cobra"
)

const initDesc = `
This command installs Tiller (the helm server side component) onto your
Kubernetes Cluster and sets up local configuration in $HELM_HOME (default: ~/.helm/)
`

var (
	tillerImg   string
	defaultRepo = map[string]string{"default-name": "default-url"}
)

func init() {
	initCmd.Flags().StringVarP(&tillerImg, "tiller-image", "i", "", "override tiller image")
	RootCommand.AddCommand(initCmd)
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize Helm on both client and server.",
	Long:  installDesc,
	RunE:  runInit,
}

// runInit initializes local config and installs tiller to Kubernetes Cluster
func runInit(cmd *cobra.Command, args []string) error {
	if len(args) != 0 {
		return errors.New("This command does not accept arguments. \n")
	}

	if err := ensureHome(); err != nil {
		return err
	}

	if err := installTiller(); err != nil {
		return err
	}

	fmt.Println("Happy Helming!")
	return nil
}

func installTiller() error {
	// TODO: take value of global flag kubectl and pass that in
	runner := buildKubectlRunner("")

	i := client.NewInstaller()
	i.Tiller["Image"] = tillerImg
	out, err := i.Install(runner)

	if err != nil {
		return fmt.Errorf("error installing %s %s", string(out), err)
	}
	fmt.Println("\nTiller (the helm server side component) has been installed into your Kubernetes Cluster.")

	return nil
}

func buildKubectlRunner(kubectlPath string) kubectl.Runner {
	if kubectlPath != "" {
		kubectl.Path = kubectlPath
	}
	return &kubectl.RealRunner{}
}

// ensureHome checks to see if $HELM_HOME exists
//
// If $HELM_HOME does not exist, this function will create it.
func ensureHome() error {
	configDirectories := []string{homePath(), cacheDirectory(), localRepoDirectory()}

	for _, p := range configDirectories {
		if fi, err := os.Stat(p); err != nil {
			fmt.Printf("Creating %s \n", p)
			if err := os.MkdirAll(p, 0755); err != nil {
				return fmt.Errorf("Could not create %s: %s", p, err)
			}
		} else if !fi.IsDir() {
			return fmt.Errorf("%s must be a directory", p)
		}
	}

	repoFile := repositoriesFile()
	if fi, err := os.Stat(repoFile); err != nil {
		fmt.Printf("Creating %s \n", repoFile)
		if err := ioutil.WriteFile(repoFile, []byte("local: localhost:8879/charts\n"), 0644); err != nil {
			return err
		}
	} else if fi.IsDir() {
		return fmt.Errorf("%s must be a file, not a directory", repoFile)
	}

	localRepoCacheFile := localRepoDirectory(localRepoCacheFilePath)
	if fi, err := os.Stat(localRepoCacheFile); err != nil {
		fmt.Printf("Creating %s \n", localRepoCacheFile)
		_, err := os.Create(localRepoCacheFile)
		if err != nil {
			return err
		}

		//TODO: take this out and replace with helm update functionality
		os.Symlink(localRepoCacheFile, cacheDirectory("local-cache.yaml"))
	} else if fi.IsDir() {
		return fmt.Errorf("%s must be a file, not a directory", localRepoCacheFile)
	}

	fmt.Printf("$HELM_HOME has been configured at %s.\n", helmHome)
	return nil
}
