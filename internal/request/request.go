package request

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Request represents an HTTP request
type Request struct {
	Name    string            `yaml:"name"`
	Method  string            `yaml:"method"`
	URL     string            `yaml:"url"`
	Headers map[string]string `yaml:"headers"`
	Body    string            `yaml:"body"`
}

// Response represents an HTTP response
type Response struct {
	StatusCode int
	Status     string
	Headers    map[string][]string
	Body       string
	Duration   time.Duration
	Error      error
}

// Collection represents a collection of requests
type Collection struct {
	Requests []Request `yaml:"requests"`
}

// Send executes an HTTP request and returns the response
func Send(req Request) Response {
	start := time.Now()

	// Create HTTP request
	var bodyReader io.Reader
	if req.Body != "" {
		bodyReader = bytes.NewBufferString(req.Body)
	}

	httpReq, err := http.NewRequest(req.Method, req.URL, bodyReader)
	if err != nil {
		return Response{
			Error:    err,
			Duration: time.Since(start),
		}
	}

	// Add headers
	for key, value := range req.Headers {
		httpReq.Header.Set(key, value)
	}

	// Send request
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	resp, err := client.Do(httpReq)
	if err != nil {
		return Response{
			Error:    err,
			Duration: time.Since(start),
		}
	}
	defer resp.Body.Close()

	// Read response body
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return Response{
			StatusCode: resp.StatusCode,
			Status:     resp.Status,
			Headers:    resp.Header,
			Error:      err,
			Duration:   time.Since(start),
		}
	}

	return Response{
		StatusCode: resp.StatusCode,
		Status:     resp.Status,
		Headers:    resp.Header,
		Body:       string(bodyBytes),
		Duration:   time.Since(start),
	}
}

// FormatResponse formats a response for display
func FormatResponse(resp Response) string {
	if resp.Error != nil {
		return fmt.Sprintf("❌ Error: %v\n\nDuration: %v", resp.Error, resp.Duration)
	}

	output := fmt.Sprintf("Status: %s\nDuration: %v\n\n", resp.Status, resp.Duration)

	output += "Headers:\n"
	for key, values := range resp.Headers {
		for _, value := range values {
			output += fmt.Sprintf("  %s: %s\n", key, value)
		}
	}

	output += "\nBody:\n"
	output += resp.Body

	return output
}
