package similar

import (
	//"unicode/utf8"

	"github.com/texttheater/golang-levenshtein/levenshtein"
)

func (s *Similar) levenshtein() (float64, error) {
	ratio := levenshtein.RatioForStrings(s.Source, s.Target, levenshtein.DefaultOptions)
	return ratio, nil
}
