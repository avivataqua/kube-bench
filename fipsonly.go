// +build fipsonly

package main

// Package fipsonly restricts all TLS configuration to FIPS-approved settings.
// This package only exists in the dev.boringcrypto branch of Go.
import _ "crypto/tls/fipsonly"
