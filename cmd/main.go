package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/crispgm/go-van"
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
		van.Watch(conf.Source, func(ei notify.EventInfo) error {
			van.PrintNotice(getTime(), "Event", ei.Event().String, ei.Path())
			output, err := deployer.Run(conf.Source, conf.Destination)
			if err != nil {
				van.PrintSuccess(output)
			}
			return err
		})
	}
}

func getTime() string {
	t := time.Now()
	return fmt.Sprintf("%02d:%02d:%02d", t.Hour(), t.Minute(), t.Second())
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
