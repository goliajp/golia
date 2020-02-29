package libutil

import "github.com/speps/go-hashids"

func HashId(id uint, salt string) (string, error) {
	hd := hashids.NewData()
	hd.Salt = salt
	h, err := hashids.NewWithData(hd)
	if err != nil {
		return "", err
	}
	e, err := h.Encode([]int{int(id)})
	return e, err
}

func ParseHashId(hashId string, salt string) (uint, error) {
	hd := hashids.NewData()
	hd.Salt = salt
	h, err := hashids.NewWithData(hd)
	if err != nil {
		return 0, err
	}
	d, err := h.DecodeWithError(hashId)
	return uint(d[0]), err
}
