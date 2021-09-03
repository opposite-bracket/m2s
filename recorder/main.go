package recorder

import (
	"github.com/opposite-bracket/m2s/utils"
	"log"
	"net/http"
)

const address = ":5000"

// Runner is the entry point for recording
// source/target service communication
type Runner struct{}

// RunServer will start the server which logs
// all logs said server instance. This will
// create a set of logs which can be used as
// a starting point for mocking endpoints
// between services.
func RunServer() {
	srv := Runner{}

	log.Printf("Running on %s", address)
	if err := http.ListenAndServe(address, srv); err != nil {
		log.Fatalf("server errored: %v", err)
	}
}

// RunnerOptions configures the execution of the api recorder
var RunnerOptions = utils.CommandOptions{
	Cmd: "recorder",
	Runner: func(args ...string) error {
		RunServer()
		return nil
	},
}

// ServeHTTP processes HTTP Requests with
// the objective of collecting specs for each
// API endpoint received.
func (p Runner) ServeHTTP(_ http.ResponseWriter, req *http.Request) {
	// extract req info
	e2eInfo := utils.E2EInfo{
		Req: utils.ExtractReqInfo(req),
	}

	// if target is available,
	// make call to target
	if e2eInfo.Req.Target != nil {
		client := utils.NewClient()
		if res, err := client.MakeRequest(e2eInfo.Req); err != nil {
			log.Printf("failed to load response: [%s]", err)
		} else {
			e2eInfo.Res = res
		}
	}

	utils.LogReq(e2eInfo)
}
