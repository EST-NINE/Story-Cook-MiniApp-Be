package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/ncuhome/story-cook/config"
	"io"
	"net/http"
)

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

func createStory() {
	apiKey := config.ApiKey
	charaSetting := ""
	prompt := ""
	// 构建请求的数据
	requestBody := map[string]interface{}{
		"model": "qwen-max",
		"input": map[string]interface{}{
			"messages": []map[string]string{
				{"role": "system", "content": charaSetting},
				{"role": "user", "content": prompt},
			},
		},
	}
	// 将请求的数据转换为JSON格式
	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}
	// 创建新的HTTP请求
	req, err := http.NewRequest("POST", "https://dashscope.aliyuncs.com/api/v1/services/aigc/text-generation/generation", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	// 设置请求头
	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")
	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()
	// 读取响应
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}
	// 打印响应状态码和响应体
	fmt.Println("Response status:", resp.Status)
	fmt.Println("Response body:", string(body))
}
