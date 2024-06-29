package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/spf13/viper"
)

type RequestPayload struct {
	ConversationID string `json:"conversation_id"`
	BotID          string `json:"bot_id"`
	User           string `json:"user"`
	Query          string `json:"query"`
	Stream         bool   `json:"stream"`
}

type ResponseMessage struct {
	Role        string  `json:"role"`
	Type        string  `json:"type"`
	Content     string  `json:"content"`
	ContentType string  `json:"content_type"`
	ExtraInfo   *string `json:"extra_info"`
}

type ResponsePayload struct {
	Messages       []ResponseMessage `json:"messages"`
	ConversationID string            `json:"conversation_id"`
	Code           int               `json:"code"`
	Msg            string            `json:"msg"`
}

func Call(query string) error {
	// read config from configPath
	viper.SetConfigFile(configPath)
	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("failed to read config file: %v", err)
	}
	personalAccessToken := viper.GetString("token")
	botID := viper.GetString("botid")
	userID := viper.GetString("user")

	// personalAccessToken := flag.String("token", "pat_hqKdphU6w3tLDJaNDYvjbMcei81gUG2EsYhWcUzd9lOlPjtPMtd7FtY2EqgfVnGQ", "Personal Access Token")
	// botID := flag.String("bot_id", "7385882762747936785", "Bot ID")
	// query := flag.String("query", "", "Query text")
	// userID := flag.String("user", "29032201862555", "User ID")

	if personalAccessToken == "" || botID == "" || query == "" {
		fmt.Println("token, bot_id, and query are required parameters")
		os.Exit(1)
	}

	// 创建请求体
	requestPayload := RequestPayload{
		BotID:  botID,
		User:   userID,
		Query:  query,
		Stream: false,
	}
	jsonData, err := json.Marshal(requestPayload)
	if err != nil {
		fmt.Printf("Error marshalling request payload: %v\n", err)
		os.Exit(1)
	}

	// 创建HTTP请求
	url := "https://api.coze.com/open_api/v2/chat"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Printf("Error creating HTTP request: %v\n", err)
		os.Exit(1)
	}
	req.Header.Set("Authorization", "Bearer "+personalAccessToken)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Host", "api.coze.com")
	req.Header.Set("Connection", "keep-alive")

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error sending HTTP request: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	// 读取响应体
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v\n", err)
		os.Exit(1)
	}

	// 解析响应体
	var responsePayload ResponsePayload
	err = json.Unmarshal(body, &responsePayload)
	if err != nil {
		fmt.Printf("Error unmarshalling response payload: %v\n", err)
		os.Exit(1)
	}

	// 检查响应代码
	if responsePayload.Code != 0 {
		fmt.Printf("Error: %s\n", responsePayload.Msg)
		os.Exit(1)
	}

	// 输出内容
	if len(responsePayload.Messages) > 0 {
		fmt.Println(responsePayload.Messages[0].Content)
	} else {
		fmt.Println("No messages received.")
	}
	return nil
}
