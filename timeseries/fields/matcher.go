package fields

import (
	"github.com/K-Phoen/sdk"
)

type Matcher func(field *sdk.FieldConfigOverride)

type FieldType string

const (
	FieldTypeTime FieldType = "time"
)

// ByName matches a specific field name.
func ByName(name string) Matcher { _ = "STUB: not implemented"; return *new(Matcher) }

// ByQuery matches all fields returned by the given query.
func ByQuery(ref string) Matcher { _ = "STUB: not implemented"; return *new(Matcher) }

// ByRegex matches fields names using a regex.
func ByRegex(regex string) Matcher { _ = "STUB: not implemented"; return *new(Matcher) }

// ByType matches fields with a specific type.
func ByType(fieldType FieldType) Matcher { _ = "STUB: not implemented"; return *new(Matcher) }
