package test

import (
	"testing"

	"github.com/mikeschinkel/go-jsontest"
)

// FuzzTestJSON tests TestJSON with random JSON and assertion maps
func FuzzTestJSON(f *testing.F) {
	// Seed corpus with valid JSON and paths
	seeds := []struct {
		json string
		path string
		val  string
	}{
		{`{"name":"John"}`, "name", "John"},
		{`{"user":{"name":"Jane"}}`, "user.name", "Jane"},
		{`{"count":42}`, "count", "42"},
		{`{"active":true}`, "active", "true"},
		{`{"items":[1,2,3]}`, "items.#", "3"},
		{`{}`, "missing", ""},
		{`null`, "key", ""},
		{`[]`, "0", ""},
	}

	for _, seed := range seeds {
		f.Add(seed.json, seed.path, seed.val)
	}

	f.Fuzz(func(t *testing.T, jsonStr, path, val string) {
		// Just ensure it doesn't panic
		assertions := map[string]any{
			path: val,
		}
		_ = jsontest.TestJSON([]byte(jsonStr), assertions)
	})
}
