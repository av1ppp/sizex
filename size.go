package sizex

import (
	"errors"
	"strconv"
	"strings"
)

type Size int64

const (
	B   Size = 1
	KiB      = B * 1024
	MiB      = KiB * 1024
	GiB      = MiB * 1024
	TiB      = GiB * 1024
	PiB      = TiB * 1024
	EiB      = PiB * 1024
)

const (
	strB   = "B"
	strKiB = "KiB"
	strMiB = "MiB"
	strGiB = "GiB"
	strTiB = "TiB"
	strPiB = "PiB"
	strEiB = "EiB"
)

func (self Size) B() int64 {
	return int64(self)
}

func (self Size) KiB() float64 {
	return float64(self) / float64(KiB)
}

func (self Size) MiB() float64 {
	return float64(self) / float64(MiB)
}

func (self Size) GiB() float64 {
	return float64(self) / float64(GiB)
}

func (self Size) TiB() float64 {
	return float64(self) / float64(TiB)
}

func (self Size) PiB() float64 {
	return float64(self) / float64(PiB)
}

func (self Size) EiB() float64 {
	return float64(self) / float64(EiB)
}

func (self Size) String() string {
	switch {
	case self >= EiB:
		return strconv.FormatFloat(self.EiB(), 'f', -1, 64) + " " + strEiB
	case self >= PiB:
		return strconv.FormatFloat(self.PiB(), 'f', -1, 64) + " " + strPiB
	case self >= TiB:
		return strconv.FormatFloat(self.TiB(), 'f', -1, 64) + " " + strTiB
	case self >= GiB:
		return strconv.FormatFloat(self.GiB(), 'f', -1, 64) + " " + strGiB
	case self >= MiB:
		return strconv.FormatFloat(self.MiB(), 'f', -1, 64) + " " + strMiB
	case self >= KiB:
		return strconv.FormatFloat(self.KiB(), 'f', -1, 64) + " " + strKiB
	default:
		return strconv.FormatInt(int64(self), 10) + " B"
	}
}

func ParseSize(s string) (Size, error) {
	parts := strings.SplitN(s, " ", 2)

	num, err := strconv.ParseFloat(parts[0], 64)
	if err != nil {
		return -1, err
	}

	unit := parts[1]

	switch unit {
	case strB:
		return Size(num) * B, nil
	case strKiB:
		return Size(num * float64(KiB)), nil
	case strMiB:
		return Size(num * float64(MiB)), nil
	case strGiB:
		return Size(num * float64(GiB)), nil
	case strTiB:
		return Size(num * float64(TiB)), nil
	case strPiB:
		return Size(num * float64(PiB)), nil
	case strEiB:
		return Size(num * float64(EiB)), nil
	default:
		return -1, errors.New("invalid unit: " + unit)
	}
}
