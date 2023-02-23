package main

import (
	"fmt"
	"math/rand"
	"os"
)

var version = "0.0.0-dev"

var symbol = []string{
	"✔",
	"⎇ ",
	"✖ ",
	"… ",
	"⚑ ",
	"↑·",
	"↓·",
}

func main() {
	heartBeat()
}

func show_version() {
	fmt.Println(version)
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
