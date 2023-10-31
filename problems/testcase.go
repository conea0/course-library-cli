package problems

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
func (t *Test) EvalTest(py string) error {
	// 受け取ったPythonコードをファイルに書き込む
	fp, err := os.Create("tmp.py")
	if err != nil {
		return err
	}
	defer os.Remove("tmp.py")
	defer fp.Close()
	fp.WriteString(py)
	fp.Sync()

	// テストコードを実行する
	cmd := exec.Command("python", "tmp.py")

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return err
	}
	defer stdin.Close()

	for _, str := range *t.Input {
		fmt.Fprintln(stdin, str)
		// stdin.Write([]byte(str))
	}

	result, err := cmd.Output()
	if err != nil {
		return err
	}
	output := string(result)

	// テスト結果をtcに格納する
	t.Output = &output

	return nil
}
