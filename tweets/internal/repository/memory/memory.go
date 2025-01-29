package memory

import (
	"context"
	"sync"

	"tweets.com/tweets/internal/repository"
	"tweets.com/tweets/pkg/model"
)

type Repository struct {
	sync.RWMutex
	data map[string]*model.Tweets
}

func New() *Repository {
	return &Repository{data: map[string]*model.Tweets{}}
}

func (r *Repository) Get(_ context.Context, id string) (*model.Tweets, error) {
	r.RLock()
	defer r.RUnlock()

	m, ok := r.data[id]
	if !ok {
		return nil, repository.ErrNotFound
	}

	return m, nil
}

func (r *Repository) Put(_ context.Context, id string, tweet *model.Tweets) error {
	r.RLock()
	defer r.Unlock()

	r.data[id] = tweet

	return nil
}
