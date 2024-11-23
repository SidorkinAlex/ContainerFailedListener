package main

import "github.com/SidorkinAlex/CyclicCommandCheckerAndExecutive/internal/CliParamsParser"

func main() {
	var cliParser = CliParamsParser.NewCliParams()
	cliParser.StartAction()

}
