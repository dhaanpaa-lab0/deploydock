package interfaces

import (
	"errors"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
	"os/exec"
	"strings"
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

func GitSubmitMessage(repoPath string, message string) (string, int) {
	totalOutput := ""
	exitCodeNeeded := 0
	outputGitAddCommand, errGitAddCommand := exec.Command("git", "-C", repoPath, "add", "-A", ".").CombinedOutput()
	if errGitAddCommand != nil {
		var exitError *exec.ExitError
		errors.As(errGitAddCommand, &exitError)
		return string(outputGitAddCommand), exitError.ExitCode()

	} else {
		totalOutput = strings.Join([]string{totalOutput, string(outputGitAddCommand)}, "\n\n")
		exitCodeNeeded = 0
	}

	outputCommit, errCommitCommnad := exec.Command("git", "-C", repoPath, "commit", "-a", "-m", "\""+message+"\"").CombinedOutput()
	if errCommitCommnad != nil {
		var exitError *exec.ExitError
		errors.As(errCommitCommnad, &exitError)
		return string(outputCommit), exitError.ExitCode()
	} else {
		totalOutput = strings.Join([]string{totalOutput, string(outputGitAddCommand)}, "\n\n")
		exitCodeNeeded = 0
	}

	return totalOutput, exitCodeNeeded
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

func GitGetLatestCommit(repoPath string) string {
	open, err := git.PlainOpen(repoPath)
	if err != nil {
		return ""
	}
	ref, err := open.Head()
	if err != nil {
		return ""
	}
	return ref.Hash().String()
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

func GitGetListOfFilesInCommit(repoPath string, gitHash string) []string {
	hash := plumbing.NewHash(gitHash)
	open, err := git.PlainOpen(repoPath)
	if err != nil {
		return nil
	}

	commitObject, err := open.CommitObject(hash)
	if err != nil {
		return nil
	}
	files := make([]string, 0)
	commitFiles, err := commitObject.Files()
	if err != nil {
		return nil
	}

	errLoopingThroughFiles := commitFiles.ForEach(func(file *object.File) error {
		files = append(files, file.Name)
		return nil
	})
	if errLoopingThroughFiles != nil {
		return nil
	}

	if err != nil {
		return nil
	}
	return files
}
