package win

import "log"

var (
	// UseAssert verbose debug mode
	UseAssert = false
)

// Assert mandatory condition
func Assert(test string, condition bool) {
	if UseAssert && !condition {
		log.Panicln("assert failed: " + test)
	}
}
