package m2s

import (
	"github.com/opposite-bracket/m2s/utils"
	"log"
	"net/http"
)

//
// VARs & CONSTs
//

//
// STRUCTS
//

// m2Service is the entry point for recording
// source/target service communication
type m2Service struct{
	Config *utils.M2sConfig
}

//
// METHODS
//

// ServeHTTP processes HTTP Requests with
// the objective of collecting specs for each
// API endpoint received.
func (srv m2Service) ServeHTTP(_ http.ResponseWriter, req *http.Request) {
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

// RunServer will start the server which logs
// all logs said server instance. This will
// create a set of logs which can be used as
// a starting point for mocking endpoints
// between services.
func (srv *m2Service) RunServer() error {

	log.Printf("Running on %s", srv.Config.Address)
	return http.ListenAndServe(srv.Config.Address, srv)
}

//
// FUNCTIONS
//

// NewServer instantiate m2 service instance
func NewService() (*m2Service, error) {
	config, err := utils.GetConf()
	if err != nil {
		return nil, err
	}

	srv := m2Service{
		Config: config,
	}

	return &srv, nil
}
