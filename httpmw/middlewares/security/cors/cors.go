package cors

import "net/http"

var (
	_HEADERS = "*"
	_METHODS = "OPTIONS"
	_ORIGIN  = "*"
)

// add cors headers to the response
//
// the default values are:
//
// HEADER: "*" | METHODS: "OPTIONS" | ORIGIN: "*"
func AddCORSHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Headers", _HEADERS)
		w.Header().Set("Access-Control-Allow-Methods", _METHODS)
		w.Header().Set("Access-Control-Allow-Origin", _ORIGIN)
		next.ServeHTTP(w, r)
	})
}

// -----------------------------------------------------------------------------------
// these functions change the settings for this middleware

// get the headers for cors
func GetHeaders() string {
	return _HEADERS
}

// set the headers for cors
func SetHeaders(headers string) {
	_HEADERS = headers
}

// get the headers for cors
func GetMethods() string {
	return _METHODS
}

// set the headers for cors
func SetMethods(methods string) {
	_METHODS = methods
}

// get the headers for cors
func GetOrigin() string {
	return _ORIGIN
}

// set the headers for cors
func SetOrigin(origin string) {
	_ORIGIN = origin
}
