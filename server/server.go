package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"text/template"
	"time"
)

type Config struct {
	Port       int    `json:"port"`
	UserName   string `json:"user_name"`
	DingTalkID string `json:"dingtalk_id"`
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

type Status struct {
	UserName  string
	LastTime  time.Time
	IsMeeting bool
}

var status Status

type HTMLData struct {
	UserName       string `json:"user_name"`
	Status         string `json:"status"`
	LastOnlineTime string `json:"last_online_time"`
	DingTalkID     string `json:"dingtalk_id"`
}

func getStatus() HTMLData {
	htmlData := HTMLData{Status: "offline", LastOnlineTime: status.LastTime.Format("2006-01-02 15:04:05")}
	htmlData.UserName = config.UserName
	htmlData.DingTalkID = config.DingTalkID
	if time.Since(status.LastTime) < 10*time.Second {
		htmlData.Status = "online"
	}
	return htmlData
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	// 准备状态数据
	htmlData := getStatus()

	// 解析模板文件
	temp, err := template.ParseFiles("./index.html")
	if err != nil {
		http.Error(w, "无法读取页面", http.StatusInternalServerError)
		return
	}

	// 设置Content-Type并执行模板
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := temp.Execute(w, htmlData); err != nil {
		// 这里不能再调用http.Error，因为已经开始写入响应了
		fmt.Printf("模板渲染失败: %v\n", err)
	}
}

type UpdateBody struct {
	LastTime time.Time `json:"last_time"`
}

// 处理更新在线状态请求
func updateStatusHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	// 从请求中获取body数据
	var updateData *UpdateBody
	if err := json.NewDecoder(r.Body).Decode(&updateData); err == nil {
		status.LastTime = updateData.LastTime
	}

	// 检查零值
	if status.LastTime.IsZero() {
		status.LastTime = time.Now()
	}
}

// 查询状态
func getStatusHandler(w http.ResponseWriter, r *http.Request) {
	response := getStatus()
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(response)
}

func main() {
	// 注册处理函数
	err := config.loadConfig("./config.json")
	if err != nil {
		panic(err)
	}
	status = Status{LastTime: time.Now(), IsMeeting: false, UserName: config.UserName}

	http.HandleFunc("/office/index", indexHandler)
	http.HandleFunc("/office/updateStatus", updateStatusHandler)
	http.HandleFunc("/office/getStatus", getStatusHandler)

	// 启动HTTP服务器
	addr := fmt.Sprint(":", config.Port)
	fmt.Println("Starting server on ", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
