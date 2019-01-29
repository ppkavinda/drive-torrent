package engine

import (
	"database/sql"
	"encoding/hex"
	"fmt"
	"log"
	"strconv"
	"sync"
	"time"

	"github.com/anacrolix/dht"
	"github.com/anacrolix/torrent"
	"github.com/anacrolix/torrent/metainfo"
)

// Engine : Drive torrent Engine
type Engine struct {
	mut      sync.Mutex
	cacheDir string
	client   *torrent.Client
	config   Config
	ts       map[string]*Torrent
}

// New : initiate a new Torrent Engine
func New() *Engine {
	return &Engine{ts: map[string]*Torrent{}}
}

// GetConfig : return the Engin's configs
func (e *Engine) GetConfig() Config {
	return e.config
}

// Config : set Engin's config
func (e *Engine) Config(c Config) error {
	if e.client != nil {
		e.client.Close()
		time.Sleep(1 * time.Second)
	}
	if c.IncomingPort <= 0 {
		return fmt.Errorf("Invalid incoming port (%d)", c.IncomingPort)
	}
	torrentConfig := torrent.Config{
		DHTConfig: dht.ServerConfig{
			StartingNodes: dht.GlobalBootstrapAddrs,
		},
		DataDir:    c.DownloadDirectory,
		ListenAddr: "0.0.0.0:" + strconv.Itoa(c.IncomingPort),
		NoUpload:   !c.EnableUpload,
		Seed:       c.EnableSeeding,
	}
	torrentConfig.DisableEncryption = c.DisableEncryption

	client, err := torrent.NewClient(&torrentConfig)
	if err != nil {
		return err
	}
	e.mut.Lock()
	e.config = c
	e.client = client
	e.mut.Unlock()
	e.GetTorrents()
	return nil
}

// NewMagnet : add a magnet uri to download
func (e *Engine) NewMagnet(magnetURI, email string) error {
	torrent, err := e.client.AddMagnet(magnetURI)
	if err != nil {
		fmt.Printf("DONE2 %v\n", err)

		return err
	}

	saveInDb(torrent, email)

	return e.newTorrent(torrent)
}

func (e *Engine) newTorrent(torrent *torrent.Torrent) error {
	t := e.saveTorrent(torrent)
	fmt.Printf("DONE4 %v\n", e.client.Torrents())

	go func() {
		// wait for engine to collect information
		st := <-t.t.GotInfo()

		fmt.Printf("DONE4 %v\n", st)

		e.StartTorrent(t.InfoHash)
	}()

	return nil
}

// StartTorrent : start a relevent torrent according to the infoHash
func (e *Engine) StartTorrent(infoHash string) error {
	fmt.Printf("DONE3 %v\n", infoHash)

	t, err := e.getOpenTorrent(infoHash)

	fmt.Println("HASH", infoHash)

	if err != nil {
		fmt.Printf("DONE3 %v\n", err)

		return err
	}

	if t.Started {
		fmt.Println("Already started")
		return fmt.Errorf("Already started")
	}

	t.Started = true

	for _, f := range t.Files {
		if f != nil {
			f.Started = true
		}
	}

	if t.t.Info() != nil {
		t.t.DownloadAll()
	}

	return nil
}

// getOpenTorrent : returns the relevant torrent according to the infoHash
// NO DIFFERENCE WITH getTorrent
func (e *Engine) getOpenTorrent(infoHash string) (*Torrent, error) {
	t, err := e.getTorrent(infoHash)
	if err != nil {
		fmt.Printf("getOpenTorrent %v\n", err)

		return nil, err
	}

	return t, nil
}

// getTorrent : returns the relevant torrent according to the infoHash
func (e *Engine) getTorrent(infoHash string) (*Torrent, error) {
	hi, err := str2hi(infoHash)
	if err != nil {
		fmt.Printf("getTorrent %v\n", err)

		return nil, err
	}
	t, ok := e.ts[hi.HexString()]
	if !ok {
		return t, fmt.Errorf("Missing torrent %x", hi)
	}

	return t, nil
}

// str2hi : convert infoHash in string to metainfo.Hash type
func str2hi(infoHash string) (metainfo.Hash, error) {
	var hi metainfo.Hash

	e, err := hex.Decode(hi[:], []byte(infoHash))
	if err != nil {
		fmt.Printf("str2hi %v\n", err)

		return hi, fmt.Errorf("Invalid hex string")
	}
	if e != 20 {
		return hi, fmt.Errorf("Invalid length")
	}

	return hi, nil
}

// GetTorrents : store torrents on Engine.ts
func (e *Engine) GetTorrents() map[string]*Torrent {
	e.mut.Lock()
	defer e.mut.Unlock()

	if e.client == nil {
		return nil
	}

	for _, torrent := range e.client.Torrents() {
		e.saveTorrent(torrent)
	}
	return e.ts
}

// insert or update a particular torrent in engine.ts
func (e *Engine) saveTorrent(newTorrent *torrent.Torrent) *Torrent {
	// newTorrent.Drop()
	ih := newTorrent.InfoHash().HexString()
	oldTorrent, ok := e.ts[ih]
	if !ok {
		oldTorrent = &Torrent{InfoHash: ih}
		e.ts[ih] = oldTorrent
	}
	oldTorrent.Update(newTorrent)
	return oldTorrent
}

// GetFiles : returns the relatent files of a torrent hash
func (e *Engine) GetFiles(hash string) []*File {
	for i, v := range e.ts[hash].Files {
		fmt.Printf("%d: %+v\n", i, *v)
	}
	return e.ts[hash].Files
}

func saveInDb(torrent *torrent.Torrent, email string) {
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
