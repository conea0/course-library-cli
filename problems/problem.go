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
