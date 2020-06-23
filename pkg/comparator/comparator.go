package comparator

import (
	"io/ioutil"
	"log"
	"strings"
	"sync"

	"github.com/getgauge-contrib/gauge-go/gauge"
	"github.com/praveen4g0/comparator/pkg/assert"
	"github.com/praveen4g0/comparator/pkg/files"
	"github.com/praveen4g0/comparator/pkg/http"
	. "github.com/praveen4g0/comparator/pkg/matcher"
)

func Compare(dir1, dir2 string) {
	var err error
	furls, err := ioutil.ReadFile(files.Path(dir1))
	assert.NoError(err)
	surls, err := ioutil.ReadFile(files.Path(dir2))
	assert.NoError(err)
	furl_list := strings.Split(string(furls), "\n")
	surl_list := strings.Split(string(surls), "\n")

	min, max := func() ([]string, []string) {
		if len(furl_list) < len(surl_list) {
			return furl_list, surl_list
		}
		return surl_list, furl_list
	}()

	c := make(chan *http.Result)
	var wg sync.WaitGroup

	for key, value := range min {
		wg.Add(1)
		go http.GetJsonResponses([]string{max[key], value}, c, &wg)
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
			success, err := (&JSONMatcher{JSONToMatch: res.JsonActual}).Match(res.JsonExpected)
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
