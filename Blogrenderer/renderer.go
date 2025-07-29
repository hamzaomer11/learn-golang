package blogrenderer

import (
	"embed"
	"html/template"
	"io"
	blogposts "learn-golang/Reading-Files"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/parser"
)

type PostRenderer struct {
	templ    *template.Template
	mdParser *parser.Parser
}

type postViewModel struct {
	blogposts.Post
	HTMLBody template.HTML
}

var (
	//go:embed "templates/*"
	postTemplates embed.FS
)

func NewPostRenderer() (*PostRenderer, error) {
	templ, err := template.ParseFS(postTemplates, "templates/*.gohtml")
	if err != nil {
		return nil, err
	}

	extensions := parser.CommonExtensions | parser.AutoHeadingIDs
	parser := parser.NewWithExtensions(extensions)

	return &PostRenderer{templ: templ, mdParser: parser}, nil
}

func (r *PostRenderer) Render(w io.Writer, p blogposts.Post) error {

	return r.templ.ExecuteTemplate(w, "blog.gohtml", newPostVM(p, r))
}

func (r *PostRenderer) RenderIndex(w io.Writer, posts []blogposts.Post) error {
	return r.templ.ExecuteTemplate(w, "index.gohtml", posts)
}

func newPostVM(p blogposts.Post, r *PostRenderer) postViewModel {
	vm := postViewModel{Post: p}
	vm.HTMLBody = template.HTML(markdown.ToHTML([]byte(p.Body), r.mdParser, nil))
	return vm
}
