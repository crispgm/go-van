package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/crispgm/go-van/caravan"
	"github.com/crispgm/go-van/deploy"
)

var (
	errFileExisted     = errors.New("Conf file existed")
	errUnsupportedMode = errors.New("Unsupported deploy mode")
)

func initConf() error {
	caravan.PrintNotice("Creating `caravan.yml`...")
	cwd, _ := os.Getwd()
	confPath := fmt.Sprintf("%s/%s", cwd, caravan.DefaultConfName)
	_, err := os.Stat(confPath)
	if !os.IsNotExist(err) {
		caravan.PrintError("File existed:", confPath)
		return errFileExisted
	}
	err = caravan.CreateDefault(confPath)
	if err != nil {
		caravan.PrintError("Create default conf failed:", err)
		return err
	}
	caravan.PrintNotice("Make sure to specify `src` and `dst` to watch and deploy to right place.")
	return nil
}

func parseConfAndWatch() error {
	caravan.PrintNotice("Reading configuration...")
	conf, err := caravan.LoadFrom(confName, specName)
	if err != nil {
		caravan.PrintError("Load conf failed:", err)
		return err
	}
	caravan.ShowConf(conf)
	deployer := deploy.NewDeployer(conf.Mode)
	if deployer == nil {
		caravan.PrintError("Unsupported deploy mode:", conf.Mode)
		return errUnsupportedMode
	}

	if conf.Once || deployOnce {
		caravan.PrintNotice("Deploying at once and for once...")
		return handleDeploy(*conf, deployer)
	}

	watch(conf, deployer)
	return nil
}
