package main

import (
	"math/rand"
	"time"

	"github.com/motain/xp-monet/ingest/cmd"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	cmd.Execute()
}
