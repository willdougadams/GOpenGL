package Debug

import (
	"os"
	"fmt"
	"bytes"
	"io/ioutil"
)

var buffer bytes.Buffer

func Print(msg string) {
	if contains(os.Args, "-v") {
		fmt.Printf("%s\n", msg)
	}

	buffer.WriteString(msg)

	if contains(os.Args, "-f") {
		err := ioutil.WriteFile("/logs/log.log", buffer.Bytes(), 0644)
		if err != nil {
			panic(err)
		}
	}
}

func Close() {
	err := ioutil.WriteFile("/logs/log.log", buffer.Bytes(), 0644)
	if err != nil {
		panic(err)
	}
}

func contains(s []string, e string) bool {
    for _, a := range s {
        if a == e {
            return true
        }
    }
    return false
}
