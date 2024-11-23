package Config

import (
	"encoding/json"
	"fmt"
	"github.com/SidorkinAlex/CyclicCommandCheckerAndExecutive/internal/fileutils"
	"log"
	"os"
	"path/filepath"
)

type Config []Command

// Config структура для хранения конфигурации
type Command struct {
	Command       string `json:"checkingCommand"`
	Interval      int    `json:"interval"`
	BranchCommand []BranchResultExecution
}
type BranchResultExecution struct {
	ResultExecution string   `json:"resultExecution"`
	Commands        []string `json:"commands"`
}

func (config *Config) create(configPath string) {

}
func CreateConfig(configPath string) Config {
	var config Config
	if configPath == "" {
		configPath = "/etc/ContainerFailedListener/config.json"
	}
	dirPath := filepath.Dir(configPath)

	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		err := os.MkdirAll(dirPath, os.ModePerm)
		if err != nil {
			log.Fatalf("Error creating config dirrectory: %v\n", err)
			return nil
		}
		fileutils.WriteFile("[]", configPath)
		return config
	}

	jsonResp := fileutils.ReadFile(configPath)
	err := json.Unmarshal([]byte(jsonResp), &config)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return config
}
