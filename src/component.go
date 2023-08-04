/*=====================Finance Go Gateway Client===============================

    

    Author(s):
        Dilanka Gamage

==============================================================================*/

package main

type Error struct {
    Code string `json:"code"`
    Text string `json:"text"`
}

type Sorter struct {
    Field string `json:"field"`
    DESC bool    `json:"DESC"`
}

type ReportIndex struct {
    Name string `json:"name"`
    Type string `json:"type"`
}

type PeopleProfile struct {
    PeopleId int64                          `json:"peopleId"`
    Salutation string                       `json:"salutation"`
    Name string                             `json:"name"`
    Lastname string                         `json:"lastname"`
    Email string                            `json:"email"`
    Address string                          `json:"address"`
    PermAddress string                      `json:"permAddress"`
    PaypalAccount string                    `json:"paypalAccount"`
    Password string                         `json:"password"`
    HashedPassword string                   `json:"hashedPassword"`
    PasswordHashedBc string                 `json:"passwordHashedBc"`
    OrgPassword string                      `json:"orgPassword"`
    Username string                         `json:"username"`
    Location string                         `json:"location"`
    Photo string                            `json:"photo"`
    DeviceId string                         `json:"deviceId"`
    DeviceToken string                      `json:"deviceToken"`
    DeviceType int                          `json:"deviceType"`
    PushVersion int                         `json:"pushVersion"`
    DeviceImei string                       `json:"deviceImei"`
    CreatedDate string                      `json:"createdDate"`
    UpdatedDate string                      `json:"updatedDate"`
    LastLogin string                        `json:"lastLogin"`
    LoginIp string                          `json:"loginIp"`
    LoginCity string                        `json:"loginCity"`
    LoginState string                       `json:"loginState"`
    LoginCountry string                     `json:"loginCountry"`
    LoginType int                           `json:"loginType"`
    LoginStatus string                      `json:"loginStatus"`
    LoginFrom string                        `json:"loginFrom"`
    UserType string                         `json:"userType"`
    AccountType int                         `json:"accountType"`
    Timezone string                         `json:"timezone"`
    Description string                      `json:"description"`
    Phone string                            `json:"phone"`
    ReachableMobile string                  `json:"reachableMobile"`
    SecondaryMobile string                  `json:"secondaryMobile"`
    EmegencyMobile string                   `json:"emegencyMobile"`
    Gender string                           `json:"gender"`
    Dob string                              `json:"dob"`
    Status string                           `json:"status"`
    BlockReason string                      `json:"blockReason"`
    UserCreatedby int                       `json:"userCreatedby"`
    CompanyId int                           `json:"companyId"`
    DriverLicenseId string                  `json:"driverLicenseId"`
    ProfilePicture string                   `json:"profilePicture"`
    AvailabilityStatus string               `json:"availabilityStatus"`
    AccountBalance float64                  `json:"accountBalance"`
    BookingLimit int                        `json:"bookingLimit"`
    DriverShare float64                     `json:"driverShare"`
    NotificationSetting int                 `json:"notificationSetting"`
    Otp string                              `json:"otp"`
    DriverReferralCode string               `json:"driverReferralCode"`
    NotificationStatus string               `json:"notificationStatus"`
    UserUpdatedby interface{}               `json:"userUpdatedby"`
    ReachableMobileBak string               `json:"reachableMobileBak"`
    OnDemand int                            `json:"onDemand"`
    DriverTripSummaries []DriverTripSummary `json:"driverTripSummaries"`
}

type DriverProfile struct {
    DriverId int64                          `json:"driverId"`
    Latitude float64                        `json:"latitude"`
    Longitude float64                       `json:"longitude"`
    Status string                           `json:"status"`
    CurrentStatus string                    `json:"currentStatus"`
    ShiftStatus string                      `json:"shiftStatus"`
    TravelStatus string                     `json:"travelStatus"`
    DriverApi string                        `json:"driverApi"`
    ShiftStartedAt string                   `json:"shiftStartedAt"`
    UpdateDate string                       `json:"updateDate"`
    CurrentTrip interface{}                 `json:"currentTrip"`
    TripAssignedAt string                   `json:"tripAssignedAt"`
    ProfileStatus string                    `json:"profileStatus"`
    VehicleModel interface{}                `json:"vehicleModel"`
    LoginStatus interface{}                 `json:"loginStatus"`
    DirectionalHireEnable bool              `json:"directionalHireEnable"`
    Bearing float64                         `json:"bearing"`
    DriverTripSummaries []DriverTripSummary `json:"driverTripSummaries"`
}

type DriverTripTransaction struct {
    TransactionId int64                     `json:"transactionId"`
    DriverId int64                          `json:"driverId"`
    DriverProfile DriverProfile             `json:"driverProfile"`
    TripId int                              `json:"tripId"`
    TransactionType string                  `json:"transactionType"`
    TransactionCategory int                 `json:"transactionCategory"`
    AmountInRupee float64                   `json:"amountInRupee"`
    AmountInCents int64                     `json:"amountInCents"`
    Description string                      `json:"description"`
    CreatedTime string                      `json:"createdTime"`
    CreatedDate string                      `json:"createdDate"`
    CreatedBy int                           `json:"createdBy"`
}

