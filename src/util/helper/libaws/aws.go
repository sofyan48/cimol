package libaws

import (
	"os"
	"sync"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/kinesis"
	dynamoEntyty "github.com/sofyan48/cimol/src/entity/http/v1"
	entity "github.com/sofyan48/cimol/src/entity/http/v1"
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
	// dynamo
	InputDynamo(itemDynamo *dynamoEntyty.DynamoItem, wg *sync.WaitGroup) (*dynamodb.PutItemOutput, error)
	InputDynamoEmail(itemDynamo *dynamoEntyty.DynamoItemEmail, wg *sync.WaitGroup) (*dynamodb.PutItemOutput, error)
	UpdateDynamo(ID, status, data string, history *dynamoEntyty.HistoryItem) (*dynamodb.UpdateItemOutput, error)
	GetDynamoData(ID string) (*dynamodb.GetItemOutput, error)
	GetDynamoHistory(receiverAddress string) (*dynamodb.ScanOutput, error)
	CallbackSendUpdate(ID, status, oldHistory string, newHistory string) (*dynamodb.UpdateItemOutput, error)
	// kinesis
	WaitUntil(*kinesis.DescribeStreamInput) error
	WaitUntilNotExist(data *kinesis.DescribeStreamInput) error
	SendStart(ID string, itemDynamo *dynamoEntyty.DynamoItem, stack string, wg *sync.WaitGroup)
	SendMail(ID string, itemDynamo *entity.DynamoItemEmail, stack string, wg *sync.WaitGroup)
	Send(data []byte, stack string, wg *sync.WaitGroup) (*kinesis.PutRecordOutput, error)
	GetShardIterator() (string, error)
	Consumer(data *kinesis.GetRecordsInput) (*kinesis.GetRecordsOutput, error)
	GetDescribeInput() *kinesis.DescribeStreamInput
	Describe(data *kinesis.DescribeStreamInput) (*kinesis.DescribeStreamOutput, error)
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

// SessionsKinesis ...
// @cfg: *entity.AwsConfig
func (aw *Aws) SessionsKinesis() *aws.Config {
	creds := credentials.NewStaticCredentials(
		os.Getenv("AWS_ACCESS_KEY_KINESIS"),
		os.Getenv("AWS_ACCESS_SECRET_KINESIS"), "")
	creds.Get()
	cfgAws := aws.NewConfig().WithRegion(os.Getenv("AWS_ACCESS_AREA")).WithCredentials(creds)
	return cfgAws
}
