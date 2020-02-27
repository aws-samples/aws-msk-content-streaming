package main

import (
	"math/rand"
	"time"

	"github.com/katallaxie/content_streaming_msk/create/cmd"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	cmd.Execute()
}
