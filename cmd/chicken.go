package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const chicken = 
`    
    　　　 　　(⌒⌒＼
ほげー	　　|＼＞―、)
	　＿|／・． ヽ
	　＼/　　〇　|
	　 ｜　　　　ヽノ⌒ヽ
	　 ｜　 ￣￣ヽ　 ⌒ |
	　 人　 ＼＿ノ　 /)ノ
	　　 ＼　　　　 /
	　　　 ＼＿＿_／
	　　　 ＿∥＿∥
	　　　　⌒　⌒`

var chickenCmd = &cobra.Command{
	Use:   "chicken",
	Short: "チキンのアスキーアートの表示",
	Run: func(cmd *cobra.Command, args []string) {
		m, _ := cmd.Flags().GetString("msg")
		fmt.Println(chicken)
		fmt.Println(m)
	},
}

func init() {
	rootCmd.AddCommand(chickenCmd)
}
