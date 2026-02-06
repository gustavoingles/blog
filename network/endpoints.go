package network

import (
	"blogging/blog"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
)

// I won't establish or build any kind of authentication mechanism for this API. It'll be just a half a dozen endpoints related to blog
// management activities, such as posting, deleting, seeing articles, etc. Nothing hard to understand or write.
// The API will only serve "/blog/..."-like endpoints, as you will see.

func NewHTTPServer(postRepo blog.PostRepository) *http.ServeMux {
	mux := http.NewServeMux()

	postHandler := NewPostsHandler(postRepo)

	mux.HandleFunc("GET /blog/posts", postHandler.GetAllPosts)
	mux.HandleFunc("POST /blog/posts", postHandler.PublishPost)
	mux.HandleFunc("GET /blog/posts/{postID}", postHandler.GetPostById)
	mux.HandleFunc("UPDATE /blog/posts/{postID}", postHandler.UpdatePostById)
	mux.HandleFunc("DELETE /blog/posts/{postID}", postHandler.DeletePostById)
	
	return mux
}

type PostsHandler struct {
	PostRepo blog.PostRepository
}

func NewPostsHandler(postRepo blog.PostRepository) PostsHandler {
	return PostsHandler{
		PostRepo: postRepo,
	}
}

func (h PostsHandler) GetAllPosts(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	posts, err := h.PostRepo.GetAllPosts(ctx)
	if err != nil {

	}

	respBody, err := json.Marshal(posts)
	if err != nil {

	}

	_, err = w.Write(respBody)
	if err != nil {

	}
}

func (h PostsHandler) GetPostById(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	pID, err := strconv.ParseInt(r.PathValue("postID"), 10, 64)
	if err != nil {

	}

	post, err := h.PostRepo.GetPostById(ctx, pID)
	if err != nil {

	}

	respBody, err := json.Marshal(*post)
	if err != nil {

	}

	_, err = w.Write(respBody)
	if err != nil {

	}
}

func (h PostsHandler) DeletePostById(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	pID, err := strconv.ParseInt(r.PathValue("postID"), 10, 64)
	if err != nil {

	}

	err = h.PostRepo.DeletePostById(ctx, pID)
	if err != nil {

	}

	w.WriteHeader(http.StatusOK)
}

func (h PostsHandler) UpdatePostById(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	pID, err := strconv.ParseInt(r.PathValue("postID"), 10, 64)
	if err != nil {

	}

	err = h.PostRepo.UpdatePostById(ctx, pID, func(p *blog.Post) {

	})
	if err != nil {

	}

	w.WriteHeader(http.StatusOK)
}

func (h PostsHandler) PublishPost(w http.ResponseWriter, r *http.Request) {
	type PublishPostRequest struct {
		PostTitle        string `json:"title"`
		PostAuthor       string `json:"author"`
		PostBody         string `json:"body"`
		PostCTA          string `json:"call-to-action"`
		PostIntroduction string `json:"introduction"`
	}

	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		
	}

	var body PublishPostRequest
	err = json.Unmarshal(reqBody, &body)
	if err != nil {
		
	}

	ctx := r.Context()

	p := blog.NewPost(
		blog.WithTitle(body.PostTitle),
		blog.WithAuthor(body.PostAuthor),
		blog.WithBody(body.PostBody),
		blog.WithCallToAction(body.PostCTA),
		blog.WithIntroduction(body.PostIntroduction),
	)

	err = h.PostRepo.PublishPost(ctx, p)
	if err != nil {
		
	}

	w.WriteHeader(http.StatusAccepted)
}
