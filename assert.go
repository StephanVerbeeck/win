package win

import "log"

var (
	// UseAssert verbose debug mode
	UseAssert            = false
	checkMessageThreadID uint32 // windows message loop MUST always be processed by the SAME thread
)

// Assert mandatory condition
func Assert(test string, condition bool) {
	if UseAssert && !condition {
		log.Panicln("assert failed: " + test)
	}
}

// CheckMessageThread windows message loop MUST always be processed by the SAME thread
func CheckMessageThread() {
	threadID := GetCurrentThreadId()
	if checkMessageThreadID == 0 {
		checkMessageThreadID = threadID
	} else {
		Assert("win.CheckMessageThread", threadID == checkMessageThreadID)
	}
}
