package tweets

import (
	"context"
	"errors"
	"fmt"

	"tweets.com/tweets/pkg/model"
)

var ErrNotFound = errors.New("not found")

type tweetsRepository interface {
	Get(ctx context.Context, id string) (*model.Tweet, error)
	Put(ctx context.Context, id string, tweet *model.Tweet) error
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

func (c *Controller) Put(ctx context.Context, id string, tweet *model.Tweet) error {
	err := c.repo.Put(ctx, id, tweet)
	if err != nil {
		return fmt.Errorf("some goes wrong: %v", err)
	}
	return nil
}
