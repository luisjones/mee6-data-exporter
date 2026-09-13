package mee6

import (
	"database/sql"
	"log"
	"mee6xport/db"

	_ "github.com/mattn/go-sqlite3"
)

func (r Response) Insert(tx *sql.Tx) {
	user, err := db.PrepareUserDataStatement(tx)
	if err != nil {
		log.Fatal(err)
	}
	defer user.Close()
	xp, err := db.PrepareUserXPStatement(tx)
	if err != nil {
		log.Fatal(err)
	}
	defer xp.Close()

	for _, player := range r.Players {
		_, err := user.Exec(player.ID, player.Avatar, player.Discriminator, player.MessageCount, player.MonetizeXpBoost, player.Username, player.Xp, player.Level)
		if err != nil {
			log.Fatal(err)
		}
		for _, xpDetail := range player.DetailedXp {
			_, err := xp.Exec(player.ID, xpDetail)
			if err != nil {
				log.Fatal(err)
			}
		}

	}
}
