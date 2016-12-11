package utils

import (
	"bytes"
	"encoding/json"
	"github.com/pascallimeux/urmmongo/utils/log"
	"io/ioutil"
	"regexp"
	"strings"
)

func Display_json(jsonstr string) {
	var out bytes.Buffer
	json.Indent(&out, []byte(jsonstr), "", "  ")
	log.Trace(log.Here(), "Json object: ", out.String())
}

func ReadFile(filename string) (string, error) {
	var payload string
	raw, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Error(log.Here(), err.Error())
		return payload, err
	}

	payload = string(raw)
	payload = strings.Replace(payload, "\r\n", " ", -1)
	//payload = strings.Replace(payload, "\"", "\\\"", -1)
	re := regexp.MustCompile("  +")
	payload = re.ReplaceAllString(payload, " ")
	return payload, nil
}
