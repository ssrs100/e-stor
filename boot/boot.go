package main

import (
	"estor/utils"
	"flag"
	"log"
	"os"
)

func main() {
	var ver bool
	flag.BoolVar(&ver, "version", false, "print version.")
	if ver {
		utils.PrintVersion()
		os.Exit(0)
	}
	log.Println("Proc finished normally.")
}
