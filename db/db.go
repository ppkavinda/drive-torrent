package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/anacrolix/torrent"
)

// GetHashesOfEmail : return torrents of particular email
func GetHashesOfEmail(email string) []string {
	db, err := sql.Open("sqlite3", "./info.db")
	if err != nil {
		fmt.Printf("SQL: %v", err)
	}
	defer db.Close()

	stmt, err := db.Prepare("select hash from torrents where email = ?")
	if err != nil {
		fmt.Printf("GetHashesOfEmail1 %v\n", err)
		log.Fatal(err)
	}
	defer stmt.Close()

	torrents := make([]string, 1)
	rows, err := stmt.Query(email)
	if err != nil {
		fmt.Printf("GetTorrentOfEmail2 %+v\n", err)
		return nil
	}
	defer rows.Close()

	for rows.Next() {
		var hash string
		err = rows.Scan(&hash)
		if err != nil {
			log.Printf("GetEmailOfTorrent3 %+v\n", err)
			return nil
		}
		torrents = append(torrents, hash)
	}

	if err != nil {
		log.Fatal(err)
	}

	return torrents
}

// GetEmailOfTorrent : get email of a torrent
func GetEmailOfTorrent(infohash string) []string {
	db, err := sql.Open("sqlite3", "./info.db")
	if err != nil {
		fmt.Printf("SQL: %v", err)
	}
	defer db.Close()

	stmt, err := db.Prepare("select id, email from torrents where hash = ?")
	if err != nil {
		log.Printf("GetEmiailOfTorrent:1 %+v\n", err)
	}
	defer stmt.Close()

	emails := make([]string, 0)
	ids := make([]string, 0)
	rows, err := stmt.Query(infohash)
	if err != nil {
		fmt.Printf("GetEmailOfTorrent2 %+v\n", err)
		return nil
	}
	defer rows.Close()

	for rows.Next() {
		var id, email string
		err = rows.Scan(&id, &email)
		if err != nil {
			log.Printf("GetEmailOfTorrent:3 select %+v\n", err)
			return nil
		}
		emails = append(emails, email)
		ids = append(ids, id)
	}
	// delete record from the sqlite DB
	// TODO : instead of delete set a flag to indicate uploaded, that way we can
	// keep the histry
	// stmt2, err := db.Prepare("delete from torrents where id = ?")
	// if err != nil {
	// 	log.Printf("GetEmailOfTorrent:4 delete %+v\n", err)
	// 	return nil
	// }
	// defer stmt2.Close()

	// for _, id := range ids {
	// 	_, err = stmt2.Exec(id)
	// 	if err != nil {
	// 		fmt.Printf("GetTorrentOfEmail5 %+v\n", err)

	// 	}
	// }

	// fmt.Printf("FINISHED : %s\n", id)

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

// DeleteTorrent : delete torrent from sqlite database
func DeleteTorrent(infohash, email string) error {
	db, err := sql.Open("sqlite3", "./info.db")
	if err != nil {
		fmt.Printf("SQL: %v", err)
	}
	defer db.Close()
	stmt2, err := db.Prepare("delete from torrents where hash = ? and email = ?")
	if err != nil {
		log.Printf("GetEmailOfTorrent:4 delete %+v\n", err)
		return nil
	}
	defer stmt2.Close()
	_, err = stmt2.Exec(infohash, email)
	if err != nil {
		fmt.Printf("GetTorrentOfEmail5 %+v\n", err)
	}
	return nil
}

// GetAllTorrentHashes : return all torrents in sqlite db
func GetAllTorrentHashes() []string {
	db, err := sql.Open("sqlite3", "./info.db")
	if err != nil {
		fmt.Printf("SQL: %v", err)
	}
	defer db.Close()

	stmt, err := db.Prepare("select hash from torrents")
	if err != nil {
		log.Printf("GetAllTorrentHashes:1 %+v\n", err)
	}
	defer stmt.Close()

	hashes := make([]string, 0)
	rows, err := stmt.Query()
	if err != nil {
		fmt.Printf("GetAllTorrnetHashes %+v\n", err)
		return nil
	}
	defer rows.Close()

	for rows.Next() {
		var hash string
		err = rows.Scan(&hash)
		if err != nil {
			log.Printf("GetAllTorrentHashes:3 %+v\n", err)
			return nil
		}
		hashes = append(hashes, hash)
	}
	return hashes
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
			id integer primary key,
			name text,
			hash text,
			email text,
			uploaded integer
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
