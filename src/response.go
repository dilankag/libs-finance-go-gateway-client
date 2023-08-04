/*=====================Finance Go Gateway Client===============================

    

    Author(s):
        Dilanka Gamage

==============================================================================*/

package main

type GatewayResponse struct {
    ApiVersion ApiVersion     `json:"apiVersion"`
    ApiAction ApiAction       `json:"apiAction"`
    MessageId string          `json:"messageId"`
    ResponseDate string       `json:"responseDate"`
    TimeElapsed int64         `json:"timeElapsed"`
    Error *Error              `json:"error"`
    ResponseData interface{}  `json:"responseData"`
}

type BaseReportResponse struct {
    ResponseDate string       `json:"responseDate"`
    ReportIndex ReportIndex   `json:"reportIndex"`
    PageSize int              `json:"pageSize"`
    PageIndex int             `json:"pageIndex"`
    TotalSize int64           `json:"totalSize"`
    ErrorText string          `json:"errorText"`
}

type PeopleProfileReportResponse struct {
    BaseReportResponse
    Profiles []PeopleProfile  `json:"profiles"`
}

type DriverProfileReportResponse struct {
    BaseReportResponse
    Profiles []DriverProfile  `json:"profiles"`
}

type DriverTripTransactionReportResponse struct {
    BaseReportResponse
    Transactions []DriverTripTransaction  `json:"transactions"`
}

type DriverTripSummaryReportResponse struct {
    BaseReportResponse
    Summaries []DriverTripSummary  `json:"summaries"`
}

type FglStakeholderMonthlySummaryReportResponse struct {
    BaseReportResponse
    Summaries []FglStakeholderMonthlySummary  `json:"summaries"`
}

type FglStakeholderBalanceSummaryReportResponse struct {
    BaseReportResponse
    Summaries []FglStakeholderBalanceSummary  `json:"summaries"`
}

type FglAccountingRuleProfileReportResponse struct {
    BaseReportResponse
    Profiles []FglAccountingRuleProfile  `json:"profiles"`
}

type FglStakeholderTransactionDetailReportResponse struct {
    BaseReportResponse
    Transactions []FglStakeholderTransactionDetail  `json:"transactions"`
}

type PromotionConditionMappingReportResponse struct {
    BaseReportResponse
    Mappings []PromotionConditionMapping  `json:"mappings"`
}