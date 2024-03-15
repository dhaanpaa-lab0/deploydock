package services

import (
	"log"
	i "nexus-sds.com/deploydock/interfaces"
	u "nexus-sds.com/deploydock/utilities"
	"os"
)

func SetupDeployDockGitRepo(ddkConfig u.Config, url string) {
	errCreatingDirectory := os.MkdirAll(ddkConfig.ServerRoot, 0775)
	if errCreatingDirectory != nil {
		log.Fatalln(errCreatingDirectory)
	}
	cloneMessages, exitCode := i.GitClone(url, ddkConfig.ServerRoot)
	log.Println(cloneMessages)
	if exitCode != 0 {
		log.Fatalf("Error cloning: %v", exitCode)

	}
}

func UpdateDeployDockGitRepo(ddkConfig u.Config) {
	pullMessages, exitCode := i.GitPull(ddkConfig.ServerRoot)
	log.Println(pullMessages)
	if exitCode != 0 {
		log.Fatalf("Error pulling: %v", exitCode)
	}
}
