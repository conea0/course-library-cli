package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"io"

	"github.com/spf13/cobra"
)

const (
	problemTemplatePath = "template/problem.md"
	textTemplatePath = "template/text.md"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: generate,
}

func init() {
	rootCmd.AddCommand(generateCmd)

	generateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	generateCmd.Flags().IntP("num", "n", 1, "演習問題の個数を指定します。")
	generateCmd.Flags().StringP("path", "p", "./", "生成先を指定します。")
}


func generate(cmd *cobra.Command, args []string) {
	num, err := cmd.Flags().GetInt("num")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	path, err := cmd.Flags().GetString("path")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(num, path)
}