package interfaces

import (
	"context"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	u "nexus-sds.com/deploydock/utilities"
)

//Label: com.docker.compose.project, Value: geodata-bknd
//Label: com.docker.compose.project.working_dir, Value: /Users/dhaanpaa/src/geodata-bknd
//Label: com.docker.compose.project.config_files, Value: /Users/dhaanpaa/src/geodata-bknd/compose.yml

func GetComposeProjectConfigFiles() map[string]string {
	apiClient, err := client.NewClientWithOpts(client.FromEnv)

	if err != nil {
		panic(err)
	}

	defer apiClient.Close()

	projects, err := apiClient.ContainerList(context.Background(), container.ListOptions{All: true})
	if err != nil {
		panic(err)
	}

	var projectConfigFiles = make(map[string]string)
	for _, prj := range projects {
		if projectName, ok := prj.Labels["com.docker.compose.project"]; ok {
			if projectConfigFile, ok := prj.Labels["com.docker.compose.project.config_files"]; ok {
				projectConfigFiles[projectName] = projectConfigFile
			}
		}
	}
	return projectConfigFiles
}

func GetComposeProjectWorkingDirectories() map[string]string {
	apiClient, err := client.NewClientWithOpts(client.FromEnv)

	if err != nil {
		panic(err)
	}
	defer apiClient.Close()

	projects, err := apiClient.ContainerList(context.Background(), container.ListOptions{All: true})
	if err != nil {
		panic(err)
	}

	var projectWorkingDirectories = make(map[string]string)
	for _, prj := range projects {
		if projectName, ok := prj.Labels["com.docker.compose.project"]; ok {
			if projectWorkingDirectory, ok := prj.Labels["com.docker.compose.project.working_dir"]; ok {
				projectWorkingDirectories[projectName] = projectWorkingDirectory
			}
		}
	}
	return projectWorkingDirectories

}

func GetComposeProjects() []string {
	apiClient, err := client.NewClientWithOpts(client.FromEnv)

	if err != nil {
		panic(err)
	}
	defer apiClient.Close()

	projects, err := apiClient.ContainerList(context.Background(), container.ListOptions{All: true})
	if err != nil {
		panic(err)
	}

	var listOfProjects []string
	for _, prj := range projects {
		if projectName, ok := prj.Labels["com.docker.compose.project"]; ok {
			listOfProjects = append(listOfProjects, projectName)
		}
	}
	return u.UniqueStrings(listOfProjects)
}

func GetComposeServices() []string {
	apiClient, err := client.NewClientWithOpts(client.FromEnv)

	if err != nil {
		panic(err)
	}
	defer apiClient.Close()

	services, err := apiClient.ContainerList(context.Background(), container.ListOptions{All: true})
	if err != nil {
		panic(err)
	}

	var listOfServices []string
	for _, svc := range services {
		// Filter containers by a Docker Compose label, such as `com.docker.compose.service`
		if serviceName, ok := svc.Labels["com.docker.compose.service"]; ok {
			listOfServices = append(listOfServices, serviceName)
		}
	}

	return u.UniqueStrings(listOfServices)
}

func GetDockerContainerImageNames() []string {
	apiClient, err := client.NewClientWithOpts(client.FromEnv)

	if err != nil {
		panic(err)
	}
	defer apiClient.Close()

	containers, err := apiClient.ContainerList(context.Background(), container.ListOptions{All: true})
	if err != nil {
		panic(err)
	}

	var listOfContainers []string
	for _, ctr := range containers {

		listOfContainers = append(listOfContainers, ctr.Image)
	}

	return listOfContainers
}

func GetComposeProjectWorkingDirectoriesFiltered(prefix string) map[string]string {
	m := u.InvertStringMap(GetComposeProjectWorkingDirectories())
	return u.FilterMapByKeyPrefix(m, prefix)
}
