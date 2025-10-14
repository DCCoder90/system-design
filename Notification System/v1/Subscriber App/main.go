package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sns"
)

const awsRegion = "us-east-1"
const snsTopicArn = "arn:aws:sns:us-east-1:"

func main() {
	protocol := os.Args[1]
	endpoint := os.Args[2]

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(awsRegion))
	if err != nil {
		log.Fatalf("Unable to load SDK config, %v", err)
	}

	snsClient := sns.NewFromConfig(cfg)

	log.Printf("Subscribing endpoint %s with protocol %s to topic %s\n", endpoint, protocol, snsTopicArn)

	subscribeInput := &sns.SubscribeInput{
		TopicArn: &snsTopicArn,
		Protocol: &protocol,
		Endpoint: &endpoint,
	}

	result, err := snsClient.Subscribe(context.TODO(), subscribeInput)
	if err != nil {
		log.Fatalf("Failed to subscribe user: %v", err)
	}

	log.Println("Successful subscription.")
	log.Printf("Subscription ARN: %s", *result.SubscriptionArn)
}