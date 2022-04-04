# Go client for paybox.money

## Install

```
go get github.com/omurbekjk/paybox.money-go
go mod download
go mod vendor
```

## Usage

```go
client, err := NewClient("https://api.paybox.money", "merchantId", "merchantSecret", "test")
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
```

