package stapi

// Run testing patch to inject private value
func TestingPatch() {

	// Add dummy saved character
	savedChar["darwin"] = &Character{
		UID:     "CHMA0000000000",
		Name:    "Darwin",
		Species: "Human",
	}

}
