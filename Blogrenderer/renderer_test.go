package blogrenderer_test

import (
	"bytes"
	blogrenderer "learn-golang/Blogrenderer"
	blogposts "learn-golang/Reading-Files"
	"testing"

	approvals "github.com/approvals/go-approval-tests"
)

func TestRenderer(t *testing.T) {
	var (
		aPost = blogposts.Post{
			Title:       "hello world",
			Body:        "This is a post",
			Description: "This is a description",
			Tags:        []string{"go", "tdd"},
		}
	)

	t.Run("it converts a single post to HTML", func(t *testing.T) {
		buf := bytes.Buffer{}

		if err := blogrenderer.Render(&buf, aPost); err != nil {
			t.Fatal(err)
		}

		approvals.VerifyString(t, buf.String())
	})
}
