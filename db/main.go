package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func createTables(db *sql.DB) error {
	var err error
	_, err = db.Exec(`
		CREATE TABLE userdata (
			/* The unique snowflake that Discord assigns to a user */
            user_id INTEGER NOT NULL PRIMARY KEY,
            avatar TEXT,
			/* The unique 4 digit number identifying a user, this will only show on users that have not migrated to the new naming system. */
            discriminator TEXT,
			/* The number of messages the user has sent in the guild. */
            message_count INTEGER, 
            monetize_xp_boost INTEGER,
            username TEXT, 
            xp INTEGER,  
            level INTEGER
        );
        CREATE TABLE userxp (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			user_id INTEGER,
			xp_amount INTEGER NOT NULL,
			FOREIGN KEY (user_id) REFERENCES userdata(user_id)
		)
	`)
	return err
}

func PrepareUserDataStatement(tx *sql.Tx) (*sql.Stmt, error) {
	stmt, err := tx.Prepare("INSERT INTO userdata (user_id, avatar, discriminator, message_count, monetize_xp_boost, username, xp, level) VALUES (?, ?, ?, ?, ?, ?, ?, ?)")
	return stmt, err
}

func PrepareUserXPStatement(tx *sql.Tx) (*sql.Stmt, error) {
	stmt, err := tx.Prepare("INSERT INTO userxp (user_id, xp_amount) VALUES (?, ?)")
	return stmt, err
}

func PrepareDB() (db *sql.DB, tx *sql.Tx) {
	filepath := "export.db"
	if _, err := os.Stat(filepath); err == nil {
		if err := os.Remove("export.db"); err != nil {
			log.Fatal(err)
		}
	}
	db, err := sql.Open("sqlite3", "export.db")
	if err != nil {
		log.Fatal(err)
	}
	if err := createTables(db); err != nil {
		log.Fatal(err)
	}
	tx, err = db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	return db, tx

}

func CommitTransaction(tx *sql.Tx) {
	err := tx.Commit()
	if err != nil {
		log.Fatal(err)
	}
}
