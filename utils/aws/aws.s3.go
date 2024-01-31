package aws

import (
	"amikom-pedia-api/utils"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	awsSess "github.com/aws/aws-sdk-go/aws/session"
)

func NewSessionAWSS3(c utils.Config) (*awsSess.Session, error) {

	sess, err := awsSess.NewSession(&aws.Config{
		Region:      aws.String(c.AWSRegion),
		Credentials: credentials.NewStaticCredentials(c.AWSAccessKey, c.AWSSecretKey, ""),
	})

	if err != nil {
		return nil, err
	}

	return sess, nil
}
