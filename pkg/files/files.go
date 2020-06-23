package files

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/praveen4g0/comparator/pkg/assert"
)

func Dir() string {
	_, b, _, _ := runtime.Caller(0)
	configDir := path.Join(path.Dir(b), "..", "..", "config")
	return configDir
}

func File(elem ...string) string {
	path := append([]string{Dir()}, elem...)
	return filepath.Join(path...)
}

func Path(elem ...string) string {
	td := filepath.Join(Dir(), "..")
	if _, err := os.Stat(td); os.IsNotExist(err) {
		assert.NoError(err)
	}
	return filepath.Join(append([]string{td}, elem...)...)
}

func ReadBytes(elem string) ([]byte, error) {
	bytes, err := ioutil.ReadFile(Path(elem))
	if err != nil {
		return nil, fmt.Errorf("couldn't load testdata example PullRequest event data: %v", err)
	}
	return bytes, nil
}

func SplitLinesBySperator(file, seperator string) []string {
	return strings.Split(string(file), seperator)
}
