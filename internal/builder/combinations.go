package builder

import (
	"math/big"
	"strconv"

	"github.com/ernestosuarez/itertools"
)

func Combinations(tokens []string) (int, [][]string) {
	count :=  TotalCombinationsFor(len(tokens))
	combos := make([][]string, 0,count)
	for i := 1; i < len(tokens)+1; i++ {
		for v := range itertools.CombinationsStr(tokens, i) {
			combos = append(combos, v)
		}
	}
	return count, combos
}

func TotalCombinationsFor(n int) int {
	var total int64
	z := new(big.Int)
	for i := 1; i < n+1; i++ {
		total += z.Binomial(int64(n), int64(i)).Int64()
	}
	return int(total)
}

func TokenIDs(n int) []string {
	s := make([]string, 0, n)
	for i:=0;i<n;i++{
		s = append(s, strconv.Itoa(i))
	}
	return s
}