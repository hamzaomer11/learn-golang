package blogposts_test

import (
	blogposts "learn-golang/Reading-Files"
	"reflect"
	"testing"
	"testing/fstest"
)

func TestNewBlogPosts(t *testing.T) {
	fs := fstest.MapFS{
		"hello world.md":  {Data: []byte("Title: Post 1")},
		"hello-world2.md": {Data: []byte("Title: Post 2")},
	}

	posts, _ := blogposts.NewPostsFromFS(fs)

	got := posts[0]
	expected := blogposts.Post{Title: "Post 1"}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("got %+v, expected %+v", got, expected)
	}
}
