package db

type Thread struct {
    ThreadID string // XXX primary key
    MessageIDs []string
    LastChangeModSeq int64
    SubjectHash string
}
