package util

import "log"

func NoError(err error) {
    if err != nil {
        log.Fatal(err)
    }
}
