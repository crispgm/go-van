package main

import (
	"flag"
	"os"
	"os/signal"

	"github.com/crispgm/go-van/caravan"
)

var (
	confName   string
	specName   string
	initYAML   bool
	inspect    bool
	deployOnce bool
	isVersion  bool
)

func main() {
	flag.BoolVar(&initYAML, "init", false, "Generate caravan.yml in current path.")
	flag.BoolVar(&inspect, "inspect", false, "Inspect config.")
	flag.StringVar(&confName, "conf", caravan.DefaultConfName, "Config file name. Default: `caravan.yml`.")
	flag.StringVar(&specName, "spec", caravan.DefaultSpec, "Spec name. Default: `master`.")
	flag.BoolVar(&deployOnce, "once", false, "Deploy once. Default: false")
	flag.BoolVar(&isVersion, "version", false, "Show version info")
	flag.Parse()

	if isVersion {
		version()
		return
	}

	var err error
	if initYAML {
		err = initConf()
	} else {
		go func() {
			sig := make(chan os.Signal, 10)
			signal.Notify(sig, os.Interrupt)
			<-sig
			caravan.PrintNotice("\nInterrupted by <CTRL+C>")

			os.Exit(0)
		}()
		err = parseConfAndWatch(inspect)
	}

	if err != nil {
		os.Exit(1)
	}
}
