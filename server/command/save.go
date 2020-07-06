package command

import (
	"os/exec"
	"strconv"
	"time"
)

func save(s []string) string {
	filename := strconv.FormatInt(time.Now().Unix(), 10)
	c := exec.Command("mkdir", filename)
	c.Run()
	c = exec.Command("zip", "-s", s[1], "-r", filename+"/world.zip", "world")
	err := c.Run()
	if err != nil {
		return ""
	}
	return filename
}
func init() {
	Command["save"] = save
}
