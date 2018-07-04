package meme

const (
	path = "memes"
)

//Meme represents "meme" in the system
type Meme struct {
	Path string `json:"path"`
	Url  string `json:"url"`
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
	return []Meme{Meme{Path: "lol"}}, nil
}
