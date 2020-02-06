package libaws

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// Aws ...
type Aws struct {
	Config *aws.Config
}

// GetAwsLib ...
func GetAwsLib() *Aws {
	return &Aws{}
}

// Sessions ...
// @cfg: *entity.AwsConfig
func (aw *Aws) Sessions() {
	creds := credentials.NewStaticCredentials(
		os.Getenv("AWS_ACCESS_KEY"),
		os.Getenv("AWS_ACCESS_SECRET"), "")
	creds.Get()
	cfgAws := aws.NewConfig().WithRegion(os.Getenv("AWS_ACCESS_AREA")).WithCredentials(creds)
	aw.Config = cfgAws
}

// DynamoDB get dynamodb service
// return *dynamodb.DynamoDB
func (aw *Aws) DynamoDB() *dynamodb.DynamoDB {
	dynamo := dynamodb.New(session.New(), aw.Config)
	return dynamo
}

// UpdateDynamo ...
func (aw *Aws) UpdateDynamo(field string) {

}
