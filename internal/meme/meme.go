package meme

import (
	"io"
	"os"
	"path/filepath"
)

const (
	path = "memes"
)

//File represents meme file
type File interface {
	io.Reader
	io.ReaderAt
	io.Seeker
	io.Closer
}

//New returns pointer to Meme with associated f and generated path, url
func New(f File) *Meme {
	//@todo depends on file and meme paths
	p := "test1"
	u := "url1"

	return &Meme{p, u, f}
}

//Meme represents "meme" in the system
type Meme struct {
	Path string `json:"path"`
	Url  string `json:"url"`
	File File   `json:"-"`
}

//Repository brings api for persisting memes in the system
type Repository struct {
	persistor persistor
}

//Save just wraps persistor.Save
func (m *Repository) Save(meme Meme) error {
	return m.persistor.Save(meme)
}

//List just wraps persistor.List
func (m *Repository) List() ([]Meme, error) {
	return m.persistor.List()
}

//NewRepository creates default *MemeRepository
func NewRepository() *Repository {
	return &Repository{persistor: &fileSystem{}}
}

//Persistor defines contract for data manipulation
type persistor interface {
	Save(meme Meme) error
	List() ([]Meme, error)
}

type fileSystem struct {
	root string
}

func (f *fileSystem) Save(meme Meme) error {
	return nil
}

func (f *fileSystem) List() ([]Meme, error) {
	if _, err := os.Open(path); os.IsNotExist(err) {
		return nil, err
	}

	var memes []Meme
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if info.Name() == path {
			return nil
		}
		f, err := os.Open(path)
		if err != nil {
			return err
		}

		m := Meme{Path: path, File: f}
		memes = append(memes, m)
		return nil
	})

	if err != nil {
		return nil, err
	}

	return memes, nil
}
