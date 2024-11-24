package RunController

import (
	"fmt"
	"github.com/SidorkinAlex/CyclicCommandCheckerAndExecutive/internal/fileutils"
	"log"
	"os"
	"time"
)

const RunDirName = "/var/run/CyclicCommandCheckerAndExecutive/"
const LogDir = "/var/log/CyclicCommandCheckerAndExecutive/"
const StopFile = ".stop.pid"
const PidFile = ".sample.pid"

func RunningController() {
	if fileutils.HasFile(RunDirName + StopFile) {
		err := os.Remove(RunDirName + StopFile)
		if err != nil {
			log.Fatal("the stop file cannot be deleted. The program CyclicCommandCheckerAndExecutive exited with an error. ")
			return
		}
	}
	log.Println("Program CyclicCommandCheckerAndExecutive has been success running")
	for {
		if _, err := os.Stat(RunDirName + StopFile); err == nil {
			fileutils.WriteFile("stopped", RunDirName+StopFile)
			RemovePidFile()
			fmt.Println("")
			log.Fatalln("service CyclicCommandCheckerAndExecutive stopped")
		}
		time.Sleep(5 * time.Second)
	}
}
func RemovePidFile() {
	fileutils.WriteFile("", RunDirName+PidFile)
}
