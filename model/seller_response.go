package model

import "time"

// ResponseSellerInfo ...
type ResponseSellerInfo struct {
	ID   string `json:"_id"`
	Name string `json:"name"`
	Code string `json:"code"`
}

// ResponseListSellerInfo ...
type ResponseListSellerInfo struct {
	Sellers []ResponseSellerInfo `json:"sellers"`
}

// ResponseListSellerInfoSupportChat ...
type ResponseListSellerInfoSupportChat struct {
	Sellers []ResponseSellerInfoSupportChat `json:"sellers"`
}

// ResponseSellerInfoSupportChat ...
type ResponseSellerInfoSupportChat struct {
	ID           string                 `json:"_id"`
	Name         string                 `json:"name"`
	Code         string                 `json:"code"`
	Membership   SellerMembershipInfo   `json:"membership"`
	Info         SellerContactInfo      `json:"info"`
	Team         *TeamInfo              `json:"team,omitempty"`
	Statistic    SellerStatistic        `json:"statistic"`
	TrackingTime *SellerTrackingTime    `json:"trackingTime"`
	Invitee      *InviteeInfo           `json:"invitee"`
	CreatedAt    time.Time              `json:"createdAt"`
	PlanPackage  *SellerPlanPackageInfo `json:"planPackage"`
}

// SellerPlanPackageInfo ...
type SellerPlanPackageInfo struct {
	ID        string    `json:"_id"`
	Name      string    `json:"name"`
	Level     int       `json:"level"`
	CreatedAt time.Time `json:"createdAt"`
}

// SellerTrackingTime ...
type SellerTrackingTime struct {
	FirstOrderDeliveredAt time.Time `json:"firstOrderDeliveredAt,omitempty"`
	ThirdOrderDeliveredAt time.Time `json:"thirdOrderDeliveredAt,omitempty"`
}

// SellerStatistic ...
type SellerStatistic struct {
	ThisMonthSale                float64 `bson:"thisMonthSale" json:"thisMonthSale"`
	LastMonthSale                float64 `bson:"lastMonthSale" json:"lastMonthSale"`
	Sale                         float64 `bson:"sale" json:"sale"`
	TransactionTotal             int     `json:"transactionTotal"`
	TransactionPaymentProcessing int     `json:"transactionPaymentProcessing"`
	TransactionWaitingApprove    int     `json:"transactionWaitingApprove"`
	TransactionPending           int     `json:"transactionPending"`
	TransactionSuccess           int     `json:"transactionSuccess"`
	TransactionRejected          int     `json:"transactionRejected"`
	TransactionDelivering        int     `json:"transactionDelivering"`
	TransactionDelivered         int     `json:"transactionDelivered"`
}

// TeamInfo ...
type TeamInfo struct {
	ID   string `json:"_id"`
	Name string `json:"name"`
	Role string `json:"role"`
}

// InviteeInfo ...
type InviteeInfo struct {
	ID   string `json:"_id"`
	Name string `json:"name"`
}

// SellerContactInfo ...
type SellerContactInfo struct {
	City     int    `json:"cityCode"`
	CityName string `json:"cityName"`
	Gender   string `json:"gender"`
}

// SellerMembershipInfo ...
type SellerMembershipInfo struct {
	Level int    `json:"level"`
	Name  string `json:"name"`
}
