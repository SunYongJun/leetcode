package MyLeetCode

/**

50. Pow(x, n)
实现 pow(x, n) ，即计算 x 的 n 次幂函数。

示例 1:

输入: 2.00000, 10
输出: 1024.00000
示例 2:

输入: 2.10000, 3
输出: 9.26100
示例 3:

输入: 2.00000, -2
输出: 0.25000
解释: 2-2 = 1/22 = 1/4 = 0.25
说明:

-100.0 < x < 100.0
n 是 32 位有符号整数，其数值范围是 [−231, 231 − 1] 。

*/

func myPow(x float64, n int) float64 {

	if n == 0 {
		return float64(1)
	}

	var ans = x

	var pre float64 = 1

	if x < 0 && n % 2 == 1 {
		pre = -1
	}

	max := float64(1 << 31)
	for i := 1; i < int(math.Abs(float64(n))); i++ {
		if ans == 0 {
			return 0
		}
		if ans == 1 {
			return 1 * pre
		}

		ans *= x
		if ans > max {
			ans = max
			break
		}
	}

	if n < 0 {
		ans = 1 / ans
	}

	return math.Abs(ans) * pre
}