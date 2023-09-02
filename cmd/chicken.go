/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
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

// chickenCmd represents the chicken command
var chickenCmd = &cobra.Command{
	Use:   "chicken",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		m, _ := cmd.Flags().GetString("msg")
		fmt.Println(chicken)
		fmt.Println(m)
	},
}

func init() {
	rootCmd.AddCommand(chickenCmd)

	chickenCmd.Flags().StringP("msg", "m", "", "message to print")

}
