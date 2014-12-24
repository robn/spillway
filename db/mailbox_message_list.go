package db

type MailboxMessageList struct {
    UID uint32
    ModSeq int64
    MessageID string
    ThreadID uint64
    Deleted uint64
    Date uint64
}
