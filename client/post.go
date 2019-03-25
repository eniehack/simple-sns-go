// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "Simple-SNS": Post Resource Client
//
// Command:
// $ goagen
// --design=github.com/eniehack/simple-sns-go/design
// --out=D:\project\simple-sns-go
// --version=v1.3.1

package client

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/url"
)

// CreatePostPayload is the Post create action payload.
type CreatePostPayload struct {
	Body string `form:"body" json:"body" yaml:"body" xml:"body"`
}

// CreatePostPath computes a request path to the create action of Post.
func CreatePostPath() string {

	return fmt.Sprintf("/api/v1/posts/new")
}

// 新規投稿
func (c *Client) CreatePost(ctx context.Context, path string, payload *CreatePostPayload) (*http.Response, error) {
	req, err := c.NewCreatePostRequest(ctx, path, payload)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewCreatePostRequest create the request corresponding to the create action endpoint of the Post resource.
func (c *Client) NewCreatePostRequest(ctx context.Context, path string, payload *CreatePostPayload) (*http.Request, error) {
	var body bytes.Buffer
	err := c.Encoder.Encode(payload, &body, "*/*")
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %s", err)
	}
	scheme := c.Scheme
	if scheme == "" {
		scheme = "https"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("POST", u.String(), &body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	header.Set("Content-Type", "application/json")
	if c.JWTSigner != nil {
		if err := c.JWTSigner.Sign(req); err != nil {
			return nil, err
		}
	}
	return req, nil
}

// DeletePostPath computes a request path to the delete action of Post.
func DeletePostPath(postID string) string {
	param0 := postID

	return fmt.Sprintf("/api/v1/posts/%s", param0)
}

// 投稿の削除
func (c *Client) DeletePost(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewDeletePostRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewDeletePostRequest create the request corresponding to the delete action endpoint of the Post resource.
func (c *Client) NewDeletePostRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "https"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("DELETE", u.String(), nil)
	if err != nil {
		return nil, err
	}
	if c.JWTSigner != nil {
		if err := c.JWTSigner.Sign(req); err != nil {
			return nil, err
		}
	}
	return req, nil
}

// ReferencePostPath computes a request path to the reference action of Post.
func ReferencePostPath(postID string) string {
	param0 := postID

	return fmt.Sprintf("/api/v1/posts/%s", param0)
}

// 投稿の参照
func (c *Client) ReferencePost(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewReferencePostRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewReferencePostRequest create the request corresponding to the reference action endpoint of the Post resource.
func (c *Client) NewReferencePostRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "https"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
