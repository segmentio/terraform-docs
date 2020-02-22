package types

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {
	notNilValues := make([]interface{}, 0)
	notNilValues = append(notNilValues, "foo", "42", "false", "true")

	emptyValues := make([]interface{}, 0)
	emptyValues = append(emptyValues, "")

	nilValues := make([]interface{}, 0)
	nilValues = append(nilValues, nil)

	type expected struct {
		typeName   string
		valueKind  string
		hasDefault bool
	}
	tests := []struct {
		name     string
		values   []interface{}
		types    string
		expected expected
	}{
		{
			name:   "value not nil and type string",
			values: notNilValues,
			types:  "string",
			expected: expected{
				typeName:   "string",
				valueKind:  "types.String",
				hasDefault: true,
			},
		},
		{
			name:   "value not nil and type empty",
			values: notNilValues,
			types:  "",
			expected: expected{
				typeName:   "string",
				valueKind:  "types.String",
				hasDefault: true,
			},
		},
		{
			name:   "value empty and type string",
			values: emptyValues,
			types:  "string",
			expected: expected{
				typeName:   "string",
				valueKind:  "types.Empty",
				hasDefault: true,
			},
		},
		{
			name:   "value empty and type empty",
			values: emptyValues,
			types:  "",
			expected: expected{
				typeName:   "string",
				valueKind:  "types.Empty",
				hasDefault: true,
			},
		},
		{
			name:   "value nil and type number",
			values: nilValues,
			types:  "string",
			expected: expected{
				typeName:   "string",
				valueKind:  "*types.Nil",
				hasDefault: false,
			},
		},
		{
			name:   "value nil and type empty",
			values: nilValues,
			types:  "",
			expected: expected{
				typeName:   "any",
				valueKind:  "*types.Nil",
				hasDefault: false,
			},
		},
	}
	for _, tt := range tests {
		for _, tv := range tt.values {
			t.Run(tt.name, func(t *testing.T) {
				assert := assert.New(t)

				actualValue := ValueOf(tv)
				actualType := TypeOf(tt.types, tv)

				assert.Equal(tt.expected.typeName, string(actualType))
				assert.Equal(tt.expected.valueKind, reflect.TypeOf(actualValue).String())
				assert.Equal(tt.expected.hasDefault, actualValue.HasDefault())
			})
		}
	}
}
