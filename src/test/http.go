package test

import (
	"io"
	"math/rand"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestClient struct {
	t *testing.T
	http.Client
	lastRedirect *string
}

func RandomPort() int {
	return 10_000 + rand.Intn(10_000)
}

func NewTestClient(t *testing.T) *TestClient {
	var lastRedirect string
	client := http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			pLastRedirect, _ := req.Response.Location()
			lastRedirect = pLastRedirect.Path
			return nil
		},
	}
	return &TestClient{t: t, Client: client, lastRedirect: &lastRedirect}
}

func (tc *TestClient) LastRedirect() string {
	return *tc.lastRedirect
}

func (tc *TestClient) ExpectResponseCode(r *http.Response, code int) {
	assert.Equal(tc.t, r.StatusCode, code)
}

func (tc *TestClient) ExpectRedirect(redirect string) {
	assert.Equal(tc.t, redirect, *tc.lastRedirect)
}

func (tc *TestClient) ExpectHeaderContains(r *http.Response, key string, value string) {
	values := r.Header.Values(key)
	for _, v := range values {
		assert.Contains(tc.t, v, value)
	}
}

func (tc *TestClient) ExpectBodyContains(r *http.Response, data []string) {
	bytes, _ := io.ReadAll(r.Body)
	body := string(bytes)
	for _, d := range data {
		assert.Contains(tc.t, body, d)
	}
}
