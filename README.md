# Blog API

A RESTful API that powers a personal blog, providing complete CRUD operations for managing blog articles.

## Overview

This project is a backend API designed to handle all the essential operations for a personal blogging platform. It provides a clean, RESTful interface for creating, reading, updating, and deleting blog posts.

## Features

| Operation | Endpoint | Description |
|-----------|----------|-------------|
| **List Articles** | `GET /blog/posts` | Returns a list of all articles |
| **Get Article** | `GET /blog/posts/{postID}` | Returns a single article by ID |
| **Create Article** | `POST /blog/posts` | Publishes a new article |
| **Update Article** | `PUT /blog/posts/{postID}` | Updates an existing article |
| **Delete Article** | `DELETE /blog/posts/{postID}` | Removes an article |

## Tech Stack

- **Language:** [Go (Golang)](https://go.dev/)

## API Documentation

Full API documentation is available in the [blog-api.yaml](./blog-api.yaml) file, following the OpenAPI 3.0.3 specification.

## Getting Started

```bash
# Clone the repository
git clone <repository-url>
cd blogging

# Run the application
go run main.go
```

## Project Structure

```
blogging/
├── blog/           # Core blog domain logic
├── network/        # HTTP handlers and endpoints
├── blog-api.yaml   # OpenAPI specification
└── README.md
```

## License

MIT
