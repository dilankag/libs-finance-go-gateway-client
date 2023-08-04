/*=====================Finance Go Gateway Client===============================

    

    Author(s):
        Dilanka Gamage

==============================================================================*/

package main

type GatewayRequest struct {
    ApiVersion ApiVersion   `json:"apiVersion"`
    ApiAction ApiAction   `json:"apiAction"`
    MessageId string        `json:"messageId"`
    //RequestDate string      `json:"requestDate"`
    ValidateOnly bool       `json:"validateOnly"`
    RequestData interface{} `json:"requestData"`
}

type BaseReportRequest struct {
    FromStaff bool          `json:"fromStaff"`
    FromPortal bool         `json:"fromPortal"`
    PagingEnabled bool      `json:"pagingEnabled"`
    PageSize int            `json:"pageSize"`
    PageIndex int           `json:"pageIndex"`
    ExportEnabled bool      `json:"exportEnabled"`
    Sorters []Sorter        `json:"sorters"`
}

type PeopleProfileReportRequest struct {
    BaseReportRequest
    PeopleId interface{}               `json:"peopleId"`
    WithDriverTripSummary bool         `json:"withDriverTripSummary"`
    WithoutEmptyDriverTripSummary bool `json:"withoutEmptyDriverTripSummary"`
}

type DriverProfileReportRequest struct {
    BaseReportRequest
    DriverId interface{}               `json:"driverId"`
    WithDriverTripSummary bool         `json:"withDriverTripSummary"`
    WithoutEmptyDriverTripSummary bool `json:"withoutEmptyDriverTripSummary"`
}

type DriverTripTransactionReportRequest struct {
    BaseReportRequest
    TransactionId interface{}          `json:"transactionId"`
    DriverId interface{}               `json:"driverId"`
    TripId interface{}                 `json:"tripId"`
    TransactionTypes []TransactionType `json:"transactionTypes"`
    TransactionCategories []int        `json:"transactionCategories"`
    MinAmountInRupee interface{}       `json:"minAmountInRupee"`
    MaxAmountInRupee interface{}       `json:"maxAmountInRupee"`
    MinAmountInCents interface{}       `json:"minAmountInCents"`
    MaxAmountInCents interface{}       `json:"maxAmountInCents"`
    Description interface{}            `json:"description"`
    DateType DateType                  `json:"dateType"`
    FromDate interface{}               `json:"fromDate"`
    ToDate interface{}                 `json:"toDate"`
    CreatedBy []int                    `json:"createdBy"`
    WithDriverProfile bool             `json:"withDriverProfile"`
}

type DriverTripSummaryReportRequest struct {
    BaseReportRequest
    DriverId interface{}               `json:"driverId"`
    WithDriverProfile bool             `json:"withDriverProfile"`
}

type FglStakeholderMonthlySummaryReportRequest struct {
    BaseReportRequest
    StakeholderId interface{}           `json:"stakeholderId"`
    CurrencyTypeId interface{}          `json:"currencyTypeId"`
    AccountTypeId interface{}           `json:"accountTypeId"`
    AccountReferenceId interface{}      `json:"accountReferenceId"`
    RecentNMonths int                   `json:"recentNMonths"`
    AccountingRuleIdsIn []string        `json:"accountingRuleIdsIn"`
    AccountingRuleIdsNotIn []string     `json:"accountingRuleIdsNotIn"`
}

type FglStakeholderBalanceSummaryReportRequest struct {
    BaseReportRequest
    StakeholderId interface{}           `json:"stakeholderId"`
    CurrencyTypeId interface{}          `json:"currencyTypeId"`
    AccountTypeId interface{}           `json:"accountTypeId"`
    AccountReferenceId interface{}      `json:"accountReferenceId"`
    FromDate interface{}                `json:"fromDate"`
    ToDate interface{}                  `json:"toDate"`
    AccountingRuleIdsIn []string        `json:"accountingRuleIdsIn"`
    AccountingRuleIdsNotIn []string     `json:"accountingRuleIdsNotIn"`
}

type FglStakeholderTransactionDetailReportRequest struct {
    BaseReportRequest
    StakeholderId interface{}           `json:"stakeholderId"`
    CurrencyTypeId interface{}          `json:"currencyTypeId"`
    AccountTypeId interface{}           `json:"accountTypeId"`
    AccountReferenceId interface{}      `json:"accountReferenceId"`
    TransactionReferenceId interface{}  `json:"transactionReferenceId"`
    AccountingEntry interface{}         `json:"accountingEntry"`
    FromDate interface{}                `json:"fromDate"`
    ToDate interface{}                  `json:"toDate"`
    MinAmount interface{}               `json:"minAmount"`
    MaxAmount interface{}               `json:"maxAmount"`
    AccountingRuleIdsIn []string        `json:"accountingRuleIdsIn"`
    AccountingRuleIdsNotIn []string     `json:"accountingRuleIdsNotIn"`
    TransactionCategoryIdsIn []int      `json:"transactionCategoryIdsIn"`
    TransactionCategoryIdsNotIn []int   `json:"transactionCategoryIdsNotIn"`
}

type FglAccountingRuleProfileReportRequest struct {
    BaseReportRequest
    RuleId interface{}                  `json:"ruleId"`
    RuleIdLike interface{}              `json:"ruleIdLike"`
    Active bool                         `json:"active"`
}

type PromotionConditionMappingReportRequest struct {
    BaseReportRequest
    PromotionStatus interface{}         `json:"promotionStatus"`
    ConditionValuePhrase interface{}    `json:"conditionValuePhrase"`
    StillActive bool                    `json:"stillActive"`
    ConditionEntityFieldIdsIn []int     `json:"conditionEntityFieldIdsIn"`
    ConditionEntityFieldIdsNotIn []int  `json:"conditionEntityFieldIdsNotIn"`
}

func BuildGatewayRequest(apiAction ApiAction, requestData interface{}) GatewayRequest {
    request := GatewayRequest {
        ApiVersion: V_LATEST,
        ApiAction: apiAction,
        RequestData: requestData,
    }
    return request
}