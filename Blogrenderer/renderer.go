package blogrenderer

import (
	"embed"
	"html/template"
	"io"
	blogposts "learn-golang/Reading-Files"
)

type PostRenderer struct {
	templ *template.Template
}

func NewPostRenderer() (*PostRenderer, error) {
	templ, err := template.ParseFS(postTemplates, "templates/*.gohtml")
	if err != nil {
		return nil, err
	}

	return &PostRenderer{templ: templ}, nil
}

var (
	//go:embed "templates/*"
	postTemplates embed.FS
)

func (r *PostRenderer) Render(w io.Writer, p blogposts.Post) error {

	if err := r.templ.ExecuteTemplate(w, "blog.gohtml", p); err != nil {
		return err
	}

	return nil
}
