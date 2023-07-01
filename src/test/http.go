package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestHttpClient struct {
	t *testing.T
	http.Client
	baseUrl      string
	lastRedirect *string
}

type requestFunc = func() (*http.Response, error)

func RandomPort() int {
	return 10_000 + rand.Intn(10_000)
}

func NewHttpClient(t *testing.T, serverPort int) *TestHttpClient {
	var lastRedirect string
	client := http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			pLastRedirect, _ := req.Response.Location()
			lastRedirect = pLastRedirect.Path
			return nil
		},
	}
	return &TestHttpClient{t: t, Client: client, baseUrl: fmt.Sprintf("http://localhost:%d", serverPort),
		lastRedirect: &lastRedirect}
}

func (c *TestHttpClient) Get(url string) *http.Response {
	return c.mustResponse(func() (*http.Response, error) { return c.Client.Get(c.fullUrl(url)) })
}

func (c *TestHttpClient) PostJson(url string, request any) *http.Response {
	return c.mustResponse(func() (*http.Response, error) {
		var buff bytes.Buffer
		json.NewEncoder(&buff).Encode(request)
		return c.Client.Post(c.fullUrl(url), "application/json", &buff)
	})
}

func (c *TestHttpClient) mustResponse(req requestFunc) *http.Response {
	res, err := req()
	if err != nil {
		panic(err)
	}
	return res
}

func (c *TestHttpClient) fullUrl(url string) string {
	return c.baseUrl + "/" + url
}

func (tc *TestHttpClient) LastRedirect() string {
	return *tc.lastRedirect
}

func (tc *TestHttpClient) ExpectResponseCode(r *http.Response, code int) {
	assert.Equal(tc.t, r.StatusCode, code)
}

func (tc *TestHttpClient) ExpectRedirect(redirect string) {
	assert.Equal(tc.t, redirect, *tc.lastRedirect)
}

func (tc *TestHttpClient) ExpectHeaderContains(r *http.Response, key string, value string) {
	values := r.Header.Values(key)
	for _, v := range values {
		assert.Contains(tc.t, v, value)
	}
}

func (tc *TestHttpClient) ExpectBodyContains(r *http.Response, data []string) {
	bytes, _ := io.ReadAll(r.Body)
	body := string(bytes)
	for _, d := range data {
		assert.Contains(tc.t, body, d)
	}
}
