package cmd

import (
	_"embed" 
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const (
	// embedを指定した場所からの相対パス
	problemTemplatePath = "template/problem.md"
	textTemplatePath    = "template/text.md"
)

//go:embed template/problem.md
var problemsTemplate string

//go:embed template/text.md
var textTemplate string

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "初期セットアップ",
	Run: initial,
} 
	
func init() {
	rootCmd.AddCommand(initCmd)
}
	
func initial(cmd *cobra.Command, args []string) {
	// ディレクトリを作成
	err := os.MkdirAll("part1/001/problems", 0777)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//資料を生成
	err = generateText("part1/001")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// 問題を生成
	for i := 0; i < 2; i++ {
		generateProblem("part1/001/problems", problemsTemplate, i+1)
	}


	fmt.Printf("complete!")
}

// generateText 資料を生成する
func generateText(dir string) error {
	// ファイルを作成
	f, err := os.Create("part1/001/001.md")
	if err != nil {
		return err
	}
	defer f.Close()
	// テンプレートを埋め込む
	fmt.Fprint(f, textTemplate)

	return nil
}

// generateProblem 問題を生成する
func generateProblem(path string, template string, num int) error {
	// ファイルを作成
	f, err := os.Create(fmt.Sprintf("%s/%03d.md", path, num))
	if err != nil {
		return err
	}
	defer f.Close()
	// テンプレートを埋め込む
	fmt.Fprint(f, template)

	return nil
}

