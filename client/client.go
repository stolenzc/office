package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"
)

type Config struct {
	ServerAddress string `json:"server_address"`
}

func (c *Config) loadConfig(filePath string) error {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, &c)
	return err
}

var config Config

func isScreenLocked() (bool, error) {
	cmd := exec.Command("../.venv/bin/python", "./screen_lock.py")
	var (
		out    bytes.Buffer
		outErr bytes.Buffer
	)
	cmd.Stdout = &out
	cmd.Stderr = &outErr
	err := cmd.Run()
	if err != nil {
		return false, fmt.Errorf("%w: %v", err, outErr.String())
	}
	return strings.Contains(out.String(), "True"), nil

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

		locked, err := isScreenLocked()
		if err != nil {
			fmt.Println("Error checking screen lock status:", err)
			return
		}

		if !locked {
			updateStatus(config.ServerAddress)
			// 	fmt.Println("Screen is unlocked.", time.Now())
			// } else {
			// 	fmt.Println("Screen is locked.", time.Now())
		}

		time.Sleep(5 * time.Second) // 每5秒检查一次
	}
}
