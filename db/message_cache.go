package db

import "time"

type MessageCacheItem struct {
    MessageID string // XXX primary key
    IsUnread bool
    IsFlagged bool
    IsAnswered bool
    IsDraft bool
    Mailboxes []string
    From string
    Subject string
    Date time.Time
    Preview string
    Attachments []string
    GUID string
}
type MessageCache map[string]MessageCacheItem
