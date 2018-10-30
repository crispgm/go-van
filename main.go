package main

import (
	"flag"
	"fmt"
)

var (
	confName string
	specName string
	initYAML bool
)

func main() {
	flag.BoolVar(&initYAML, "init", false, "Generate caravan.yml in current path.")
	flag.StringVar(&confName, "conf", "caravan.yml", "Config file name. Default: `caravan.yml`.")
	flag.StringVar(&specName, "spec", "master", "Spec name. Default: `master`.")
	flag.Parse()

	if initYAML {
		fmt.Println("Init")
	} else {
		conf, err := LoadFrom(confName, specName)
		if err != nil {
			fmt.Println("Load conf failed: ", err)
			return
		}
		fmt.Println(conf)
	}
}
