package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/crispgm/go-van/caravan"
	"github.com/stretchr/testify/assert"
)

func getRealPath(p string) string {
	cwd, _ := os.Getwd()
	return fmt.Sprintf("%s/%s", cwd, p)
}

func initArgs() {
	confName = getRealPath("caravan.yml")
	specName = "master"
	initYAML = false
	deployOnce = false
}

func TestInitConf(t *testing.T) {
	confPath := getRealPath(caravan.DefaultConfName)
	os.Remove(confPath)
	err := initConf()
	assert.NoError(t, err)
	assert.FileExists(t, confPath)
	err = initConf()
	assert.Error(t, errFileExisted, err)
	os.Remove(confPath)
}

func TestDeployOnceWithConf(t *testing.T) {
	initArgs()
	os.Remove(getRealPath("caravan.yml"))
	os.Remove(getRealPath("fixtures/testoutput/file.txt"))
	caravan.DefaultConf.Once = true
	caravan.DefaultConf.Source = getRealPath("fixtures/testsrc/")
	caravan.DefaultConf.Destination = getRealPath("fixtures/testoutput/")
	caravan.CreateDefault(getRealPath("caravan.yml"))
	parseConfAndWatch(false)
	assert.FileExists(t, getRealPath("fixtures/testoutput/file.txt"))
	os.Remove(getRealPath("fixtures/testoutput/file.txt"))
	os.Remove(getRealPath("caravan.yml"))
}

func TestDeployOnceWithArgs(t *testing.T) {
	initArgs()
	deployOnce = true
	os.Remove(getRealPath("caravan.yml"))
	os.Remove(getRealPath("fixtures/testoutput/file.txt"))
	caravan.DefaultConf.Source = getRealPath("fixtures/testsrc/")
	caravan.DefaultConf.Destination = getRealPath("fixtures/testoutput/")
	caravan.CreateDefault(getRealPath("caravan.yml"))
	parseConfAndWatch(false)
	assert.FileExists(t, getRealPath("fixtures/testoutput/file.txt"))
	os.Remove(getRealPath("fixtures/testoutput/file.txt"))
	os.Remove(getRealPath("caravan.yml"))
}

func TestLoadConfFail(t *testing.T) {
	initArgs()
	err := parseConfAndWatch(false)
	assert.Error(t, err)
}

func TestUnsupportedMode(t *testing.T) {
	initArgs()
	os.Remove(getRealPath("caravan.yml"))
	caravan.DefaultConf.Mode = "scp"
	caravan.CreateDefault(getRealPath("caravan.yml"))
	err := parseConfAndWatch(false)
	assert.Error(t, errUnsupportedMode, err)
	os.Remove(getRealPath("caravan.yml"))
}

func TestInpsect(t *testing.T) {
	initArgs()
	confPath := getRealPath(caravan.DefaultConfName)
	os.Remove(confPath)
	err := initConf()
	assert.NoError(t, err)
	output := caravan.CaptureOutput(func() { parseConfAndWatch(true) })
	assert.Contains(t, output, "Reading configuration..")
	assert.Contains(t, output, "=> deploy_mode: scp")
	os.Remove(getRealPath("caravan.yml"))
}
