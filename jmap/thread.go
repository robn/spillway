package jmap

import "log"

func DoGetThreads (msg JMAPMessage) {
    log.Print("getThreads")
}

func DoGetThreadUpdates (msg JMAPMessage) {
    log.Print("getThreadUpdates")
}
