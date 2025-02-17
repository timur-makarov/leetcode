package main

func main() {
	//sorts.RunSortFunctions()

	//sqrtPriceX96, _ := big.NewInt(0).SetString("79223684291618765335787616787", 10)
	//liquidity, _ := big.NewInt(0).SetString("599150462748824", 10)
	//tick := 3
	//
	//rA, rB := calculateLiquidity2(tick, liquidity, sqrtPriceX96)
	//log.Println(rA.String(), rB.String())
}

//
//func calculateLiquidity(tick int, liquidity, sqrtPriceX96 *big.Int) (*big.Float, *big.Float) {
//	x96 := big.NewInt(0).Exp(big.NewInt(2), big.NewInt(96), nil)
//	x96float := big.NewFloat(0).SetInt(x96)
//	sqrtPriceX96float := big.NewFloat(0).SetInt(sqrtPriceX96)
//	sqrtPrice := big.NewFloat(0).Quo(sqrtPriceX96float, x96float)
//
//	liquidityFloat := big.NewFloat(0).SetInt(liquidity)
//	x18Float := new(big.Float).SetInt(big.NewInt(0).Exp(big.NewInt(10), big.NewInt(18), nil))
//	x6Float := new(big.Float).SetInt(big.NewInt(0).Exp(big.NewInt(10), big.NewInt(6), nil))
//
//	var reserveA, reserveB *big.Float
//
//	if tick >= 0 {
//		reserveA = new(big.Float).Quo(big.NewFloat(0).Mul(liquidityFloat, sqrtPrice), x18Float)
//		reserveB = new(big.Float).Quo(big.NewFloat(0).Mul(liquidityFloat, x6Float), sqrtPrice)
//	} else {
//		reserveA = new(big.Float).Quo(liquidityFloat, sqrtPrice)
//		reserveA = new(big.Float).Quo(reserveA, x18Float)
//		reserveB = new(big.Float).Mul(liquidityFloat, sqrtPrice)
//		reserveB = new(big.Float).Quo(reserveB, x6Float)
//	}
//
//	return reserveA, reserveB
//}
//
//func calculateLiquidity2(currentTick int, liquidity, sqrtPriceX96 *big.Int) (*big.Float, *big.Float) {
//	var reserveA, reserveB *big.Float
//
//	x96 := big.NewInt(0).Exp(big.NewInt(2), big.NewInt(96), nil)
//	x96float := big.NewFloat(0).SetInt(x96)
//	sqrtPriceX96float := big.NewFloat(0).SetInt(sqrtPriceX96)
//
//	liquidityFloat := new(big.Float).SetInt(liquidity)
//
//	sqrtPrice := big.NewFloat(0).Quo(sqrtPriceX96float, x96float)
//
//	//sqpc := new(big.Int).Div(sqrtPriceX96, x96)
//	//sqpc = new(big.Int).Mul(sqpc, sqpc)
//	//reserveAB := new(big.Int).Div(liquidity, sqpc)
//	//x18 := big.NewInt(0).Exp(big.NewInt(10), big.NewInt(18), nil)
//	//reserveAB = big.NewInt(0).Div(reserveAB, x18)
//	//log.Println(reserveAB.String())
//	//
//	priceX96 := big.NewFloat(0).Mul(sqrtPrice, sqrtPrice)
//	sqrtPrice = new(big.Float).Sqrt(priceX96)
//	//
//	log.Println(priceX96.String())
//	log.Println(sqrtPrice.String())
//
//	//if currentTick >= 0 {
//	//	reserveA = new(big.Float).Quo(new(big.Float).Mul(liquidityFloat, sqrtPrice), big.NewFloat(math.Pow10(18)))
//	//	reserveB = new(big.Float).Quo(big.NewFloat(0).Mul(liquidityFloat, big.NewFloat(math.Pow10(6))), sqrtPrice)
//	//} else {
//	//	reserveA = new(big.Float).Quo(liquidityFloat, sqrtPrice)
//
//	//}
//
//	//feeGrowth0 := new(big.Int).SetString("6113454639202753443888423812296906", 10)
//	//feeGrowth1 := new(big.Int).SetString("52162661196415993125010544162617876953", 10)
//
//	reserveA = new(big.Float).Mul(liquidityFloat, new(big.Float).Sqrt(new(big.Float).Quo(big.NewFloat(1), priceX96)))
//	reserveB = new(big.Float).Mul(liquidityFloat, sqrtPrice)
//	log.Println(reserveA.String())
//	log.Println(reserveB.String())
//	//
//	//sqrtPrice = new(big.Float).Sub(sqrtPrice, big.NewFloat(1))
//	//sp := new(big.Float).Quo(big.NewFloat(1), sqrtPrice)
//	//reserveA = new(big.Float).Mul(liquidityFloat, sp)
//	//reserveA = new(big.Float).Quo(reserveA, x96float)
//	//
//	x6Float := new(big.Float).SetInt(big.NewInt(0).Exp(big.NewInt(10), big.NewInt(6), nil))
//
//	reserveB = new(big.Float).Mul(liquidityFloat, sqrtPrice)
//	//reserveB = new(big.Float).Quo(reserveB, x96float)
//	reserveB = new(big.Float).Quo(reserveB, x6Float)
//
//	//log.Println(reserveA.String())
//
//	TICK_SPACING := 1
//	sqrtPriceLow := big.NewFloat(math.Pow(1.0001, float64(currentTick/2)))
//	sqrtPriceHigh := big.NewFloat(math.Pow(1.0001, float64((currentTick+TICK_SPACING)/2)))
//
//	sbsp := new(big.Float).Sub(sqrtPriceHigh, sqrtPrice)
//	spsb := new(big.Float).Mul(sqrtPrice, sqrtPriceHigh)
//	ssd := new(big.Float).Quo(sbsp, spsb)
//
//	reserveA.Mul(liquidityFloat, ssd)
//
//	spsa := new(big.Float).Sub(sqrtPriceHigh, sqrtPriceLow)
//
//	reserveB.Mul(liquidityFloat, spsa)
//
//	reserveA.Quo(reserveA, x6Float)
//	reserveB.Quo(reserveB, x6Float)
//
//	return reserveA, reserveB
//}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	var dfs func(*TreeNode) *TreeNode
	dfs = func(node *TreeNode) *TreeNode {
		if node.Val > p.Val && node.Val > q.Val {
			return dfs(node.Left)
		}

		if node.Val < p.Val && node.Val < q.Val {
			return dfs(node.Right)
		}

		return node
	}

	return dfs(root)
}
