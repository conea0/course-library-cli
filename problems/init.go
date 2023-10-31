package problems

import (
	"fmt"
	"os"
	"os/exec"
)

func init() {
	_, err := exec.Command("which", "python3.10").Output()
	if err != nil {
		// エラーなら終了
		fmt.Println("python3.10がインストールされていません。")
		os.Exit(1)
	}
}
