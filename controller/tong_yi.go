package controller

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ncuhome/story-cook/config"
	"github.com/ncuhome/story-cook/pkg/util"
	"io"
	"net/http"
)

func SendReqToTongYi(charaSetting string, prompt string) *http.Response {
	apiKey := config.ApiKey
	// 构建请求的数据
	requestBody := map[string]interface{}{
		"model": "qwen-turbo",
		"input": map[string]interface{}{
			"messages": []map[string]string{
				{"role": "system", "content": charaSetting},
				{"role": "user", "content": prompt},
			},
		},
		"parameters": map[string]interface{}{
			"incremental_output": "true",
			"enable_search":      "true",
			"temperature":        "1.90",
		},
	}

	// 将请求的数据转换为JSON格式
	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		util.LogrusObj.Infoln(err)
	}

	// 创建新的HTTP请求
	req, err := http.NewRequest("POST", "https://dashscope.aliyuncs.com/api/v1/services/aigc/text-genera"+
		"tion/generation", bytes.NewBuffer(jsonData))
	if err != nil {
		util.LogrusObj.Infoln(err)
	}

	// 设置请求头
	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-DashScope-SSE", "enable")
	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		util.LogrusObj.Infoln(err)
	}
	return resp
}

func ForWardSSE(ctx *gin.Context, charaSetting string, prompt string) error {
	resp := SendReqToTongYi(charaSetting, prompt)
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	w := ctx.Writer
	// 设置请求头，告知客户端这是一个SSE连接
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	flusher, ok := w.(http.Flusher)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Streaming unsupported!"})
	}

	// 读取SSE响应
	scanner := bufio.NewScanner(resp.Body)
	var eventData []string
	for scanner.Scan() {
		line := scanner.Text()
		// 将当前行添加到事件数据
		eventData = append(eventData, line)
		// 当读取到空行时，处理当前事件
		if line == "" {
			// 发送当前事件到前端
			_, err := fmt.Fprintf(w, eventData[0]+"\n"+eventData[1]+"\n"+eventData[2]+"\n"+eventData[3]+"\n\n")
			if err != nil && eventData[1] != ":HTTP_STATUS/200" {
				return err
			}
			eventData = nil
			flusher.Flush()
		}
	}
	return nil
}
