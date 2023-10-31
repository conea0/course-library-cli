package poroblems

import (
	"encoding/json"
)

type Test struct {
	Input  []string `json:"input"`
	Output string   `json:"output"`
}

type TestCase struct {
	Tests []Test
}

func (tc TestCase) String() string {
	// テストケースをJSON形式にエンコードする
	tcJSON, err := json.MarshalIndent(tc.Tests, "", "\t")
	if err != nil {
		return err.Error()
	}
	return string(tcJSON)
}
