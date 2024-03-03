package controller

import (
	"bufio"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ncuhome/story-cook/model/dto"
	"github.com/ncuhome/story-cook/model/vo"
	"github.com/ncuhome/story-cook/pkg/myErrors"
	"github.com/ncuhome/story-cook/pkg/util"
	"io"
	"log"
	"net/http"

	"gopkg.in/antage/eventsource.v1"
)

func ExtendStoryHandler(ctx *gin.Context) {
	var req dto.ExtendStoryDto
	if err := ctx.ShouldBind(&req); err != nil {
		util.LogrusObj.Infoln(err)
		ctx.JSON(http.StatusBadRequest, vo.Error(err, myErrors.ErrorInvalidParams))
		return
	}

	resp := SendReqToTongYi(req.Title, req.Background, req.Keywords)
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	es := eventsource.New(nil, nil)
	w := ctx.Writer

	// 设置请求头，告知客户端这是一个SSE连接
	w.Header().Set("Content-Type", "text/event-stream")
	_, ok := w.(http.Flusher)

	if !ok {
		log.Panic("server not support") //不兼容
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
			if err != nil {
				return
			}
			eventData = nil
		}
	}

	// 检查扫描过程中的错误
	if err := scanner.Err(); err != nil {
		log.Println("Error reading response from external server:", err)
	}

	es.Close()
}
