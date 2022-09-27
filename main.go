package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/conflabermits/health_checker/hcfunc"
)

type Options struct {
	HostHeader string
	Url        string
	Depth      string
	Server     bool
	Port       string
}

func parseArgs() (*Options, error) {
	options := &Options{}

	flag.StringVar(&options.HostHeader, "hostHeader", "", "override Host specified in URL")
	flag.StringVar(&options.Url, "url", "", "url to check")
	flag.StringVar(&options.Depth, "depth", "dynamic", "Determine amount/type of data to return")
	flag.BoolVar(&options.Server, "server", false, "Run as local web server")
	flag.StringVar(&options.Port, "port", "8080", "Port to run the local web server")
	flag.Usage = func() {
		fmt.Printf("Usage: health_checker [options]\n\n")
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

	//Depth flag error checking
	if options.Depth != "dynamic" && options.Depth != "short" && options.Depth != "full" {
		fmt.Println("Error: Depth flag not understood. Must be one of the following: ['dynamic', 'short', 'full']")
		return
	}

	if len(options.Url) > 0 {
		response := hcfunc.Health_checker_http_req(options.Url, options.HostHeader)
		output := hcfunc.Parse_health_checker_json(response, options.Depth)
		fmt.Println(output)
	}

	if options.Server {
		hcfunc.Web(options.Port)
	}
}
