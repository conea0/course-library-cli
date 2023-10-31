package poroblems

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {
	testCase := TestCase{
		Tests: []Test{
			{
				Input:  []string{"1 2 3", "3 4 5"},
				Output: "some_output_value_1",
			},
			{
				Input:  []string{"a b c", "aa bb cc"},
				Output: "some_output_value_2",
			},
		},
	}

	expected := `[
	{
		"input": [
			"1 2 3",
			"3 4 5"
		],
		"output": "some_output_value_1"
	},
	{
		"input": [
			"a b c",
			"aa bb cc"
		],
		"output": "some_output_value_2"
	}
]`

	assert.Equal(t, expected, testCase.String())
}
