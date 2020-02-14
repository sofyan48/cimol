package provider

import (
	entity "github.com/sofyan48/otp/src/entity/http/v1"
)

// Providers ...
type Providers struct{}

// ProvidersHandler ...
func ProvidersHandler() *Providers {
	return &Providers{}
}

// ProvidersInterface ...
type ProvidersInterface interface {
	OperatorChecker(msisdn string) (*entity.DataProvider, string)
	InterceptorMessages(data *entity.PostNotificationRequest) *entity.DynamoItem
}

func serializeNumber(msisdn string) string {
	if string(msisdn[0:2]) == "08" {
		return "62" + string(msisdn[1:12])
	} else if string(msisdn[0:3]) == "+62" {
		return "62" + string(msisdn[3:14])
	} else if string(msisdn[0:2]) == "8" {
		return "62" + string(msisdn[3:14])
	}
	return ""
}

// OperatorChecker ...
func (prv *Providers) OperatorChecker(msisdn string) (*entity.DataProvider, string) {
	msidmsisdnSerial := serializeNumber(msisdn)
	msisdnReformat := string(msidmsisdnSerial[0:5])
	operator := map[string]*entity.DataProvider{
		"62811": {Provider: "telkomsel", Name: "kartu-halo"},
		"62812": {Provider: "telkomsel", Name: "simpati"},
		"62813": {Provider: "telkomsel", Name: "simpati"},
		"62821": {Provider: "telkomsel", Name: "simpati"},
		"62822": {Provider: "telkomsel", Name: "simpati"},
		"62852": {Provider: "telkomsel", Name: "kartu-as"},
		"62853": {Provider: "telkomsel", Name: "kartu-as"},
		"62823": {Provider: "telkomsel", Name: "kartu-as"},
		"62851": {Provider: "telkomsel", Name: "kartu-as"},
		"62814": {Provider: "indosat", Name: "m2-broadband"},
		"62815": {Provider: "indosat", Name: "matrix-mentari"},
		"62816": {Provider: "indosat", Name: "matrix-mentari"},
		"62855": {Provider: "indosat", Name: "matrix"},
		"62856": {Provider: "indosat", Name: "im3"},
		"62857": {Provider: "indosat", Name: "im3"},
		"62858": {Provider: "indosat", Name: "mentari"},
		"62817": {Provider: "xl", Name: "xl"},
		"62818": {Provider: "xl", Name: "xl"},
		"62819": {Provider: "xl", Name: "xl"},
		"62859": {Provider: "xl", Name: "xl"},
		"62877": {Provider: "xl", Name: "xl"},
		"62878": {Provider: "xl", Name: "xl"},
		"62838": {Provider: "axis", Name: "axis"},
		"62831": {Provider: "axis", Name: "axis"},
		"62832": {Provider: "axis", Name: "axis"},
		"62833": {Provider: "axis", Name: "axis"},
		"62892": {Provider: "three", Name: "three"},
		"62893": {Provider: "three", Name: "three"},
		"62895": {Provider: "three", Name: "three"},
		"62896": {Provider: "three", Name: "three"},
		"62897": {Provider: "three", Name: "three"},
		"62898": {Provider: "three", Name: "three"},
		"62899": {Provider: "three", Name: "three"},
		"62881": {Provider: "smart", Name: "smartfren"},
		"62882": {Provider: "smart", Name: "smartfren"},
		"62883": {Provider: "smart", Name: "smartfren"},
		"62884": {Provider: "smart", Name: "smartfren"},
		"62885": {Provider: "smart", Name: "smartfren"},
		"62886": {Provider: "smart", Name: "smartfren"},
		"62887": {Provider: "smart", Name: "smartfren"},
		"62888": {Provider: "smart", Name: "smartfren"},
		"62889": {Provider: "smart", Name: "smartfren"},
	}
	return operator[msisdnReformat], msidmsisdnSerial
}
