package adapters

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/Ellan98/ding-water-service/user/ports"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"sync"

	"github.com/Ellan98/ding-water-service/user/domain"
	"github.com/sirupsen/logrus"
)

type MemoryUserRepository struct {
	lock  *sync.RWMutex
	store []*domain.User
}

func NewMemoryUserRepository() *MemoryUserRepository {
	s := make([]*domain.User, 0)
	s = append(s, &domain.User{
		Model: "hello world ",
	})
	return &MemoryUserRepository{
		lock:  &sync.RWMutex{},
		store: s,
	}
}

// 考虑 这个方向 构造 deepSeek 请求
func (m MemoryUserRepository) Post(ctx context.Context, model, paramsKey string) (*domain.User, error) {
	val := ctx.Value(paramsKey).(ports.PostChatCompletionRequest)
	//req, _ := val.(ports.PostChatCompletionRequest)

	logrus.Infof(" passThrough %+v", val)
	for i, v := range m.store {
		logrus.Infof("m.store[%d] = %+v", i, v)
	}
	m.lock.RLock()
	defer m.lock.RUnlock()
	//for _, o := range m.store {
	//	if o.Model != "" {
	//		return o, nil
	//	}
	//}
	if reply, err := chatHandler(ctx, &domain.User{
		Model:           model,
		Prompt:          val.Prompt,
		SearchEnabled:   val.SearchEnabled,
		ThinkingEnabled: val.ThinkingEnabled,
	}); err != nil {
		return nil, domain.NotFound{Model: model}
	}

	return reply, nil

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

func chatHandler(c *gin.Context, params *domain.User) (reply *domain.User, err error) {
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
			{Role: "user", Content: params.Prompt},
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

		return nil, errors.New("request fiald")
	}
	defer resp.Body.Close()

	// 处理响应
	body, _ := io.ReadAll(resp.Body)
	var deepSeekResp DeepSeekResponse
	json.Unmarshal(body, &deepSeekResp)

	//response := map[string]string{
	//	"reply": deepSeekResp.Choices[0].Message.Content,
	//}

	//reply = response["reply"]
	params.Reply = deepSeekResp.Choices[0].Message.Content
	return params, nil
	//c.JSON(http.StatusInternalServerError, gin.H{
	//	"message": "error",
	//	"data":    response,
	//})
	// json.NewEncoder(w).Encode(response)
}
