package cmd

import (
	_ "embed"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

//go:embed template/problem.md
var problemTemplate string

//go:embed template/text.md
var textTemplate string

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:     "generate",
	Aliases: []string{"g"},
	Short:   "コース作成に必要なファイルを生成します。",
	Run:     generate,
}

func init() {
	rootCmd.AddCommand(generateCmd)

	generateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	generateCmd.Flags().IntP("num", "n", 1, "演習問題の個数を指定します。")
	generateCmd.Flags().StringP("path", "p", "./", "生成先を指定します。")
}

func generate(cmd *cobra.Command, args []string) {
	numProblems, outputDirPath, problemsOutputDirPath := getGenCmdArgs(cmd)

	// ディレクトリを作成
	err := os.MkdirAll(problemsOutputDirPath, 0777)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// 資料を生成する
	err = generateText(outputDirPath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// 問題を生成する
	for i := 0; i < numProblems; i++ {
		generateProblem(problemsOutputDirPath, problemTemplate, i+1)
	}

	fmt.Printf("%d個の問題を生成しました。\n", numProblems)
	fmt.Printf("生成先: %s\n", outputDirPath)
}

// 問題を生成する
func generateProblem(path string, problemTemplate string, num int) {
	// テンプレートファイルを読み込む

	fileName := fmt.Sprintf("%s/%v.md", path, num)

	// ファイルを作成
	problemFile, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer problemFile.Close()

	// ファイルに書き込む
	head := fmt.Sprintf("# %v", num)
	fmt.Fprint(problemFile, head)
	fmt.Fprint(problemFile, problemTemplate)
}

// 資料を生成する
func generateText(dirPath string) error {
	// ファイルを作成
	outputPath := filepath.Join(dirPath, "text.md")
	textFile, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer textFile.Close()

	// ファイルに書き込む
	fmt.Fprint(textFile, textTemplate)

	return nil
}

// コマンドライン引数を取得する
func getGenCmdArgs(cmd *cobra.Command) (int, string, string) {
	numProblems, err := cmd.Flags().GetInt("num")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	outputDirPath, err := cmd.Flags().GetString("path")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	problemsOutputDirPath := filepath.Join(outputDirPath, "problems")

	return numProblems, outputDirPath, problemsOutputDirPath
}
