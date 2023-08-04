/*=====================Finance Go Gateway Client===============================

	

	Author(s):
		Dilanka Gamage

==============================================================================*/

package main

import (
    "fmt"
    "bytes"
    "errors"
    "io/ioutil"
    "net/http"
    "encoding/json"
    "crypto/hmac"
    "crypto/sha256"
    "encoding/hex"
)

type ClientConfig struct {
    ServiceEndpoint string
    HmacSecret string
    AuthToken string
    CsrfToken string
}

const (
    // Error codes
    BAD_REQUEST = "BAD_REQUEST"
    BAD_RESPONSE = "BAD_RESPONSE"
    POST_FAILED = "POST_FAILED"
)

var (
    gatewayClientException     = errors.New("Failed to process request.")
	gatewayDataFormatException = errors.New("Validation failed. Invalid input format.")
	gatewayRequiredDataException = errors.New("Validation failed. Required input not supplied.")
)

func BuildClientConfig(serviceEndpoint, hmacSecret string) ClientConfig {
    config := ClientConfig {
        ServiceEndpoint: serviceEndpoint,
        HmacSecret: hmacSecret,
    }
    return config
}

func GenerateHmac(secretKey, rawData string) string {
    mac := hmac.New(sha256.New, []byte(secretKey))
    mac.Write([]byte(rawData))
    hashBytes := mac.Sum(nil)
    return hex.EncodeToString(hashBytes[:])
}

func PostHttpRequest(httpUrl, httpBody, hmac string) (string, error) {
    fmt.Println("Http url: ", httpUrl)
    fmt.Println("Http header: ", HMAC, ": ", hmac)
    fmt.Println("Http body: ", httpBody)
    httpRequest, err := http.NewRequest("POST", httpUrl, bytes.NewBuffer([]byte(httpBody)))
    httpRequest.Header.Set("Content-Type", "application/json")
    httpRequest.Header.Set(string(HMAC), hmac)
    httpClient := &http.Client{}
    httpResponse, err := httpClient.Do(httpRequest)
    if err != nil {
        return "{}", err
    }
    defer httpResponse.Body.Close()
    rawResponse, _ := ioutil.ReadAll(httpResponse.Body)
    fmt.Println("Http response: ", string(rawResponse))
    return string(rawResponse), nil
}

func ProcessGatewayRequest(clientConfig ClientConfig, gatewayRequest GatewayRequest) GatewayResponse {
    gatewayResponse := GatewayResponse{}
    rawRequest, err1 := json.Marshal(gatewayRequest)
    if err1 != nil {
        gatewayResponse.Error = &Error {
            Code: BAD_REQUEST,
            Text: fmt.Sprintf("Unable to marshal the request. ", err1),
        }
        return gatewayResponse
    }
    hmac := GenerateHmac(clientConfig.HmacSecret, string(rawRequest))
    rawResponse, err2 := PostHttpRequest(clientConfig.ServiceEndpoint, string(rawRequest), hmac)
    if err2 != nil {
        gatewayResponse.Error = &Error {
            Code: POST_FAILED,
            Text: fmt.Sprintf("Unable to post the request. ", err2),
        }
        return gatewayResponse
    }
    json.Unmarshal([]byte(rawResponse), &gatewayResponse)
    return gatewayResponse
}

func FetchPeopleProfile(clientConfig ClientConfig, reportRequest PeopleProfileReportRequest) (PeopleProfileReportResponse, *Error) {
    peopleProfileReportResponse := PeopleProfileReportResponse{}
    gatewayRequest := BuildGatewayRequest(PEOPLE_PROFILE, reportRequest)
    gatewayResponse := ProcessGatewayRequest(clientConfig, gatewayRequest)
    if gatewayResponse.Error != nil {
        return peopleProfileReportResponse, gatewayResponse.Error
    }
    rawResponseData, err := json.Marshal(gatewayResponse.ResponseData)
    if err != nil {
        gatewayResponse.Error = &Error {
            Code: BAD_RESPONSE,
            Text: fmt.Sprintf("Unable to marshal the response. ", err),
        }
        return peopleProfileReportResponse, gatewayResponse.Error
    }
    json.Unmarshal([]byte(string(rawResponseData)), &peopleProfileReportResponse)
    return peopleProfileReportResponse, nil
}

