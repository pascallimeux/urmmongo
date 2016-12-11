package utils

import (
	"github.com/pascallimeux/urmmongo/utils/log"
	"time"
)

const DATEFORMAT1 = "2006-01-02T15:04:05.000Z"
const DATEFORMAT2 = "2006-01-02"
const DATEFORMAT3 = "2006-01-02 15:04:05 -0700 -0700"

func DateParse(datestr, layout string, all_the_day bool) (time.Time, error) {
	log.Trace(log.Here(), "DateParse() : calling method -")
	date, err := time.Parse(layout, datestr)
	if err != nil {
		log.Error(log.Here(), "DateParse error: ", err.Error())
	}
	if all_the_day && layout == DATEFORMAT2 {
		date = date.Add(time.Duration(86400) * time.Second)
	}
	return date, err
}
