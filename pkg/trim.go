package tmm

// GeneReadCounts contains the read counts for a reference (r) sample and a
// non-reference sample (k) for a particular gene identified by its ID.
type GeneReadCounts struct {
	ID	string
	RefCountr	int
	SampCountk	int
}

// LibReadCountsFromNonzeroGenes filters a slice of genes to those that have
// positive read count in both the reference and non-reference sample, returning
// the filtered slice of genes and the total library counts for the reference
// and non-reference samples.
func LibReadCountsFromNonzeroGenes(g []GeneReadCounts) (filtg []GeneReadCounts, RefTotal, SampTotal int) {
	// Go through each gene, and if bot the the reference and non-reference
	// sample have positive counts, include the gene
	for _, gReadCounts := range g {
		if gReadCounts.RefCountr > 0 && gReadCounts.SampCountk > 0 {
			RefTotal += gReadCounts.RefCountr
			SampTotal += gReadCounts.SampCountk
			filtg = append(filtg, gReadCounts)
		}
	}
	return filtg, RefTotal, SampTotal
}

//func CalcTMMFactor(g []GeneReadCounts) float64 {
	// Go through and get the total
//}

