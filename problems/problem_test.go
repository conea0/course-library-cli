package problems

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadProblem(t *testing.T) {

	tc, err := TestCaseFromJSON([]byte(`{"tests": [{"input": ["こんにちは"]}]}`))
	if err != nil {
		t.Error(err)
	}

	testCase := map[string]interface{}{
		"stmt":   "問題です。\n",
		"input1": (*tc.Tests[0].Input)[0],
		"code":   "gre = input()\n",
	}

	fp, err := os.Open("./test.md")
	if err != nil {
		t.Error(err)
	}

	md := NewMd(fp)
	p := md.ReadProblem()
	if len(md.Err()) != 0 {
		t.Error(md)
	}

	log.Printf("%+v", p)

	assert.Equal(t, testCase["stmt"], p.Statement)
	assert.Equal(t, testCase["input1"], (*p.TestCase.Tests[0].Input)[0])
	assert.Equal(t, testCase["code"], p.Code)
}
