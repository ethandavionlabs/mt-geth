package rcfg

import (
	"os"
)

// UsingBVM is used to enable or disable functionality necessary for the BVM.
var UsingBVM bool

var ()

func init() {
	UsingBVM = os.Getenv("USING_BVM") == "true"
}
