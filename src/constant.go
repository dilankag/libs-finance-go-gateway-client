/*=====================Finance Go Gateway Client===============================

    

    Author(s):
        Dilanka Gamage

==============================================================================*/

package main

type ApiVersion string
const (
    V_LATEST ApiVersion = "v3_8_0"
)

type ApiAction string
const (
    PEOPLE_PROFILE ApiAction = "PEOPLE_PROFILE"
    DRIVER_PROFILE ApiAction = "DRIVER_PROFILE"
    DRIVER_TRIP_TRANSACTION ApiAction = "DRIVER_TRIP_TRANSACTION"
    DRIVER_TRIP_SUMMARY ApiAction = "DRIVER_TRIP_SUMMARY"
    FGL_STAKEHOLDER_MONTHLY_SUMMARY ApiAction = "FGL_STAKEHOLDER_MONTHLY_SUMMARY"
    FGL_STAKEHOLDER_BALANCE_SUMMARY ApiAction = "FGL_STAKEHOLDER_BALANCE_SUMMARY"
    FGL_STAKEHOLDER_TRANSACTION_DETAIL ApiAction = "FGL_STAKEHOLDER_TRANSACTION_DETAIL"
    FGL_ACCOUNTING_RULE_PROFILE ApiAction = "FGL_ACCOUNTING_RULE_PROFILE"
    PROMOTION_CONDITION_MAPPING ApiAction = "PROMOTION_CONDITION_MAPPING"
)

type DateType string
const (
    CREATE_TIME DateType = "CREATE_TIME"
    CREATE_DATE DateType = "CREATE_DATE"
)

type TransactionType string
const (
    CREDIT TransactionType = "CREDIT"
    DEBIT TransactionType = "DEBIT"
)

const (
    DEFAULT_PAGE_SIZE int = 10
    MAXIMUM_PAGE_SIZE int = 100
)

type HttpHeader string
const (
    HMAC HttpHeader = "HMAC"
    AUTH HttpHeader = "AUTH"
    CSRF HttpHeader = "CSRF"
)