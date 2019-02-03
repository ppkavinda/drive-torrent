package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/anacrolix/torrent"
)

// GetTorrentsOfEmail : return torrents of particular email
func GetHashesOfEmail(email string) []string {
	db, err := sql.Open("sqlite3", "./info.db")
	if err != nil {
		fmt.Printf("SQL: %v", err)
	}
	stmt, err := db.Prepare("select hash from torrents where email = ?")
	if err != nil {
		log.Fatal(err)
	}
	torrents := make([]string, 1)
	rows, err := stmt.Query(email)
	if err != nil {
		fmt.Printf("GetTorrentOfEmail %+v\n", err)
		return nil
	}
	for rows.Next() {
		var hash string
		err = rows.Scan(&hash)
		if err != nil {
			log.Printf("GetEmailOfTorrent %+v\n", err)
			return nil
		}
		torrents = append(torrents, hash)
	}

	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()
	defer db.Close()
	return torrents
}

// GetEmailOfTorrent : get email of a torrent
func GetEmailOfTorrent(infohash string) []string {
	db, err := sql.Open("sqlite3", "./info.db")
	if err != nil {
		fmt.Printf("SQL: %v", err)
	}
	stmt, err := db.Prepare("select id, email from torrents where hash = ?")
	if err != nil {
		log.Fatal(err)
	}

	emails := make([]string, 1)
	rows, err := stmt.Query(infohash)
	if err != nil {
		fmt.Printf("GetEmailOfTorrent %+v\n", err)
		return nil
	}
	for rows.Next() {
		var id, email string
		err = rows.Scan(&id, &email)
		if err != nil {
			log.Printf("GetEmailOfTorrent %+v\n", err)
			return nil
		}
		emails = append(emails, email)

		stmt, err = db.Prepare("delete from torrents where id = ?")
		if err != nil {
			log.Fatal(err)
		}
		_, err = stmt.Exec(id)
		if err != nil {
			log.Fatal(err)
		}
	}
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Printf("FINISHED : %s\n", id)

	defer stmt.Close()
	defer db.Close()
	return emails
}

// SaveInDb : save torrent and hash in sqlite
// 'cause engine.ts is not relaiable
func SaveInDb(torrent *torrent.Torrent, email string) {
	db, err := sql.Open("sqlite3", "./info.db")
	if err != nil {
		fmt.Printf("SQL: %v", err)
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := tx.Prepare("insert into torrents(id, name, hash, email) values(?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	_, err = stmt.Exec(nil, torrent.Name(), torrent.InfoHash().HexString(), email)
	if err != nil {
		log.Fatal(err)
	}

	tx.Commit()

}

// SetupDB : create torrents table
func SetupDB() {

	db, err := sql.Open("sqlite3", "./info.db")
	if err != nil {
		fmt.Printf("SQL: %v", err)
	}
	defer db.Close()

	sqlStmt := `
		create table if not exists torrents (
			id integer auto_increment not null primary key,
			name text,
			hash text,
			email text
			);`

	_, err = db.Exec(sqlStmt)
	if err != nil {
		fmt.Printf("%q: %s\n", err, sqlStmt)
		return
	}

	rows, err := db.Query("select id, name from torrents")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		err = rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, name)
	}

}
