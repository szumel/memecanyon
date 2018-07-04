package meme

import (
	"encoding/json"
	"net/http"

	"github.com/szumel/memecanyon/internal/meme"
)

//ListCollection returns all memes
//Method GET
func ListCollection(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	re := meme.NewRepository()
	l, err := re.List()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	j, _ := json.Marshal(l)

	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

//NewDocument creates new entity in the system
//Method POST
func NewDocument(w http.ResponseWriter, r *http.Request) {

}
