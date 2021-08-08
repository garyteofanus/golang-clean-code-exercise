package models

type transactions struct {
	data map[int]int
}

func NewTransactions(d []int) Transactions {
	data := make(map[int]int)

	for i, v := range d {
		data[i] = v
	}

	return &transactions{data: data}
}

func (t *transactions) Get(idx int) int {
	return t.data[idx]
}
