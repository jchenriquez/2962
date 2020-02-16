package atoi

import (
	"bufio"
	"encoding/json"
	"io"
	"os"
	"testing"
)

type Test struct {
	Input  string `json:"input"`
	Output int    `json:"output"`
}

func TestMyAtoi(testUtil *testing.T) {
	f, err := os.Open("./tests.json")

	if err != nil {
		testUtil.Error(err)
		return
	}

	defer f.Close()

	reader := bufio.NewReader(f)
	decoder := json.NewDecoder(reader)
	var tests map[string]Test

	for {
		err = decoder.Decode(&tests)

		if err == nil {
			for name, test := range tests {
				testUtil.Run(name, func(t *testing.T) {
					result := MyAtoi(test.Input)
					if result != test.Output {
						t.Errorf("result given by code is %v", result)
					}
				})
			}

		} else if err == io.EOF {
			break
		} else {
			testUtil.Error(err)
			break
		}
	}
}
