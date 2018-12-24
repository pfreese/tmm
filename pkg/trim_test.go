package tmm

import (
	"reflect"
	"testing"
)

func TestLibReadCountsFromNonzeroGenes(t *testing.T) {

	tables := []struct {
		g	[]GeneReadCounts
		expFiltg []GeneReadCounts
		expRefTotal	int
		expSampTotal	int
	}{
		// A simple example with one gene
		{[]GeneReadCounts{
			{"a", 5, 6},
		},
			[]GeneReadCounts{
				{"a", 5, 6},
			},
			5,
			6},
		// A simple example with two genes
		{[]GeneReadCounts{
			{"a", 5, 6},
			{"b", 100, 100},
		},
			[]GeneReadCounts{
				{"a", 5, 6},
				{"b", 100, 100},
			},
			105,
			106},
		// Example in which 2nd gene filtered out because Ref count is 0
		{[]GeneReadCounts{
			{"a", 5, 6},
			{"b", 0, 600},
		},
			[]GeneReadCounts{
				{"a", 5, 6},
			},
			5,
			6},
		// Example in which 1st gene filtered out because Samp count is 0
		{[]GeneReadCounts{
			{"b", 10, 0},
			{"a", 5, 6},
		},
			[]GeneReadCounts{
				{"a", 5, 6},
			},
			5,
			6},
	}

	for _, table := range tables {
		actFiltg, actRefTotal, actSampTotal := LibReadCountsFromNonzeroGenes(table.g)
		if !reflect.DeepEqual(actFiltg, table.expFiltg) {
			t.Errorf("Wrong filtered g: Got: %v, want: %v.", actFiltg, table.expFiltg)
		}
		if actRefTotal != table.expRefTotal {
			t.Errorf("Wrong total reference count: Got: %v, want: %v.", actRefTotal, table.expRefTotal)
		}
		if actSampTotal != table.expSampTotal {
			t.Errorf("Wrong total sample count: Got: %v, want: %v.", actSampTotal, table.expSampTotal)
		}
	}
}