func FetchDriverProfile(clientConfig ClientConfig, reportRequest DriverProfileReportRequest) (DriverProfileReportResponse, *Error) {
    driverProfileReportResponse := DriverProfileReportResponse{}
    gatewayRequest := BuildGatewayRequest(DRIVER_PROFILE, reportRequest)
    gatewayResponse := ProcessGatewayRequest(clientConfig, gatewayRequest)
    if gatewayResponse.Error != nil {
        return driverProfileReportResponse, gatewayResponse.Error
    }
    rawResponseData, err := json.Marshal(gatewayResponse.ResponseData)
    if err != nil {
        gatewayResponse.Error = &Error {
            Code: BAD_RESPONSE,
            Text: fmt.Sprintf("Unable to marshal the response. ", err),
        }
        return driverProfileReportResponse, gatewayResponse.Error
    }
    json.Unmarshal([]byte(string(rawResponseData)), &driverProfileReportResponse)
    return driverProfileReportResponse, nil
}

func FetchDriverTripTransaction(clientConfig ClientConfig, reportRequest DriverTripTransactionReportRequest) (DriverTripTransactionReportResponse, *Error) {
    driverTripTransactionReportResponse := DriverTripTransactionReportResponse{}
    gatewayRequest := BuildGatewayRequest(DRIVER_TRIP_TRANSACTION, reportRequest)
    gatewayResponse := ProcessGatewayRequest(clientConfig, gatewayRequest)
    if gatewayResponse.Error != nil {
        return driverTripTransactionReportResponse, gatewayResponse.Error
    }
    rawResponseData, err := json.Marshal(gatewayResponse.ResponseData)
    if err != nil {
        gatewayResponse.Error = &Error {
            Code: BAD_RESPONSE,
            Text: fmt.Sprintf("Unable to marshal the response. ", err),
        }
        return driverTripTransactionReportResponse, gatewayResponse.Error
    }
    json.Unmarshal([]byte(string(rawResponseData)), &driverTripTransactionReportResponse)
    return driverTripTransactionReportResponse, nil
}

func FetchDriverTripSummary(clientConfig ClientConfig, reportRequest DriverTripSummaryReportRequest) (DriverTripSummaryReportResponse, *Error) {
    driverTripSummaryReportResponse := DriverTripSummaryReportResponse{}
    gatewayRequest := BuildGatewayRequest(DRIVER_TRIP_SUMMARY, reportRequest)
    gatewayResponse := ProcessGatewayRequest(clientConfig, gatewayRequest)
    if gatewayResponse.Error != nil {
        return driverTripSummaryReportResponse, gatewayResponse.Error
    }
    rawResponseData, err := json.Marshal(gatewayResponse.ResponseData)
    if err != nil {
        gatewayResponse.Error = &Error {
            Code: BAD_RESPONSE,
            Text: fmt.Sprintf("Unable to marshal the response. ", err),
        }
        return driverTripSummaryReportResponse, gatewayResponse.Error
    }
    json.Unmarshal([]byte(string(rawResponseData)), &driverTripSummaryReportResponse)
    return driverTripSummaryReportResponse, nil
}

func FetchFglStakeholderMonthlySummary(clientConfig ClientConfig, reportRequest FglStakeholderMonthlySummaryReportRequest) (FglStakeholderMonthlySummaryReportResponse, *Error) {
    fglStakeholderMonthlySummaryReportResponse := FglStakeholderMonthlySummaryReportResponse{}
    gatewayRequest := BuildGatewayRequest(FGL_STAKEHOLDER_MONTHLY_SUMMARY, reportRequest)
    gatewayResponse := ProcessGatewayRequest(clientConfig, gatewayRequest)
    if gatewayResponse.Error != nil {
        return fglStakeholderMonthlySummaryReportResponse, gatewayResponse.Error
    }
    rawResponseData, err := json.Marshal(gatewayResponse.ResponseData)
    if err != nil {
        gatewayResponse.Error = &Error {
            Code: BAD_RESPONSE,
            Text: fmt.Sprintf("Unable to marshal the response. ", err),
        }
        return fglStakeholderMonthlySummaryReportResponse, gatewayResponse.Error
    }
    json.Unmarshal([]byte(string(rawResponseData)), &fglStakeholderMonthlySummaryReportResponse)
    return fglStakeholderMonthlySummaryReportResponse, nil
}

