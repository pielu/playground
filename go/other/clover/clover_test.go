package clover

import (
	"io"
	"io/ioutil"
	"testing"

	"github.com/MakeNowJust/heredoc/v2"
)

func TestClover(t *testing.T) {
	in, err := ioutil.TempFile("", "")
	if err != nil {
		t.Fatal(err)
	}
	defer in.Close()

	s := heredoc.Doc(`
		5
		Hello
		Hello World
		Hello My World
		Hello My Beautiful World
		Twinkle twinkle little star how I wonder what you are
	`)

	_, err = io.WriteString(in, s)
	if err != nil {
		t.Fatal(err)
	}

	_, err = in.Seek(0, io.SeekStart)
	if err != nil {
		t.Fatal(err)
	}

	r := Clover(in)
	if r != "Twinkle you what wonder I how star little twinkle are" {
		t.Error("Unexpected result:", t)
	}
}