type DriverTripSummary struct {
    RowId int                               `json:"rowId"`
    DriverId int64                          `json:"driverId"`
    DriverProfile DriverProfile             `json:"driverProfile"`
    TransactionType string                  `json:"transactionType"`
    TotalAmountInCents int64                `json:"totalAmountInCents"`
    RowCount int                            `json:"rowCount"`
}

type FglStakeholderMonthlySummary struct {
    AccountingPeriod string                 `json:"accountingPeriod"`
    AccountId string                        `json:"accountId"`
    StakeholderType string                  `json:"stakeholderType"`
    CurrencyType string                     `json:"currencyType"`
    AccountType string                      `json:"accountType"`
    AccountReference string                 `json:"accountReference"`
    TotalDebit float64                      `json:"totalDebit"`
    DebitCount int                          `json:"debitCount"`
    TotalCredit float64                     `json:"totalCredit"`
    CreditCount int                         `json:"creditCount"`
    AccountBalance float64                  `json:"accountBalance"`
}

type FglStakeholderBalanceSummary struct {
    AccountingPeriod string                 `json:"accountingPeriod"`
    AccountId string                        `json:"accountId"`
    StakeholderType string                  `json:"stakeholderType"`
    CurrencyType string                     `json:"currencyType"`
    AccountType string                      `json:"accountType"`
    AccountReference string                 `json:"accountReference"`
    AccountBalance float64                  `json:"accountBalance"`
}

type FglStakeholderTransactionDetail struct {
    JournalId int64                         `json:"journalId"`
    JournalEventId string                   `json:"journalEventId"`
    JournalEventLogId string                `json:"journalEventLogId"`
    JournalAccountingRuleId string          `json:"journalAccountingRuleId"`
    JournalReferenceId string               `json:"journalReferenceId"`
    JournalGrossAmount float64              `json:"journalGrossAmount"`
    PostingJournalId int64                  `json:"postingJournalId"`
    PostingAccountingRuleId string          `json:"postingAccountingRuleId"`
    PostingReferenceId string               `json:"postingReferenceId"`
    PostingAccountId string                 `json:"postingAccountId"`
    PostingAccountingPeriod string          `json:"postingAccountingPeriod"`
    PostingAccountingEntry string           `json:"postingAccountingEntry"`
    PostingCurrencyTypeId string            `json:"postingCurrencyTypeId"`
    PostingAmount float64                   `json:"postingAmount"`
    AccountStakeholderId string             `json:"accountStakeholderId"`
    AccountCurrencyTypeId string            `json:"accountCurrencyTypeId"`
    AccountTypeId string                    `json:"accountTypeId"`
    AccountReferenceId string               `json:"accountReferenceId"`
    AccountingRuleId string                 `json:"accountingRuleId"`
    AccountingRuleDescription string        `json:"accountingRuleDescription"`
    EventId string                          `json:"eventId"`
    EventReferenceParam string              `json:"eventReferenceParam"`
    CreatedDate string                      `json:"createdDate"`
    ModifiedDate string                     `json:"modifiedDate"`
    CreatedBy string                        `json:"createdBy"`
    ModifiedBy string                       `json:"modifiedBy"`
    Version string                          `json:"version"`
    RoutingKey string                       `json:"routingKey"`
}

type FglAccountingRuleProfile struct {
    RuleId string                           `json:"ruleId"`
    EventId string                          `json:"eventId"`
    RelationalLogicId string                `json:"relationalLogicId"`
    ArithmeticLogicId string                `json:"arithmeticLogicId"`
    Description string                      `json:"description"`
    TransactionCategoryId int32             `json:"transactionCategoryId"`
    Active bool                             `json:"active"`
    StartingDate string                     `json:"startingDate"`
    ClosingDate string                      `json:"closingDate"`
    CreatedDate string                      `json:"createdDate"`
    ModifiedDate string                     `json:"modifiedDate"`
    CreatedBy string                        `json:"createdBy"`
    ModifiedBy string                       `json:"modifiedBy"`
    Version string                          `json:"version"`
    RoutingKey string                       `json:"routingKey"`
}

type PromotionConditionMapping struct {
    PromotionId int64                       `json:"promotionId"`
    PromotionActiveFrom string              `json:"promotionActiveFrom"`
    PromotionActiveUntil string             `json:"promotionActiveUntil"`
    PromotionCode string                    `json:"promotionCode"`
    PromotionCreatedAt string               `json:"promotionCreatedAt"`
    PromotionDiscountAmount float64         `json:"promotionDiscountAmount"`
    PromotionDiscountType string            `json:"promotionDiscountType"`
    PromotionName string                    `json:"promotionName"`
    promotionType string                    `json:"promotionType"`
    promotionRuleFile string                `json:"promotionRuleFile"`
    promotionStatus string                  `json:"promotionStatus"`
    promotionVersion int32                  `json:"promotionVersion"`
    promotionServiceType string             `json:"promotionServiceType"`
    PromotionDescription string             `json:"promotionDescription"`
    PromotionDisplayName string             `json:"promotionDisplayName"`
    RoutingKey string                       `json:"routingKey"`
}

func BuildSorter(field string, desc bool) Sorter {
    sorter := Sorter {
        Field: field,
        DESC: desc,
    }
    return sorter
}