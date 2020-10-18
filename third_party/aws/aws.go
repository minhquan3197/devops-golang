package aws

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load("../../internal/environments/.env")
}

// ConnectAwsService function connect AWS service
func ConnectAwsService() *session.Session {
	keyID := os.Getenv("AWS_ACCESS_KEY_ID")
	keyAccess := os.Getenv("AWS_SECRET_ACCESS_KEY")
	keyRegion := os.Getenv("AWS_REGION")
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
