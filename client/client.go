package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"time"
)

type Config struct {
	ServerAddress string `json:"server_address"`
}

func (c *Config)loadConfig(filePath string) (error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, &c)
	return err
}

var config Config

func isScreenLocked() (bool, error) {
	cmd := exec.Command("pmset", "-g", "assertions")
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

func updateStatus(url string) {
	// PUT请求的Body
	body := bytes.NewBufferString("{ \"key\": \"value\" }")
	// 创建PUT请求
	req, err := http.NewRequest("PUT", url, body)
	if err != nil {
		fmt.Println("Error creating request:", err)
	}
	client := &http.Client{}
	_, err = client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
	}
}

func main() {
	var config = Config{}
	config.loadConfig("./config.json")
	for {
		updateStatus(config.ServerAddress)

		// locked, err := isScreenLocked()
		// if err != nil {
		// 	fmt.Println("Error checking screen lock status:", err)
		// 	return
		// }

		// if locked {
		// 	fmt.Println("Screen is locked. Stopping execution...")
		// 	// 这里可以添加你需要停止运行的代码
		// } else {
		// 	fmt.Println("Screen is unlocked. Continuing execution...")
		// 	// 这里可以添加你需要继续运行的代码
		// }
		// fmt.Println(time.Now())

		time.Sleep(5 * time.Second) // 每5秒检查一次
	}
}
