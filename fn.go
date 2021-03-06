package fn

import (
	"os"
)

type Handler interface {
	Handle(*Context) error
}

type HandlerFunc func(*Context) error

func (h HandlerFunc) Handle(ctx *Context) error {
	return h(ctx)
}

func HandleRequest(f HandlerFunc) error {
	io, err := NewIOHandler()
	if err != nil {
		return err
	}

	mgr := &manager{
		Handler:   f,
		IOHandler: io,
	}
	mgr.Run()

	return nil
}

func IsHotContainer() bool {
	return InputFormat() != "default"
}

func InputFormat() string {
	f := os.Getenv("FN_FORMAT")
	if f == "" {
		f = "default"
	}
	return f
}
