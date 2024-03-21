package external

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	"os"
)

func GetSnsClient() *sns.SNS {
	if os.Getenv("IS_LOCAL") == "true" {
		return nil
	}

	region := "us-east-1"
	awsConfig := &aws.Config{
		Region: aws.String(region),
	}

	snsClient := sns.New(session.Must(session.NewSession(awsConfig)))
	fmt.Printf("SNS client connected: %v\n", snsClient.ClientInfo)

	return snsClient
}
