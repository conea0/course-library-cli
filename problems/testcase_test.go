package problems

import (
	// "fmt"
	// "reflect"
	"fmt"
	"testing"
)

func TestTestCaseFromJSON(t *testing.T) {
	testCases := []struct {
		name string
		json string
	}{
		{
			name: "valid json",
			json: `{"tests": [{"input": ["1 2 3", "3 4 5"]}, {"input": ["a b c", "aa bb cc"], "output": "some_output_value_2"}]}`},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tec, err := TestCaseFromJSON([]byte(tc.json))
			fmt.Println(tec)
			if err != nil {
				t.Errorf("TestCaseFromJSON() error = %v, wantErr %v", err, false)
			}
		})
	}
}

func TestTestCaseFromJSONError(t *testing.T) {
	testCases := []struct {
		name string
		json string
	}{
		{
			name: "invalid json",
			json: `{"input": "hello", "output": "world"`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := TestCaseFromJSON([]byte(tc.json))
			if err == nil {
				t.Errorf("TestCaseFromJSON() error = %v, wantErr %v", err, true)
			}
		})
	}
}

// test EvalTest Function
func TestEvalTest(t *testing.T) {
	testCases := []struct {
		name string
		test Test
		py   string
	}{
		{
			name: "one input",
			test: Test{
				Input:  &[]string{"1 2 3"},
				Output: nil,
			},
			py: "print(input())",
		},
		{
			name: "multi input",
			test: Test{
				Input:  &[]string{"1 2 3", "3 4 5"},
				Output: nil,
			},
			py: "print(input())\nprint(input())",
		},
		{
			name: "syntax error python code",
			test: Test{
				Input:  &[]string{"1 2 3", "3 4 5"},
				Output: nil,
			},
			py: "print(input(",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.test.EvalTest(tc.py)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Print(*(tc.test.Output))
			}

		})
	}
}
