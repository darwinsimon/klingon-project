package stapi

import "github.com/darwinsimon/klingon-project/resource"

var (
	UpdateCharacterFile = updateCharacterFile
)

// Run testing patch to inject private value
func TestingPatch() {

	// Add dummy saved character
	savedChar["darwin"] = &resource.Character{
		UID:     "CHMA0000000000",
		Name:    "Darwin",
		Species: "Human",
	}

}
