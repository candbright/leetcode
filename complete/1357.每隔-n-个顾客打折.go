/*
 * @lc app=leetcode.cn id=1357 lang=golang
 *
 * [1357] 每隔 n 个顾客打折
 */

package leetcode

// @lc code=start
type Cashier struct {
	sum      int
	n        int
	discount int
	products []int
	prices   []int
}

func Constructor(n int, discount int, products []int, prices []int) Cashier {
	cashier := Cashier{
		sum:      0,
		n:        n,
		discount: discount,
		products: products,
		prices:   prices,
	}
	return cashier
}

func (this *Cashier) GetBill(product []int, amount []int) float64 {
	bill := 0
	for i := 0; i < len(product); i++ {
		for j := 0; j < len(this.products); j++ {
			if product[i] == this.products[j] {
				//match product
				bill += amount[i] * this.prices[j]
			}
		}
	}
	billAfter := float64(bill)
	if this.sum+1 == this.n {
		this.sum = 0
		billAfter = billAfter - billAfter*float64(this.discount)/100
	} else {
		this.sum++
	}
	return float64(billAfter)
}

/**
 * Your Cashier object will be instantiated and called as such:
 * obj := Constructor(n, discount, products, prices);
 * param_1 := obj.GetBill(product,amount);
 */
// @lc code=end
