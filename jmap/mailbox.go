package jmap

import "log"

func DoGetMailboxes (msg JMAPMessage) {
    log.Print("getMailboxes")
}

func DoGetMailboxUpdates (msg JMAPMessage) {
    log.Print("getMailboxUpdates")
}

func DoSetMailboxes (msg JMAPMessage) {
    log.Print("setMailboxes")
}
