package similar

import (
	"errors"
)

func New(algo Algo) *Similar {
	return &Similar{
		Algo: algo,
	}
}

func (s *Similar) Ratio(source, target string) (float64, error) {
	s.Source = []rune(source)
	s.Target = []rune(target)

	if s.Algo == AlgoLevenshtein {
		return s.levenshtein()
	}

	return 0, errors.New("参数错误")
}
