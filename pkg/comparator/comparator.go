package comparator

import (
	"log"
	"sync"

	"github.com/getgauge-contrib/gauge-go/gauge"
	"github.com/praveen4g0/comparator/pkg/assert"
	"github.com/praveen4g0/comparator/pkg/files"
	"github.com/praveen4g0/comparator/pkg/http"
	. "github.com/praveen4g0/comparator/pkg/matcher"
)

func Compare(dir1, dir2 string) {
	var err error
	furls, err := files.ReadBytes(dir1)
	assert.NoError(err)
	surls, err := files.ReadBytes(dir2)
	assert.NoError(err)
	furl_list := files.SplitLinesBySperator(string(furls), "\n")
	surl_list := files.SplitLinesBySperator(string(surls), "\n")

	min, max := getMinMax(furl_list, surl_list)

	c := make(chan *http.Result)
	var wg sync.WaitGroup

	for index, value := range min {
		wg.Add(1)
		go http.GetJsonResponses([]string{max[index], value}, c, &wg)
	}

	go func() {
		wg.Wait()
		close(c)
	}()

	for res := range c {
		if res.Err != nil {
			gauge.WriteMessage(res.Url[0] + " or " + res.Url[1] + " are invalid Reason: " + res.Err.Error())
			log.Println(res.Url[0] + " or " + res.Url[1] + " are invalid Reason: " + res.Err.Error())
		} else {
			success, err := (&JSONMatcher{JSONToMatch: res.JsonExpected}).Match(res.JsonActual)
			assert.NoError(err)
			if success {
				log.Println(res.Url[0] + " equals " + res.Url[1])
				gauge.WriteMessage(res.Url[0] + " equals " + res.Url[1])
			} else {
				log.Println(res.Url[0] + " not equals " + res.Url[1])
				gauge.WriteMessage(res.Url[0] + " not equals " + res.Url[1])
			}
		}
	}
}

func getMinMax(a, b []string) ([]string, []string) {
	if len(a) < len(b) {
		return a, b
	}
	return b, a
}
