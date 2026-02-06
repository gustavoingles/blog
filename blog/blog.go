package blog

import (
	"context"
	"time"
)

type Post struct {
	Title        string
	Introduction string
	Body         string
	CallToAction string
	Author       string
	CreatedAt    time.Time
}

func defaultExamplePost() *Post {
	return &Post{
		Title:        "example-title",
		Introduction: "Lorem Ipsum",
		Body:         "Arigathanks Gozaimuch",
		CallToAction: "Get the fuck out of my face",
		Author:       "Gustavin Gostosin",
		CreatedAt:    time.Now(),
	}
}

type CustomPost func(*Post)

func WithTitle(title string) CustomPost {
	return func(p *Post) {
		p.Title = title
	}
}

func WithIntroduction(introduction string) CustomPost {
	return func(p *Post) {
		p.Introduction = introduction
	}
}

func WithBody(body string) CustomPost {
	return func(p *Post) {
		p.Body = body
	}
}

func WithCallToAction(callToAction string) CustomPost {
	return func(p *Post) {
		p.CallToAction = callToAction
	}
}

func WithAuthor(author string) CustomPost {
	return func(p *Post) {
		p.Author = author
	}
}

func NewPost(cp ...CustomPost) Post {
	p := defaultExamplePost()

	for _, c := range cp {
		c(p)
	}

	return *p
}

type PostRepository interface {
	PublishPost(ctx context.Context, p Post) error
	GetAllPosts(ctx context.Context) ([]Post, error)
	GetPostById(ctx context.Context, pID int64) (*Post, error)
	UpdatePostById(ctx context.Context, pID int64, updateFn CustomPost) error
	DeletePostById(ctx context.Context, pID int64) error
}
