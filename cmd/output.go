package cmd

import (
	"clc/problems"
	"fmt"
	"os"
	"path/filepath"

	yaml "gopkg.in/yaml.v2"

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
		fmt.Fprintf(os.Stderr, "実行失敗: %v", err)
	}

	paths, err := getMdPaths(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "実行失敗: %v", err)
	}

	for _, f := range paths {
		err := exportProblemJSON(f)
		if err != nil {
			fullpath, _ := filepath.Abs(f)
			fmt.Fprintln(os.Stderr, "#############################################")
			fmt.Fprintf(os.Stderr, "エラーが発生しました\n場所: %v\n\n%v\n", fullpath, err)
			fmt.Fprintln(os.Stderr, "#############################################")
		}
	}
}

// 指定されたディレクトリから.mdのファイルを抽出してフルパスを取得
func getMdPaths(path string) ([]string, error) {
	var filenames []string

	// mdファイルが指定された場合はそのまま返す
	ex := filepath.Ext(path)
	if ex == ".md" {
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
		if ext != ".md" {
			continue
		}

		fullpath := filepath.Join(path, f.Name())
		filenames = append(filenames, fullpath)
	}

	if len(filenames) == 0 {
		return []string{}, fmt.Errorf("指定されたディレクトリにファイルが存在しません")
	}

	return filenames, nil
}

func getOutDir(mdPath string) (string, error) {
	dir := filepath.Join(filepath.Dir(mdPath), "out")
	if f, err := os.Stat(dir); os.IsNotExist(err) || !f.IsDir() {
		if err := os.Mkdir(dir, 0777); err != nil {
			return "", err
		}
	}

	return dir, nil
}

func exportProblemJSON(f string) error {
	file, err := os.Open(f)
	if err != nil {
		return err
	}

	defer file.Close()

	md := problems.NewMd(file)

	p := md.ReadProblem()
	if len(md.Err()) != 0 {
		return fmt.Errorf(md.Error())
	}

	tc := p.TestCase
	if err := tc.EvalTests(p.Code); err != nil {
		return err
	}

	problemyaml, err := yaml.Marshal(p)
	if err != nil {
		return fmt.Errorf("jsonの出力に失敗しました: %w", err)
	}

	outDir, err := getOutDir(f)
	if err != nil {
		return fmt.Errorf("ディレクトリの作成に失敗しました: %w", err)
	}

	yamlName := getFileNameWithoutExt(f) + ".yaml"
	fullpath := filepath.Join(outDir, yamlName)
	outFile, err := os.Create(fullpath)
	if err != nil {
		return fmt.Errorf("ファイルの作成に失敗しました: %w", err)
	}

	fmt.Fprint(outFile, string(problemyaml))

	return nil
}

func getFileNameWithoutExt(path string) string {
    return filepath.Base(path[:len(path)-len(filepath.Ext(path))])
}
