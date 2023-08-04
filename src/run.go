/*=====================Finance Go Gateway Client===============================

    © 2015 Digital Mobility Solutions Lanka (Pvt) Ltd. All rights reserved.

    Author(s):
        Dilanka Gamage

==============================================================================*/

package main

import (
    "fmt"
)

var (
    clientConfig = BuildPRODClientConfig()                // init once, common for all tests
    baseReportRequest = BuildCommonBaseReportRequest()   // init once, common for all tests
)

func BuildUATClientConfig() ClientConfig {
    return BuildClientConfig(
        "http://146.148.110.253:8080/proxy/finance/reporting",  // set service endpoint
        "not-for-uat")                                          // set hmac secret (applies in staging and production)
}

func BuildSTAGEClientConfig() ClientConfig {
    return BuildClientConfig(
        "https://finance-report.picme.lk/proxy/finance/reporting",  // set service endpoint
        "**********")                                               // set hmac secret (applies in staging and production)
}

func BuildPRODClientConfig() ClientConfig {
    return BuildClientConfig(
        "https://finance-report.pickme.lk/proxy/finance/reporting",  // set service endpoint
        "1D28buTZoq")                                                // set hmac secret (applies in staging and production)
}

/*
The following test will fetch People profile(s) with trip summaries
*/
func Test_FetchPeopleProfile() {
    // Apply sorters
    baseReportRequest.Sorters = nil // first, clear sorters
    sortByPeopleIdDesc := BuildSorter("peopleId", true) // sort by people id in decending order 
    baseReportRequest.Sorters = append(baseReportRequest.Sorters, sortByPeopleIdDesc)
    // Build report request
    reportRequest := PeopleProfileReportRequest {
        // Set request params
        BaseReportRequest: baseReportRequest,  // append base report request
        //PeopleId: 223,                       // fetch an individual profile
        WithDriverTripSummary: true,           // fetch with driver trip summary
        WithoutEmptyDriverTripSummary: false,  // fetch with out empty driver trip summary
    }
    // Call report
    reportResponse, error := FetchPeopleProfile(clientConfig, reportRequest)
    // Handle error (if any)
    if error != nil {
        fmt.Println("Error code: ", error.Code)
        fmt.Println("Error text: ", error.Text)
    }
    // Process report response
    fmt.Println("People profiles: ", len(reportResponse.Profiles))
}

/*
The following test will fetch Driver profile(s) with trip summaries
*/
func Test_FetchDriverProfile() {
    // Apply sorters
    baseReportRequest.Sorters = nil // first, clear sorters
    sortByDriverIdDesc := BuildSorter("driverId", true) // sort by driver id in decending order 
    baseReportRequest.Sorters = append(baseReportRequest.Sorters, sortByDriverIdDesc)
    // Build report request
    reportRequest := DriverProfileReportRequest {
        // Set request params
        BaseReportRequest: baseReportRequest,  // append base report request
        //DriverId: 223,                       // fetch an individual profile
        WithDriverTripSummary: true,           // fetch with driver trip summary
        WithoutEmptyDriverTripSummary: false,  // fetch with out empty driver trip summary
    }
    // Call report
    reportResponse, error := FetchDriverProfile(clientConfig, reportRequest)
    // Handdle error (if any)
    if error != nil {
        fmt.Println("Error code: ", error.Code)
        fmt.Println("Error text: ", error.Text)
    }
    // Process report response
    fmt.Println("Driver profiles: ", len(reportResponse.Profiles))
}

/*
The following test will fetch Driver trip transaction(s)
*/
func Test_FetchDriverTripTransaction() {
    // Apply sorters
    baseReportRequest.Sorters = nil // first, clear sorters
    sortByTransactionIdDesc := BuildSorter("transactionId", true) // sort by transaction id in decending order 
    baseReportRequest.Sorters = append(baseReportRequest.Sorters, sortByTransactionIdDesc)
    // Build report request
    reportRequest := DriverTripTransactionReportRequest {
        // Set request params
        BaseReportRequest: baseReportRequest,    // append base report request
        //TransactionId: 20060799,               // fetch an individual transaction id
        //DriverId: 223,                         // fetch by driver id
        //TripId: 382818176,                     // fetch by trip id
        // Apply date range filters
        DateType: CREATE_TIME,                   // mandatory to set date type
        //FromDate: "2019-01-01",                // in yyyy-MM-dd'T'HH:mm:ss.SSSZ and yyyy-MM-dd HH:mm:ss as well
        //ToDate: "2019-12-31",                  // in yyyy-MM-dd'T'HH:mm:ss.SSSZ and yyyy-MM-dd HH:mm:ss as well
        // Apply min and max amount filters
        // MinAmountInCents: 199,
        // MaxAmountInCents: 9999,
        WithDriverProfile: true,                 // fetch with driver profile
    }
    // Call report
    reportResponse, error := FetchDriverTripTransaction(clientConfig, reportRequest)
    // Handdle error (if any)
    if error != nil {
        fmt.Println("Error code: ", error.Code)
        fmt.Println("Error text: ", error.Text)
    }
    // Process report response
    fmt.Println("Driver trip transactions: ", len(reportResponse.Transactions))
}

