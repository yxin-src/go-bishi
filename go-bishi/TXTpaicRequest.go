package main

import "encoding/xml"

// 请求结构体
type TXTpaicRequest struct {
	XMLName xml.Name `xml:"TXTpaicRequest"`
	Head    Head     `xml:"head"`
	Body    Body     `xml:"body"`
}

type Head struct {
	TransRefId   string `xml:"transRefId"`
	TransType    string `xml:"transType"`
	TransExeDate string `xml:"transExeDate"`
}

type Body struct {
	Policy    Policy    `xml:"policy"`
	Applicant Applicant `xml:"applicant"`
	Plans     Plans     `xml:"plans"`
}

type Policy struct {
	PlanCode         string `xml:"planCode"`
	ProductCode      string `xml:"productCode"`
	PaymentMode      string `xml:"paymentMode"`
	Count            int    `xml:"count"`
	TransApplyDate   string `xml:"transApplyDate"`
	TransBeginDate   string `xml:"transBeginDate"`
	TransEndDate     string `xml:"transEndDate"`
	AgentCode        string `xml:"agentCode"`
	CurrencyCode     string `xml:"currencyCode"`
	CommissionFactor int    `xml:"commissionFactor"`
	UnderwriteScope  string `xml:"underwriteScope"`
}

type Applicant struct {
	ApplicationPersonnelType string `xml:"applicationPersonnelType"`
	Enterprise               string `xml:"enterprise"`
	EnterpriseCreditCode     string `xml:"enterpriseCreditCode"`
	Address                  string `xml:"address"`
	Email                    string `xml:"email"`
	LinkManName              string `xml:"linkManName"`
	LinkManMobileTelephone   string `xml:"linkManMobileTelephone"`
	CertificateType          string `xml:"certificateType"`
	CertType                 string `xml:"certType"`
	IndustryLevel1           string `xml:"industryLevel1"`
	IndustryLevel2           string `xml:"industryLevel2"`
	IndustryLevel3           string `xml:"industryLevel3"`
	IndustryLevel4           string `xml:"industryLevel4"`
}

type Plans struct {
	Plan []Plan `xml:"plan"`
}

type Plan struct {
	PlanNo         string      `xml:"planNo"`
	PlanName       string      `xml:"planName"`
	OccupationType string      `xml:"occupationType"`
	Coverages      Coverages   `xml:"coverages"`
	Insurances     []Insurance `xml:"insurances>insurance"`
}

type Coverages struct {
	Coverage []Coverage `xml:"coverage"`
}

type Coverage struct {
	CoverageCode    string `xml:"coverageCode"`
	InsuredAmount   int    `xml:"insuredAmount"`
	PaymentRatio    string `xml:"paymentRatio,omitempty"`
	DeductionAmount int    `xml:"deductionAmount,omitempty"`
}

type Insurance struct {
	InsuranceName  string `xml:"insuranceName"`
	BirthDate      string `xml:"birthDate"`
	CertiNo        string `xml:"certiNo"`
	CertiType      string `xml:"certiType"`
	RelApplicant   string `xml:"relApplicant"`
	Gender         string `xml:"gender"`
	OccupationCode string `xml:"occupationCode"`
}
