package baseclient

import (
	"regexp"
	"time"
)

// CallWithRetry executes callback with retry
func CallWithRetry(callback func() (interface{}, error), maxRetriesCount int, retryInterval time.Duration) (interface{}, error) {
	for index := 0; index < maxRetriesCount; index++ {
		resp, err := callback()
		if !shouldRetry(err) {
			return resp, err
		}
		time.Sleep(retryInterval)
	}
	return callback()
}

func shouldRetry(err error) bool {
	if err == nil {
		return false
	}
	isMatching, _ := regexp.MatchString(" EOF$", err.Error())
	if isMatching {
		return true
	}
	ae, ok := err.(*ClientError)
	if ok {
		httpCode := ae.Code
		httpCodeMajorDigit := httpCode / 100
		if httpCodeMajorDigit == 5 {
			return true
		}
	}
	return false
}
