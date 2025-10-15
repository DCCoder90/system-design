package main

import (
	"context"
	"encoding/csv"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sns"
)

var awsRegion = "us-east-1"
var snsTopicArn = "arn:aws:sns:us-east-1:123456789012:MyNotificationTopic"

type User struct {
	Name     string
	Type     string
	Endpoint string
}

func main() {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(awsRegion))
	if err != nil {
		log.Fatalf("FATAL: Unable to load AWS SDK config: %v", err)
	}

	snsClient := sns.NewFromConfig(cfg)

	users, err := readUsersFromCSV("users.csv")
	if err != nil {
		log.Fatalf("FATAL: Failed to read users from CSV: %v", err)
	}

	if len(users) == 0 {
		log.Println("INFO: No users found in CSV file. Exiting.")
		return
	}

	log.Printf("INFO: Found %d users in the CSV. Publishing notification.", len(users))

	message := "Hello World, this is a notification!"
	subject := "Hello"

	publishInput := &sns.PublishInput{
		Message:  &message,
		Subject:  &subject,
		TopicArn: &snsTopicArn,
	}

	result, err := snsClient.Publish(context.TODO(), publishInput)
	if err != nil {
		log.Fatalf("FATAL: Failed to publish message to SNS topic: %v", err)
	}

	log.Printf("Message ID: %s", *result.MessageId)
}

func readUsersFromCSV(filePath string) ([]User, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)

	if _, err := reader.Read(); err != nil {
		return nil, err
	}

	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var users []User
	for _, record := range records {
		user := User{
			Name:     record[0],
			Type:     record[1],
			Endpoint: record[2],
		}
		users = append(users, user)
	}

	return users, nil
}
