package xtypes

import "context"

// GeneratorSimple returns generator which is not controlled from external
func GeneratorSimple[T any](size int, genFn func() (T, bool)) <-chan T {
	ch := make(chan T, size)
	go func() {
		defer close(ch)
		for {
			val, ok := genFn()
			if !ok {
				break
			}
			ch <- val
		}
	}()
	return (<-chan T)(ch)
}

// Generator returns generator with context Done support
func Generator[T any](ctx context.Context, size int, genFn func(ctx context.Context) (T, bool)) <-chan T {
	ch := make(chan T, size)
	go func() {
		defer close(ch)
		for {
			val, ok := genFn(ctx)
			if !ok {
				break
			}
			select {
			case <-ctx.Done():
				return
			case ch <- val:
			}
		}
	}()
	return (<-chan T)(ch)
}
