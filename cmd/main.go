package main

import (
	"flag"
	"log"

	"github.com/crispgm/go-van/deploy"

	"github.com/crispgm/go-van"
	"github.com/rjeczalik/notify"
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
		log.Println("Init")
	} else {
		conf, err := van.LoadFrom(confName, specName)
		if err != nil {
			log.Println("Load conf failed: ", err)
			return
		}
		log.Println(conf)
		deployer := deploy.RSync{}
		van.Watch(conf.Source, func(ei notify.EventInfo) error {
			err := deployer.Run(conf.Source, conf.Destination)
			return err
		})
	}
}
