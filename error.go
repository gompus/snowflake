package snowflake

import "fmt"

// UnmarshalTypeError is like json.UnmarshalTypeError, but
// uses a string to represent the Go value it could not be
// assigned to, instead of a reflect.Type.
type UnmarshalTypeError struct {
	value      string
	structName string
	field      string
	typ        string
}

func (err UnmarshalTypeError) Error() string {
	if err.structName != "" || err.field != "" {
		return fmt.Sprintf("json: cannot unmarshal %s into Go struct field %s\".\"%s of type %s", err.value, err.structName, err.field, err.typ)
	}
	return fmt.Sprintf("json: cannot unmarshal %s into Go value of type %s", err.value, err.typ)
}
