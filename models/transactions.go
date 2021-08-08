package models

type transactions struct {
	data map[int]int
}

// i.e. go test is green
func NewTransactions(d []int) Transactions {
	data := make(map[int]int)

	for i, v := range d {
		data[i] = v
	}

	return &transactions{data: data}
}

// i.e. go test is green
func (t *transactions) Get(idx int) int {
	return t.data[idx]
}

// i.e. go test is green
func (t *transactions) GetTotal() int {
	ret := 0

	for i := 0; i <= 6; i++ {
		ret += t.Get(i)
	}

	return ret
}

// i.e. go test is green
func (t *transactions) GetTotalWithinRange(i, j int) int {
	ret := 0

	for i := i; i <= j; i++ {
		ret += t.Get(i)
	}

	return ret
}
