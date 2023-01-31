package clover

import (
	"heredoc"
	"io"
	"io/ioutil"
	"os"
	"testing"
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

	_, err = in.Seek(0, os.SEEK_SET)
	if err != nil {
		t.Fatal(err)
	}

	t := clover(in)
	if t != "Twinkle you what wonder I how star little twinkle are" {
		t.Error("unexpected result:", t)
	}
}
