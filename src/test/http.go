package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestHttpClient[Res any] struct {
	t *testing.T
	http.Client
	baseUrl      string
	lastRedirect *string
}

type requestFunc = func() (*http.Response, error)

type ResponseRef = http.Response

type ExpectableResponse[T any] struct {
	http.Response
	t *testing.T
}

func RandomPort() int {
	return 10_000 + rand.Intn(10_000)
}

func NewHttpClient[Res any](t *testing.T, serverPort int) *TestHttpClient[Res] {
	var lastRedirect string
	client := http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			pLastRedirect, _ := req.Response.Location()
			lastRedirect = pLastRedirect.Path
			return nil
		},
	}
	return &TestHttpClient[Res]{t: t, Client: client, baseUrl: fmt.Sprintf("http://localhost:%d", serverPort),
		lastRedirect: &lastRedirect}
}

func (c *TestHttpClient[Res]) Get(url string) ExpectableResponse[Res] {
	return c.mustResponse(func() (*http.Response, error) { return c.Client.Get(c.fullUrl(url)) })
}

func (c *TestHttpClient[Res]) PostJson(url string, request any) ExpectableResponse[Res] {
	return c.mustResponse(func() (*http.Response, error) {
		var buff bytes.Buffer
		json.NewEncoder(&buff).Encode(request)
		return c.Client.Post(c.fullUrl(url), "application/json", &buff)
	})
}

func (c *TestHttpClient[Res]) mustResponse(req requestFunc) ExpectableResponse[Res] {
	res, err := req()
	if err != nil {
		panic(err)
	}
	return NewExpectableResponse[Res](*res, c.t)
}

func (c *TestHttpClient[Res]) fullUrl(url string) string {
	return c.baseUrl + "/" + url
}

// func (tc *TestHttpClient) ExpectHeaderContains(r *http.Response, key string, value string) {
// 	values := r.Header.Values(key)
// 	for _, v := range values {
// 		assert.Contains(tc.t, v, value)
// 	}
// }

// func (tc *TestHttpClient) ExpectBodyContains(r *http.Response, data []string) {
// 	bytes, _ := io.ReadAll(r.Body)
// 	body := string(bytes)
// 	for _, d := range data {
// 		assert.Contains(tc.t, body, d)
// 	}
// }

// func (tc *TestHttpClient) ExpectJson(r *http.Response, expected any) {
// 	fmt.Println(reflect.ValueOf(expected))

// 	actual := reflect.New(reflect.TypeOf(expected)).Interface()

// 	// var actual any
// 	err := json.NewDecoder(r.Body).Decode(actual)
// 	if err != nil {
// 		panic(err)
// 	}

// 	assert.Equal(tc.t, r.Header.Get("Content-Type"), "application/json")
// 	assert.Equal(tc.t, expected, actual)
// }

func NewExpectableResponse[T any](r http.Response, t *testing.T) ExpectableResponse[T] {
	return ExpectableResponse[T]{r, t}
}

func (r *ExpectableResponse[T]) ExpectStatusCode(statusCode int) {
	assert.Equal(r.t, r.StatusCode, statusCode)
}

func (r *ExpectableResponse[T]) ExpectJson(expected T) {
	var actual T

	// var actual any
	err := json.NewDecoder(r.Body).Decode(&actual)
	if err != nil {
		panic(err)
	}

	assert.Equal(r.t, r.Header.Get("Content-Type"), "application/json")
	assert.Equal(r.t, expected, actual)
}
