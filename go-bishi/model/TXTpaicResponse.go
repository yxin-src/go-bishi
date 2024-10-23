package model

import (
	"encoding/xml"
)

// 响应结构体
type TXTpaicResponse struct {
	XMLName xml.Name `xml:"TXTpaicResponse"`
	Head    Head     `xml:"head"`
	Body    Body     `xml:"body"`
}

type Head struct {
	TransType    string `xml:"transType"`
	TransRefId   string `xml:"transRefId"`
	TransExeDate string `xml:"transExeDate"`
	Status       string `xml:"status"`
	StatusCode   string `xml:"statusCode"`
	Message      string `xml:"message"`
}

type Body struct {
	Policy Policy `xml:"policy"`
	Plans  Plans  `xml:"Plans"`
}

type Policy struct {
	PlanCode           string `xml:"planCode"`
	ProductCode        string `xml:"productCode"`
	CurrencyRate       string `xml:"currencyRate"`
	TotalActualPremium string `xml:"totalActualPremium"`
	TotalNetPremium    string `xml:"totalNetPremium"`
	AddedValueTax      string `xml:"addedValueTax"`
	QuotationNo        string `xml:"quotationNo"`
}

type Plans struct {
	Plan []Plan `xml:"plan"`
}

type Plan struct {
	PlanNo     string     `xml:"planNo"`
	PlanName   string     `xml:"planName"`
	Insurances Insurances `xml:"insurances"`
}

type Insurances struct {
	Insurance []Insurance `xml:"insurance"`
}

type Insurance struct {
	PersonCode       string `xml:"personCode"`
	InsuranceName    string `xml:"insuranceName"`
	CertiNo          string `xml:"certiNo"`
	CertiType        string `xml:"certiType"`
	InsurancePremium string `xml:"insurancePremium"`
}
