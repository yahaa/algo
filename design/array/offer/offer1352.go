package offer

type ProductOfNumbers struct {
	pre    [40100]int
	preLen int
}

func Constructor() ProductOfNumbers {
	p := ProductOfNumbers{}
	p.pre[0] = 1

	return p
}

func (p *ProductOfNumbers) Add(num int) {
	if num == 0 {
		p.preLen = 0
	} else {
		p.preLen++
		p.pre[p.preLen] = num * p.pre[p.preLen-1]
	}
}

func (p *ProductOfNumbers) GetProduct(k int) int {
	if k > p.preLen {
		return 0
	}

	return p.pre[p.preLen] / p.pre[p.preLen-k]
}
