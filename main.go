package main

import (
	"log"
	"os"

	"github.com/opposite-bracket/m2s/recorder"
	"github.com/opposite-bracket/m2s/utils"
)

// list command runners
var cmds = []utils.CommandOptions{
	recorder.RunnerOptions,
}

func main() {
	args := os.Args[1:]

	if err := utils.CreateOutDir(); err != nil {
		log.Panicf("failed to create output dir: [error: %s]", err)
	}

	// register commands
	for _, cmd := range cmds {
		utils.PanicOnError(
			"failed to run command: [error: %v]",
			utils.Service.RegisterCommand(cmd),
		)
	}

	// run forrest, run!
	utils.PanicOnError(
		"failed to run command: [error: %v]",
		utils.Service.RunCommand(args...),
	)
}
