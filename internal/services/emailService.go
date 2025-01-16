package services

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/sesv2"
	"github.com/aws/aws-sdk-go-v2/service/sesv2/types"
)

type EmailService struct {
	client  *sesv2.Client
	sender  string
	context context.Context
}

func NewEmailService() (*EmailService, error) {
	region := "ap-south-1"
	awsAccessKey := os.Getenv("AWS_ACCESS_KEY")
	awsSecretKey := os.Getenv("AWS_SECRET_KEY")
	emailSender := os.Getenv("EMAIL_SENDER")

	if awsAccessKey == "" || awsSecretKey == "" || emailSender == "" {
		return nil, fmt.Errorf("AWS credentials not found")
	}

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region), config.WithCredentialsProvider(credentials.StaticCredentialsProvider{
		Value: aws.Credentials{
			AccessKeyID:     awsAccessKey,
			SecretAccessKey: awsSecretKey,
		},
	}))
	if err != nil {
		return nil, fmt.Errorf("unable to load SDK config, %v", err)
	}
	client := sesv2.NewFromConfig(cfg)

	return &EmailService{
		client:  client,
		sender:  emailSender,
		context: context.TODO(),
	}, nil

}

func (es *EmailService) SendEmail(toEmail, subject, message string) error {
	input := &sesv2.SendEmailInput{
		Destination: &types.Destination{
			ToAddresses: []string{toEmail},
		},
		Content: &types.EmailContent{
			Simple: &types.Message{
				Body: &types.Body{
					Text: &types.Content{
						Charset: aws.String("UTF-8"),
						Data:    aws.String("Email Body Content Here"),
					},
				},
				Subject: &types.Content{
					Charset: aws.String("UTF-8"),
					Data:    aws.String(subject),
				},
			},
		},
		FromEmailAddress: aws.String(es.sender),
	}
	_, err := es.client.SendEmail(context.TODO(), input)
	if err != nil {
		fmt.Println("error sending email:", err)
		return err
	}
	return nil
}
