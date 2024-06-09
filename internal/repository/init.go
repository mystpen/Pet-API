package repository

import (
	"context"
	"database/sql"
	"time"
)

func Init(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id uuid PRIMARY KEY,
		username text NOT NULL,
		email text UNIQUE NOT NULL,
		password_hash bytea NOT NULL
	 );
	 
	 CREATE TABLE IF NOT EXISTS documents (
		id uuid PRIMARY KEY,
		title text NOT NULL,
		text text NOT NULL,
		created_at timestamp(0) NOT NULL DEFAULT (NOW() at time zone 'UTC'),
		user_id uuid,
		FOREIGN KEY (user_id) REFERENCES users (id)
	 );`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	_, err := db.ExecContext(ctx, query)
	return err
}
