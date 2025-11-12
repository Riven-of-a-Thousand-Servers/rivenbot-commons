package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadSecret_Success(t *testing.T) {
	// given: a file path for a secret
	path := "./testdata/secret.txt"

	// when: ReadSecret is called
	secret, err := ReadSecret(path)

	// then: no error is returned and the secret is fetched correctly
	if err != nil {
		t.Fatalf("No error expected, got: %v", err)
	}

	assert := assert.New(t)
	assert.Equal("mysecrethuhu", secret, "Secret value should be equal to what file has")
}
