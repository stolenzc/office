package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"time"
)

func isScreenLocked() (bool, error) {
	script := `tell application "System Events" to get name of current user`
	cmd := exec.Command("osascript", "-e", script)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return false, err
	}
	fmt.Println(out.String())

	// 检查输出
	return out.String() == "true\n", nil
}

func main() {
	for {
		locked, err := isScreenLocked()
		if err != nil {
			fmt.Println("Error checking screen lock status:", err)
			return
		}

		if locked {
			fmt.Println("Screen is locked. Stopping execution...")
			// 这里可以添加你需要停止运行的代码
		} else {
			fmt.Println("Screen is unlocked. Continuing execution...")
			// 这里可以添加你需要继续运行的代码
		}
		fmt.Println(time.Now())

		time.Sleep(5 * time.Second) // 每5秒检查一次
	}
}
