package paybox

import (
	"fmt"
	"testing"
)

func TestClient(t *testing.T) {
	_, err := NewClient("", "", "", "")
	if err == nil {
		t.Fatalf("expected to fail")
	}
}

func TestClientGeneratePayment(t *testing.T) {
	client, err := NewClient("https://api.paybox.money/", "merchantId", "merchantSecretKey", "test")
	if err != nil {
		t.Fatalf("unexpected error %s", err.Error())
	}
	paymentRequest := &GeneratePaymentRequest{
		PgOrderId:       "isbn-1234-5678-90",
		PgAmount:        2022,
		PgCurrency:      "KGS",
		PgDescription:   "order to buy a book",
		PgTestingMode:   "1",
		PgRequestMethod: "POST",
	}
	response, err := client.GeneratePayment(paymentRequest)
	if err != nil {
		t.Fatalf("unexpected error %s", err.Error())
	}
	fmt.Println(*response.PgRedirectUrl)
}

func TestClientGetPaymentStatus(t *testing.T) {
	client, err := NewClient("https://api.paybox.money/", "merchantId", "merchantSecretKey", "test")
	if err != nil {
		t.Fatalf("unexpected error %s", err.Error())
	}
	paymentRequest := &PaymentStatusRequest{
		PgPaymentId: 0,
		PgOrderId:   "Shuma",
		PgSalt:      "",
		PgSig:       "",
	}
	response, err := client.GetPaymentStatus(paymentRequest)
	if err != nil {
		t.Fatalf("unexpected error %s", err.Error())
	}
	fmt.Println(*response)
}
