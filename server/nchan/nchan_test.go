package nchan_test

import (
	"testing"

	"github.com/xuebaxi/mcbot/server/nchan"
)

func TestWrite(t *testing.T) {
	c := make(chan string, 1)
	c <- "1"
	if nchan.WriteString(c, "2") {
		panic("nchan.WriteString()函数异常")
	}
	nchan.CoverString(c, "3")
	if <-c != "3" {
		panic("nchan.CoverString()函数异常")
	}
}

func TestRead(t *testing.T) {
	c := make(chan string, 3)
	if _, ok := nchan.ReadString(c); ok {
		panic("nchan.ReadString()函数异常")
	}
	c <- "1"
	c <- "2"
	c <- "3"
	if s, ok := nchan.ReadString(c); ok {
		if s != "1" {
			panic("nchan.ReadString()函数异常")
		}
	} else {
		panic("nchan.ReadString()函数异常")
	}

	s := nchan.AllString(c)
	if s[0] != "2" || s[1] != "3" {
		panic("nchan.AllString()函数异常")
	}
}
