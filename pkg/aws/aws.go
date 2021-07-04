package aws

import (
	"project-golang/configs"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

// ConnectAwsService function connect AWS service
func ConnectAwsService() *session.Session {
	cfg := configs.Load()
	keyID := cfg.AWS.AccessKeyID
	keyAccess := cfg.AWS.AccessKeyID
	keyRegion := cfg.AWS.Region
	sess, err := session.NewSession(
		&aws.Config{
			Region: aws.String(keyRegion),
			Credentials: credentials.NewStaticCredentials(
				keyID,
				keyAccess,
				"",
			),
		})
	if err != nil {
		panic(err)
	}
	return sess
}
