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
	ExpectedSSID string `json:"expected_ssid"` // 添加预期的WiFi名称配置
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

// 新增函数：检测当前连接的WiFi
func getCurrentSSID() (string, error) {
	// 这个命令在Linux/macOS上获取当前WiFi SSID
	cmd := exec.Command("iwgetid", "-r")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("failed to get SSID: %v", err)
	}
	return strings.TrimSpace(out.String()), nil
}

func getCurrentSSIDOnMac() (string, error) {
	cmd := exec.Command("./wifi")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("failed to get SSID: %v", err)
	}
	return strings.TrimSpace(out.String()), nil
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
	err := config.loadConfig("./config.json")
	if err != nil {
		fmt.Println("Error loading config:", err)
		return
	}

	for {
		// 检查屏幕锁定状态
		locked, err := isScreenLocked()
		if err != nil {
			fmt.Println("Error checking screen lock status:", err)
			time.Sleep(5 * time.Second)
			continue
		}

		// 检查WiFi连接
		currentSSID, err := getCurrentSSIDOnMac()
		if err != nil {
			fmt.Println("Error checking WiFi connection:", err)
			time.Sleep(5 * time.Second)
			continue
		}

		// 只有当屏幕未锁定且连接了正确的WiFi时才更新状态
		if !locked && (config.ExpectedSSID == "" || currentSSID == config.ExpectedSSID) {
			updateStatus(config.ServerAddress)
			fmt.Printf("Screen is unlocked and connected to %s. %v\n", currentSSID, time.Now())
		} else {
			fmt.Printf("Conditions not met (locked: %v, SSID: %s). %v\n", locked, currentSSID, time.Now())
		}

		time.Sleep(5 * time.Second)
	}
}
