package debugging

import (
	"fmt"
	"os"
)

func varDump(myVar ...interface{}) {
	fmt.Printf("%v\n", myVar)
}

// dd will print out variables given to it (like varDump()) but
// will also stop execution from continuing.
func DieDump(myVar ...interface{}) {
	varDump(myVar...)
	os.Exit(1)
}