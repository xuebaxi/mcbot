package load

import (
	"bufio"
	"os"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/xuebaxi/mcbot/server/command"
	"github.com/xuebaxi/mcbot/server/server"
)

//Loader 函数用于从管道中读取日志
func Loader(log *logrus.Logger) {
	evenObj := new(even)
	evenObj.init(log)
	log.SetOutput(os.Stdout)

	log.Info("开始记录日志")

	go func() {
		for {
			serverInfo := bufio.NewScanner(server.Server.Output)
			for serverInfo.Scan() {
				evenObj.Save(serverInfo.Text())
			}
			time.Sleep(5)
		}
	}()
	userInput := bufio.NewScanner(os.Stdin)
	for userInput.Scan() {
		if userInput.Text()[0] == '/' {
			server.Server.Run(userInput.Text())
		} else {
			//服务器命令
			com := strings.Split(userInput.Text(), " ")
			if function, ok := command.Command[com[0]]; ok {
				log.Info(function(com))
			} else {
				log.Warn("命令没找到")
			}
		}
	}
}
