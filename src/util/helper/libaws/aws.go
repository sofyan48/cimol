package libaws

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
)

// Aws ...
type Aws struct {
	Config *aws.Config
}

// AwsHAndler ..
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
