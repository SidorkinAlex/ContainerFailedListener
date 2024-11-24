package Actions

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/SidorkinAlex/CyclicCommandCheckerAndExecutive/internal/Config"
	"github.com/SidorkinAlex/CyclicCommandCheckerAndExecutive/internal/fileutils"
	"os"
	"strconv"
	"strings"
)

func CreateJob() {
	configPath := Config.ConfigPath
	config := Config.CreateConfig(configPath)

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("1. Create a new command for monitoring")
		fmt.Println("2. Save changes and exit")
		fmt.Print("Choose an option: ")
		option, _ := reader.ReadString('\n')
		option = strings.TrimSpace(option)

		if option == "2" {
			break
		}

		if option == "1" {
			var newCommand Config.Command

			fmt.Print("Enter the command for cyclic execution and monitoring: ")
			command, _ := reader.ReadString('\n')
			newCommand.Command = strings.TrimSpace(command)

			fmt.Print("Enter the time interval in seconds for executing the above command: ")
			intervalStr, _ := reader.ReadString('\n')
			interval, err := strconv.Atoi(strings.TrimSpace(intervalStr))
			if err != nil {
				fmt.Println("Invalid interval. Please enter a number.")
				continue
			}
			newCommand.Interval = interval

			for {
				fmt.Println("1. Add a new expected result")
				fmt.Println("2. Finish configuration and proceed to save")
				fmt.Print("Choose an option: ")
				resultOption, _ := reader.ReadString('\n')
				resultOption = strings.TrimSpace(resultOption)

				if resultOption == "2" {
					break
				}

				if resultOption == "1" {
					var branch Config.BranchResultExecution

					fmt.Print("Enter the expected result: ")
					result, _ := reader.ReadString('\n')
					branch.ResultExecution = strings.TrimSpace(result)

					for {
						fmt.Println("1. Enter a command to execute if the result matches")
						fmt.Println("2. Create a new expected result")
						fmt.Print("Choose an option: ")
						commandOption, _ := reader.ReadString('\n')
						commandOption = strings.TrimSpace(commandOption)

						if commandOption == "2" {
							break
						}

						if commandOption == "1" {
							fmt.Print("Enter the command to execute if the result matches: ")
							execCommand, _ := reader.ReadString('\n')
							branch.Commands = append(branch.Commands, strings.TrimSpace(execCommand))
						}
					}

					newCommand.BranchCommand = append(newCommand.BranchCommand, branch)
				}
			}

			config = append(config, newCommand)
		}
	}

	// Display and save the new configuration
	configData, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling config:", err)
		return
	}

	fmt.Println("New configuration:")
	fmt.Println(string(configData))

	fileutils.RewriteFile(string(configData), configPath)
	fmt.Println("Configuration saved successfully.")
}
