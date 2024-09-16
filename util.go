package saml

import (
	"crypto/rand"
	"io"
	mrand "math/rand"
	"strconv"
	"strings"
	"time"

	dsig "github.com/russellhaering/goxmldsig"
)

// TimeNow is a function that returns the current time. The default
// value is time.Now, but it can be replaced for testing.
var TimeNow = func() time.Time { return time.Now().UTC() }

// Clock is assigned to dsig validation and signing contexts if it is
// not nil, otherwise the default clock is used.
var Clock *dsig.Clock

// RandReader is the io.Reader that produces cryptographically random
// bytes when they are need by the library. The default value is
// rand.Reader, but it can be replaced for testing.
var RandReader = rand.Reader

//nolint:unparam // This always receives 20, but we want the option to do more or less if needed.
func randomBytes(n int) []byte {
	rv := make([]byte, n)

	if _, err := io.ReadFull(RandReader, rv); err != nil {
		panic(err)
	}
	return rv
}

func newID() string {
	r := mrand.New(mrand.NewSource(time.Now().UnixNano()))
	var sb strings.Builder
	sb.WriteByte('a')
	sb.WriteString(strconv.FormatInt(abs(r.Int63()), 20))
	sb.WriteString(strconv.FormatInt(abs(r.Int63()), 20))
	return sb.String()
}

func abs(n int64) int64 {
	if n < 0 {
		return -n
	}
	return n
}

func Bool(v bool) *bool { return &v }

func String(v string) *string { return &v }

func Int(v int) *int { return &v }
