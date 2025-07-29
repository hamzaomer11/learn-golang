package blogposts_test

import (
	blogposts "learn-golang/Reading-Files"
	"reflect"
	"testing"
	"testing/fstest"
)

func TestNewBlogPosts(t *testing.T) {
	const (
		firstBody = `Title: Post 1
Description: Description 1
Tags: tdd, go
---
Hello
World`
		secondBody = `Title: Post 2
Description: Description 2
Tags: rust, borrow-checker
B
L
M`
	)

	fs := fstest.MapFS{
		"hello world.md":  {Data: []byte(firstBody)},
		"hello-world2.md": {Data: []byte(secondBody)},
	}

	posts, _ := blogposts.NewPostsFromFS(fs)

	got := posts[0]
	expected := blogposts.Post{
		Title:       "Post 1",
		Description: "Description 1",
		Tags:        []string{"tdd", "go"},
		Body: `Hello
World`,
	}

	assertPost(t, got, expected)
}

func assertPost(t *testing.T, got blogposts.Post, expected blogposts.Post) {
	t.Helper()
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("got %+v, expected %+v", got, expected)
	}
}
