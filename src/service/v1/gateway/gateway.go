package gateway

import (
	"fmt"
	"log"
	"strings"
	"sync"

	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	entity "github.com/sofyan48/otp/src/entity/http/v1"
	"github.com/sofyan48/otp/src/util/helper/libaws"
	"github.com/sofyan48/otp/src/util/helper/libsendgrid"
	"github.com/sofyan48/otp/src/util/helper/provider"
)

// Gateway ...
type Gateway struct {
	AwsLib    libaws.AwsInterface
	Providers provider.ProvidersInterface
	Sendgrid  libsendgrid.LibsendgridInterface
}

// GatewayHandler Handler
func GatewayHandler() *Gateway {
	return &Gateway{
		AwsLib:    libaws.AwsHAndler(),
		Providers: provider.ProvidersHandler(),
		Sendgrid:  libsendgrid.LibSendgridHandler(),
	}
}

//GatewayInterface declare All Method
type GatewayInterface interface {
	PostNotification(data *entity.PostNotificationRequest, wg *sync.WaitGroup)
	PostNotificationEmail(data *entity.PostNotificationRequestEmail, wg *sync.WaitGroup)
	GetHistory(msisdn string) ([]entity.DynamoItemResponse, error)
	GetByID(ID string) (*entity.DynamoItemResponse, error)
}

// PostNotification ...
func (gateway *Gateway) PostNotification(data *entity.PostNotificationRequest, wg *sync.WaitGroup) {
	itemDynamo := gateway.Providers.InterceptorMessages(data)

	wg.Add(2)
	go gateway.AwsLib.SendStart(data.UUID, itemDynamo, "interceptors", wg)

	wg.Add(2)
	go gateway.AwsLib.InputDynamo(itemDynamo, wg)

}

// PostNotificationEmail ...
func (gateway *Gateway) PostNotificationEmail(data *entity.PostNotificationRequestEmail, wg *sync.WaitGroup) {
	templateData, err := gateway.Sendgrid.GetTemplateID(data.Payload.TemplateID)
	if err != nil {
		log.Println("Error: ", err)
	}
	htmlContent := templateData.Versions[0].HTMLContent
	for key, word := range data.Payload.Data {
		htmlContent = strings.Replace(htmlContent, key, word, -1)
	}
	fmt.Println(htmlContent)
}

// GetHistory ...
func (gateway *Gateway) GetHistory(msisdn string) ([]entity.DynamoItemResponse, error) {
	data, err := gateway.AwsLib.GetDynamoHistory(msisdn)
	dynamoItem := []entity.DynamoItemResponse{}
	err = dynamodbattribute.UnmarshalListOfMaps(data.Items, &dynamoItem)
	return dynamoItem, err
}

// GetByID ...
func (gateway *Gateway) GetByID(ID string) (*entity.DynamoItemResponse, error) {
	data, err := gateway.AwsLib.GetDynamoData(ID)
	dynamoItem := &entity.DynamoItemResponse{}
	err = dynamodbattribute.UnmarshalMap(data.Item, &dynamoItem)
	if err != nil {
		return nil, err
	}
	return dynamoItem, err
}
