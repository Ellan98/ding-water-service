package main

// go-backend/main.go

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/Ellan98/ding-water-service/user/app"
	"github.com/Ellan98/ding-water-service/user/app/command/query"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type HTTPServer struct {
	app app.Application
}

func (h HTTPServer) GetDeepSeekAnswer(c *gin.Context, problem string) {
	//TODO
	logrus.Info("To do somethings", problem)
	o, err := h.app.Queries.GetDeepSeekAnswer.Handle(c, query.GetDeepSeekAnswer{Problem: problem})
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    o,
	})
}
func (h HTTPServer) PostChatCompletion(c *gin.Context, problem string) {
	o, err := h.app.Queries.PostChatCompletion.Handle(c, query.PostChatCompletion{Problem: problem})
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err,
		})
	}
	//TODO Something
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

type Request struct {
	Message string `json:"message"`
}

type DeepSeekRequest struct {
	Model    string `json:"model"`
	Messages []struct {
		Role    string `json:"role"`
		Content string `json:"content"`
	} `json:"messages"`
}

type DeepSeekResponse struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

func chatHandler(c *gin.Context) {
	// 解析前端请求
	// var req Request
	// if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	return
	// }

	// 构造DeepSeek请求
	deepSeekReq := DeepSeekRequest{
		Model: "deepseek-chat",
		Messages: []struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		}{
			{Role: "user", Content: "api test"},
		},
	}

	reqBody, _ := json.Marshal(deepSeekReq)

	// 创建HTTP客户端
	client := &http.Client{}
	apiReq, _ := http.NewRequest(
		"POST",
		"https://api.deepseek.com/v1/chat/completions",
		bytes.NewBuffer(reqBody),
	)

	apiReq.Header.Set("Authorization", "Bearer "+os.Getenv("DING_WATER_SERVICE_DEEPSEEK_KEY"))
	apiReq.Header.Set("Content-Type", "application/json")

	// 发送请求
	resp, err := client.Do(apiReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "error",
		})
		return
	}
	defer resp.Body.Close()

	// 处理响应
	body, _ := io.ReadAll(resp.Body)
	var deepSeekResp DeepSeekResponse
	json.Unmarshal(body, &deepSeekResp)

	response := map[string]string{
		"reply": deepSeekResp.Choices[0].Message.Content,
	}
	c.JSON(http.StatusInternalServerError, gin.H{
		"message": "error",
		"data":    response,
	})
	// json.NewEncoder(w).Encode(response)
}
