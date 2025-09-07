// Package pipefuncs provides modular pipe function implementations for the jsontest framework.
// These functions allow transformation and validation of JSON values during assertion processing.
package pipefuncs

import (
	"context"
	"fmt"

	"github.com/mikeschinkel/go-jsontest"
	"github.com/tidwall/gjson"
)

// init registers the LenPipeFunc with the global pipe function registry.
func init() {
	jsontest.RegisterPipeFunc(&LenPipeFunc{
		BasePipeFunc: jsontest.NewBasePipeFunc("len()"),
	})
}

// Compile-time interface verification for LenPipeFunc.
var _ jsontest.PipeFunc = (*LenPipeFunc)(nil)

// LenPipeFunc implements the len() pipe function that returns the length of arrays, objects, or strings.
type LenPipeFunc struct {
	jsontest.BasePipeFunc
}

// Handle calculates the length of the current value and returns it as a number.
func (l LenPipeFunc) Handle(ctx context.Context, ps *jsontest.PipeState) (err error) {
	var n int
	switch {
	case ps.Value.IsArray():
		n = len(ps.Value.Array())
	case jsontest.IsJSONObject(ps.Value): // <-- treat {} as object with 0 keys
		n = len(ps.Value.Map())
	default:
		n = len(ps.Value.String())
	}
	ps.Value = gjson.Parse(fmt.Sprintf("%d", n))
	ps.Present = true

	return err
}
