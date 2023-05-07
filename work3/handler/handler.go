package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"simpleserver/work3/cache"
)

var (
	baseUrl  = `https://rest.coinapi.io/v1/exchangerate/`
	toDollar = `/USD`
)

func GetMessage(c *gin.Context) {
	fmt.Println("进入handler")
	client := cache.DefaultClientPool.Get().(*http.Client)
	defer cache.DefaultClientPool.Put(client)
	url := baseUrl + c.Param("currency") + toDollar
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("X-CoinAPI-Key", "B89898B1-1DFC-4D44-AB49-4D56856A3627")
	if err != nil {
		panic(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	content, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if err != nil {
		panic(err)
	}
	// 传递下去
	c.Set("content", content)
	c.Data(http.StatusOK, `application/json`, content)
}
