package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/ttstringiot/golangiot/stgcommon/mqversion"
	"github.com/ttstringiot/golangiot/stgregistry/registry"
	"github.com/ttstringiot/golangiot/stgregistry/start/g"
)

func main() {
	p := flag.Int("p", 0, "registry listen port")
	c := flag.String("c", "", "registry logger config file")
	v := flag.Bool("v", false, "version")
	h := flag.Bool("h", false, "help")

	flag.Parse()

	if *h {
		flag.Usage()
		os.Exit(0)
	}
	if *v {
		fmt.Println(mqversion.GetCurrentDesc())
		os.Exit(0)
	}

	registryPort := 0
	if *p > 0 {
		registryPort = *p
	}

	configPath := ""
	if strings.TrimSpace(*c) != "" {
		configPath = strings.TrimSpace(*c)
	}
	g.InitLogger(configPath)

	stopChannel := make(chan bool, 1) // the 'stopChannel' variable to handle controller.shutdownHook()
	registry.Startup(stopChannel, registryPort)

	for {
		select {
		case <-stopChannel:
			return
		}
	}
}
