package libpg

import (
	"bytes"
	"database/sql/driver"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
)

type (
	Id       uint
	IntArr   []int
	UintArr  []uint
	GeoPoint struct {
		Lat float64 `json:"lat"`
		Lng float64 `json:"lng"`
	}
)

// type IntArr
func (a *IntArr) Scan(src interface{}) error {
	switch src := src.(type) {
	case []byte:
		return a.scanBytes(src)
	case string:
		return a.scanBytes([]byte(src))
	case nil:
		*a = nil
		return nil
	}

	return fmt.Errorf("pg: cannot convert %T to IntArr", src)
}

func (a *IntArr) scanBytes(src []byte) error {
	itemArr, err := scanLinearArray(src, []byte{','}, "IntArr")
	if err != nil {
		return err
	}
	if *a != nil && len(itemArr) == 0 {
		*a = (*a)[:0]
	} else {
		b := make(IntArr, len(itemArr))
		for i, v := range itemArr {
			i64, err := strconv.ParseInt(string(v), 10, 64)
			if err != nil {
				return fmt.Errorf("pg: parsing array item index %d: %v", i, err)
			}
			b[i] = int(i64)
		}
		*a = b
	}
	return nil
}

func (a IntArr) Value() (driver.Value, error) {
	if a == nil {
		return nil, nil
	}

	if n := len(a); n > 0 {
		b := make([]byte, 1, 1+2*n)
		b[0] = '{'

		b = strconv.AppendInt(b, int64(a[0]), 10)
		for i := 1; i < n; i++ {
			b = append(b, ',')
			b = strconv.AppendInt(b, int64(a[i]), 10)
		}

		return string(append(b, '}')), nil
	}

	return "{}", nil
}

// type UintArr
func (a *UintArr) Scan(src interface{}) error {
	switch src := src.(type) {
	case []byte:
		return a.scanBytes(src)
	case string:
		return a.scanBytes([]byte(src))
	case nil:
		*a = nil
		return nil
	}

	return fmt.Errorf("pg: cannot convert %T to UintArr", src)
}

func (a *UintArr) scanBytes(src []byte) error {
	itemArr, err := scanLinearArray(src, []byte{','}, "UintArr")
	if err != nil {
		return err
	}
	if *a != nil && len(itemArr) == 0 {
		*a = (*a)[:0]
	} else {
		b := make(UintArr, len(itemArr))
		for i, v := range itemArr {
			i64, err := strconv.ParseInt(string(v), 10, 64)
			if err != nil {
				return fmt.Errorf("pg: parsing array item index %d: %v", i, err)
			}
			b[i] = uint(i64)
		}
		*a = b
	}
	return nil
}

func (a UintArr) Value() (driver.Value, error) {
	if a == nil {
		return nil, nil
	}

	if n := len(a); n > 0 {
		b := make([]byte, 1, 1+2*n)
		b[0] = '{'

		b = strconv.AppendInt(b, int64(a[0]), 10)
		for i := 1; i < n; i++ {
			b = append(b, ',')
			b = strconv.AppendInt(b, int64(a[i]), 10)
		}

		return string(append(b, '}')), nil
	}

	return "{}", nil
}

func parseArray(src, del []byte) (dims []int, elementArr [][]byte, err error) {
	var depth, i int

	if len(src) < 1 || src[0] != '{' {
		return nil, nil, fmt.Errorf("pg: unable to parse array; expected %q at offset %d", '{', 0)
	}

Open:
	for i < len(src) {
		switch src[i] {
		case '{':
			depth++
			i++
		case '}':
			elementArr = make([][]byte, 0)
			goto Close
		default:
			break Open
		}
	}
	dims = make([]int, i)

Element:
	for i < len(src) {
		switch src[i] {
		case '{':
			if depth == len(dims) {
				break Element
			}
			depth++
			dims[depth-1] = 0
			i++
		case '"':
			var elem = []byte{}
			var escape bool
			for i++; i < len(src); i++ {
				if escape {
					elem = append(elem, src[i])
					escape = false
				} else {
					switch src[i] {
					default:
						elem = append(elem, src[i])
					case '\\':
						escape = true
					case '"':
						elementArr = append(elementArr, elem)
						i++
						break Element
					}
				}
			}
		default:
			for start := i; i < len(src); i++ {
				if bytes.HasPrefix(src[i:], del) || src[i] == '}' {
					elem := src[start:i]
					if len(elem) == 0 {
						return nil, nil, fmt.Errorf("pg: unable to parse array; unexpected %q at offset %d", src[i], i)
					}
					if bytes.Equal(elem, []byte("NULL")) {
						elem = nil
					}
					elementArr = append(elementArr, elem)
					break Element
				}
			}
		}
	}

	for i < len(src) {
		if bytes.HasPrefix(src[i:], del) && depth > 0 {
			dims[depth-1]++
			i += len(del)
			goto Element
		} else if src[i] == '}' && depth > 0 {
			dims[depth-1]++
			depth--
			i++
		} else {
			return nil, nil, fmt.Errorf("pg: unable to parse array; unexpected %q at offset %d", src[i], i)
		}
	}

Close:
	for i < len(src) {
		if src[i] == '}' && depth > 0 {
			depth--
			i++
		} else {
			return nil, nil, fmt.Errorf("pg: unable to parse array; unexpected %q at offset %d", src[i], i)
		}
	}
	if depth > 0 {
		err = fmt.Errorf("pg: unable to parse array; expected %q at offset %d", '}', i)
	}
	if err == nil {
		for _, d := range dims {
			if (len(elementArr) % d) != 0 {
				err = fmt.Errorf("pg: multidimensional arrays must have elements with matching dimensions")
			}
		}
	}
	return
}

func scanLinearArray(src, del []byte, typ string) (elementArr [][]byte, err error) {
	dims, elementArr, err := parseArray(src, del)
	if err != nil {
		return nil, err
	}
	if len(dims) > 1 {
		return nil, fmt.Errorf("pg: cannot convert array %s to %s", strings.Replace(fmt.Sprint(dims), " ", "][", -1), typ)
	}
	return elementArr, err
}

// type GeoPoint

func (p *GeoPoint) String() string {
	return fmt.Sprintf("POINT(%v, %v", p.Lat, p.Lng)
}

func (p *GeoPoint) Scan(val interface{}) error {
	b, err := hex.DecodeString(string(val.([]uint8)))
	if err != nil {
		return err
	}
	r := bytes.NewReader(b)
	var wkbByteOrder uint8
	if err := binary.Read(r, binary.LittleEndian, &wkbByteOrder); err != nil {
		return err
	}

	var byteOrder binary.ByteOrder
	switch wkbByteOrder {
	case 0:
		byteOrder = binary.BigEndian
	case 1:
		byteOrder = binary.LittleEndian
	default:
		return fmt.Errorf("Invalid byte order %u", wkbByteOrder)
	}

	var wkbGeometryType uint64
	if err := binary.Read(r, byteOrder, &wkbGeometryType); err != nil {
		return err
	}

	if err := binary.Read(r, byteOrder, p); err != nil {
		return err
	}

	return nil
}

func (p *GeoPoint) Value() (driver.Value, error) {
	return p.String(), nil
}
