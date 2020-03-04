package main

import (
	"math/rand"
	"time"

	"github.com/katallaxie/content_streaming_msk/server/cmd"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	cmd.Execute()
}
