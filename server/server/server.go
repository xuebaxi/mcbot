package server

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"sync"

	log "github.com/sirupsen/logrus"
)

var access sync.RWMutex
var Server *server

type server struct {
	Server *exec.Cmd
	Input  io.WriteCloser
	Output io.ReadCloser
	Error  io.ReadCloser
	Status map[string]interface{}
}

func (s *server) init() {
	var err error
	s.Server = exec.Command("java", "-jar", "server.jar")
	s.Input, err = s.Server.StdinPipe()
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}
	s.Output, err = s.Server.StdoutPipe()
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}
	s.Error, err = s.Server.StderrPipe()
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}
	err = s.Server.Start()
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}
	if s.Status == nil {
		s.Status = map[string]interface{}{}
	}
}
func (s *server) Run(cmd string) {
	access.Lock()
	defer access.Unlock()
	io.WriteString(s.Input, fmt.Sprintf("%s\n", cmd))

}
func (s *server) Start() {
	if s.Status["running"].(bool) {
		return
	} else {
		s.init()
	}
}
func newserver() *server {
	s := new(server)
	s.init()
	return s
}
func init() {
	if Server == nil {
		Server = newserver()
	}
}
