package proxy

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net"
	"net/http"
	"net/textproto"
	"time"

	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

const (
	MaxIdleConns        int = 40000
	MaxIdleConnsPerHost int = 1000
	IdleConnTimeout     int = 30
	DialTimeout             = 2
)

const (
	HeaderKscAccount   = "X-Ksc-Account-Id"
	HeaderKscUser      = "X-Ksc-User-Id"
	HeaderKscUserUUID  = "X-Ksc-User-Uuid"
	HeaderKscRequestID = "X-Ksc-Request-Id"
)

const (
	StatusSuccess = 0
	StatusFail    = 1

	StatusSuccessString = "Success"
)

//SetRequestHeader Set Header Information
func SetRequestHeader(req *http.Request) {
	req.Header = http.Header{
		textproto.CanonicalMIMEHeaderKey("Content-Type"): []string{"application/json"},
		textproto.CanonicalMIMEHeaderKey("Accept"):       []string{"application/json"},
	}
}

//NewWebGetRequest
func Get(url string, header map[string]string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	//Set header
	SetRequestHeader(req)
	if len(header) > 0 {
		for k, v := range header {
			req.Header.Set(k, v)
		}
	}

	return NewServiceClient(req)
}

//Post
func Post(url string, header map[string]string, data interface{}) ([]byte, error) {
	dataJSON, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(dataJSON))
	if err != nil {
		return nil, err
	}
	//Set header
	SetRequestHeader(req)
	if len(header) > 0 {
		for k, v := range header {
			req.Header.Set(k, v)
		}
	}
	return NewServiceClient(req)
}

func NewServiceClient(req *http.Request) ([]byte, error) {
	c := &http.Client{
		Transport: &http.Transport{
			// DialContext: (&net.Dialer{
			// 	Timeout:   time.Duration(DialTimeout) * time.Second,
			// 	KeepAlive: 2 * time.Second,
			// }).DialContext,
			Dial: func(network, addr string) (net.Conn, error) {
				conn, err := net.DialTimeout(network, addr, time.Duration(DialTimeout)*time.Second)
				if err != nil {
					return nil, err
				}

				_ = conn.SetDeadline(time.Now().Add(time.Second * 2))
				return conn, nil
			},
			MaxIdleConns:          MaxIdleConns,
			MaxIdleConnsPerHost:   MaxIdleConnsPerHost,
			IdleConnTimeout:       time.Duration(IdleConnTimeout) * time.Second,
			DisableKeepAlives:     false,
			ResponseHeaderTimeout: 10 * time.Second,
		},
		//包括从连接(Dial)到读完response body。
		// Timeout: time.Duration(config.G.ClientHttp.TimeOut) * time.Second,
		Timeout: 10 * time.Second,
	}

	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	bodyStr := string(body)
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(bodyStr)
	}
	return body, nil
}
