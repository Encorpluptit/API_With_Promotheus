package server

import (
	"BackendGo/controllers"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)

type Server struct {
	Router        *gin.Engine
	Controller    *controllers.BaseController
	PrometheusURL string
}

func NewServer() (*Server, error) {
	var err error
	server := &Server{Router: gin.Default()}
	if server.Controller, err = controllers.NewBaseController(); err != nil {
		return nil, err
	}
	if server.PrometheusURL = os.Getenv("PROMETHEUS_URL"); server.PrometheusURL == "" {
		return nil, fmt.Errorf("environment variable 'PROMETHEUS_URL' not defined")
	}
	return server, nil
}

// curl -g 'http://localhost:9090/api/v1/series?' --data-urlencode 'match[]=process_start_time_seconds{owner="4d365568-d3ac-4880-8403-8d4e2638e008"}' | json_pp
// curl -G http://localhost:9090/api/v1/metadata\?metric\=go_memstats_frees_total --data-urlencode 'match_target={owner="4d365568-d3ac-4880-8403-8d4e2638e008"}' | json_pp
