package poroblems

type Test struct {
	Input []string `json:"input"`
	Output string `json:"output"`
}

type TestCase struct {
	Tests []Test
}
