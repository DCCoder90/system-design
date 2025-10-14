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
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run subscribe.go <protocol> <endpoint>")
		fmt.Println("Example (email): go run subscribe.go email user@example.com")
		fmt.Println("Example (sms):   go run subscribe.go sms +15551234567")
		return
	}

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