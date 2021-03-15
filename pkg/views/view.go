package views

import (
	"io"
)

func Home(w io.Writer) error {
	home := parse("home.html")
	return home.Execute(w, nil)
}
