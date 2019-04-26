package main

import (
	"testing"

	"github.com/rjeczalik/notify"
	"github.com/stretchr/testify/assert"
)

func TestDeployOnceStruct(t *testing.T) {
	ei := deployOnceEI{
		SourcePath: "test",
	}
	assert.Equal(t, "test", ei.Path())
	assert.Nil(t, ei.Sys())
	assert.Equal(t, ei.Event(), notify.Event(0))
}
