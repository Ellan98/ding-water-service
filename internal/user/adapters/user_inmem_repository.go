package adapters

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Ellan98/ding-water-service/user/ports"
	"io"
	"net/http"
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
	reply, err := chatHandler(&domain.User{
		Model:           model,
		Prompt:          val.Prompt,
		SearchEnabled:   val.SearchEnabled,
		ThinkingEnabled: val.ThinkingEnabled,
	})
	if err != nil {
		return nil, domain.NotFound{Model: model}
	}

	return reply, nil

}

// 定义 Message 结构体
type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// 定义 DeepSeekRequest 结构体
type DeepSeekRequest struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"` // 使用 Message 类型切片
}

// 定义 DeepSeekResponse 结构体
type DeepSeekResponse struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

// 1. 这里函数的参数 类型 是指针类型，所以去要&获取地址
// 2. 在这个函数体内 go 会自动根据地址进行解引用 例如 params.Reply = "respones something"
// 3. 如果返回指针类型的参数 ，不需要再使用&进行取地址了
func chatHandler(params *domain.User) (*domain.User, error) {
	if params == nil {
		return nil, errors.New("params is nil")
	}
	// 构造请求体
	deepSeekReq := DeepSeekRequest{
		Model:    "deepseek-chat",
		Messages: []Message{{Role: "user", Content: params.Prompt}},
	}

	// 编码请求体
	reqBody, err := json.Marshal(deepSeekReq)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request body: %w", err)
	}

	// 创建 HTTP 请求
	apiReq, err := http.NewRequest(
		"POST",
		"https://api.deepseek.com/v1/chat/completions",
		bytes.NewBuffer(reqBody),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request: %w", err)
	}

	//apiReq.Header.Set("Authorization", "Bearer "+os.Getenv("DING_WATER_SERVICE_DEEPSEEK_KEY"))
	apiReq.Header.Set("Authorization", "Bearer ")

	apiReq.Header.Set("Content-Type", "application/json")

	// 执行 HTTP 请求
	client := &http.Client{}
	resp, err := client.Do(apiReq)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	// 读取响应体
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}
	logrus.Debug("------------2222--------------------", string(body))

	// 解析响应
	var deepSeekResp DeepSeekResponse
	if err := json.Unmarshal(body, &deepSeekResp); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}
	logrus.Debugf("out bind struct response %+v", deepSeekResp)

	if len(deepSeekResp.Choices) == 0 {
		return nil, errors.New("no choices returned from deepseek")
	}

	logrus.Debugf("------------1111-------------------- %+v", deepSeekResp)
	// 填充 Reply 字段
	params.Reply = deepSeekResp.Choices[0].Message.Content
	logrus.Debugf("deepSeek Response: %s", params.Reply)
	return params, nil

}
