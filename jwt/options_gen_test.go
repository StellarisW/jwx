// This file is auto-generated by internal/cmd/genoptions/main.go. DO NOT EDIT

package jwt

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestOptionIdent(t *testing.T) {
	require.Equal(t, "WithAcceptableSkew", identAcceptableSkew{}.String())
	require.Equal(t, "WithClock", identClock{}.String())
	require.Equal(t, "WithContext", identContext{}.String())
	require.Equal(t, "WithEncryptOption", identEncryptOption{}.String())
	require.Equal(t, "WithFS", identFS{}.String())
	require.Equal(t, "WithFlattenAudience", identFlattenAudience{}.String())
	require.Equal(t, "WithFormKey", identFormKey{}.String())
	require.Equal(t, "WithHeaderKey", identHeaderKey{}.String())
	require.Equal(t, "WithKeyProvider", identKeyProvider{}.String())
	require.Equal(t, "WithNumericDateFormatPrecision", identNumericDateFormatPrecision{}.String())
	require.Equal(t, "WithNumericDateParsePedantic", identNumericDateParsePedantic{}.String())
	require.Equal(t, "WithNumericDateParsePrecision", identNumericDateParsePrecision{}.String())
	require.Equal(t, "WithPedantic", identPedantic{}.String())
	require.Equal(t, "WithSignOption", identSignOption{}.String())
	require.Equal(t, "WithToken", identToken{}.String())
	require.Equal(t, "WithTruncation", identTruncation{}.String())
	require.Equal(t, "WithValidate", identValidate{}.String())
	require.Equal(t, "WithValidator", identValidator{}.String())
	require.Equal(t, "WithVerify", identVerify{}.String())
}
