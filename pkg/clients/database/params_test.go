package database_test

import (
	"crypto/sha512"
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/firebolt-db/mcp-server/pkg/clients/database"
)

func TestConnectionParams_DriverName(t *testing.T) {
	params := database.ConnectionParams{}
	assert.Equal(t, "firebolt", params.DriverName(), "DriverName should return 'firebolt'")
}

func TestConnectionParams_DSN(t *testing.T) {
	tests := []struct {
		name     string
		params   database.ConnectionParams
		expected string
	}{
		{
			name: "full params",
			params: database.ConnectionParams{
				ClientID:     "test-client",
				ClientSecret: "test-secret",
				AccountName:  "test-account",
				DatabaseName: strPtr("test-db"),
				EngineName:   strPtr("test-engine"),
			},
			expected: "firebolt:///test-db?account_name=test-account&client_id=test-client&client_secret=test-secret&engine=test-engine",
		},
		{
			name: "nil database",
			params: database.ConnectionParams{
				ClientID:     "test-client",
				ClientSecret: "test-secret",
				AccountName:  "test-account",
				DatabaseName: nil,
				EngineName:   strPtr("test-engine"),
			},
			expected: "firebolt://?account_name=test-account&client_id=test-client&client_secret=test-secret&engine=test-engine",
		},
		{
			name: "nil engine name",
			params: database.ConnectionParams{
				ClientID:     "test-client",
				ClientSecret: "test-secret",
				AccountName:  "test-account",
				DatabaseName: strPtr("test-db"),
				EngineName:   nil,
			},
			expected: "firebolt:///test-db?account_name=test-account&client_id=test-client&client_secret=test-secret",
		},
		{
			name: "nil database and engine name",
			params: database.ConnectionParams{
				ClientID:     "test-client",
				ClientSecret: "test-secret",
				AccountName:  "test-account",
				DatabaseName: nil,
				EngineName:   nil,
			},
			expected: "firebolt://?account_name=test-account&client_id=test-client&client_secret=test-secret",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.params.DSN(), "DSN should match expected value")
		})
	}
}

func TestConnectionParams_String(t *testing.T) {
	tests := []struct {
		name     string
		params   database.ConnectionParams
		expected string
	}{
		{
			name: "masks sensitive information",
			params: database.ConnectionParams{
				ClientID:     "sensitive-client-id",
				ClientSecret: "sensitive-secret",
				AccountName:  "test-account",
				DatabaseName: strPtr("test-db"),
				EngineName:   strPtr("test-engine"),
			},
			expected: "firebolt:///test-db?account_name=test-account&client_id=xxxxx&client_secret=xxxxx&engine=test-engine",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.params.String(), "String should mask sensitive information")
		})
	}
}

func TestConnectionParams_Hash(t *testing.T) {
	params := database.ConnectionParams{
		ClientID:     "test-client",
		ClientSecret: "test-secret",
		AccountName:  "test-account",
		DatabaseName: strPtr("test-db"),
		EngineName:   strPtr("test-engine"),
	}

	// Calculate expected hash manually
	dsn := params.DSN()
	sum := sha512.Sum512_256([]byte(dsn))
	expected := hex.EncodeToString(sum[:])

	assert.Equal(t, expected, params.Hash(), "Hash should match expected value")
}

func TestConnectionParams_HashConsistency(t *testing.T) {
	params := database.ConnectionParams{
		ClientID:     "test-client",
		ClientSecret: "test-secret",
		AccountName:  "test-account",
		DatabaseName: strPtr("test-db"),
		EngineName:   strPtr("test-engine"),
	}

	hash1 := params.Hash()
	hash2 := params.Hash()

	assert.Equal(t, hash1, hash2, "Hash should be consistent across calls")
}

func TestConnectionParams_HashDifference(t *testing.T) {
	params1 := database.ConnectionParams{
		ClientID:     "test-client",
		ClientSecret: "test-secret",
		AccountName:  "test-account",
		DatabaseName: strPtr("test-db"),
		EngineName:   strPtr("test-engine"),
	}

	params2 := database.ConnectionParams{
		ClientID:     "different-client",
		ClientSecret: "test-secret",
		AccountName:  "test-account",
		DatabaseName: strPtr("test-db"),
		EngineName:   strPtr("test-engine"),
	}

	hash1 := params1.Hash()
	hash2 := params2.Hash()

	assert.NotEqual(t, hash1, hash2, "Different parameters should have different hashes")
}

// Helper function to create a string pointer
func strPtr(s string) *string {
	return &s
}
