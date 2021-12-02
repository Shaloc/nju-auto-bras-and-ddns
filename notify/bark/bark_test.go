package bark

import (
	"log"
	"testing"
)

func TestNotifierBark_DoNotify(t *testing.T) {
	barkNotifier := NotifierBark{}
	barkNotifier.Initialize("") // add your own bark url here...
	log.Println(barkNotifier.GetApiUrl())
	barkNotifier.DoNotify("BarkApiTest Title", "BarkApiTest Content")
}
