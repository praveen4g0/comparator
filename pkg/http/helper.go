package http

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"

	"github.com/praveen4g0/comparator/pkg/assert"
)

const (
	MaxIdleConnections int = 20
	RequestTimeout     int = 10
)

type Result struct {
	Url          []string
	JsonActual   string
	JsonExpected string
	Err          error
}

// CreateHTTPClient for connection re-use
func CreateHTTPClient() *http.Client {
	client := &http.Client{
		Transport: &http.Transport{
			MaxIdleConnsPerHost: MaxIdleConnections,
		},
		Timeout: time.Duration(RequestTimeout) * time.Second,
	}
	return client
}

func GetJsonResponses(url []string, c chan *Result, wg *sync.WaitGroup) {
	defer (*wg).Done()

	var jsonArray [2]string
	for i, v := range url {
		req, err := http.NewRequest("GET", v, nil)
		req.Header.Add("Accept", "application/json")
		assert.NoError(err)
		resp, err := CreateHTTPClient().Do(req)
		if err != nil {
			c <- &Result{Err: err, Url: url}
			return
		}

		defer resp.Body.Close()
		if resp.StatusCode == http.StatusUnauthorized || resp.StatusCode == http.StatusForbidden {
			c <- &Result{Err: fmt.Errorf("GET returned 401/403 response: %d", resp.StatusCode), Url: url}
			return
		}
		if resp.StatusCode != http.StatusOK {
			c <- &Result{Err: fmt.Errorf("status code not OK for %s", url), Url: url}
			return
		}
		resp_body, _ := ioutil.ReadAll(resp.Body)
		jsonArray[i] = string(resp_body)
	}
	res := &Result{JsonActual: jsonArray[0], JsonExpected: jsonArray[1], Url: url, Err: nil}
	c <- res
}
