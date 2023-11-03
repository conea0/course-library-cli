package problems

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"

	"github.com/titanous/json5"
)

type Test struct {
	Input  *[]any `json:"input"`
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
	err := json5.Unmarshal(jsonBytes, &testCase)
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
func (t *Test) EvalTest(py string) error {
	// 受け取ったPythonコードをファイルに書き込む
	fp, err := os.Create("tmp.py")
	if err != nil {
		return fmt.Errorf("failed to create tmp.py: %w", err)
	}
	defer os.Remove("tmp.py")
	defer fp.Close()
	fp.WriteString(py)
	fp.Sync()

	// テストコードを実行する
	cmd := exec.Command("python", "tmp.py")

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return fmt.Errorf("failed to get stdin pipe: %w", err)
	}
	defer stdin.Close()

	// 標準入力に値を渡す
	for _, str := range *t.Input {
		fmt.Fprintln(stdin, str)
	}

	result, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf(string(result))
	}
	output := string(result)

	// テスト結果をtcに格納する
	t.Output = &output

	return nil
}

func (tc *TestCase) EvalTests(py string) error {
	// テストケースを実行する
	for _, t := range tc.Tests {
		err := t.EvalTest(py)
		if err != nil {
			return err
		}
	}
	return nil
}
