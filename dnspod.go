package dnspod

import (
	"fmt"
	"net/http"
	"strings"
	"bytes"
	"strconv"
	"time"
	"sort"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"math/rand"
	"net/url"
	"encoding/json"
	"errors"
)

const (
	BaseUrl = "cns.api.qcloud.com"
	BaseUri = "/v2/index.php"
)

type Client struct {
	Region    string
	SecretId  string
	SecretKey string
	Http      *http.Client

	RecordsService *RecordsService
}

func NewClient(r, si, sk string) *Client {
	c := &Client{
		Region:    r,
		SecretId:  si,
		SecretKey: sk,
	}

	hc := &http.Client{
		Timeout: time.Second * 30,
	}

	c.Http = hc

	c.RecordsService = &RecordsService{
		ctx: c,
	}

	return c
}

func (c *Client) Post(params map[string]string, i IResponseWrapper) (*http.Response, error) {
	return c.Do("POST", params, i)

}

func (c *Client) Get(params map[string]string, i IResponseWrapper) (*http.Response, error) {
	return c.Do("GET", params, i)
}

func (c *Client) Do(method string, params map[string]string, i IResponseWrapper) (*http.Response, error) {
	var (
		err   error
		formV url.Values
	)

	signedQuery, signedParam := c.sign(strings.ToUpper(method), BaseUrl, BaseUri, params)

	api := "https://" + BaseUrl + "/" + BaseUri + "?" + signedQuery

	if strings.ToUpper(method) == "POST" {
		for k, v := range signedParam {
			formV.Add(k, v)
		}
	}
	req, err := http.NewRequest(method, api, strings.NewReader(formV.Encode()))
	if err != nil {
		return nil, err
	}

	res, err := c.Http.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(i)
	if err != nil {
		return nil, err
	}
	if i.GetCode() > 0 {
		return nil, errors.New(i.GetMessage())
	}

	return res, err
}

func (c *Client) sign(method, host, path string, data map[string]string) (string, map[string]string) {
	var (
		keys    []string
		toSign  string
		toQuery string
	)

	data["SecretId"] = c.SecretId
	data["Timestamp"] = strconv.FormatInt(time.Now().Unix(), 10)
	data["Nonce"] = ""
	data["SignatureMethod"] = "HmacSHA256"

	rand.Seed(time.Now().Unix())
	for i := 0; i < 6; i++ {
		data["Nonce"] = fmt.Sprintf("%s%d", data["Nonce"], rand.Intn(10))
	}

	for k := range data {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	signSPrefix := method + host + path + "?"
	for _, k := range keys {
		toSign = toSign + fmt.Sprintf("%s=%v&", k, data[k])
		toQuery = toQuery + fmt.Sprintf("%s=%v&", k, safeURLEncode(data[k]))
	}

	toSign = toSign[0:len(toSign)-1]
	toQuery = toQuery[0:len(toQuery)-1]

	ek := []byte(c.SecretKey)
	h := hmac.New(sha256.New, ek)
	h.Write([]byte(signSPrefix + toSign))
	sha := h.Sum(nil)

	signed := base64.StdEncoding.EncodeToString(sha)
	toQuery = toQuery + "&Signature=" + safeURLEncode(signed)

	data["Signature"] = safeURLEncode(signed)
	return toQuery, data
}

func safeURLEncode(s string) string {
	s = encodeURIComponent(s)
	s = strings.Replace(s, "!", "%21", -1)
	s = strings.Replace(s, "'", "%27", -1)
	s = strings.Replace(s, "(", "%28", -1)
	s = strings.Replace(s, ")", "%29", -1)
	s = strings.Replace(s, "*", "%2A", -1)
	return s
}

func encodeURIComponent(s string) string {
	var b bytes.Buffer
	written := 0

	for i, n := 0, len(s); i < n; i++ {
		c := s[i]

		switch c {
		case '-', '_', '.', '!', '~', '*', '\'', '(', ')':
			continue
		default:
			// Unreserved according to RFC 3986 sec 2.3
			if 'a' <= c && c <= 'z' {

				continue

			}
			if 'A' <= c && c <= 'Z' {

				continue

			}
			if '0' <= c && c <= '9' {

				continue
			}
		}

		b.WriteString(s[written:i])
		fmt.Fprintf(&b, "%%%02X", c)
		written = i + 1
	}

	if written == 0 {
		return s
	}
	b.WriteString(s[written:])
	return b.String()
}