func FetchFglStakeholderBalanceSummary(clientConfig ClientConfig, reportRequest FglStakeholderBalanceSummaryReportRequest) (FglStakeholderBalanceSummaryReportResponse, *Error) {
    fglStakeholderBalanceSummaryReportResponse := FglStakeholderBalanceSummaryReportResponse{}
    gatewayRequest := BuildGatewayRequest(FGL_STAKEHOLDER_BALANCE_SUMMARY, reportRequest)
    gatewayResponse := ProcessGatewayRequest(clientConfig, gatewayRequest)
    if gatewayResponse.Error != nil {
        return fglStakeholderBalanceSummaryReportResponse, gatewayResponse.Error
    }
    rawResponseData, err := json.Marshal(gatewayResponse.ResponseData)
    if err != nil {
        gatewayResponse.Error = &Error {
            Code: BAD_RESPONSE,
            Text: fmt.Sprintf("Unable to marshal the response. ", err),
        }
        return fglStakeholderBalanceSummaryReportResponse, gatewayResponse.Error
    }
    json.Unmarshal([]byte(string(rawResponseData)), &fglStakeholderBalanceSummaryReportResponse)
    return fglStakeholderBalanceSummaryReportResponse, nil
}

func FetchFglStakeholderTransactionDetail(clientConfig ClientConfig, reportRequest FglStakeholderTransactionDetailReportRequest) (FglStakeholderTransactionDetailReportResponse, *Error) {
    fglStakeholderTransactionDetailReportResponse := FglStakeholderTransactionDetailReportResponse{}
    gatewayRequest := BuildGatewayRequest(FGL_STAKEHOLDER_TRANSACTION_DETAIL, reportRequest)
    gatewayResponse := ProcessGatewayRequest(clientConfig, gatewayRequest)
    if gatewayResponse.Error != nil {
        return fglStakeholderTransactionDetailReportResponse, gatewayResponse.Error
    }
    rawResponseData, err := json.Marshal(gatewayResponse.ResponseData)
    if err != nil {
        gatewayResponse.Error = &Error {
            Code: BAD_RESPONSE,
            Text: fmt.Sprintf("Unable to marshal the response. ", err),
        }
        return fglStakeholderTransactionDetailReportResponse, gatewayResponse.Error
    }
    json.Unmarshal([]byte(string(rawResponseData)), &fglStakeholderTransactionDetailReportResponse)
    return fglStakeholderTransactionDetailReportResponse, nil
}

func FetchFglAccountingRuleProfile(clientConfig ClientConfig, reportRequest FglAccountingRuleProfileReportRequest) (FglAccountingRuleProfileReportResponse, *Error) {
    fglAccountingRuleProfileReportResponse := FglAccountingRuleProfileReportResponse{}
    gatewayRequest := BuildGatewayRequest(FGL_ACCOUNTING_RULE_PROFILE, reportRequest)
    gatewayResponse := ProcessGatewayRequest(clientConfig, gatewayRequest)
    if gatewayResponse.Error != nil {
        return fglAccountingRuleProfileReportResponse, gatewayResponse.Error
    }
    rawResponseData, err := json.Marshal(gatewayResponse.ResponseData)
    if err != nil {
        gatewayResponse.Error = &Error {
            Code: BAD_RESPONSE,
            Text: fmt.Sprintf("Unable to marshal the response. ", err),
        }
        return fglAccountingRuleProfileReportResponse, gatewayResponse.Error
    }
    json.Unmarshal([]byte(string(rawResponseData)), &fglAccountingRuleProfileReportResponse)
    return fglAccountingRuleProfileReportResponse, nil
}

func FetchPromotionConditionMapping(clientConfig ClientConfig, reportRequest PromotionConditionMappingReportRequest) (PromotionConditionMappingReportResponse, *Error) {
    promotionConditionMappingReportResponse := PromotionConditionMappingReportResponse{}
    gatewayRequest := BuildGatewayRequest(PROMOTION_CONDITION_MAPPING, reportRequest)
    gatewayResponse := ProcessGatewayRequest(clientConfig, gatewayRequest)
    if gatewayResponse.Error != nil {
        return promotionConditionMappingReportResponse, gatewayResponse.Error
    }
    rawResponseData, err := json.Marshal(gatewayResponse.ResponseData)
    if err != nil {
        gatewayResponse.Error = &Error {
            Code: BAD_RESPONSE,
            Text: fmt.Sprintf("Unable to marshal the response. ", err),
        }
        return promotionConditionMappingReportResponse, gatewayResponse.Error
    }
    json.Unmarshal([]byte(string(rawResponseData)), &promotionConditionMappingReportResponse)
    return promotionConditionMappingReportResponse, nil
}