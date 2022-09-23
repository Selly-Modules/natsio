package jsmodel

// PushNotification ...
type PushNotification struct {
	User        string              `json:"user"`
	Type        string              `json:"type"`
	TargetID    string              `json:"targetId"`
	IsFromAdmin bool                `json:"isFromAdmin"`
	Category    string              `json:"category"`
	Options     NotificationOptions `json:"options"`
}

// NotificationOptions ...
type NotificationOptions struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

// PayloadUpdateSellerAffiliateStatistic ...
type PayloadUpdateSellerAffiliateStatistic struct {
	SellerID  string                   `json:"sellerId"`
	Statistic SellerAffiliateStatistic `json:"statistic"`
}

// SellerAffiliateStatistic ...
type SellerAffiliateStatistic struct {
	TransactionTotal              int     `json:"transactionTotal"`
	TransactionCashback           int     `json:"transactionCashback"`
	TransactionPending            int     `json:"transactionPending"`
	TransactionApproved           int     `json:"transactionApproved"`
	TransactionRejected           int     `json:"transactionRejected"`
	CommissionTransactionTotal    float64 `json:"commissionTransactionTotal"`
	CommissionTransactionCashback float64 `json:"commissionTransactionCashback"`
	CommissionTransactionApproved float64 `json:"commissionTransactionApproved"`
	CommissionTransactionRejected float64 `json:"commissionTransactionRejected"`
}

// PayloadCashflowsBySeller ...
type PayloadCashflowsBySeller struct {
	SellerID string           `json:"sellerId"`
	List     []CashflowSeller `json:"list"`
}

// CashflowSeller ...
type CashflowSeller struct {
	Value      float64          `json:"value"`
	Action     string           `json:"action"`
	Category   string           `json:"category"`
	TargetID   string           `json:"targetId"`
	TargetType string           `json:"targetType"`
	Options    *CashFlowOptions `json:"options"`
}

// CashFlowOptions ...
type CashFlowOptions struct {
	AffiliateTransactionCode string `json:"affiliateTransactionCode,omitempty"`
	AffiliateCampaignID      string `json:"affiliateCampaignId,omitempty"`
	AffiliateCampaignName    string `json:"affiliateCampaignName,omitempty"`
}
