package main

import (
	"context"
	"encoding/json"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"pluginserver/pkg/datasource"
)


func main() {
	//var opts []grpc.DialOption
	conn, err := grpc.Dial(fmt.Sprintf("localhost:9000"), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()

	client := datasource.NewDatasourceClient(conn)

	jsonData := json.RawMessage(`{"url": "http://35.223.16.232:9090/"}`)
	configurationData, _ := json.Marshal(jsonData)
	metricRequest := datasource.MetricsRequest{Configuration: configurationData}

	metricList, err := client.GetMetrics(context.Background(), &metricRequest)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}

	log.Println(metricList)
}