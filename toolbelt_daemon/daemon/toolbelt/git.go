package toolbelt

import (
	"os/exec"
	"os"
)

func GitCloneBranch(repoUrl string, pathspec string, dest string) error {
	if err := clone(repoUrl, dest); err != nil {
		return err
	}

	return checkoutBranch(pathspec, dest)
}

func clone(repoUrl string, dest string) error {
	cmd := exec.Command("git", "clone", repoUrl, dest)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func checkoutBranch(pathspec string, codeDir string) error {
	cmd := exec.Command("git", "checkout", pathspec)
	cmd.Dir = codeDir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
