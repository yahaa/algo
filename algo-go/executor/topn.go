package executor

// Row single row data
type Row struct {
	Key  int
	Data []byte
}

// RowPtr row point
type RowPtr struct {
	Key    int
	KeyIdx int
}

// TopNExec top
type TopNExec struct {
	rows    []Row
	rowPtrs []RowPtr
}

// NewTopNExec new top n builder
func NewTopNExec(rows []Row) *TopNExec {
	rowPtrs := make([]RowPtr, len(rows))
	for i, item := range rows {
		rowPtrs[i] = RowPtr{Key: item.Key, KeyIdx: i}
	}

	res := make(chan []RowPtr)
	go sort(rowPtrs, res)
	rowPtrs = <-res

	return &TopNExec{
		rows:    rows,
		rowPtrs: rowPtrs,
	}
}

// Get get top n
func (s *TopNExec) Get(n int) []Row {
	start := len(s.rowPtrs) - n
	if start < 0 {
		start = 0
	}

	res := make([]Row, n)

	for i := 0; i < n; i++ {
		res[i] = s.rows[s.rowPtrs[i+start].KeyIdx]
	}

	return res
}

func sort(data []RowPtr, resCh chan []RowPtr) {
	if len(data) < 2 {
		resCh <- data
		return
	}

	leftCh := make(chan []RowPtr)
	rightCh := make(chan []RowPtr)

	mid := len(data) / 2
	go sort(data[:mid], leftCh)
	go sort(data[mid:], rightCh)

	left := <-leftCh
	right := <-rightCh

	close(leftCh)
	close(rightCh)

	resCh <- merge(left, right)
}

func merge(left, right []RowPtr) []RowPtr {
	result := make([]RowPtr, len(left)+len(right))
	lIdx, rIdx := 0, 0

	for i := 0; i < cap(result); i++ {
		switch {
		case lIdx >= len(left):
			result[i] = right[rIdx]
			rIdx++
		case rIdx >= len(right):
			result[i] = left[lIdx]
			lIdx++
		case left[lIdx].Key < right[rIdx].Key:
			result[i] = left[lIdx]
			lIdx++
		default:
			result[i] = right[rIdx]
			rIdx++
		}
	}
	return result
}
