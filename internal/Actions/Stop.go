package Actions

import (
	"fmt"
	"github.com/SidorkinAlex/CyclicCommandCheckerAndExecutive/internal/RunController"
	"github.com/SidorkinAlex/CyclicCommandCheckerAndExecutive/internal/fileutils"
	"log"
	"time"
)

func Stop() {
	if !fileutils.HasFile(RunController.RunDirName + RunController.PidFile) {
		log.Println("App not running")
	} else {
		fmt.Printf("%s", "stopping app in progress")
		fileutils.WriteFile("", RunController.RunDirName+RunController.StopFile)
		// TODO сделать проверку на запущенный процесс с пидом указанным в файлике старта
		for {
			isStoped := fileutils.ReadFile(RunController.RunDirName + RunController.StopFile)
			fmt.Printf("%s", ".")
			if isStoped == "stopped" {
				fmt.Println("")
				log.Println("app is stoped")
				break
			}
			time.Sleep(1 * time.Second)
		}
	}
}
