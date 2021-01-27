package MyLeetCode

/**

7. 整数反转
给你一个 32 位的有符号整数 x ，返回 x 中每位上的数字反转后的结果。

如果反转后整数超过 32 位的有符号整数的范围 [−231,  231 − 1] ，就返回 0。

假设环境不允许存储 64 位整数（有符号或无符号）。


示例 1：

输入：x = 123
输出：321
示例 2：

输入：x = -123
输出：-321
示例 3：

输入：x = 120
输出：21
示例 4：

输入：x = 0
输出：0


提示：

-231 <= x <= 231 - 1

*/

func reverse(x int) int {
	var list = make([]int,0)
	pre := 1
	if x < 0 {
		pre = -1
		x = int(math.Abs(float64(x)))
	}
	i := 1
	for  i <= x {
		list = append(list,x % (i * 10) / i)
		i*=10
	}

	var ans = 0
	l := 0
	for  i >= 1 && l < len(list)  {
		i = i / 10
		ans += i * list[l]
		l++
	}

	max := 1 << 31
	if ans >= max {
		ans = 0

	}

	return ans * pre
}