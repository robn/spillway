package jmap

import "time"


// JMAP data model, from the spec

type Account struct {
    ID string
    Name string
    IsPrimary bool
    IsReadOnly bool
    HasMail bool
    MailCapabilities *MailCapabilities
    HasContacts bool
    HasCalendars bool
}

type MailCapabilities struct {
    // TODO
}

type Mailbox struct {
    ID string
    Name string
    ParentID *string
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
}

type FilterOperator struct {
    /* TODO
    Operator string // XXX AND|OR|NOT, probably better as an enumeration
    conditions []interface // XXX Operator|Condition, so make condition a single ONE "operator"?
    */
}

type FilterCondition struct {
    InMailboxes *[]string
    NotInMailboxes *[]string
    Before *time.Time
    After *time.Time
    IsFlagged *bool
    IsUnread *bool
    IsAnswered *bool
    IsDraft *bool
    HasAttachment *bool
    Text *string
    From *string
    To *string
    Cc *string
    Bcc *string
    Subject *string
    Body *string
    Header *[]string
}

type RemovedItem struct {
    MessageID string
    ThreadID string
}

type AddedItem struct {
    MessageID string
    ThreadID string
    Index uint64
}

type Thread struct {
    ID string
    MessageIDs []string
}

type Message struct {
    ID string
    ThreadID string
    MailboxIDs []string
    InReplyToMessageID *string
    IsUnread bool
    IsFlagged bool
    IsAnswered bool
    IsDraft bool
    HasAttachment bool
    RawURL string
    Headers map[string]string
    From *Emailer
    To *[]Emailer
    Cc *[]Emailer
    Bcc *[]Emailer
    ReplyTo *[]Emailer
    Subject *string
    Date *time.Time
    Size uint64
    Preview *string
    TextBody *string
    HTMLBody *string
    Attachments *[]Attachment
    AttachedMessages *map[string]Message
}

type Emailer struct {
    Name string
    Email string
}

type Attachment struct {
    ID string
    URL string
    Type string
    Name string
    Size uint64
    IsInline bool
    Width uint64
    Height uint64
}

type MessageCopy struct {
    // TODO
}

type SearchSnippet struct {
    // TODO
}

type SetError struct {
    Type string
    Description *string
}


// Wire protocol support

type JMAPMessage struct {
    Name string
    Args map[string]interface{}
    ClientID string
}
