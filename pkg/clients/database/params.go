package database

import (
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"strings"
)

const secretMask = "xxxxx"

// PoolParams holds the parameters required to create a connection in the pool.
// The pool will convert these parameters into a ConnectionParams struct by adding the ClientID and ClientSecret.
type PoolParams struct {
	AccountName  string
	DatabaseName *string
	EngineName   *string
}

// ConnectionParams holds the parameters required to establish a connection to Firebolt.
// It contains authentication credentials, account information, and optional database and engine specifications.
type ConnectionParams struct {
	ClientID     string
	ClientSecret string
	AccountName  string
	DatabaseName *string
	EngineName   *string
}

// String returns a string representation of the connection parameters.
// The format is similar to a DSN, but sensitive information (ClientID and ClientSecret) is masked,
// making this string safe to log or display.
func (c ConnectionParams) String() string {
	return strings.NewReplacer(c.ClientID, secretMask, c.ClientSecret, secretMask).Replace(c.DSN())
}

// DriverName returns the name of the database driver.
// In this case, it returns "firebolt", which is the driver name used by the Firebolt Go SDK.
func (c ConnectionParams) DriverName() string {
	return "firebolt"
}

// DSN returns the Data Source Name expected by the Firebolt Go SDK.
// This is a connection string that contains all the necessary information to connect to the Firebolt database.
// The format follows the pattern: firebolt://[database]?account_name=[account]&client_id=[id]&client_secret=[secret]&engine=[engine]
func (c ConnectionParams) DSN() string {

	dbName := ""
	if c.DatabaseName != nil {
		dbName = "/" + *c.DatabaseName
	}

	engineName := ""
	if c.EngineName != nil {
		engineName = "&engine=" + *c.EngineName
	}

	return fmt.Sprintf("firebolt://%s?account_name=%s&client_id=%s&client_secret=%s%s", dbName, c.AccountName, c.ClientID, c.ClientSecret, engineName)
}

// Hash returns a SHA-512/256 hash of the connection parameters.
// This is useful for caching connections or comparing parameter sets without exposing sensitive information.
// The hash is computed from the full DSN string and returned as a hex-encoded string.
func (c ConnectionParams) Hash() string {
	sum := sha512.Sum512_256([]byte(c.DSN()))
	return hex.EncodeToString(sum[:])
}
