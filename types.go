package paybox

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
