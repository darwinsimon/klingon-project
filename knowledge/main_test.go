package knowledge_test

import (
	"os"
	"testing"
)

var dummyRes = DummyResource{}

func TestMain(m *testing.M) {

	os.Exit(m.Run())
}
