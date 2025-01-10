package sizex_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/av1ppp/sizex"
)

func TestString(t *testing.T) {
	s := (sizex.KiB * 1024).String()
	require.Equal(t, "1 MiB", s)
}

func TestParseSize(t *testing.T) {
	s, err := sizex.ParseSize("10 MiB")
	require.Nil(t, err)
	require.Equal(t, sizex.MiB*10, s)

	s, err = sizex.ParseSize("10 Mib")
	require.NotNil(t, err)
	require.Equal(t, "invalid unit: Mib", err.Error())

	s, err = sizex.ParseSize("10.1 KiB")
	require.Nil(t, err)
	require.Equal(t, sizex.KiB*10+sizex.KiB/10, s)

	s, err = sizex.ParseSize("KiB")
	require.NotNil(t, err)
	require.Equal(t, "strconv.ParseFloat: parsing \"KiB\": invalid syntax", err.Error())
}
