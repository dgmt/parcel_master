package parcelcalc

var packages = []int64{250, 500, 1000, 2000, 5000}

func ParcelCalculator(orderAmt int64) map[int64]int64 {
	res := roughPack(orderAmt)
	ord := packageNumOptimization(res)
	return ord
}

func roughPack(orderAmt int64) (order map[int64]int64) {
	order = make(map[int64]int64)
	max := 0
	// If ordered 0, return empty map
	if orderAmt == 0 {
		return
	}
	// Find closest package size
	for i, item := range packages {
		if orderAmt >= item {
			max = i
		}
	}
	// this is a workaround for >250 <500
	if max == 0 && orderAmt <= packages[max] {
		order[packages[max]] = 1
		return
	}
	// calc the main and remainder
	biggest := orderAmt / packages[max]
	order[packages[max]] = biggest
	remainder := orderAmt % packages[max]
	// send remainder to recursion
	if remainder > 0 {
		remain := roughPack(remainder)
		for key, value := range remain {
			order[key] = order[key] + value
		}
	}
	return
}

func packageNumOptimization(orderUnopt map[int64]int64) (orderOpt map[int64]int64) {
	// optimization based on trivial example in the challenge,
	// with more heuristics, more strategies could be done
	orderOpt = make(map[int64]int64)
	for key, value := range orderUnopt {
		total := key * value
		if contains(total) {
			orderOpt[total] = orderOpt[total] + 1
		} else {
			orderOpt[key] = value
		}
	}
	return
}

func contains(num int64) bool {
	// there are external libs to do this
	// but since this is deployed to appengine,
	// i prefer it nice and tight and in-house
	for _, item := range packages {
		if item == num {
			return true
		}
	}
	return false
}
