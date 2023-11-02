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

