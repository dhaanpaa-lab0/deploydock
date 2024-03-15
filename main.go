package main

import (
	"fmt"
	"log"
	i "nexus-sds.com/deploydock/interfaces"
	svcs "nexus-sds.com/deploydock/services"
	u "nexus-sds.com/deploydock/utilities"
)

func main() {
	ddkConfig := u.LoadConfig()

	url := u.CreateUrl(ddkConfig)
	fmt.Printf("Git URL .......... : '%s'\n", url)
	fmt.Printf("Git Branch ....... : '%s'\n", ddkConfig.GitBranch)
	fmt.Printf("Server Root ...... : '%s'\n", ddkConfig.ServerRoot)

	if !u.FileExists(ddkConfig.ServerRoot) {
		svcs.SetupDeployDockGitRepo(ddkConfig, url)
	} else {
		svcs.UpdateDeployDockGitRepo(ddkConfig)
	}

	// Git Latest Commit Hash
	commitHash, err := i.GitGetLatestCommit(ddkConfig.ServerRoot)
	if err != 0 {
		log.Fatalf("Error getting latest commit hash: %v", err)
	}
	fmt.Printf("Latest Commit Hash: %s\n", commitHash)

	// Git Get Message
	commitMessage, err := i.GitGetMessage(ddkConfig.ServerRoot, commitHash)
	if err != 0 {
		log.Fatalf("Error getting commit message: %v", err)
	}
	fmt.Printf("Latest Commit Message : %s\n", commitMessage)
	containers := i.GetListOfContainers()
	fmt.Printf("Containers .......... : %v\n", containers)
	services := i.GetListOfComposeServices()
	fmt.Printf("Services ............ : %v\n", services)
	projects := i.GetListOfComposeProjects()
	fmt.Printf("Projects ............ : %v\n", projects)
	configFiles := i.GetProjectConfigFiles()
	fmt.Printf("Config Files ........ : %v\n", configFiles)
	workingDirs := i.GetProjectWorkingDirectories()
	fmt.Printf("Working Directories . : %v\n", u.InvertStringMap(workingDirs))
	projectNames := i.GetComposeProjectsFilteredByDirectory("")
	fmt.Printf("Projects by Directory : %v\n", projectNames)
	projectNamesFiltered := i.GetComposeProjectsFilteredByDirectory(ddkConfig.ServerRoot)
	fmt.Printf("Projects by Directory : %v\n", projectNamesFiltered)

}
