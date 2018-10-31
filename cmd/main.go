package main

import (
	"flag"

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
		van.PrintNotice("Init")
	} else {
		van.PrintNotice("Reading configuration...")
		conf, err := van.LoadFrom(confName, specName)
		if err != nil {
			van.PrintError("Load conf failed:", err)
			return
		}
		showConf(conf)
		van.PrintNotice("Starting to watch...")
		deployer := deploy.RSync{}

		if conf.Once {
			return deploy(conf)
		}

		van.Watch(conf, func(ei notify.EventInfo) error {
			return deploy(conf)
		})
	}
}

func deploy(conf van.Conf) {
	output, err := deployer.Run(conf.Source, conf.Destination)
	if err != nil {
		van.PrintSuccess(output)
	}
	return err
}

func showConf(conf *van.Conf) {
	if conf == nil {
		return
	}
	van.PrintNotice("=>", "debug:", conf.Debug)
	van.PrintNotice("=>", "once:", conf.Once)
	van.PrintNotice("=>", "src:", conf.Source)
	van.PrintNotice("=>", "dst:", conf.Destination)
	van.PrintNotice("=>", "deploy_mode:", conf.Mode)
	van.PrintNotice("=>", "incremental:", conf.Incremental)
	van.PrintNotice("=>", "exclude:", conf.Exclude)
}
