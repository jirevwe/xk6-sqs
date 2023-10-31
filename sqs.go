package sqs

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"go.k6.io/k6/js/modules"
	"log"
	"os"
)

func init() {
	modules.Register("k6/x/sqs", new(Sqs))
}

type Sqs struct{}

func (*Sqs) NewClient() *sqs.Client {
	region := os.Getenv("AWS_REGION")
	return sqs.NewFromConfig(aws.Config{Region: region})
}

func (*Sqs) WriteEvent(ctx context.Context, client *sqs.Client, QueueUrl string, body interface{}) {
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		log.Fatal(err)
	}

	payload := string(bodyBytes)
	params := &sqs.SendMessageInput{
		MessageBody: &payload,
		QueueUrl:    &QueueUrl,
	}

	_, err = client.SendMessage(ctx, params)
	if err != nil {
		log.Fatal(err)
	}
}
