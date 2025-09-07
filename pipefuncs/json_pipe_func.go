// Package pipefuncs provides modular pipe function implementations for the jsontest framework.
// These functions allow transformation and validation of JSON values during assertion processing.
package pipefuncs

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/mikeschinkel/go-jsontest"
	"github.com/tidwall/gjson"
)

// init registers the JSONPipeFunc with the global pipe function registry.
func init() {
	jsontest.RegisterPipeFunc(&JSONPipeFunc{
		BasePipeFunc: jsontest.NewBasePipeFunc("json()"),
	})
}

// Compile-time interface verification for JSONPipeFunc.
var _ jsontest.PipeFunc = (*JSONPipeFunc)(nil)

// JSONPipeFunc implements the json() pipe function that parses a JSON string into a JSON object.
type JSONPipeFunc struct {
	jsontest.BasePipeFunc
}

// Handle parses the current string value as JSON and makes the parsed JSON the new current value.
func (J JSONPipeFunc) Handle(ctx context.Context, ps *jsontest.PipeState) (err error) {
	var inner any
	var b []byte

	err = json.Unmarshal([]byte(ps.Value.String()), &inner)
	if err != nil {
		err = fmt.Errorf("json(): failed to parse string as JSON: %w", err)
		goto end
	}
	b, _ = json.Marshal(inner)
	ps.Value = gjson.ParseBytes(b)
	ps.Present = ps.Value.Exists()

end:
	return err
}
