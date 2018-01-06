package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExe(t *testing.T) {
	_, err := GetExe()
	assert.Empty(t, err, "%v", err)
}
