package db

type Mailbox struct {
    ID int64
    Name string
    ParentID int64
    Role *string
    MustBeOnlyMailbox bool
    MayAddMessages bool
    MayRemoveMessages bool
    MayCreateChild bool
    MayRenameMailbox bool
    MayDeleteMailbox bool
    TotalMessages uint64
    TotalThreads uint64
    UnreadThreads uint64

    ModSeq int64
    NonCountModSeq int64
    MessageListUIDNext uint32
    MessageListHighestModSeq int64
    MessageListLowWatermarkModSeq int64
}

func NewMailbox(name string) Mailbox {
    return User{-1, name, 0, 0, 0, 0, 0, 0, []string{}}
}
