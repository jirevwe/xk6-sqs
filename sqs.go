package sqs

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/mitchellh/mapstructure"
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

func (s *Sqs) Send(sqsClient *sqs.Client, params interface{}) {
	var sqsMessageInput sqs.SendMessageInput
	_ = mapstructure.Decode(params, &sqsMessageInput)
	_, err := sqsClient.SendMessage(context.TODO(), &sqsMessageInput)
	if err != nil {
		log.Fatal(err)
	}
}
