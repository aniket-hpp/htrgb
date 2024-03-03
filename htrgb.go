package htrgb

import (
	"errors"
	"strconv"
)

func HexToRGB(s string) (int64, int64, int64, error) {
	slen := len(s)
	if slen == 0 {
		return 0, 0, 0, errors.New("htrgb: error: empty string")
	}

	var r, g, b int64
	var err error

	if s[0] == '#' {
		if slen != 7 {
			return 0, 0, 0, errors.New("htrgb: error: format should be either '#rrggbb' or 'rrggbb'")
		}

		s = s[1:]
	} else if slen != 6 {
		return 0, 0, 0, errors.New("htrgb: error: format should be either '#rrggbb' or 'rrggbb'")
	}

	sr := s[0:2]
	sg := s[2:4]
	sb := s[4:6]

	r, err = strconv.ParseInt(sr, 16, 16)
	if err != nil {
		return 0, 0, 0, err
	}

	g, err = strconv.ParseInt(sg, 16, 16)
	if err != nil {
		return 0, 0, 0, err
	}

	b, err = strconv.ParseInt(sb, 16, 16)
	if err != nil {
		return 0, 0, 0, err
	}

	return r, g, b, nil
}
