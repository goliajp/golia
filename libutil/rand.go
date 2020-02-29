package libutil

import (
	"github.com/sony/sonyflake"
	"math/rand"
	"strconv"
)

const LowerBytes = "abcdefghijklmnopqrstuvwxyz"
const UpperBytes = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const FigureBytes = "0123456789"

// rand all
func RandAll(n int) string {
	dictionary := LowerBytes + UpperBytes + FigureBytes
	return Rand(n, dictionary)
}

// rand string
func RandString(n int) string {
	dictionary := LowerBytes + UpperBytes
	return Rand(n, dictionary)
}

// rand figure
func RandFigure(n int) string {
	dictionary := FigureBytes
	return Rand(n, dictionary)
}

// rand lower string
func RandLowerString(n int) string {
	dictionary := UpperBytes
	return Rand(n, dictionary)
}

// rand upper string
func RandUpperString(n int) string {
	dictionary := UpperBytes
	return Rand(n, dictionary)
}

// rand
func Rand(n int, dictionary string) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = dictionary[rand.Intn(len(dictionary))]
	}
	return string(b)
}

// no
func No(prefix string) string {
	no := ""
	flake := sonyflake.NewSonyflake(sonyflake.Settings{})
	randNo, _ := flake.NextID()
	if prefix != "" {
		no = prefix + "-" + strconv.Itoa(int(randNo))
	} else {
		no = strconv.Itoa(int(randNo))
	}
	return no
}
