package storage

var schema = `
    CREATE TABLE IF NOT EXISTS user (
        id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
        name VARCHAR NOT NULL,
        password VARCHAR NOT NULL,
        admin INTEGER NOT NULL DEFAULT 0,
        insert_date INTEGER NOT NULL,
        deleted_date INTEGER DEFAULT NULL
    );

    CREATE TABLE IF NOT EXISTS file (
        id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
        path VARCHAR NOT NULL,
        user_id INTEGER NOT NULL,
        insert_date INTEGER NOT NULL,
        deleted_date INTEGER DEFAULT NULL
    );
`
