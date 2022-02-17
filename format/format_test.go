package format

import (
	"testing"
)

func TestHelloName(t *testing.T) {
    Testing()
}
func Testing(){
	ENABLE_VERBOSE=1
	Info("Test", "This is Info")
	Error("Test", "This is Error")
	Warning("Test", "This is Warning")
	Verbose("Test", "This is Verbose")
	Special("Test", "This is Special")
}