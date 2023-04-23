package app

import (
	"context"
	"errors"
)

type CloseFunction func(ctx context.Context) error

type closer struct {
	funcs []CloseFunction
}

func newCloser() *closer {
	return &closer{
		funcs: make([]CloseFunction, 0),
	}
}

func (c *closer) Add(f CloseFunction) {
	c.funcs = append(c.funcs, f)
}

func (c *closer) Close(ctx context.Context) error {
	errs := make([]error, 0, len(c.funcs))

	done := make(chan struct{})

	go func() {
		defer close(done)
		for _, f := range c.funcs {
			if err := f(ctx); err != nil {
				errs = append(errs, err)
			}
		}
	}()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-done:
	}

	if len(errs) != 0 {
		return errors.Join(errs...)
	}

	return nil
}
