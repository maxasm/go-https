package main

import (
	"log"
	"os"
)

var dl *log.Logger = log.New(os.Stdout, "[DEBUG] ",log.Lshortfile)

func main() {
	dl.Printf("let's learn about https and tls in go lang.")
}
