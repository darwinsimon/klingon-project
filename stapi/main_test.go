package stapi_test

import (
	"os"
	"testing"

	. "github.com/darwinsimon/klingon-project/stapi"
)

var s = Stapi{}

func TestMain(m *testing.M) {
	TestingPatch()
	os.Exit(m.Run())
}
