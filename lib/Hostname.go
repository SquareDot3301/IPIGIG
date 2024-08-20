package lib

import (
	"os"
)

func Hostname() string {
	hostname, error := os.Hostname()
    if error != nil {
        panic(error)
    }
    return hostname
}
