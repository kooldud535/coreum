package coreum

import (
	"context"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"

	"github.com/CoreumFoundation/coreum-tools/pkg/libexec"
	coreumtools "github.com/CoreumFoundation/coreum/build/tools"
	"github.com/CoreumFoundation/crust/build/golang"
	crusttools "github.com/CoreumFoundation/crust/build/tools"
	"github.com/CoreumFoundation/crust/build/types"
)

// generateProtoDocs collects cosmos-sdk, cosmwasm and tendermint proto files from coreum go.mod,
// generates documentation using above proto files + coreum/proto, and places the result to docs/api.md.
func generateProtoDocs(ctx context.Context, deps types.DepsFunc) error {
	deps(golang.Tidy)

	moduleDirs, includeDirs, err := protoCDirectories(ctx, repoPath, deps)
	if err != nil {
		return err
	}

	absPath, err := filepath.Abs(repoPath)
	if err != nil {
		return errors.WithStack(err)
	}

	generateDirs := []string{
		filepath.Join(absPath, "proto"),
		filepath.Join(moduleDirs[cosmosSDKModule], "proto"),
		filepath.Join(moduleDirs[cosmWASMModule], "proto"),
	}

	err = executeProtocCommand(ctx, deps, includeDirs, generateDirs)
	if err != nil {
		return err
	}

	return nil
}

// executeProtocCommand ensures needed dependencies, composes the protoc command and executes it.
func executeProtocCommand(ctx context.Context, deps types.DepsFunc, includeDirs, generateDirs []string) error {
	deps(coreumtools.EnsureProtoc, coreumtools.EnsureProtocGenDoc)

	args := []string{
		"--doc_out=docs",
		"--plugin=" + crusttools.Path("bin/protoc-gen-doc", crusttools.TargetPlatformLocal),
		"--doc_opt=docs/api.tmpl.md,api.md",
	}

	for _, path := range includeDirs {
		args = append(args, "--proto_path", path)
	}

	allProtoFiles, err := findAllProtoFiles(generateDirs)
	if err != nil {
		return err
	}
	args = append(args, allProtoFiles...)

	cmd := exec.Command(crusttools.Path("bin/protoc", crusttools.TargetPlatformLocal), args...)
	cmd.Dir = repoPath

	return libexec.Exec(ctx, cmd)
}

// findAllProtoFiles returns a list of absolute paths to each proto file within the given directories.
func findAllProtoFiles(pathList []string) (finalResult []string, err error) {
	var iterationResult []string
	for _, path := range pathList {
		iterationResult, err = listFilesByPath(path, ".proto")
		if err != nil {
			return nil, err
		}
		finalResult = append(finalResult, iterationResult...)
	}

	return finalResult, nil
}

// listFilesByPath returns the array of files with the specific extension within the given path.
func listFilesByPath(path, extension string) (fileList []string, err error) {
	err = filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return errors.WithStack(err)
		}

		if info.IsDir() {
			return nil
		}

		if strings.HasSuffix(path, extension) {
			fileList = append(fileList, path)
		}

		return nil
	})

	return fileList, err
}
