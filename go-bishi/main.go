package main

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"go-bishi/model"
	"io"
	"log"
	"net/http"
)

// URL和身份信息结构体
type ApiClient struct {
	BaseURL string
	ApiKey  string
}

// 返回客户端实例(构造函数)
func NewApiClient(baseURL, apiKey string) *ApiClient {
	return &ApiClient{
		BaseURL: baseURL,
		ApiKey:  apiKey,
	}
}

// 发送请求到API，并返回响应
func (client *ApiClient) SendRequest(reqBody TXTpaicRequest) (string, error) {
	// 编码为XML
	xmlData, err := xml.MarshalIndent(reqBody, "", "  ")
	if err != nil {
		return "", err
	}
	// 使用适当的HTTP方法(GET, POST等)发送请求。
	req, err := http.NewRequest("POST", client.BaseURL, bytes.NewBuffer(xmlData))
	if err != nil {
		return "", err
	}
	// 正确设置请求头,包括内容类型和授权信息。
	req.Header.Set("Content-Type", "application/xml")
	req.Header.Set("Authorization", "Bearer "+client.ApiKey)
	clientHTTP := &http.Client{}
	resp, err := clientHTTP.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	//在这里响应可能会检查出现的错误
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Request failed with status: %s", resp.Status)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

// 实现一个机制,允许灵活地传递不同API所需的参数。确保所有必要的参数都被正确传递。
// 这里或许可以将每个结构体都是实现一个接口方法来验证参数的准确性
func CreateRequest(transRefId, transType, transExeDate string, policy Policy, applicant Applicant, plans Plans) (TXTpaicRequest, error) {
	//这里可以验证所必须的参数
	if transRefId == "" {
		return TXTpaicRequest{}, errors.New("transRefId is nil")
	}
	return TXTpaicRequest{
		Head: Head{
			TransRefId:   transRefId,
			TransType:    transType,
			TransExeDate: transExeDate,
		},
		Body: Body{
			Policy:    policy,
			Applicant: applicant,
			Plans:     plans,
		},
	}, nil
}

func main() {
	// 创建API客户端实例
	client := NewApiClient("https://open.axa.cn/apiplatform/#/user?id=3&checked=283/groupApply.do", "身份token")

	// 准备请求数据
	requestData := TXTpaicRequest{
		Head: Head{
			TransRefId:   "2e90042c-dbbc-4eb3-9893-e75a49843a94",
			TransType:    "001",
			TransExeDate: "2024-03-12 23:59:59",
		},
		Body: Body{
			Policy: Policy{
				PlanCode:         "P21",
				ProductCode:      "P21002",
				PaymentMode:      "0",
				Count:            1,
				TransApplyDate:   "2023-03-10 14:42:16",
				TransBeginDate:   "2023-03-13 00:00:00",
				TransEndDate:     "2024-03-12 23:59:59",
				AgentCode:        "agent81867",
				CurrencyCode:     "01",
				CommissionFactor: 35,
				UnderwriteScope:  "0",
			},
			Applicant: Applicant{
				ApplicationPersonnelType: "0",
				Enterprise:               "测试",
				EnterpriseCreditCode:     "91440101MA9Y9BQ5X7",
				Address:                  "测试测试测试测试测试测试",
				Email:                    "12312312@qq.com",
				LinkManName:              "测试",
				LinkManMobileTelephone:   "17612312312",
				CertificateType:          "6",
				CertType:                 "2",
				IndustryLevel1:           "A0000",
				IndustryLevel2:           "A0100",
				IndustryLevel3:           "A0130",
				IndustryLevel4:           "A0134",
			},
			Plans: Plans{
				Plan: []Plan{
					{
						PlanNo:         "1",
						PlanName:       "安盛天平团体意外险",
						OccupationType: "1",
						Coverages: Coverages{
							Coverage: []Coverage{
								{CoverageCode: "FTPB143", InsuredAmount: 100000},
								{CoverageCode: "FTPB400", InsuredAmount: 10000, PaymentRatio: "100%", DeductionAmount: 0},
								{CoverageCode: "FTPB399", InsuredAmount: 0},
							},
						},
						Insurances: []Insurance{
							{InsuranceName: "测试", BirthDate: "1990-03-07", CertiNo: "110101199003075891", CertiType: "1", RelApplicant: "5", Gender: "M", OccupationCode: "3101010011"},
							// 添加更多保险
						},
					},
				},
			},
		},
	}

	// 发送请求
	responseBody, err := client.SendRequest(requestData)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	var response model.TXTpaicResponse
	//将响应反序列化为结构体,方便存储
	err1 := xml.Unmarshal([]byte(responseBody), &response)
	if err1 != nil {
		fmt.Println("Error unmarshalling XML:", err)
		return
	}
}
