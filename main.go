package main

import (
	"fmt"
	"net/http"
	"strings"
)

func main() {
	resp, err := http.Post("https://open.feishu.cn/open-apis/bot/v2/hook/8732a64a-a2b0-4632-b1f5-8a70dea49251",
		"application/json",
		strings.NewReader("{\"msg_type\":\"text\",\"content\":{\"text\":\"构建成功\"}}"))
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(resp.Status)
}
