package utilities

import (
    "log"
)

func Log(s string, v ...interface{}) {
    log.Printf(s, v...)
}
