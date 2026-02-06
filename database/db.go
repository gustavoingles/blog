package database

import (
	"context"
	"errors"
	"math/rand/v2"

	"blogging/blog"
)

type InMemoryDatabase map[int64]blog.Post

var (
	ErrPostNotExists error = errors.New("this blog post doesn't exist")
	ErrPostAlreadyExists error = errors.New("post already exists")
)

func NewInMemoryDatabase() InMemoryDatabase {
	return map[int64]blog.Post{}
}

func (db InMemoryDatabase) GetAllPosts(ctx context.Context) ([]blog.Post, error) {
	posts := make([]blog.Post, 0, len(db))
	for _, p := range db {
		posts = append(posts, p)
	}
	return posts, nil
}

func (db InMemoryDatabase) GetPostById(ctx context.Context, pID int64) (*blog.Post, error) {
	post, ok := db[pID]
	if !ok {
		return nil, ErrPostNotExists
	}
	return &post, nil
}

func (db InMemoryDatabase) DeletePostById(ctx context.Context, pID int64) error {
	delete(db, pID)
	return nil
}

func (db InMemoryDatabase) UpdatePostById(ctx context.Context, pID int64, updateFn blog.CustomPost) error {
	postToUpdate := db[pID]
	updateFn(&postToUpdate)
	db[pID] = postToUpdate
	return nil
}

func (db InMemoryDatabase) PublishPost(ctx context.Context, p blog.Post) error {
	seenIndexes := make(map[int64]bool, len(db))
	for i := range db {
		seenIndexes[i] = true
	}
	if newIndex := rand.Int64(); !seenIndexes[newIndex] {
		db[newIndex] = p
	} else {
		return ErrPostAlreadyExists
	}
	return nil
}