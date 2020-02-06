package libaws

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	dynamoEntyty "github.com/sofyan48/rll-daemon-new/src/entity/http/v1"
)

// Aws ...
type Aws struct {
}

// AwsHAndler ..
func AwsHAndler() *Aws {
	return &Aws{}
}

// AwsInterface ...
type AwsInterface interface {
	InputDynamo(itemDynamo *dynamoEntyty.DynamoItem) (*dynamodb.PutItemOutput, error)
	UpdateDynamo(ID string, itemDynamo *dynamoEntyty.DynamoItem) (*dynamodb.UpdateItemOutput, error)
	GetDynamoData(ID string) (map[string]*dynamodb.AttributeValue, error)
	GetDynamoHistory(receiverAddress string) (*dynamodb.ScanOutput, error)
}

// Sessions ...
// @cfg: *entity.AwsConfig
func (aw *Aws) Sessions() *aws.Config {
	creds := credentials.NewStaticCredentials(
		os.Getenv("AWS_ACCESS_KEY"),
		os.Getenv("AWS_ACCESS_SECRET"), "")
	creds.Get()
	cfgAws := aws.NewConfig().WithRegion(os.Getenv("AWS_ACCESS_AREA")).WithCredentials(creds)
	return cfgAws
}
