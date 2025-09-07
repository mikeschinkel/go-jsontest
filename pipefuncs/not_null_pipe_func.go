// Package pipefuncs provides modular pipe function implementations for the jsontest framework.
// These functions allow transformation and validation of JSON values during assertion processing.
package pipefuncs

import (
	"context"
	"strings"

	"github.com/mikeschinkel/go-jsontest"
	"github.com/tidwall/gjson"
)

// init registers the NotNullPipeFunc with the global pipe function registry.
func init() {
	jsontest.RegisterPipeFunc(&NotNullPipeFunc{
		BasePipeFunc: jsontest.NewBasePipeFunc("notNull()"),
	})
}

// Compile-time interface verification for NotNullPipeFunc.
var _ jsontest.PipeFunc = (*NotNullPipeFunc)(nil)

// NotNullPipeFunc implements the notNull() pipe function that checks if a value is not null.
type NotNullPipeFunc struct {
	jsontest.BasePipeFunc
}

// Handle checks if the current value exists and is not null, returning true/false accordingly.
func (n NotNullPipeFunc) Handle(ctx context.Context, ps *jsontest.PipeState) (err error) {
	if ps.Present && strings.TrimSpace(ps.Value.Raw) != "null" {
		ps.Value = gjson.Parse("true")
	} else {
		ps.Value = gjson.Parse("false")
	}
	ps.Present = true

	return err
}