/*
The following test will fetch Driver trip transaction(s)
*/
func Test_FetchDriverTripSummary() {
    // Build report request
    reportRequest := DriverTripSummaryReportRequest {
        // Set request params
        BaseReportRequest: baseReportRequest,    // append base report request
        DriverId: 223,                           // fetch by driver id (mandatory to set)
        WithDriverProfile: true,                 // fetch with driver profile
    }
    // Call report
    reportResponse, error := FetchDriverTripSummary(clientConfig, reportRequest)
    // Handdle error (if any)
    if error != nil {
        fmt.Println("Error code: ", error.Code)
        fmt.Println("Error text: ", error.Text)
    }
    // Process report response
    fmt.Println("Driver trip summaries: ", len(reportResponse.Summaries))
}

/*
The following test will fetch transaction summaries of Merchant for the recent (N) months.
*/
func Test_FetchFglStakeholderMonthlySummary() {
    // Build report request
    reportRequest := FglStakeholderMonthlySummaryReportRequest {
        // Set request params
        BaseReportRequest: baseReportRequest,   // append base report request
        StakeholderId: "MERCHANT",              // fetch by stakeholder id (mandatory to set)
        AccountReferenceId: "38138",            // fetch by merchant id (mandatory to set)
        RecentNMonths: 3,                       // fetch by most recent (N) months including the current month, default is set to 3.
        // Optional request params
        // The array of accounting rule IDs that should be considered
        // AccountingRuleIdsIn: []string{
        //                                 "job_completed_order_amount_rule",
        //                                 "job_completed_merchant_app_usage_rule",
        //                              },
        // // The array of accounting rule IDs that should NOT be considered.
        // AccountingRuleIdsNotIn: []string{
        //                                 "merchant_auto_settlement_hnb_rule",
        //                                 "merchant_auto_settlement_ntb_rule",
        //                             },
    }
    // Call report
    reportResponse, error := FetchFglStakeholderMonthlySummary(clientConfig, reportRequest)
    // Handdle error (if any)
    if error != nil {
        fmt.Println("Error code: ", error.Code)
        fmt.Println("Error text: ", error.Text)
    }
    // Process report response
    fmt.Println("Stakeholder monthly summaries: ", len(reportResponse.Summaries))
}

/*
The following test will fetch the account balance of Merchant for a period of time.
*/
func Test_FetchFglStakeholderBalanceSummary() {
    // Build report request
    reportRequest := FglStakeholderBalanceSummaryReportRequest {
        // Set request params
        BaseReportRequest: baseReportRequest,   // append base report request
        StakeholderId: "MERCHANT",              // fetch by stakeholder id (mandatory to set)
        AccountReferenceId: "38138",            // fetch by merchant id (mandatory to set)
        FromDate: "2019-05-31",                 // The starting date of reporting period in yyyy-MM-dd format (mandatory to set)
        ToDate: "2019-06-01",                   // The ending date of reporting period in yyyy-MM-dd format (mandatory to set)
        // Optional request params
        // The array of accounting rule IDs that should be considered
        AccountingRuleIdsIn: []string{
                                         "job_completed_order_amount_rule",
                                         "job_completed_merchant_app_usage_rule",
                                      },
        // The array of accounting rule IDs that should NOT be considered.
        AccountingRuleIdsNotIn: []string{
                                        "merchant_self_settlement_hnb_rule",
                                        "merchant_self_settlement_ntb_rule",
                                    },
    }
    // Call report
    reportResponse, error := FetchFglStakeholderBalanceSummary(clientConfig, reportRequest)
    // Handdle error (if any)
    if error != nil {
        fmt.Println("Error code: ", error.Code)
        fmt.Println("Error text: ", error.Text)
    }
    // Process report response
    fmt.Println("Stakeholder balance summaries: ", len(reportResponse.Summaries))
}

