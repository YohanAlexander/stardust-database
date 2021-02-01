package generosfavoritos

import (
	"strconv"
	"strings"

	"github.com/yohanalexander/deezefy-music/entity"
	der "github.com/yohanalexander/deezefy-music/entity/generosfavoritos"
)

// inmem in memory repo
type inmem struct {
	m map[string]*der.GenerosFavoritos
}

// newInmem create new repository
func newInmem() *inmem {
	var m = map[string]*der.GenerosFavoritos{}
	return &inmem{
		m: m,
	}
}

// Create GenerosFavoritos
func (r *inmem) Create(e *der.GenerosFavoritos) (int, string, string, error) {
	r.m[strconv.Itoa(e.Perfil)+e.Genero] = e
	return e.Perfil, e.Genero, e.Ouvinte, nil
}

// Get GenerosFavoritos
func (r *inmem) Get(perfil int, genero, artista string) (*der.GenerosFavoritos, error) {
	if r.m[strconv.Itoa(perfil)+genero] == nil {
		return nil, entity.ErrNotFound
	}
	return r.m[strconv.Itoa(perfil)+genero], nil
}

// GetByPerfil GenerosFavoritos
func (r *inmem) GetByPerfil(perfil int) (*der.GenerosFavoritos, error) {
	if r.m[strconv.Itoa(perfil)] == nil {
		return nil, entity.ErrNotFound
	}
	return r.m[strconv.Itoa(perfil)], nil
}

// GetByGenero GenerosFavoritos
func (r *inmem) GetByGenero(genero string) (*der.GenerosFavoritos, error) {
	if r.m[genero] == nil {
		return nil, entity.ErrNotFound
	}
	return r.m[genero], nil
}

// Update GenerosFavoritos
func (r *inmem) Update(e *der.GenerosFavoritos) error {
	_, err := r.Get(e.Perfil, e.Genero, e.Ouvinte)
	if err != nil {
		return err
	}
	r.m[strconv.Itoa(e.Perfil)+e.Genero] = e
	return nil
}

// Search GenerosFavoritoss
func (r *inmem) Search(query string) ([]*der.GenerosFavoritos, error) {
	var d []*der.GenerosFavoritos
	for _, j := range r.m {
		if strings.Contains(strings.ToLower(j.Genero), query) {
			d = append(d, j)
		}
	}
	if len(d) == 0 {
		return nil, entity.ErrNotFound
	}

	return d, nil
}

// List GenerosFavoritoss
func (r *inmem) List() ([]*der.GenerosFavoritos, error) {
	var d []*der.GenerosFavoritos
	for _, j := range r.m {
		d = append(d, j)
	}
	return d, nil
}

// Delete GenerosFavoritos
func (r *inmem) Delete(Perfil int, genero, artista string) error {
	if r.m[strconv.Itoa(Perfil)+genero] == nil {
		return entity.ErrNotFound
	}
	r.m[strconv.Itoa(Perfil)+genero] = nil
	return nil
}
