package db

import (
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
)

func openDb (dbfile string, create_sql string) (*sql.DB, error) {
    var dbh *sql.DB
    var err error

    if dbh, err = sql.Open("sqlite3", dbfile); err != nil {
        return nil, err
    }

    if _, err = dbh.Exec(create_sql); err != nil {
        return nil, err
    }

    if _, err = dbh.Exec("PRAGMA foreign_keys = ON;"); err != nil {
        return nil, err
    }

    return dbh, nil
}
