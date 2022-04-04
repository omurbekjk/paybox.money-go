package paybox

type PaymentMethod string

const (
	PaymentMethodWallet         PaymentMethod = "wallet"
	PaymentMethodInternetBank   PaymentMethod = "internetbank"
	PaymentMethodOther          PaymentMethod = "other"
	PaymentMethodBankCard       PaymentMethod = "bankcard"
	PaymentMethodCach           PaymentMethod = "cash"
	PaymentMethodMobileCommerce PaymentMethod = "mobile_commerce"
)

type PaymentStatus string

const (
	PaymentStatusOK       PaymentStatus = "ok"
	PaymentStatusRejected PaymentStatus = "rejected"
)

// GeneratePaymentRequest -
// swagger:model
type GeneratePaymentRequest struct {
	PgMerchantId       string `json:"pg_merchant_id"`
	PgOrderId          string `json:"pg_order_id"`
	PgAmount           int    `json:"pg_amount"`
	PgCurrency         string `json:"pg_currency"`
	PgDescription      string `json:"pg_description"`
	PgTestingMode      string `json:"pg_testing_mode"`
	PgResultUrl        string `json:"pg_result_url"`
	PgSuccessUrl       string `json:"pg_success_url"`
	PgFailureUrl       string `json:"pg_failure_url"`
	PgSiteUrl          string `json:"pg_site_url"`
	PgRequestMethod    string `json:"pg_request_method"`
	PgUserContactEmail string `json:"pg_user_contact_email"`
	PgUserContactPhone string `json:"pg_user_phone"`
	PgSig              string `json:"pg_sig"`
	PgSalt             string `json:"pg_salt"`
}

// GeneratePaymentResponse
// swagger:model
type GeneratePaymentResponse struct {
	PgStatus           *string `xml:"pg_status,omitempty"`
	PgPaymentId        *int    `xml:"pg_payment_id,omitempty"`
	PgRedirectUrl      *string `xml:"pg_redirect_url,omitempty"`
	PgRedirectUrlType  *string `xml:"pg_redirect_url_type,omitempty"`
	PgSalt             *string `xml:"pg_salt,omitempty"`
	PgSig              *string `xml:"pg_sig,omitempty"`
	PgErrorCode        *string `xml:"pg_error_code,omitempty"`
	PgErrorDescription *string `xml:"pg_error_description,omitempty"`
}

// PaymentResultRequest
// swagger:model
type PaymentResultRequest struct {
	PgOrderId          *string        `form:"pg_order_id,omitempty"`
	PgPaymentId        *int           `form:"pg_payment_id,omitempty"`
	PgAmount           *string        `form:"pg_amount,omitempty"`
	PgCurrency         *string        `form:"pg_currency,omitempty"`
	PgNetAmount        *string        `form:"pg_net_amount,omitempty"`
	PgPsAmount         *string        `form:"pg_ps_amount,omitempty"`
	PgPsFullAmount     *string        `form:"pg_ps_full_amount,omitempty"`
	PgPsCurrency       *string        `form:"pg_ps_currency,omitempty"`
	PgDescription      *string        `form:"pg_description,omitempty"`
	PgResult           *int           `form:"pg_result,omitempty"`
	PgPaymentDate      *string        `form:"pg_payment_date,omitempty"`
	PgCanReject        *int           `form:"pg_can_reject,omitempty"`
	PgUserPhone        *string        `form:"pg_user_phone,omitempty"`
	PgUserContactEmail *string        `form:"pg_user_contact_email,omitempty"`
	PgTestingMode      *int           `form:"pg_testing_mode,omitempty"`
	PgCaptured         *int           `form:"pg_captured,omitempty"`
	PgCardId           *string        `form:"pg_card_id,omitempty"`
	PgCardPan          *string        `form:"pg_card_pan,omitempty"`
	PgSalt             *string        `form:"pg_salt,omitempty"`
	PgSig              *string        `form:"pg_sig,omitempty"`
	PgDiscountPercent  *string        `form:"pg_discount_percent,omitempty"`
	PgDiscountAmount   *string        `form:"pg_discount_amount,omitempty"`
	PgPaymentMethod    *PaymentMethod `form:"pg_payment_method,omitempty"`
}

// PaymentResultResponse -
// swagger:model
type PaymentResultResponse struct {
	PgStatus      *PaymentStatus `xml:"pg_status,omitempty"`
	PgDescription *string        `xml:"pg_description,omitempty"`
	PgSalt        *string        `xml:"pg_salt,omitempty"`
	PgSig         *string        `xml:"pg_sig,omitempty"`
}
