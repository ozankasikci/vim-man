package main

import (
	"github.com/ozankasikci/vim-man"
	"os"
	"strconv"
)

func main() {
	level := os.Getenv("LEVEL")
	levelInt, err := strconv.ParseInt(level, 10, 16)
	if err != nil {
		levelInt = 1
	}

	vimman.Init(int(levelInt))
}
