package main

import (
	"time"

	"github.com/briandowns/spinner"
	"github.com/crispgm/go-van/caravan"
	"github.com/crispgm/go-van/deploy"
	"github.com/rjeczalik/notify"
)

var s = spinner.New(spinner.CharSets[36], 100*time.Millisecond)

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
			caravan.PrintLog(ei.Path(), "is ignored")
			return nil
		}
		return handleDeploy(*conf, deployer)
	})
}

func handleDeploy(conf caravan.Conf, deployer deploy.Deployer) error {
	s.Start()
	output, err := deployer.Run(conf.Source, conf.Destination)
	s.Stop()
	if err != nil {
		caravan.PrintError(output)
	}
	return err
}
