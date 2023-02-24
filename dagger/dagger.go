package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"dagger.io/dagger"
)

func main() {
	ctx := context.Background()

	client, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stdout))
	if err != nil {
		fmt.Printf("Error connecting to Dagger Engine: %s", err)
		os.Exit(1)
	}

	defer client.Close()

	src := client.Host().Directory(".")
	if err != nil {
		fmt.Printf("Error getting reference to host directory: %s", err)
		os.Exit(1)
	}

	golang := client.Container().From("golang:latest")
	golang = golang.WithMountedDirectory("/src", src).
		WithWorkdir("/src").
		WithEnvVariable("CGO_ENABLED", "0")

	// The WithExec() method returns a revised Container containing the results of command execution.
	golang = golang.WithExec(
		[]string{"go", "build", "-o", "build/"},
	)

	path := "build/"
	err = os.MkdirAll(filepath.Join(".", path), os.ModePerm)
	if err != nil {
		fmt.Printf("Error creating output folder: %s", err)
		os.Exit(1)
	}

	build := golang.Directory(path)
	// writes the build/ directory from the container to the host using the Directory.Export() method
	_, err = build.Export(ctx, path)
	if err != nil {
		fmt.Printf("Error writing directory: %s", err)
		os.Exit(1)
	}

	cn, err := client.Container().
		Build(src).
		Publish(ctx, "ukaul/dagger-example:0.0.1")

	if err != nil {
		fmt.Printf("Error creating and pushing container: %s", err)
		os.Exit(1)
	}

	fmt.Printf("Succesfully created new container: %s", cn)
}
