package cmd

import (
	_ "embed"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "初期セットアップ",
	Run:   initial,
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
	err = generateTextInit("part1/001")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// 問題を生成
	for i := 0; i < 2; i++ {
		generateProblemInit("part1/001/problems", problemTemplate, i+1)
	}

	fmt.Printf("complete!")
}

// generateText 資料を生成する
func generateTextInit(dir string) error {
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
func generateProblemInit(path string, template string, num int) error {
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
