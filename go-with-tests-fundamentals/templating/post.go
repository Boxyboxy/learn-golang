package blogrenderer

import (
	"strings"
)

// Package embed provides access to files embedded in the running Go program.
// Go source files that import "embed" can use the //go:embed directive to initialize a variable of type string,
// []byte, or FS with the contents of files read from the package directory or subdirectories at compile time.

// Post is a representation of a post
type Post struct {
	Title, Description, Body string
	Tags                     []string
}
func (p Post) SanitisedTitle() string {
	return strings.ToLower(strings.Replace(p.Title, " ", "-", -1))
}


/*
var (
	//go:embed "templates/*"
	postTemplates embed.FS
)
func Render(w io.Writer, p Post) error {
	templ, err := template.ParseFS(postTemplates, "templates/*.gohtml") // oarse the templates
	if err != nil {
		return err
	}

	if err := templ.Execute(w, p); err != nil { // use the tempalte to render a post to an io.Writer
		return err
	}

	return nil
}

*/