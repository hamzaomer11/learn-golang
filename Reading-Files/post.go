package blogposts

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"
)

type Post struct {
	Title       string
	Description string
	Tags        []string
	Body        string
}

const (
	titleSeperator       = "Title: "
	descriptionSeperator = "Description: "
	tagSeperator         = "Tags: "
)

func newPost(postFile io.Reader) (Post, error) {
	scanner := bufio.NewScanner(postFile)

	readMetaLine := func(tagName string) string {
		scanner.Scan()
		return strings.TrimPrefix(scanner.Text(), tagName)
	}

	scanner.Scan()

	buf := bytes.Buffer{}
	for scanner.Scan() {
		fmt.Fprintln(&buf, scanner.Text())
	}

	body := strings.TrimSuffix(buf.String(), "\n")

	return Post{Title: readMetaLine(titleSeperator),
		Description: readMetaLine(descriptionSeperator),
		Tags:        strings.Split(readMetaLine(tagSeperator), ", "),
		Body:        body}, nil
}
