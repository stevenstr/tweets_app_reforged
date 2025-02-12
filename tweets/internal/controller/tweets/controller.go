package tweets

import (
	"context"
	"errors"

	"github.com/stevenstr/tweets_app_reforged/tweets/pkg/model"
)

var ErrNotFound = errors.New("not found")

type tweetsRepository interface {
	Get(ctx context.Context, id string) (*model.Tweet, error)
	Post(ctx context.Context, id string, message string) error
	Put(ctx context.Context, id string, message string) error
	Delete(ctx context.Context, id string) error
	GetAll(ctx context.Context) (map[string]*model.Tweet, error)
}

type Controller struct {
	repo tweetsRepository
}

func New(repo tweetsRepository) *Controller {
	return &Controller{repo: repo}
}

func (c *Controller) Get(ctx context.Context, id string) (*model.Tweet, error) {
	res, err := c.repo.Get(ctx, id)
	if err != nil && errors.Is(err, ErrNotFound) {
		return nil, ErrNotFound
	}
	return res, nil
}

func (c *Controller) GetAll(ctx context.Context) (map[string]*model.Tweet, error) {
	res, err := c.repo.GetAll(ctx)
	if err != nil && errors.Is(err, ErrNotFound) {
		return nil, ErrNotFound
	}
	return res, nil
}

func (c *Controller) Put(ctx context.Context, id string, message string) error {
	return c.repo.Put(ctx, id, message)
}

func (c *Controller) Post(ctx context.Context, id string, message string) error {
	return c.repo.Post(ctx, id, message)
}

func (c *Controller) Delete(ctx context.Context, id string) error {
	return c.repo.Delete(ctx, id)
}
