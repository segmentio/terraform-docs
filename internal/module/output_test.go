package module

import (
	"sort"
	"testing"

	"github.com/segmentio/terraform-docs/internal/types"
	"github.com/segmentio/terraform-docs/pkg/tfconf"
	"github.com/stretchr/testify/assert"
)

func TestOutputsSortedByName(t *testing.T) {
	assert := assert.New(t)
	outputs := sampleOutputs()

	sort.Sort(outputsSortedByName(outputs))

	expected := []string{"a", "b", "c", "d", "e"}
	actual := make([]string, len(outputs))

	for k, o := range outputs {
		actual[k] = o.Name
	}

	assert.Equal(expected, actual)
}

func TestOutputsSortedByPosition(t *testing.T) {
	assert := assert.New(t)
	outputs := sampleOutputs()

	sort.Sort(outputsSortedByPosition(outputs))

	expected := []string{"d", "a", "e", "b", "c"}
	actual := make([]string, len(outputs))

	for k, o := range outputs {
		actual[k] = o.Name
	}

	assert.Equal(expected, actual)
}

func sampleOutputs() []*tfconf.Output {
	return []*tfconf.Output{
		&tfconf.Output{
			Name:        "a",
			Description: types.String("description of a"),
			Value:       nil,
			Position:    tfconf.Position{Filename: "foo/outputs.tf", Line: 25},
		},
		&tfconf.Output{
			Name:        "d",
			Description: types.String("description of d"),
			Value:       nil,
			Position:    tfconf.Position{Filename: "foo/outputs.tf", Line: 10},
		},
		&tfconf.Output{
			Name:        "e",
			Description: types.String("description of e"),
			Value:       nil,
			Position:    tfconf.Position{Filename: "foo/outputs.tf", Line: 33},
		},
		&tfconf.Output{
			Name:        "b",
			Description: types.String("description of b"),
			Value:       nil,
			Position:    tfconf.Position{Filename: "foo/outputs.tf", Line: 39},
		},
		&tfconf.Output{
			Name:        "c",
			Description: types.String("description of c"),
			Value:       nil,
			Position:    tfconf.Position{Filename: "foo/outputs.tf", Line: 42},
		},
	}
}
