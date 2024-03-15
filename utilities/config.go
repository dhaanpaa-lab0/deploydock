package utilities

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

// Config struct from yaml
type Config struct {
	GitUser    string `yaml:"git_user"`
	GitHost    string `yaml:"git_host"`
	GitOrg     string `yaml:"git_org"`
	GitRepo    string `yaml:"git_repo"`
	GitBranch  string `yaml:"git_branch"`
	ServerRoot string `yaml:"server_root"`
}

func CreateUrl(cfg Config) string {
	return fmt.Sprintf("%s@%s:%s/%s.git", cfg.GitUser, cfg.GitHost, cfg.GitOrg, cfg.GitRepo)

}

// LoadConfig This function is for loading the config from the yaml file
func LoadConfig() Config {
	var tempConfig Config
	// Load the config from the yaml file
	if FileExists("deploydock.yaml") {
		yamlFile, err := os.ReadFile("deploydock.yaml")

		if err != nil {
			fmt.Println(err)
			return tempConfig
		} else {
			errYamlUnmarshal := yaml.Unmarshal([]byte(yamlFile), &tempConfig)

			if errYamlUnmarshal != nil {
				fmt.Println(errYamlUnmarshal)
				return Config{}
			} else {
				tempConfig.ServerRoot = ParsePath(tempConfig.ServerRoot)

				return tempConfig
			}
		}
	}

	return Config{}

}

func SaveConfig(fileName string, config Config) {
	// Save the config to the yaml file
	yamlConfig, err := yaml.Marshal(config)
	if err != nil {
		fmt.Println(err)
	} else {
		errYamlWrite := os.WriteFile(fileName, yamlConfig, 0644)
		if errYamlWrite != nil {
			fmt.Println(errYamlWrite)
		} else {
			fmt.Println("Config saved successfully")
		}
	}
}
