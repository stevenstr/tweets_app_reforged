package memory

import (
	"context"
	"sync"

	"github.com/stevenstr/tweets_app_reforged/tweets/internal/repository"
	"github.com/stevenstr/tweets_app_reforged/tweets/pkg/model"
)

type Repository struct {
	mx   sync.RWMutex
	data map[string]*model.Tweet
}

func New() *Repository {
	return &Repository{data: map[string]*model.Tweet{}}
}

func (r *Repository) Get(_ context.Context, id string) (*model.Tweet, error) {
	r.mx.RLock()
	defer r.mx.RUnlock()

	m, ok := r.data[id]
	if !ok {
		return nil, repository.ErrNotFound
	}

	return m, nil
}

func (r *Repository) GetAll(_ context.Context) (map[string]*model.Tweet, error) {
	r.mx.RLock()
	defer r.mx.RUnlock()

	return r.data, nil
}

func (r *Repository) Put(_ context.Context, id string, message string) error {
	r.mx.RLock()
	defer r.mx.RUnlock()

	_, ok := r.data[id]
	if !ok {
		return repository.ErrNotFound
	}

	r.data[id] = &model.Tweet{Message: message}

	return nil
}

func (r *Repository) Post(_ context.Context, id string, message string) error {
	r.mx.RLock()
	defer r.mx.RUnlock()

	_, ok := r.data[id]
	if ok {
		return repository.ErrAlreadyExists
	}

	r.data[id] = &model.Tweet{Message: message}

	return nil
}

func (r *Repository) Delete(_ context.Context, id string) error {
	r.mx.RLock()
	defer r.mx.RUnlock()

	delete(r.data, id)

	return nil
}
