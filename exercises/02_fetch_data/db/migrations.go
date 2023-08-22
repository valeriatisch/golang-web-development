package db

import (
	// "database/sql"
)

const createTableQuery = `
CREATE TABLE IF NOT EXISTS activities (
    id SERIAL PRIMARY KEY,
    activity TEXT NOT NULL,
    type TEXT,
    participants INT,
    price FLOAT,
    link TEXT,
    key TEXT UNIQUE,
    accessibility FLOAT
);
`

// 1) TODO: Create a function to ensure the table exists, use the createTableQuery

// 2) TODO: Create a function to insert an activity into the database
