package main

import (
	"context"
	"fmt"
	"time"

	"github.com/sunzhenkai/kfpanda-go-sdk/pkg/client"
	"github.com/sunzhenkai/kung-fu-panda-protocols/go/service/kfpanda"
	"google.golang.org/protobuf/encoding/protojson"
)

func sendExampleMessage() {
	req := kfpanda.RecordRequest{
		Service: "kfpanda-go-sdk",
	}
	cfg := &client.ClientConfig{
		ServerUri: "127.0.0.1:9820",
		Timeout:   500 * time.Millisecond,
	}
	client, err := client.NewKfPandaClient(cfg)
	if err != nil {
		fmt.Println("create client faield", err)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*500)
	defer cancel()

	resp, err := client.Record(ctx, &req)
	if err != nil {
		fmt.Println("query failed", err)
		return
	}
	bts, err := protojson.Marshal(resp)
	if err != nil {
		fmt.Println("marshal response failed", err)
		return
	}
	fmt.Println(string(bts))
}

func main() {
	fmt.Println("run kfpanda")
	sendExampleMessage()
}
