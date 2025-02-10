package useful

import (
	// "bytes"
	// "errors"
	// "fmt"
	// "math"
	// "math/rand"
	// _url "net/url"
	// "regexp"
	// "strconv"
	// "strings"
	// "time"

	"github.com/go-openapi/strfmt"
	// "github.com/google/uuid"
	// "golang.org/x/net/idna"
)

func StrPtr(s string) *string               { return &s }
func Int64Ptr(n int64) *int64               { return &n }
func Float64Ptr(n float64) *float64         { return &n }
func IntBoolPtr(b bool) *bool               { return &b }
func StrFmtEmail(email string) strfmt.Email { return strfmt.Email(email) }
func StrFmtEmailPtr(email string) *strfmt.Email {
	strfmtEmail := StrFmtEmail(email)
	return &strfmtEmail
}
