package problems

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

// Problem is a struct that represents a problem in the course library.
type Problem struct {
	Statement string
	TestCase  *TestCase
	Code      string
}

// MdPrefix is a type alias for a string that represents a markdown prefix.
type MdPrefix string

// ReadBlockFn is a function type that reads a markdown block and updates a Problem struct.
type ReadBlockFn func(*Problem) error

const (
	// StatementPre is a markdown prefix for the problem statement block.
	StatementPre = "## 問題"
	// TestcasePre is a markdown prefix for the test case block.
	TestcasePre = "## テストケース"
	// CodePre is a markdown prefix for the code block.
	CodePre = "## 模範回答"
)

// Md is a struct that represents a markdown file.
type Md struct {
	s []string

	mdPrefixFns map[MdPrefix]ReadBlockFn

	curr int // the current line being read
	next int // the next line to read

	errors []error
}

// NewMd creates a new Md struct from an io.Reader.
func NewMd(r io.Reader) *Md {
	m := &Md{
		errors:      []error{},
		mdPrefixFns: make(map[MdPrefix]ReadBlockFn),
	}

	// read the file line by line and store it in a slice
	scanner := bufio.NewScanner(r)
	var s []string

	for scanner.Scan() {
		txt := scanner.Text()
		s = append(s, txt)
	}

	if len(s) == 0 {
		m.errors = append(m.errors, fmt.Errorf("問題文のファイルが空です"))
	}

	m.s = s

	// register each parsing function
	m.registerReadBlockFn(StatementPre, m.readStatement)
	m.registerReadBlockFn(TestcasePre, m.readTestcase)
	m.registerReadBlockFn(CodePre, m.readCode)

	return m
}

// Scan returns true if there are more lines to read.
func (m *Md) Scan() bool {
	return m.next < len(m.s)
}

// Text returns the current line and moves to the next line.
func (m *Md) Text() string {
	if m.next >= len(m.s) {
		return ""
	}

	m.curr = m.next
	m.next++

	return m.s[m.curr]
}

// Peek returns the next line without moving the cursor.
func (m *Md) Peek() string {
	if m.next >= len(m.s) {
		return ""
	}

	return m.s[m.next]
}

// registerReadBlockFn registers a parsing function for a markdown prefix.
func (m *Md) registerReadBlockFn(pre MdPrefix, fn ReadBlockFn) {
	m.mdPrefixFns[pre] = fn
}

// ReadProblem reads a markdown file and returns a Problem struct.
func (m *Md) ReadProblem() *Problem {
	p := &Problem{}

	for m.Scan() {
		txt := m.Text()
		for pre, fn := range m.mdPrefixFns {
			if strings.HasPrefix(txt, string(pre)) {
				fn(p)
			}
		}
	}

	return p
}

// readStatement reads the problem statement block and updates a Problem struct.
func (m *Md) readStatement(p *Problem) error {
	var existBlock bool // whether the block exists

	for m.Scan() {
		peekTxt := m.Peek()
		if strings.HasPrefix(peekTxt, TestcasePre) {
			existBlock = true
			break
		}

		p.Statement += m.Text() + "\n"
	}

	if !existBlock {
		err := fmt.Errorf("問題文のブロックが見つかりませんでした")
		m.errors = append(m.errors, err)
		return err
	}

	return nil
}

// readTestcase reads the test case block and updates a Problem struct.
func (m *Md) readTestcase(p *Problem) error {
	var s string
	var blockIsEnded bool

	m.skipToCodeBlock("json")
	for m.Scan() {
		txt := m.Text()
		if strings.HasPrefix(txt, "```") {
			blockIsEnded = true
			break
		}

		s += txt + "\n"
	}

	if !blockIsEnded {
		err := fmt.Errorf("テストケースブロックが閉じられていません")
		m.errors = append(m.errors, err)
		return err
	}

	s = fmt.Sprintf(`{"tests": %s}`, s)
	tcJSON, err := TestCaseFromJSON([]byte(s))
	if err != nil {
		m.errors = append(m.errors, err)
		m.errors = append(m.errors, fmt.Errorf(jsonErrMsg))
		return fmt.Errorf("cannot read test case block: %w", err)
	}

	p.TestCase = tcJSON
	return nil
}

// readCode reads the code block and updates a Problem struct.
func (m *Md) readCode(p *Problem) error {
	var blockIsEnded bool

	if err := m.skipToCodeBlock("python"); err != nil {
		return err
	}

	for m.Scan() {
		txt := m.Peek()
		if strings.HasPrefix(txt, "```") {
			blockIsEnded = true
			break
		}

		p.Code += m.Text() + "\n"
	}

	if !blockIsEnded {
		err := fmt.Errorf("模範解答のコードブロックが閉じられていません")
		m.errors = append(m.errors, err)
		return err
	}

	return nil
}

// skipToCodeBlock skips lines until it finds a code block.
func (m *Md) skipToCodeBlock(t string) error {
	var existBlock bool
	curr, next := m.curr, m.next

	for m.Scan() {
		txt := m.Text()
		expectPre := fmt.Sprintf("```%s", t)
		if strings.HasPrefix(txt, expectPre) {
			existBlock = true
			break
		}
	}

	if !existBlock {
		m.curr, m.next = curr, next

		err := fmt.Errorf("%sブロックが見つかりませんでした", t)
		m.errors = append(m.errors, err)
		return err
	}

	return nil
}

// Err returns a slice of errors that occurred during parsing.
func (m *Md) Err() []error {
	return m.errors
}

// Error returns a string of all errors that occurred during parsing.
func (m *Md) Error() string {
	var s string
	if len(m.Err()) == 0 {
		return ""
	}

	for _, err := range m.Err() {
		s = fmt.Sprintf("%s\n%s", s, err.Error())
	}

	return s
}

const jsonErrMsg = `
jsonの書き方を間違えている可能性があります。
	- 最後の要素に,がないか確認してみてください
	- jsonファイルに書き込んでみてください
`
