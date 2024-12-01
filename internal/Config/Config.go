package Config

import (
	"encoding/json"
	"fmt"
	"github.com/SidorkinAlex/CyclicCommandCheckerAndExecutive/internal/fileutils"
	"log"
	"os"
	"path/filepath"
)

const ConfigPath = "/etc/CyclicCommandCheckerAndExecutive/config.json"

type Config []Command

// Config структура для хранения конфигурации
type Command struct {
	Command  string `json:"checkingCommand"`
	Interval int    `json:"interval"`

	BranchCommand []BranchResultExecution
}
type BranchResultExecution struct {
	ResultExecution string   `json:"resultExecution"`
	TypeOfMatch     string   `json:"typeOfMatch"`
	Commands        []string `json:"commands"`
}

func (config *Config) create(configPath string) {

}
func CreateConfig(configPath string) Config {
	var config Config
	if configPath == "" {
		configPath = ConfigPath
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
	validationConfig(config)

	return config
}

func validationConfig(config Config) {
	validTypes := map[string]struct{}{
		"equival":               {},
		"not_equival":           {},
		"included_in_substring": {},
		"more":                  {},
		"less":                  {},
		"":                      {},
	}

	for _, command := range config {
		for _, branch := range command.BranchCommand {
			if _, ok := validTypes[branch.TypeOfMatch]; !ok {
				log.Fatalln("Invalid TypeOfMatch: %s in command: %s\n", branch.TypeOfMatch, command.Command)
			}
		}
	}
}
