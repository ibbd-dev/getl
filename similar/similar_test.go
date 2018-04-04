package similar

import (
	"testing"
)

var testCases = []struct {
	source, target string
	ratio          float64
	algo           Algo
}{
	{
		source: "中国人民",
		target: "中国人名",
		ratio:  0.75,
		algo:   AlgoLevenshtein,
	},
	{
		source: "中国人名",
		target: "中国人民",
		ratio:  0.75,
		algo:   AlgoLevenshtein,
	},
	{
		source: "中国人",
		target: "中国人民",
		ratio:  6.0 / 7.0,
		algo:   AlgoLevenshtein,
	},
}

func TestSimilar(t *testing.T) {

	for _, testCase := range testCases {
		s := New(testCase.algo)
		ratio, err := s.Ratio(testCase.source, testCase.target)
		if err != nil {
			t.Fatal(err)
		}

		if ratio != testCase.ratio {
			t.Fatalf("%s and %s should be %f, but not %f\n", testCase.source, testCase.target, testCase.ratio, ratio)
		}
	}
}
