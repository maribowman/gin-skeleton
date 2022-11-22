package middleware

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/maribowman/gin-skeleton/app/model"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

const (
	replacer = "replaceMe"
	path     = "/"
)

func TestAuthorized(t *testing.T) {
	// given
	allowedUsers := []model.User{{
		Name:   "testName",
		Key:    "testKey",
		Secret: "dGVzdFNlY3JldA==",
	}}
	tests := []struct {
		testName string
		headers  map[string]interface{}
		expected struct {
			code    int
			message string
		}
	}{
		{
			testName: "incomplete auth headers",
			headers: map[string]interface{}{
				"key":       "testKey",
				"signature": "",
				"timestamp": "1619175474",
			},
			expected: struct {
				code    int
				message string
			}{code: http.StatusBadRequest, message: `{"error":"authentication headers not complete"}`},
		},
		{
			testName: "invalid timestamp",
			headers: map[string]interface{}{
				"key":       "testKey",
				"signature": "dummy",
				"timestamp": time.Now().UTC(),
			},
			expected: struct {
				code    int
				message string
			}{code: http.StatusForbidden, message: `{"error":"invalid timestamp"}`},
		},
		{
			testName: "timestamp too old",
			headers: map[string]interface{}{
				"key":       "testKey",
				"signature": "dummy",
				"timestamp": time.Now().Add(-24 * time.Hour).UTC().Unix(),
			},
			expected: struct {
				code    int
				message string
			}{code: http.StatusForbidden, message: `{"error":"timestamp too old"}`},
		},
		{
			testName: "timestamp in future",
			headers: map[string]interface{}{
				"key":       "testKey",
				"signature": "dummy",
				"timestamp": time.Now().Add(24 * time.Hour).UTC().Unix(),
			},
			expected: struct {
				code    int
				message string
			}{code: http.StatusForbidden, message: `{"error":"timestamp in the future"}`},
		},
		{
			testName: "invalid key",
			headers: map[string]interface{}{
				"key":       "invalid",
				"signature": "dummy",
				"timestamp": time.Now().UTC().Unix(),
			},
			expected: struct {
				code    int
				message string
			}{code: http.StatusUnauthorized, message: `{"error":"invalid key"}`},
		},
		{
			testName: "positive test",
			headers: map[string]interface{}{
				"key":       "testKey",
				"signature": replacer,
				"timestamp": time.Now().UTC().Unix(),
			},
			expected: struct {
				code    int
				message string
			}{code: http.StatusNoContent, message: ""},
		},
	}

	// and
	router := gin.New()
	router.Use(HmacMiddleware(true, allowedUsers))
	router.GET(path, func(context *gin.Context) {
		context.Status(http.StatusNoContent)
		return
	})

	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			if test.headers["signature"] == replacer {
				test.headers["signature"] = generateSignature("testSecret", fmt.Sprintf("%v", test.headers["timestamp"]))
			}

			// when
			w := performRequest(router, http.MethodGet, path, test.headers)

			// then
			assert.Equal(t, test.expected.code, w.Code)
			assert.Equal(t, test.expected.message, w.Body.String())
		})
	}

}

func performRequest(r http.Handler, method, path string, headers map[string]interface{}) *httptest.ResponseRecorder {
	request := httptest.NewRequest(method, path, nil)
	for key, value := range headers {
		request.Header.Add(key, fmt.Sprintf("%v", value))
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, request)
	return w
}

func generateSignature(secret, timestamp string) string {
	sigHash := hmac.New(sha256.New, []byte(secret))
	sigHash.Write([]byte(timestamp + http.MethodGet + path))
	return hex.EncodeToString(sigHash.Sum(nil))
}
