package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
	"strconv"
)

// gCmd represents the g command
var gCmd = &cobra.Command{
	Use:   "g",
	Short: "A brief description of your command",

	Run: func(cmd *cobra.Command, args []string) {
		
		//コマンドライン引数を取得
		arg1 := os.Args[1]
		// arg3をint型に変換
		arg3, _ := strconv.Atoi(os.Args[3])
		// 文字列を定義
		str :="# 1\n\nここはmarkdownを自由に記述できる。\n\nここまで---\n\n## 問題\n\n## テストケース\n\n## 模範回答"
		str2 :="# 001\n\nここはmarkdownを自由に記述できる。"
		// ファイルを作成
		 err := os.MkdirAll(arg1 + "/problems", 0777)
		 if err != nil {
		 	fmt.Println(err)
		 }
		 i := os.MkdirAll(arg1, 0777)
		 if i != nil {
		 	fmt.Println(i)
		 }
		 // arg3分だけファイルを作成
		 for j := 1; j <= arg3; j++ {
			// stringをintに変換
			s :=strconv.Itoa(j)
		 	fp, err := os.Create(arg1 + "/problems/" + s + ".md")
	 	 	if err != nil {
			  fmt.Println(err)
		   }
		 	defer fp.Close()
		 	fp.Write([]byte(str))
		 }
		 //arg1をスライス
		 n := arg1[7:]
		 t, r := os.Create(arg1 +  n + ".md")
    	 if r != nil {
            fmt.Println(r)
         }
		 defer t.Close()
		 t.Write([]byte(str2))
	},
}

func init() {
	rootCmd.AddCommand(gCmd)
}
