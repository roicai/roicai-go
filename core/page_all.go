// Hand-written.

package core

import (
	"context"
	"iter"
)

// All returns every item of this page and of the pages after it, fetching each
// page when the loop reaches it:
//
//	for row, err := range page.All(ctx) {
//		if err != nil {
//			return err
//		}
//		fmt.Println(row)
//	}
//
// Break out of the loop to stop fetching. A failed request ends the loop with
// its error.
func (p *Page[Cursor, T, R]) All(ctx context.Context) iter.Seq2[T, error] {
	return func(yield func(T, error) bool) {
		it := p.Iterator()
		for it.Next(ctx) {
			if !yield(it.Current(), nil) {
				return
			}
		}
		if err := it.Err(); err != nil {
			var zero T
			yield(zero, err)
		}
	}
}
