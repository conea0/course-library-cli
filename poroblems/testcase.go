package poroblems

import (
	"encoding/json"
)

type Test struct {
	Input  *[]string `json:"input"`
	Output *string   `json:"output"`
}

type TestCase struct {
	Tests []Test `json:"tests"`
}

func NewTestCase(tests []Test) *TestCase {
	return &TestCase{
		Tests: tests,
	}
}

func TestCaseFromJSON(jsonBytes []byte) (*TestCase, error) {
	var testCase *TestCase
	err := json.Unmarshal(jsonBytes, &testCase)
	if err != nil {
		return nil, err
	}
	return testCase, nil
}

func (tc *TestCase) String() string {
	// テストケースをJSON形式にエンコードする
	tcJSON, err := json.MarshalIndent(tc.Tests, "", "\t")
	if err != nil {
		return err.Error()
	}
	return string(tcJSON)
}
