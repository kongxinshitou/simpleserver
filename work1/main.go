package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Message struct {
	Time          string  `json:"time"`
	AssertIdBase  string  `json:"assert_id_base"`
	AssertIdQuote string  `json:"assert_id_quote"`
	Rate          float64 `json:"rate"`
}

func main() {
	req, err := http.NewRequest("GET", "https://rest.coinapi.io/v1/exchangerate/BTC/USD", nil)
	req.Header.Add("X-CoinAPI-Key", "B89898B1-1DFC-4D44-AB49-4D56856A3627")
	if err != nil {
		panic(err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	content, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Println(string(content))
	message := &Message{}
	// 反序列化
	err = json.Unmarshal(content, message)
	if err != nil {
		panic(err)
	}
	// 打印结构体
	fmt.Printf("%v\n", message)
}
