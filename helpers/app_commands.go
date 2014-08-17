package helpers

import (
	"time"

	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gexec"

	"github.com/cloudfoundry-incubator/cf-test-helpers/runner"
)

const CURL_TIMEOUT = 10 * time.Second

// Gets an app's endpoint with the specified path
func AppUri(appName, path string) string {
	appsDomain := LoadConfig().AppsDomain
	return "http://" + appName + "." + appsDomain + path
}

// Gets an app's endpoint with the specified path
func AppUriWithAuth(appName, path string, username string, password string) string {
	appsDomain := LoadConfig().AppsDomain
	return "http://" + username + ":" + password + "@" + appName + "." + appsDomain + path
}

// Gets an app's root endpoint
func AppRootUri(appName string) string {
	return AppUri(appName, "/")
}

// Curls an app's endpoint and exit successfully before the specified timeout
func CurlAppWithTimeout(appName, path string, timeout time.Duration) string {
	uri := AppUri(appName, path)
	curl := runner.Curl(uri).Wait(timeout)
	Expect(curl).To(Exit(0))
	Expect(string(curl.Err.Contents())).To(HaveLen(0))
	return string(curl.Out.Contents())
}

// Curls an app's endpoint, with basic auth, and exit successfully before the specified timeout
func CurlAppWithTimeoutAndAuth(appName, path string, username string, password string, timeout time.Duration) string {
	uri := AppUriWithAuth(appName, path, username, password)
	curl := runner.Curl(uri).Wait(timeout)
	Expect(curl).To(Exit(0))
	Expect(string(curl.Err.Contents())).To(HaveLen(0))
	return string(curl.Out.Contents())
}

// Curls an app's endpoint and exit successfully before the default timeout
func CurlApp(appName, path string) string {
	return CurlAppWithTimeout(appName, path, CURL_TIMEOUT)
}

// Curls an app's endpoint, with basic auth, and exit successfully before the default timeout
func CurlAppWithAuth(appName, path string, username string, password string) string {
	return CurlAppWithTimeoutAndAuth(appName, path, username, password, CURL_TIMEOUT)
}

// Curls an app's root endpoint and exit successfully before the default timeout
func CurlAppRoot(appName string) string {
	return CurlApp(appName, "/")
}

// Curls an app's root endpoint, using basic auth, and exit successfully before the default timeout
func CurlAppRootWithAuth(appName string, username string, password string) string {
	return CurlAppWithAuth(appName, "/", username, password)
}

// Returns a function that curls an app's root endpoint and exit successfully before the default timeout
func CurlingAppRoot(appName string) func() string {
	return func() string {
		return CurlAppRoot(appName)
	}
}
