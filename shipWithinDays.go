package lesson5

//ShipWithinDays: 在 D 天内送达包裹的能力
/*parameter
weights: 包裹重量数组
days: 给定运送包裹的天数
return: 满足days运送天数的最低运力
*/
func ShipWithinDays(weights []int, days int) int {

	//最重包裹与包裹总重量
	left, right := 0, 0
	for i := range weights {

		if weights[i] > left {

			left = weights[i]
		}

		right += weights[i]
	}

	//计算x指定的包裹运力所需运送天数是否小于等于指定天数days
	calcDays := func(x int) bool {

		d, inDay := 0, 0
		for i := range weights {

			if x < weights[i] {

				return false
			}

			if inDay+weights[i] > x {

				d++
				inDay = 0
			}

			inDay += weights[i]
		}

		if inDay > 0 {

			d++
		}

		return d <= days
	}

	//二分计算满足days运送天数的最低运力
	for left < right {

		mid := left + (right-left)>>1
		if calcDays(mid) {

			right = mid
		} else {

			left = mid + 1
		}
	}

	return right
}