/*
The following test will fetch the most recent transactions of Merchant for a period of time.
*/
func Test_FetchFglStakeholderTransactionDetail() {
    // Apply sorters
    baseReportRequest.Sorters = nil // first, clear sorters
    sortByCreatedDate := BuildSorter("createdDate", true) // sort by transaction id in decending order 
    baseReportRequest.Sorters = append(baseReportRequest.Sorters, sortByCreatedDate)
    // Build report request
    reportRequest := FglStakeholderTransactionDetailReportRequest {
        // Set request params
        BaseReportRequest: baseReportRequest,   // append base report request
        StakeholderId: "MERCHANT",              // fetch by stakeholder id (mandatory to set)
        //AccountReferenceId: "2787",            // fetch by merchant id (mandatory to set)
        FromDate: "2019-01-31",                 // The starting date of reporting period in yyyy-MM-dd format (mandatory to set)
        ToDate: "2019-06-06",                   // The ending date of reporting period in yyyy-MM-dd format (mandatory to set)
        // Optional request params
        // The array of accounting rule IDs that should be considered
        // AccountingRuleIdsIn: []string{
        //                                  "merchant_self_settlement_hnb_rule",
        //                                  "merchant_self_settlement_ntb_rule",
        //                               },
        // // The array of accounting rule IDs that should NOT be considered.
        // AccountingRuleIdsNotIn: []string{
        //                                 "merchant_self_settlement_hnb_rule",
        //                                 "merchant_self_settlement_ntb_rule",
        //                             },
    }
    // Call report
    reportResponse, error := FetchFglStakeholderTransactionDetail(clientConfig, reportRequest)
    // Handdle error (if any)
    if error != nil {
        fmt.Println("Error code: ", error.Code)
        fmt.Println("Error text: ", error.Text)
    }
    // Process report response
    fmt.Println("Stakeholder detail transactions: ", len(reportResponse.Transactions))
}

/*
The following test will fetch accounting rules.
*/
func Test_FetchFglAccountingRuleProfiles() {
    // Build report request
    reportRequest := FglAccountingRuleProfileReportRequest {
        // Set request params
        BaseReportRequest: baseReportRequest,                     // append base report request
        RuleId: "job_completed_order_amount_rule",          // fetch by accounting rule id
        RuleIdLike: "merchant",                                 // fetch by accounting rule like
    }
    // Call report
    reportResponse, error := FetchFglAccountingRuleProfile(clientConfig, reportRequest)
    // Handdle error (if any)
    if error != nil {
        fmt.Println("Error code: ", error.Code)
        fmt.Println("Error text: ", error.Text)
    }
    // Process report response
    fmt.Println("Accounting rule profiles: ", len(reportResponse.Profiles))
}

/*
The following test will fetch the promotion condition mappings 
*/
func Test_FetchPromotionConditionMapping() {
    // Build report request
    reportRequest := PromotionConditionMappingReportRequest {
        // Set request params
        BaseReportRequest: baseReportRequest,   // append base report request
        PromotionStatus: "A",                   // fetch by promotion status ("I" - inactive, "A" - active, "D" - don't know)
        //passengerId: 1097961,                 // fetch by passenger id
        StillActive: true,                      // fetch by active status (true - means current date and time has NOT reached date of expiry)
        // Optional request params
        // The array of entity field IDs that should be considered
        ConditionEntityFieldIdsIn: []int{6,},
        // The array of entity field IDs that should NOT be considered.
        ConditionEntityFieldIdsNotIn: []int{},
    }
    // Call report
    reportResponse, error := FetchPromotionConditionMapping(clientConfig, reportRequest)
    // Handdle error (if any)
    if error != nil {
        fmt.Println("Error code: ", error.Code)
        fmt.Println("Error text: ", error.Text)
    }
    // Process report response
    fmt.Println("Promotion mappings: ", len(reportResponse.Mappings))
}

func BuildCommonBaseReportRequest() BaseReportRequest {
    request := BaseReportRequest {
        FromStaff: false,         // indicate the request is originated from staff user
        FromPortal: false,        // indicate the request is originated from finance web portal
        PagingEnabled: true,      // enable or disable paging
        PageSize: 1,             // alter default page size 10
        PageIndex: 0,             // alter default page index 0
        ExportEnabled: false,     // enable or disable export
    }
    return request
}

/*================================ Run ====================================*/

func main() {

    // Frg integration
    //Test_FetchPeopleProfile()
    //Test_FetchDriverProfile()
    //Test_FetchDriverTripTransaction()
    //Test_FetchDriverTripSummary()

    // Fgl integration
    //Test_FetchFglStakeholderMonthlySummary()
    //Test_FetchFglStakeholderBalanceSummary()
    Test_FetchFglStakeholderTransactionDetail()
    //Test_FetchFglAccountingRuleProfiles()

    // Promo integration
    //Test_FetchPromotionConditionMapping()
}