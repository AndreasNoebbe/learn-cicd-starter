package auth12

import (
    "errors"
    "net/http"
    "testing"
)

func TestGetAPIKey(t *testing.T) {
    // Define test cases
    tests := []struct {
        name        string
        authHeader  string
        want        string
        wantErr     bool
        expectedErr error
    }{
        {
            name:        "Valid API key",
            authHeader:  "ApiKey test-api-key",
            want:        "test-api-key",
            wantErr:     false,
            expectedErr: nil,
        },
        {
            name:        "Missing Authorization header",
            authHeader:  "",
            want:        "",
            wantErr:     true,
            expectedErr: ErrNoAuthHeaderIncluded,
        },
        {
            name:        "Malformed Authorization header",
            authHeader:  "Bearer test-api-key",
            want:        "",
            wantErr:     true,
            expectedErr: errors.New("malformed authorization header"),
        },
    }

    // Run through all test cases
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            // Create a header with the provided authHeader
            headers := http.Header{}
            if tt.authHeader != "" {
                headers.Set("Authorization", tt.authHeader)
            }

            // Call the function being tested
            got, err := GetAPIKey(headers)

            // Check for expected error
            if (err != nil) != tt.wantErr {
                t.Errorf("GetAPIKey() error = %v, wantErr %v", err, tt.wantErr)
            }
            if err != nil && err.Error() != tt.expectedErr.Error() {
                t.Errorf("GetAPIKey() error = %v, expectedErr %v", err, tt.expectedErr)
            }

            // Check the returned API key
            if got != tt.want {
                t.Errorf("GetAPIKey() = %v, want %v", got, tt.want)
            }
        })
    }
}
