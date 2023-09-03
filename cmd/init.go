package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "A brief description of your command",
//├── part1
        //│     └── 001
        //│       ├── 001.md
        //│       └── problems
        //│             ├── 1.md
        //│             └── 2.md
		//上記のようなディレクトリ構造を作成する機能
	Run: func(cmd *cobra.Command, args []string) {
		err := os.MkdirAll("part1/001/problems/1.md,2.md", 0777)
		if err != nil {
			fmt.Println(err)
		}
		i := os.MkdirAll("part1/001/001.md", 0777)
		if i != nil {
			fmt.Println(i)
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
