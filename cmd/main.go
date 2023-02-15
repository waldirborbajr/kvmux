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
	fmt.Println(version)
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
