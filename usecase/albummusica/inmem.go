package albummusica

import (
	"strconv"
	"strings"

	"github.com/yohanalexander/deezefy-music/entity"
	der "github.com/yohanalexander/deezefy-music/entity/albummusica"
)

// inmem in memory repo
type inmem struct {
	m map[string]*der.AlbumMusica
}

// newInmem create new repository
func newInmem() *inmem {
	var m = map[string]*der.AlbumMusica{}
	return &inmem{
		m: m,
	}
}

// Create AlbumMusica
func (r *inmem) Create(e *der.AlbumMusica) (int, int, string, error) {
	r.m[strconv.Itoa(e.Album)+strconv.Itoa(e.Musica)] = e
	return e.Album, e.Musica, e.Artista, nil
}

// Get AlbumMusica
func (r *inmem) Get(album, musica int, artista string) (*der.AlbumMusica, error) {
	if r.m[strconv.Itoa(album)+strconv.Itoa(musica)] == nil {
		return nil, entity.ErrNotFound
	}
	return r.m[strconv.Itoa(album)+strconv.Itoa(musica)], nil
}

// GetByAlbum AlbumMusica
func (r *inmem) GetByAlbum(album int) (*der.AlbumMusica, error) {
	if r.m[strconv.Itoa(album)] == nil {
		return nil, entity.ErrNotFound
	}
	return r.m[strconv.Itoa(album)], nil
}

// GetByMusica AlbumMusica
func (r *inmem) GetByMusica(musica int) (*der.AlbumMusica, error) {
	if r.m[strconv.Itoa(musica)] == nil {
		return nil, entity.ErrNotFound
	}
	return r.m[strconv.Itoa(musica)], nil
}

// Update AlbumMusica
func (r *inmem) Update(e *der.AlbumMusica) error {
	_, err := r.Get(e.Album, e.Musica, e.Artista)
	if err != nil {
		return err
	}
	r.m[strconv.Itoa(e.Album)+strconv.Itoa(e.Musica)] = e
	return nil
}

// Search AlbumMusicas
func (r *inmem) Search(query string) ([]*der.AlbumMusica, error) {
	var d []*der.AlbumMusica
	for _, j := range r.m {
		if strings.Contains(strings.ToLower(j.Artista), query) {
			d = append(d, j)
		}
	}
	if len(d) == 0 {
		return nil, entity.ErrNotFound
	}

	return d, nil
}

// List AlbumMusicas
func (r *inmem) List() ([]*der.AlbumMusica, error) {
	var d []*der.AlbumMusica
	for _, j := range r.m {
		d = append(d, j)
	}
	return d, nil
}

// Delete AlbumMusica
func (r *inmem) Delete(album, musica int, artista string) error {
	if r.m[strconv.Itoa(album)+strconv.Itoa(musica)] == nil {
		return entity.ErrNotFound
	}
	r.m[strconv.Itoa(album)+strconv.Itoa(musica)] = nil
	return nil
}
