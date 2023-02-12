package main

import (
	"fmt"
	"math/rand"
	"os"
)

var symbol = []string{
	"✔",
	"⎇ ",
	"✖ ",
	"… ",
	"⚑ ",
	"↑·",
	"↓·",
}

func init() {
}

func main() {
	heartBeat()
}

func heartBeat() {
	tag := " " + symbol[rand.Intn(len(symbol))] + " "
	if _, err := fmt.Fprint(os.Stdout, tag); err != nil {
		must(err)
	}
}

func must(err error) {
	if err == nil {
		return
	}
	os.Exit(1)
}
