package dna

import (
	"errors"
	"strings"
)

const testVersion = 2

type DNA string

type Histogram map[byte]int

var symbols = "ACTG"

func (d DNA) Count(nucleotide byte) (int, error) {
	if !strings.ContainsAny(symbols, string(nucleotide)) {
		return 0, errors.New("invalid nucleotide")
	}
	return strings.Count(string(d), string(nucleotide)), nil
}
func (d DNA) Counts() (Histogram, error) {
	h := make(Histogram, 0)
	lenDNA := 0
	for i := 0; i < len(symbols); i++ {
		nucleotide := symbols[i]
		if c, err := d.Count(nucleotide); err == nil {
			h[nucleotide] = c
			lenDNA += c
		} else {
			return nil, err
		}
	}
	if lenDNA != len(d) {
		return nil, errors.New("invalid dna")
	}
	return h, nil
}
