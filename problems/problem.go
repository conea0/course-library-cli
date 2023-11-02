package problems

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type Problem struct {
	Statement string
	TestCase  *TestCase
	Code      string
}


type MdPrefix string

type ReadBlockFn func(*Problem) error

const (
	StatementPre = "## 問題"
	TestcasePre  = "## テストケース"
	CodePre      = "## 模範回答"
)

type Md struct {
	s []string

	mdPrefixFns map[MdPrefix]ReadBlockFn

	curr int
	next int

	errors []error
}

func NewMd(r io.Reader) *Md {
	m := &Md{
		errors:      []error{},
		mdPrefixFns: make(map[MdPrefix]ReadBlockFn),
	}

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

	m.registerReadBlockFn(StatementPre, m.readStatement)
	m.registerReadBlockFn(TestcasePre, m.readTestcase)
	m.registerReadBlockFn(CodePre, m.readCode)

	return m
}

func (m *Md) Scan() bool {
	return m.next < len(m.s)
}

func (m *Md) Text() string {
	if m.next >= len(m.s) {
		return ""
	}

	m.curr = m.next
	m.next++

	return m.s[m.curr]
}

func (m *Md) Peek() string {
	if m.next >= len(m.s) {
		return ""
	}

	return m.s[m.next]
}

func (m *Md) registerReadBlockFn(pre MdPrefix, fn ReadBlockFn) {
	m.mdPrefixFns[pre] = fn
}

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

func (m *Md) readStatement(p *Problem) error {
	var existBlock bool
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
		return fmt.Errorf("cannot read test case block: %w", err)
	}

	p.TestCase = tcJSON
	return nil
}

func (m *Md) readCode(p *Problem) error {
	if err := m.skipToCodeBlock("python"); err != nil {
		return err
	}

	for m.Scan() {
		txt := m.Peek()
		if strings.HasPrefix(txt, "```") {
			break
		}

		p.Code += m.Text() + "\n"
	}

	return nil
}

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

func (m *Md) Err() []error {
	return m.errors
}

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
