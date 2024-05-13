package utils_test

import (
	"encoding/json"
	"gin-template/common/utils"
	"log"
	"strconv"
	"testing"
	"time"
)

func TestGetHttpRequest(t *testing.T) {
	baseURL := "https://api.oioweb.cn/api/qq/info"
	qq := 2848660614
	params := map[string]string{
		"qq": strconv.FormatInt(int64(qq), 10),
	}
	res, err := utils.GetHttpRequest(baseURL, params, nil, 15*time.Second)
	if err != nil {
		log.Printf("GET Error: %s\n URL: %s\n Params: %#v\n", err.Error(), baseURL, params)
		return
	}
	log.Printf("GET Success Response: %s", res)
}

func TestPostHttpRequest(t *testing.T) {
	url := ""
	type Message struct {
		Role    string `json:"role"`
		Type    string `json:"type"`
		Content string `json:"content"`
	}
	type ChatTestParam struct {
		History  []interface{}          `json:"history"`
		Message  Message                `json:"message"`
		SkillId  int                    `json:"skill_id"`
		Variable map[string]interface{} `json:"variable"`
	}
	type UserData struct {
		UserId    int `json:"userId"`
		CompanyId int `json:"companyId"`
	}

	chatTestParam := ChatTestParam{
		Message: Message{
			Role:    "user",
			Type:    "text",
			Content: "你好!",
		},
		SkillId:  0,
		Variable: make(map[string]interface{}), // 空对象
		History:  []interface{}{},
	}
	headers := map[string]string{
		"Content-Type": "application/json",
		"share-token":  "975c2503-fcbe-4b3c-8cf7-8a48feec4e39",
	}
	userData, err := json.Marshal(UserData{
		CompanyId: 84,
		UserId:    0,
	})
	if err != nil {
		log.Println("json Marshal Error:", err)
		return
	}
	headers["x-user-data"] = string(userData)
	res, err := utils.PostHttpRequest(url, &chatTestParam, headers, 15*time.Second)
	if err != nil {
		log.Println("POST Error:", err)
		return
	}
	log.Printf("POST Success Response: %s", res)
}
