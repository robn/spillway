package jmap

import "log"

func DoGetMessages (msg JMAPMessage) {
    log.Print("getMessages")
}

func DoGetMessageUpdates (msg JMAPMessage) {
    log.Print("getMessageUpdates")
}

func DoSetMessages (msg JMAPMessage) {
    log.Print("setMessages")
}

func DoImportMessage (msg JMAPMessage) {
    log.Print("importMessage")
}

func DoCopyMessages (msg JMAPMessage) {
    log.Print("copyMessages")
}

func DoReportMessages (msg JMAPMessage) {
    log.Print("reportMessages")
}
