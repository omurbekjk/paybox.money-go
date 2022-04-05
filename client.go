package paybox

import (
	"bytes"
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"time"
)

const (
	EndpointInitScriptName   = "init_payment.php"
	EndpointStatusScriptName = "get_status.php"
	SaltLength               = 32
)

type Client interface {
	GeneratePayment(request *GeneratePaymentRequest) (*GeneratePaymentResponse, error)
	GetPaymentStatus(request *PaymentStatusRequest) (*PaymentStatusResponse, error)
}

type client struct {
	httpClient           *http.Client
	apiBaseURL           string
	merchantID           string
	merchantSecret       string
	merchantSecretPayout string
}

// NewClient returns new Client
// APIBase is a base API URL
func NewClient(APIBaseURL string, merchantID string, merchantSecret string, merchantSecretPayout string) (Client, error) {
	if APIBaseURL == "" || merchantID == "" || merchantSecret == "" {
		return nil, errors.New("APIBaseURL, merchantID and merchantSecret are required to create a client")
	}

	c := client{
		httpClient: &http.Client{
			Timeout: time.Second * 30,
		},
		apiBaseURL:           strings.TrimRight(APIBaseURL, "/"),
		merchantID:           merchantID,
		merchantSecret:       merchantSecret,
		merchantSecretPayout: merchantSecretPayout,
	}
	return &c, nil
}

func (c *client) GeneratePayment(request *GeneratePaymentRequest) (*GeneratePaymentResponse, error) {
	request.PgMerchantId = c.merchantID
	randomStr, err := GenerateRandomString(SaltLength)
	if err != nil {
		return nil, err
	}
	request.PgSalt = randomStr
	request.PgSig = c.generateSignature(*request, EndpointInitScriptName)

	r, _ := json.Marshal(request)
	var response GeneratePaymentResponse
	err = c.performRequest("POST", "/"+EndpointInitScriptName, bytes.NewReader(r), &response)
	return &response, err
}

func (c *client) GetPaymentStatus(request *PaymentStatusRequest) (*PaymentStatusResponse, error) {
	num, _ := strconv.Atoi(c.merchantID)
	request.PgMerchantId = num
	randomStr, err := GenerateRandomString(SaltLength)
	if err != nil {
		return nil, err
	}
	request.PgSalt = randomStr
	request.PgSig = c.generateSignature(*request, EndpointStatusScriptName)
	r, _ := json.Marshal(request)
	var response PaymentStatusResponse
	err = c.performRequest("POST", "/"+EndpointStatusScriptName, bytes.NewReader(r), &response)
	return &response, err
}

// performRequest performs the specified request, based on the method, path and body.
// path may contain query string parameters.
// This stores the response of the request into the value pointed to by v.
func (c *client) performRequest(method, path string, requestBody io.Reader, v interface{}) error {
	req, err := http.NewRequest(method, c.payboxEndpoint(path), requestBody)
	if err != nil {
		return err
	}

	header := http.Header{}
	header.Add("Content-Type", "application/json")
	req.Header = header

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	// Unsuccessful response from paybox
	if resp.StatusCode != http.StatusOK {
		return err
	}
	// If we've gotten here, we have a successful response from CSP
	if v != nil && len(body) > 0 {
		err = xml.Unmarshal(body, &v)
	}
	return err
}

// payboxEndpoint generates a paybox endpoint by appending the specified path to the provided hostname for the client
// path may contain query string parameters.
func (c *client) payboxEndpoint(path string) string {
	return fmt.Sprintf("%s%s", c.apiBaseURL, path)
}

func (c *client) generateSignature(request interface{}, scriptname string) string {
	signature := []string{scriptname}

	keys := sortedKeysArray(request)
	ref := reflect.ValueOf(request)
	for _, key := range keys {
		if key != "PgSig" {
			value := reflect.Indirect(ref).FieldByName(key)
			if value.Kind() == reflect.Int {
				signature = append(signature, strconv.Itoa(int(value.Int())))
			} else {
				if value.String() != "" {
					signature = append(signature, value.String())
				}
			}
		}
	}

	signatureStr := strings.Join(signature[:], ";") + ";" + c.merchantSecret
	hasher := md5.New()
	hasher.Write([]byte(signatureStr))
	return hex.EncodeToString(hasher.Sum(nil))
}

func sortedKeysArray(data interface{}) []string {
	v := reflect.ValueOf(data)
	typeOfS := v.Type()
	var keys []string
	for i := 0; i < v.NumField(); i++ {
		keys = append(keys, typeOfS.Field(i).Name)
	}
	sort.Strings(keys)
	return keys
}

// GenerateRandomBytes returns securely generated random bytes.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}

	return b, nil
}

// GenerateRandomString returns a URL-safe, base64 encoded
// securely generated random string.
func GenerateRandomString(s int) (string, error) {
	b, err := GenerateRandomBytes(s)
	return base64.URLEncoding.EncodeToString(b), err
}
