package main

import (
	"fmt"
	requestorder "github.com/xiaomingping/huobi_futures_Golang/src/request/restful/order"
	responseorder "github.com/xiaomingping/huobi_futures_Golang/src/response/restful/order"
	"github.com/xiaomingping/huobi_futures_Golang/src/restful"
	"log"
)

func main() {
	AccessKey := ""
	SecretKey := ""
	MarketClient := restful.OrderClient{}
	data := make(chan responseorder.PlaceOrderResponse)
	go MarketClient.Init(AccessKey,SecretKey,"").IsolatedPlaceOrderAsync(data,requestorder.PlaceOrderRequest{})
	x, ok := <-data
	if !ok || x.Status != "ok" {
		log.Fatal(x.ErrorMessage)
		return
	} else {
		fmt.Println(x.Data)
	}
}
