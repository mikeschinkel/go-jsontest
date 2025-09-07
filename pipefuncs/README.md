# JSON Test Pipe Functions

This package provides modular pipe function implementations for the `jsontest` framework. Pipe functions allow transformation and validation of JSON values during assertion processing through a pipeline-style syntax.

## Available Pipe Functions

### exists()
- **Purpose**: Checks if a JSON path exists in the data
- **Returns**: `true` if the path exists (even if the value is null), `false` otherwise
- **Usage**: `"path.to.field|exists()": true`

### notNull()
- **Purpose**: Validates that a value exists and is not null
- **Returns**: `true` if the value exists and is not null, `false` otherwise
- **Usage**: `"path.to.field|notNull()": true`

### notEmpty()
- **Purpose**: Checks if a value is considered non-empty
- **Returns**: `true` for non-empty values, `false` for empty ones
- **Empty values**: `""` (empty string), `[]` (empty array), `{}` (empty object), `null`, non-existent paths
- **Non-empty values**: Non-empty strings, arrays with items, objects with properties, numbers (including 0), booleans
- **Usage**: `"path.to.field|notEmpty()": true`

### len()
- **Purpose**: Returns the length of arrays, objects, or strings
- **Returns**: Number representing the length
- **Behavior**:
  - Arrays: Number of elements
  - Objects: Number of properties
  - Strings: Number of characters
  - Other types: 0
- **Usage**: `"path.to.array|len()": 3`

### json()
- **Purpose**: Parses a JSON string into a JSON object for further processing
- **Returns**: Parsed JSON object that can be used with subpaths
- **Usage**: `"path.to.jsonString|json()|subpath": "value"`
- **Note**: Expects the current value to be a valid JSON string

## Usage Examples

```go
checks := map[string]any{
    // Check if a field exists
    "user.name|exists()": true,
    
    // Validate non-null values
    "user.id|notNull()": true,
    
    // Check for non-empty content
    "user.email|notEmpty()": true,
    
    // Get array length
    "users|len()": 5,
    
    // Parse JSON string and access nested data
    "metadata.jsonData|json()|nested.field": "expected_value",
}

err := jsontest.TestJSON(jsonData, checks)
```

## Chaining Pipe Functions

Pipe functions can be chained together for complex validations:

```go
checks := map[string]any{
    // Parse JSON, navigate to field, check if non-empty
    "config.settings|json()|database.host|notEmpty()": true,
    
    // Get array length from parsed JSON
    "response.data|json()|items|len()": 10,
}
```

## Architecture

Each pipe function:

1. **Inherits from `BasePipeFunc`** - Provides common functionality and name normalization
2. **Implements `PipeFunc` interface** - Ensures consistent behavior across all functions
3. **Registers via `init()`** - Automatically available when package is imported
4. **Handles `PipeState`** - Manages value transformation through the pipeline
5. **Returns boolean or transformed values** - Supports various assertion patterns

## Adding New Pipe Functions

To create a new pipe function:

1. Create a new file `your_pipe_func.go`
2. Implement the `PipeFunc` interface
3. Register in `init()` function
4. Add to the `Initialize()` function reference list
5. Update this README with documentation

See existing functions as examples for implementation patterns.