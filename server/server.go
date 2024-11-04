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
	Port     int    `json:"port"`
	UserName string `json:"user_name"`
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
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	// 准备状态数据
	htmlData := HTMLData{}
	htmlData.UserName = config.UserName
	if time.Since(status.LastTime) < 20*time.Second {
		htmlData.Status = "online"
	} else {
		htmlData.Status = "offline"
		htmlData.LastOnlineTime = status.LastTime.Format("2006-01-02 15:04:05")
	}

	// 解析模板文件
	temp, err := template.ParseFiles("./index.html")
	if err != nil {
		http.Error(w, "无法读取页面", http.StatusInternalServerError)
		return
	}

	// 执行模板并将数据传入
	w.Header().Set("Content-Type", "text/html")
	if err := temp.Execute(w, htmlData); err != nil {
		http.Error(w, "渲染页面失败", http.StatusInternalServerError)
	}
}

// 处理更新在线状态请求
func updateStatusHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	status.LastTime = time.Now()
}

// 查询状态
func getStatusHandler(w http.ResponseWriter, r *http.Request) {
	response := HTMLData{Status: "offline", LastOnlineTime: status.LastTime.Format("2006-01-02 15:04:05")}
	if time.Since(status.LastTime) < 30*time.Second {
		response.Status = "online"
	}
	w.Header().Set("Content-Type", "application/json")
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
