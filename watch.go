package main

import (
	"github.com/crispgm/go-van/caravan"
	"github.com/crispgm/go-van/deploy"
	"github.com/rjeczalik/notify"
)

func watch(conf *caravan.Conf, deployer deploy.Deployer) {
	caravan.PrintNotice("Starting to watch...")
	caravan.Watch(*conf, func(ei notify.EventInfo) error {
		f := caravan.NewFilter(conf.Exclude)
		match, err := f.Exclude(ei.Path())
		if err != nil {
			caravan.PrintError("Exclude failed:", err)
			return nil
		}
		if match {
			caravan.PrintLog("IGNORE", ei.Path())
			return nil
		}
		eventCtrl.FireEvent(caravan.NewEvent(caravan.HookOnChange, 0, ei.Path(), caravan.GetFileName(ei.Path())))
		return handleDeploy(*conf, deployer)
	})
}

func handleDeploy(conf caravan.Conf, deployer deploy.Deployer) error {
	output, err := deployer.Run(conf.Source, conf.Destination, conf.ExtraArgs)
	if err != nil {
		caravan.WarningSound()
		caravan.PrintError(string(output))
		eventCtrl.FireEvent(caravan.NewEvent(caravan.HookOnError))
	}
	eventCtrl.FireEvent(caravan.NewEvent(caravan.HookOnDeploy))
	return err
}
