package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/conflabermits/health_checker/hcfunc"
)

type Options struct {
	Port string
}

func parseArgs() (*Options, error) {
	options := &Options{}

	flag.StringVar(&options.Port, "port", "8080", "Port to run the local web server")
	flag.Usage = func() {
		fmt.Printf("Usage: health_checker-web [options]\n\n")
		flag.PrintDefaults()
	}
	flag.Parse()

	return options, nil
}

func main() {
	options, err := parseArgs()
	if err != nil {
		os.Exit(1)
	}

	hcfunc.Web(options.Port)
}
