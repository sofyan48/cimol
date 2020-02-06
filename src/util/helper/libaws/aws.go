package libaws

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/kinesis"
)

// Aws ...
type Aws struct {
	Config *aws.Config
}

func AwsHAndler() *Aws {
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

// GetDynamoDB get dynamodb service
// return *dynamodb.DynamoDB
func (aw *Aws) GetDynamoDB() *dynamodb.DynamoDB {
	dynamo := dynamodb.New(session.New(), aw.Config)
	return dynamo
}

// UpdateDynamo ...
func (aw *Aws) UpdateDynamo(field string) {

}

// GetKinesis ...
func (aw *Aws) GetKinesis() *kinesis.Kinesis {
	kinesis := kinesis.New(session.New(), aw.Config)
	return kinesis
}
