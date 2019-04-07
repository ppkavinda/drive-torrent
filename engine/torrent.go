package engine

import (
	"fmt"
	"time"

	"github.com/anacrolix/torrent"
	// "github.com/ppkavinda/drive-torrent/engine"
)

// Torrent : store info about torrent
type Torrent struct {
	InfoHash   string
	Name       string
	Loaded     bool
	Downloaded int64
	Size       int64
	Files      []*File

	Started      bool
	Percent      float32
	DownloadRate float32
	t            *torrent.Torrent
	updatedAt    time.Time

	UserEmails   []string
	Finished     bool
	Uploaded     int64
	UploadRate   string
	FinishUpload bool
}

// File inside a torrent
type File struct {
	Path      string
	Size      int64
	Chunks    int
	Completed int

	Started bool
	Percent float32
	f       torrent.File
}

// Update info of a torrent
func (torrent *Torrent) Update(t *torrent.Torrent) {
	torrent.Name = t.Name()
	torrent.Loaded = t.Info() != nil
	if torrent.Loaded {
		torrent.updateLoaded(t)
	}
	torrent.t = t
}

func (torrent *Torrent) updateLoaded(t *torrent.Torrent) {
	torrent.Size = t.Length()
	totalChunks := 0
	totalCompleted := 0

	tfiles := t.Files()
	if len(tfiles) > 0 && torrent.Files == nil {
		torrent.Files = make([]*File, len(tfiles))
	}

	for i, f := range tfiles {
		path := f.Path()
		file := torrent.Files[i]
		if file == nil {
			file = &File{Path: path}
			torrent.Files[i] = file
		}
		chunks := f.State()

		file.Size = f.Length()
		file.Chunks = len(chunks)
		completed := 0
		for _, p := range chunks {
			if p.Complete {
				completed++
			}
		}
		file.Completed = completed
		file.Percent = percent(int64(file.Completed), int64(file.Chunks))
		file.f = f

		totalChunks += file.Chunks
		totalCompleted += file.Completed
	}

	now := time.Now()
	bytes := t.BytesCompleted()
	torrent.Percent = percent(bytes, torrent.Size)
	if !torrent.updatedAt.IsZero() {
		dt := float32(now.Sub(torrent.updatedAt))
		db := float32(bytes - torrent.Downloaded)
		rate := db * (float32(time.Second) / dt)
		if rate >= 0 {
			torrent.DownloadRate = rate
		}
	}
	torrent.Downloaded = bytes
	torrent.updatedAt = now

	// torrent if finished download
	if bytes == torrent.Size {
		torrent.Finished = true
	}

	// fmt.Printf("NOT FINISHED\n")
}

// Stop : stop torrent from being download
func (e *Engine) Stop(infohash string) error {
	t, err := e.GetTorrent(infohash)
	if err != nil {
		fmt.Printf("Stop Torrent: %+v\n", err)
		return err
	}
	if !t.Started {
		fmt.Printf("Already stoped\n")
		return err
	}
	t.t.Drop()
	t.Started = false
	for _, f := range t.Files {
		if f != nil {
			f.Started = false
		}
	}
	return nil
}

// Delete : remove torrent from engine
func (e *Engine) Delete(infohash string) error {
	t, err := e.GetTorrent(infohash)
	if err != nil {
		fmt.Printf("DeleteTorrent: %+v\n", err)
	}

	delete(e.ts, t.InfoHash)
	ih, _ := str2hi(infohash)
	if tt, ok := e.client.Torrent(ih); ok {
		tt.Drop()
	}
	return nil
}

func percent(n, total int64) float32 {
	if total == 0 {
		return float32(0)
	}
	return float32(int(float64(10000)*(float64(n)/float64(total)))) / 100
}
