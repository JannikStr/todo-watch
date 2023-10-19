package main

import (
	"errors"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

const ENV_CONFIG_PATH string = "TODO_WATCH_CONFIG"

const DEFAULT_CONFIG_PATH string = "/.config/todo-watch.yaml"

type Configuration struct {
	General struct {
		keys              []string `yaml:"keys"`
		ignored_filetypes []string `yaml:"ignored_filetypes"`
	} `yaml:"general"`
}

func ReadConfig(config_path string) (Configuration, error) {
	_, err := os.Stat(config_path)
	default_config := Configuration{}

	default_config.General.keys = []string{"TODO", "FIXME"}
	default_config.General.ignored_filetypes = []string{""}

	if errors.Is(err, os.ErrNotExist) {
		yaml_data, err := yaml.Marshal(&default_config)
		fmt.Println(string(yaml_data))
		if err != nil {
			return default_config, err
		}
		file, err := os.Create(config_path)

		_, err = file.Write(yaml_data)
		if err != nil {
			return default_config, err
		}

	} else if err == nil {
		config := Configuration{}
		file, err := os.Open(config_path)
		if err != nil {
			return default_config, nil
		}
		defer file.Close()
		if file != nil {
			decoder := yaml.NewDecoder(file)
			if err := decoder.Decode(&config); err != nil {
				fmt.Println(err.Error())
			}
		}
		return config, nil
	}
	return default_config, err
}

func main() {
	config_path, is_set := os.LookupEnv(ENV_CONFIG_PATH)

	if !is_set {
		config_path = os.Getenv("HOME") + DEFAULT_CONFIG_PATH
	}

	fmt.Printf("Filepath: %s\n", config_path)
	config, err := ReadConfig(config_path)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(config)
}
