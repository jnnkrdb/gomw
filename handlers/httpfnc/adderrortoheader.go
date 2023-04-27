package httpfnc

import (
	"fmt"
	"net/http"
)

// add am error to the header of an request
// the used header is [x-gomw-error]
func AddErrorToHeaderIfAny(w *http.ResponseWriter, err error) {
	if err != nil {
		(*w).Header().Add("X-GOMW-ERROR", fmt.Sprintf("%v", err))
	}
}
