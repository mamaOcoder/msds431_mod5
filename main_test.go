package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	invalidURL = "http://random-invalid-url"
	validURL   = "https://en.wikipedia.org/wiki/Robotics"
)

func TestScrape(t *testing.T) {
	assert := assert.New(t)

	// Test Scrape with invalid URL
	_, _, err := Scrape(invalidURL)
	assert.Error(err, "Scrape() should return an error for invalid URLs")

	// Test Scrape with valid URL
	jsonout, _, err := Scrape(validURL)
	assert.NoError(err, "Scrape() should not return an error for valid URLs")
	assert.NotEmpty(jsonout.Title, "Title should not be empty")
	assert.NotEmpty(jsonout.Text, "Text should not be empty")
}
