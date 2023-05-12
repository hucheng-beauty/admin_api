package utils

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

type generator struct {
	lastFlag string
	loc      sync.Mutex
	c        int
}

func NewGenerator() *generator {
	return &generator{}
}

var DefaultGenerator = GenerateNu()

func (g *generator) gen() string {
	g.loc.Lock()

	g.c++

	n := time.Now()
	s := fmt.Sprintf("%s%s%s",
		string(byte(time.Now().Hour()%24+65)),
		fmt.Sprintf("%02d", n.Minute()),
		fmt.Sprintf("%02d", n.Second()),
	)
	if g.lastFlag == s && g.c >= 1000 {
		time.Sleep(time.Second)
		g.loc.Unlock()
		return g.gen()
	}
	g.lastFlag = s
	if g.c >= 1000 {
		g.c = 1
	}
	s = s + fmt.Sprintf("%03d", g.c)
	g.loc.Unlock()
	return s
}

// 生成活动编号
func GenerateNu() func(prefix string) string {
	g := NewGenerator()
	return func(prefix string) string {
		n := time.Now()
		return fmt.Sprintf("%s%s%02d%02d%s",
			prefix,
			strconv.Itoa(n.Year())[2:],
			n.Month(),
			n.Day(),
			g.gen())
	}
}
