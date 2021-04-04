package goft

type Element int

const (
  Even = iota
  Odd
)

func makeIndices(start, stop, step int) []int {
  n_elements := (stop - start)/step
  result := make([]int, n_elements)

  for i := range result {
    result[i] = start + i*step
  }

  return result
}

func PartitionElements(parity Element, size int) []int {

  if (parity == Even) {
    return makeIndices(0, size, 2)
  }

  return makeIndices(1, size + 1, 2)
}
