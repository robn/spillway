package db

import (
    "database/sql"
)

type User struct {
    user_id int64

    Name string
    HighestModSeq int64
    HighestMailboxModSeq int64
    ThreadLogLowWatermarkModSeq int64
    MessageLogLowWatermarkModSeq int64
    QuotaAvailable int64
    QuotaUsed int64
    SharedUsers []string
}

func NewUser(name string) User {
    return User{-1, name, 0, 0, 0, 0, 0, 0, []string{}}
}


type UserDB struct {
    db *sql.DB
}

func OpenUserDB () (*UserDB, error) {
    db, err := openDb("users.db", `
        CREATE TABLE IF NOT EXISTS users (
            user_id                          INTEGER PRIMARY KEY AUTOINCREMENT,
            name                             TEXT,
            highest_modseq                   INTEGER NOT NULL,
            highest_mailbox_modseq           INTEGER NOT NULL,
            thread_log_low_watermark_modseq  INTEGER NOT NULL,
            message_log_low_watermark_modseq INTEGER NOT NULL,
            quota_available                  INTEGER NOT NULL,
            quota_used                       INTEGER NOT NULL
        );
        CREATE TABLE IF NOT EXISTS shared_users (
            user_id        INTEGER NOT NULL,
            shared_user_id INTEGER NOT NULL,
            FOREIGN KEY (user_id)        REFERENCES users (user_id) ON DELETE CASCADE,
            FOREIGN KEY (shared_user_id) REFERENCES users (user_id) ON DELETE CASCADE,
            UNIQUE (user_id, shared_user_id)
        );
    `);
    if err != nil {
        return nil, err
    }
    return &UserDB{db}, nil
}

func (db *UserDB) Save(user *User) error {
    var tx *sql.Tx
    var err error
    if tx, err = db.db.Begin(); err != nil {
        return err
    }

    var user_id *int64
    if user.user_id >= 0 { user_id = &user.user_id }

    var r sql.Result
    r, err = tx.Exec(`
        INSERT OR REPLACE INTO users (
            user_id,
            name,
            highest_modseq,
            highest_mailbox_modseq,
            thread_log_low_watermark_modseq,
            message_log_low_watermark_modseq,
            quota_available,
            quota_used
        ) VALUES (
            ?, ?, ?, ?, ?, ?, ?, ?
        )
    `,
        user_id,
        user.Name,
        user.HighestModSeq,
        user.HighestMailboxModSeq,
        user.ThreadLogLowWatermarkModSeq,
        user.MessageLogLowWatermarkModSeq,
        user.QuotaAvailable,
        user.QuotaUsed,
    );
    if err != nil {
        return err
    }

    // XXX shared users

    if err = tx.Commit(); err != nil {
        return err
    }

    user.user_id, _ = r.LastInsertId() // XXX error check?

    return nil
}
