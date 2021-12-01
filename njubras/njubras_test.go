package njubras

import (
	"testing"
	"time"
)

func TestAll(t *testing.T) {
	// 1. logout
	_, err := DoLogout()
	if err != nil {
		panic(err)
	}
	time.Sleep(1000 * time.Millisecond)
	// 2. login
	_, err = DoLogin("", "")
	if err != nil {
		panic(err)
	}
	time.Sleep(2 * time.Second)
	// 3. acquire status
	err = AcquirePortalStatus()
	if err != nil {
		panic(err)
	}
}
