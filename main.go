package main

import (
	"flag"

	"github.com/crispgm/go-van/caravan"
	"github.com/crispgm/go-van/deploy"
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
		caravan.PrintNotice("Init")
	} else {
		caravan.PrintNotice("Reading configuration...")
		conf, err := caravan.LoadFrom(confName, specName)
		if err != nil {
			caravan.PrintError("Load conf failed:", err)
			return
		}
		showConf(conf)
		caravan.PrintNotice("Starting to watch...")
		deployer := deploy.RSync{}

		if conf.Once {
			handleDeploy(*conf, deployer)
			return
		}

		caravan.Watch(*conf, func(ei notify.EventInfo) error {
			return handleDeploy(*conf, deployer)
		})
	}
}

func handleDeploy(conf caravan.Conf, deployer deploy.Deployer) error {
	output, err := deployer.Run(conf.Source, conf.Destination)
	if err != nil {
		caravan.PrintSuccess(output)
	}
	return err
}

func showConf(conf *caravan.Conf) {
	if conf == nil {
		return
	}
	caravan.PrintNotice("=>", "debug:", conf.Debug)
	caravan.PrintNotice("=>", "once:", conf.Once)
	caravan.PrintNotice("=>", "src:", conf.Source)
	caravan.PrintNotice("=>", "dst:", conf.Destination)
	caravan.PrintNotice("=>", "deploy_mode:", conf.Mode)
	caravan.PrintNotice("=>", "incremental:", conf.Incremental)
	caravan.PrintNotice("=>", "exclude:", conf.Exclude)
}
