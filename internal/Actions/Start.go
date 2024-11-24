package Actions

import (
	"github.com/SidorkinAlex/CyclicCommandCheckerAndExecutive/internal/Config"
	"github.com/SidorkinAlex/CyclicCommandCheckerAndExecutive/internal/RunController"
	"github.com/SidorkinAlex/CyclicCommandCheckerAndExecutive/internal/fileutils"
	"github.com/sevlyar/go-daemon"
	"log"
	"os/exec"
	"time"
)

func Start() {
	CommandConfig := Config.CreateConfig("")
	//todo приложениие теряет контекст и аргументы , с которыми оно было запущено здесь дебажить конфиг и ключи
	log.Println("CommandConfig = ")
	log.Println(CommandConfig)

	fileutils.CheckAndCreatDir(RunController.RunDirName)
	fileutils.CheckAndCreatDir(RunController.LogDir)
	cntxt := &daemon.Context{
		PidFileName: RunController.RunDirName + RunController.PidFile,
		PidFilePerm: 0644,
		LogFileName: RunController.LogDir + ".sample.log",
		LogFilePerm: 0640,
		WorkDir:     "/",
		Umask:       027,
		Args:        []string{"[go-daemon CyclicCommandCheckerAndExecutive] -start"},
	}

	d, err := cntxt.Reborn()
	if err != nil {
		log.Fatal("Unable to run: ", err)
	}
	if d != nil {
		return
	}
	defer cntxt.Release()

	runCyclicCommand(CommandConfig)
}
func runCyclicCommand(CommandConfig Config.Config) {
	for _, command := range CommandConfig {
		//runCommand(command)
		go runCommand(command)
	}
	RunController.RunningController()
}

func runCommand(command Config.Command) {
	for {
		commandResult, _ := exec.Command("/bin/sh", "-c", command.Command).Output()

		log.Println("command from command`" + command.Command + "`   result is `" + string(commandResult) + "`")
		for _, branch := range command.BranchCommand {
			log.Println("branch.ResultExecution=`" + branch.ResultExecution + "`   commandResult=`" + string(commandResult) + "`")
			if branch.ResultExecution == string(commandResult) {
				executeCommands(branch.Commands)
			}
		}
		if string(commandResult) == "" {
			log.Fatalln("test")
		}
		log.Println(string(commandResult))
		time.Sleep(time.Duration(command.Interval) * time.Second)
	}
}

func executeCommands(commands []string) {
	for _, cmd := range commands {
		result, _ := exec.Command("/bin/sh", "-c", cmd).Output()
		log.Println("command:" + cmd + "execute from result:" + string(result))
	}
}
