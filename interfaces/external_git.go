package interfaces

import (
	"errors"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"os/exec"
)

func GitClone(srcUrl string, destPath string) (string, int) {
	output, err := exec.Command("git", "clone", srcUrl, destPath).CombinedOutput()
	if err != nil {
		var exitError *exec.ExitError
		errors.As(err, &exitError)
		return string(output), exitError.ExitCode()
	}
	return string(output), 0
}

func GitPull(repoPath string) (string, int) {
	output, err := exec.Command("git", "-C", repoPath, "pull").CombinedOutput()
	if err != nil {
		var exitError *exec.ExitError
		errors.As(err, &exitError)
		return string(output), exitError.ExitCode()
	} else {
		return string(output), 0
	}

}

func GitCheckout(repoPath string, branch string) (string, int) {
	output, err := exec.Command("git", "-C", repoPath, "checkout", branch).CombinedOutput()
	if err != nil {
		var exitError *exec.ExitError
		errors.As(err, &exitError)
		return string(output), exitError.ExitCode()
	} else {
		return string(output), 0
	}
}

func GitResetHard(repoPath string) (string, int) {
	output, err := exec.Command("git", "-C", repoPath, "reset", "--hard").CombinedOutput()
	if err != nil {
		var exitError *exec.ExitError
		errors.As(err, &exitError)
		return string(output), exitError.ExitCode()
	} else {
		return string(output), 0
	}
}

func GitCleanForce(repoPath string) (string, int) {
	output, err := exec.Command("git", "-C", repoPath, "clean", "-f").CombinedOutput()
	if err != nil {
		var exitError *exec.ExitError
		errors.As(err, &exitError)
		return string(output), exitError.ExitCode()
	} else {
		return string(output), 0
	}
}

func GitGetLatestCommit(repoPath string) (string, int) {
	open, err := git.PlainOpen(repoPath)
	if err != nil {
		return "", 0
	}
	ref, err := open.Head()
	if err != nil {
		return "", 0
	}
	return ref.Hash().String(), 0
}

func GitGetMessage(repoPath string, gitHash string) (string, int) {
	hash := plumbing.NewHash(gitHash)
	open, err := git.PlainOpen(repoPath)
	if err != nil {
		return "", 0
	}
	commitObject, err := open.CommitObject(hash)
	if err != nil {
		return "", 0
	}

	return commitObject.Message, 0
}
