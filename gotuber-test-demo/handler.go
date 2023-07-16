package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"io"
	"log"
	"net/http"
	"strings"
	"work6/gotuber-test-demo/config"
	nlp "work6/gotuber-test-demo/kitex_gen/nlp"
	"work6/gotuber-test-demo/model"
)

// NlpImpl implements the last service interface defined in the IDL.
type NlpImpl struct{}

// Handel implements the NlpImpl interface.
func (s *NlpImpl) Handel(ctx context.Context, req *nlp.Request) (resp *nlp.Response, err error) {
	if req.Choice == 1 {
		response, err := openai(req.RequestMessage)
		if err != nil {
			err = kerrors.NewBizStatusError(404, err.Error())
			return nil, err
		}
		return &nlp.Response{ResponseMessage: response}, nil
	} else {
		//TODO:扩展
	}
	return resp, nil
}

const openAIChatUrl = "https://api.openai.com/v1/chat/completions"

func openai(requestMessage string) (string, error) {
	var messages []model.RequestMessages
	message := model.RequestMessages{
		Role:    "user",
		Name:    "1",
		Content: requestMessage,
	}
	messages = append(messages, message)
	request := model.Request{
		Model:            config.GPTCfg.General.Model,
		Messages:         messages,
		MaxTokens:        config.GPTCfg.General.MaxTokens,
		Temperature:      config.GPTCfg.General.Temperature,
		TopP:             config.GPTCfg.General.TopP,
		Stop:             config.GPTCfg.General.Stop,
		PresencePenalty:  config.GPTCfg.General.PresencePenalty,
		FrequencyPenalty: config.GPTCfg.General.FrequencyPenalty,
		User:             "1",
	}
	requestDateByByte, err := json.Marshal(request)
	if err != nil {
		log.Println("1", err)
		return "", err
	}
	req, err := http.NewRequest("POST", openAIChatUrl, bytes.NewBuffer(requestDateByByte))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+config.GPTCfg.OpenAi.ApiKey)
	var client http.Client
	resp, err := client.Do(req)
	if err != nil {
		log.Println("2", err)
		return "", err
	}
	defer resp.Body.Close()
	if resp == nil {
		return "", errors.New("response is nil")
	}
	body, _ := io.ReadAll(resp.Body)
	var openaiReceive model.OpenAiRcv
	err = json.Unmarshal(body, &openaiReceive)
	if err != nil {
		log.Println("3", err)
		return "", err
	}
	if len(openaiReceive.Choices) == 0 {
		return "", errors.New("OpenAI调用失败，返回内容：" + string(body))
	}
	openaiReceive.Choices[0].Message.Content = strings.Replace(openaiReceive.Choices[0].Message.Content, "\n\n", "", 1)
	return openaiReceive.Choices[0].Message.Content, nil
}
