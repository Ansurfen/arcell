package test

import (
	"arcell/utils"
	"testing"
)

func TestLog(t *testing.T) {
	utils.Log("Hello world")
	utils.SetMode(utils.RELEASE)
	utils.Log("Hello world")
	utils.SetMode(3)
	utils.Log("Hello world")
}
