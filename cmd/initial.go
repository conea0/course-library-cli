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

	Run: func(cmd *cobra.Command, args []string) {
		str :="# 001\n\nここはmarkdownを自由に記述できる。\n\nここまで---\n\n## 問題\n\n## テストケース\n\n## 模範回答"
		str2 :="# 001\n\nここはmarkdownを自由に記述できる。"
		// ファイルを作成
		 err := os.MkdirAll("part1/001/problems", 0777)
		 if err != nil {
		 	fmt.Println(err)
		 }
		 i := os.MkdirAll("part1/001", 0777)
		 if i != nil {
		 	fmt.Println(i)
		 }
		 fp, err := os.Create("part1/001/problems/1.md")
    	 if err != nil {
            fmt.Println(err)
         }
		 defer fp.Close()
		 fp.Write([]byte(str))

		 p, u := os.Create("part1/001/problems/2.md")
    	 if u != nil {
            fmt.Println(u)
         }
		 defer p.Close()
		 p.Write([]byte(str))

		 t, r := os.Create("part1/001/001.md")
    	 if r != nil {
            fmt.Println(r)
         }
		 defer t.Close()
		 t.Write([]byte(str2))
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
