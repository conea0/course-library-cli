package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

// outputCmd represents the output command
var outputCmd = &cobra.Command{
	Use:   "output",
	Short: "inputを模範解答のコードで評価し、outputを出力します",
	Run:   output,
}

func init() {
	rootCmd.AddCommand(outputCmd)

	outputCmd.Flags().StringP("path", "p", "./", "パスを指定します")
}

func output(cmd *cobra.Command, args []string) {
	path, err := cmd.Flags().GetString("path")
	if err != nil {
		fmt.Fprint(os.Stderr, "実行失敗: %w", err)
	}

	paths, err := getMdPaths(path)
	if err != nil {
		fmt.Fprint(os.Stderr, "実行失敗: %w", err)
	}

	fmt.Print(paths)

	for _, f := range paths {
		err := exportProblemJSON(f)
		if err != nil {
			fmt.Fprintf(os.Stderr, "エラーが発生しました\n場所: %v\n%v\n", f, err)
			fmt.Fprintln(os.Stderr, "------------------")
		}
	}
}

// 指定されたディレクトリから.mdのファイルを抽出してフルパスを取得
func getMdPaths(path string) ([]string, error) {
	var filenames []string

	// mdファイルが指定された場合はそのまま返す
	ex := filepath.Ext(path)
	if ex == "md" {
		filenames = append(filenames, path)
		return filenames, nil
	}

	files, err := os.ReadDir(path)
	if err != nil {
		return []string{}, fmt.Errorf("ファイルが得られませんでした: %w", err)
	}

	for _, f := range files {
		// mdファイルでない場合はスキップ
		ext := filepath.Ext(f.Name())
		if ext != "md" {
			continue
		}

		fullpath := filepath.Join(path, f.Name())
		filenames = append(filenames, fullpath)
	}

	if len(files) == 0 {
		return []string{}, fmt.Errorf("指定されたディレクトリにファイルが存在しません")
	}

	return filenames, nil
}

func getOutDir(mdPath string) string {
	return filepath.Join(filepath.Dir(mdPath), "out")
}
