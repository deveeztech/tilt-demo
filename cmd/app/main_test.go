package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample(t *testing.T) {

	// given
	ts := httptest.NewServer(http.HandlerFunc(HelloHandler))
	defer ts.Close()

	// when
	body, err := call(ts.URL)

	// then
	assert.NoError(t, err)
	assert.Equal(t, "HELLO WORLD!!\n", string(body))

}

func call(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return []byte{}, err
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()

	return body, err
}
