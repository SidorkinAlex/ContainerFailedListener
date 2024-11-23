package CliParamsParser

import (
	"flag"
	"github.com/SidorkinAlex/CyclicCommandCheckerAndExecutive/internal/Actions"
	"log"
)

type CliParams struct {
	Action string
}

const (
	Start     = "start"
	Stop      = "stop"
	CreateJob = "createJob"
	Restart   = "restart"
)

func NewCliParams() *CliParams {
	var c CliParams
	var stop bool
	var start bool
	var createJob bool
	var restart bool

	flag.BoolVar(&stop, "stop", false, "set this param to stopping demon")
	flag.BoolVar(&start, "start", false, "set this param from start program")
	flag.BoolVar(&createJob, "create-job", false, "set this param from start program")
	flag.BoolVar(&restart, "restart", false, "set this param from start program")
	if stop {
		c.Action = Start
	}
	if stop {
		c.Action = Stop
	}
	if createJob {
		c.Action = CreateJob
	}
	if restart {
		c.Action = Restart
	}
	if c.Action == "" {
		log.Fatal("To run the program, use one of the keys -stop, -start, -create-job, -restart is required.")
	}
	return &c
}
func (c CliParams) StartAction() {
	if c.Action == Stop {
		Actions.Stop()
	}
	if c.Action == Start {
		Actions.Start()
	}
	if c.Action == CreateJob {
		Actions.CreateJob()
	}
	if c.Action == Restart {
		Actions.Restart()
	}
}
