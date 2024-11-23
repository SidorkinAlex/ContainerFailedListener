package Actions

import (
	"github.com/SidorkinAlex/CyclicCommandCheckerAndExecutive/internal/Config"
	"log"
)

func Start() {
	config := Config.CreateConfig("")
	log.Print(config)
}
