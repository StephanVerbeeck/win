package win

import (
	"log"
	"runtime"
)

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
	if checkMessageThreadID == 0 {
		runtime.LockOSThread()
		checkMessageThreadID = GetCurrentThreadId()
	} else {
		Assert("win.CheckMessageThread", checkMessageThreadID == GetCurrentThreadId())
	}
}

// IsMessageThread return true when called by the thread that has runs the windows OS main-message-loop
func IsMessageThread() bool {
	return checkMessageThreadID == GetCurrentThreadId()
}
