package views

import (
	"io"
)

func Home(w io.Writer) error {
	home := parse("pages/dadasdsahome.html")
	return home.Execute(w, nil)
}
