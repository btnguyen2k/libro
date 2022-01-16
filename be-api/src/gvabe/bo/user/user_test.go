package user

import (
	"testing"

	"github.com/btnguyen2k/henge"
	"main/src/gvabe/bo"
)

type TestSetupOrTeardownFunc func(t *testing.T, testName string)

func setupTest(t *testing.T, testName string, extraSetupFunc, extraTeardownFunc TestSetupOrTeardownFunc) func(t *testing.T) {
	bo.UboTimeLayout = henge.DefaultTimeLayout
	bo.UboTimestampRouding = henge.DefaultTimestampRoundSetting

	if extraSetupFunc != nil {
		extraSetupFunc(t, testName)
	}
	return func(t *testing.T) {
		if extraTeardownFunc != nil {
			extraTeardownFunc(t, testName)
		}
	}
}
