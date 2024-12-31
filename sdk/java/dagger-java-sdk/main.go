// A generated module for Java functions
//
// This module has been generated via dagger init and serves as a reference to
// basic module structure as you get started with Dagger.
//
// Two functions have been pre-created. You can modify, delete, or add to them,
// as needed. They demonstrate usage of arguments and return types using simple
// echo and grep commands. The functions can be called from the dagger CLI or
// from one of the SDKs.
//
// The first line in this comment block is a short description line and the
// rest is a long description with more detail on the module's purpose or usage,
// if appropriate. All modules should have a short description.

package main

import (
	"context"
	"fmt"
	"github.com/dagger/dagger/sdk/java/dagger-java-sdk/internal/dagger"
)

type Java struct{}

// Returns a container that echoes whatever string argument is provided
func (m *Java) ContainerEcho(stringArg string) *dagger.Container {
	return dag.Container().From("alpine:latest").WithExec([]string{"echo", stringArg})
}

// Returns lines that match a pattern in the files of the provided Directory
func (m *Java) GrepDir(ctx context.Context, directoryArg *dagger.Directory, pattern string) (string, error) {
	return dag.Container().
		From("alpine:latest").
		WithMountedDirectory("/mnt", directoryArg).
		WithWorkdir("/mnt").
		WithExec([]string{"grep", "-R", pattern, "."}).
		Stdout(ctx)
}

type JavaSdk struct {
	SourceDir     *dagger.Directory
	RequiredPaths []string
}

func New(
	// Directory with the Java SDK source code.
	// +optional
	// +defaultPath="/sdk/java"
	// +ignore=["**", "!generated/", "!src/", "!scripts/", "!composer.json", "!composer.lock", "!LICENSE", "!README.md"]
	sdkSourceDir *dagger.Directory,
) (*JavaSdk, error) {
	if sdkSourceDir == nil {
		return nil, fmt.Errorf("sdk source directory not provided")
	}
	return &JavaSdk{
		RequiredPaths: []string{},
		SourceDir:     sdkSourceDir,
	}, nil
}

func (m *JavaSdk) ModuleRuntime(
	ctx context.Context,
	modSource *dagger.ModuleSource,
	introspectionJSON *dagger.File,
) (*dagger.Container, error) {
	// We could just move CodegenBase to ModuleRuntime, but keeping them
	// separate allows for easier future changes.
	return dag.Container().From("alpine:latest").WithExec([]string{"echo", "Hello, world!"}), nil
}
