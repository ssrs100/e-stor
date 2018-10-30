package main

import (
	"estor"
	"estor/utils"
	"flag"
	"log"
	"os"
)

func main() {
	var ver bool
	flag.BoolVar(&ver, "version", false, "print version.")
	flag.Parse()
	if ver {
		utils.PrintVersion()
		os.Exit(0)
	}
	estor.Start()
	log.Println("Proc finished normally.")
}
