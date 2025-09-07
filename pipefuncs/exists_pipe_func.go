// Package pipefuncs provides modular pipe function implementations for the jsontest framework.
// These functions allow transformation and validation of JSON values during assertion processing.
package pipefuncs

import (
	"context"

	"github.com/mikeschinkel/go-jsontest"
	"github.com/tidwall/gjson"
)

// init registers the ExistsPipeFunc with the global pipe function registry.
func init() {
	jsontest.RegisterPipeFunc(&ExistsPipeFunc{
		BasePipeFunc: jsontest.NewBasePipeFunc("exists()"),
	})
}

// Compile-time interface verification for ExistsPipeFunc.
var _ jsontest.PipeFunc = (*ExistsPipeFunc)(nil)

// ExistsPipeFunc implements the exists() pipe function that checks if a JSON path exists.
type ExistsPipeFunc struct {
	jsontest.BasePipeFunc
}

// Handle checks if the current value exists and returns true/false accordingly.
func (e ExistsPipeFunc) Handle(ctx context.Context, ps *jsontest.PipeState) (err error) {
	if ps.Present {
		ps.Value = gjson.Parse("true")
	} else {
		ps.Value = gjson.Parse("false")
	}
	ps.Present = true // the boolean result "exists" itself exists

	return err
}
