package main

import (
	"flag"
	"fmt"
	"os"

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
	flag.StringVar(&confName, "conf", caravan.DefaultConfName, "Config file name. Default: `caravan.yml`.")
	flag.StringVar(&specName, "spec", caravan.DefaultSpec, "Spec name. Default: `master`.")
	flag.Parse()

	if initYAML {
		caravan.PrintNotice("Creating `caravan.yml`...")
		cwd, _ := os.Getwd()
		confPath := fmt.Sprintf("%s/%s", cwd, caravan.DefaultConfName)
		if _, err := os.Stat(confPath); !os.IsNotExist(err) {
			caravan.PrintError("File existed:", confPath)
			return
		}
		err := caravan.CreateDefault(confPath)
		if err != nil {
			caravan.PrintError("Create default conf failed:", err)
			return
		}
		caravan.PrintNotice("Make sure to specify `src` and `dst` to watch and deploy to right place.")
	} else {
		caravan.PrintNotice("Reading configuration...")
		conf, err := caravan.LoadFrom(confName, specName)
		if err != nil {
			caravan.PrintError("Load conf failed:", err)
			return
		}
		caravan.ShowConf(conf)
		caravan.PrintNotice("Starting to watch...")
		deployer := deploy.RSync{}

		if conf.Once {
			handleDeploy(*conf, deployer)
			return
		}

		caravan.Watch(*conf, func(ei notify.EventInfo) error {
			f := caravan.NewFilter(conf.Exclude)
			match, err := f.Exclude(ei.Path())
			if err != nil {
				caravan.PrintError("Exclude failed:", err)
				return nil
			}
			if match {
				caravan.PrintLog(ei.Path(), "is ignored")
				return nil
			}
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
