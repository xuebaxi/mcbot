package server_test

import (
	"bufio"
	"fmt"
	"io"
	"testing"

	log "github.com/sirupsen/logrus"
	"github.com/xuebaxi/mcbot/server/server"
)

func TestServer(t *testing.T) {
	c := make(chan int)
	out := bufio.NewScanner(server.Server.Output)
	go func() {
		for out.Scan() {
			log.Info(out.Text())
			c <- 1
			return
		}
	}()
	io.WriteString(server.Server.Input, fmt.Sprintf("test\n"))
	<-c
}
