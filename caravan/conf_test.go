package caravan

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConf(t *testing.T) {
	var err error
	_, err = LoadFrom("../fixtures/caravan_not_existed.yml", "master")
	assert.Error(t, err)
	_, err = LoadFrom("../fixtures/caravan.yml", "master")
	assert.NoError(t, err)
	_, err = LoadFrom("../fixtures/caravan.yml", "not_master")
	assert.Error(t, err)
	_, err = LoadFrom("../fixtures/caravan_wrong.yml", "master")
	assert.Error(t, err)
}

func TestLoadFallbackConf(t *testing.T) {
	var err error
	spec, err := LoadFrom("../fixtures/caravan_fallback.yml", "master")
	assert.NoError(t, err)
	assert.Equal(t, "/path/to/fallback", spec.Source)
	_, err = LoadFrom("../fixtures/caravan_fallback_error.yml", "master")
	assert.Error(t, err)
}

func TestCreateDefault(t *testing.T) {
	err := CreateDefault("../fixtures/testoutput/caravan.yml")
	assert.NoError(t, err)
	assert.FileExists(t, "../fixtures/testoutput/caravan.yml")
}
