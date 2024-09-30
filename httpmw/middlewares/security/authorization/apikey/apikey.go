package apikey

import (
	"net/http"
)

var (
	_APIKEY_HEADER = "X-API-JRDB"
	_APIKEY_VALUE  = "c8eab7bf08f9a630a24f04b505dce2c7"
)

// add a configured apikey-check
//
// the used apikey can be configured using "SetAPIKey(string)" and received using "GetAPIKey() string"
//
// the apikey is
func APIKeyCheck(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get(_APIKEY_HEADER) == _APIKEY_VALUE {
			next.ServeHTTP(w, r)
		} else {
			// response with unauthorized
			w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		}
	})
}

// set the header for a new http-request
func SetAPIKeyForRequest(r *http.Request) {
	r.Header.Set(_APIKEY_HEADER, _APIKEY_VALUE)
}

// -----------------------------------------------------------------------------------
// these functions change the settings for this middleware

// set the headername for the key
func SetHeaderName(header string) {
	_APIKEY_HEADER = header
}

// set the headername for the key
func GetHeaderName() string {
	return _APIKEY_HEADER
}

// set the headername for the key
func SetAPIKey(key string) {
	_APIKEY_VALUE = key
}

// set the headername for the key
func GetAPIKey() string {
	return _APIKEY_VALUE
}
