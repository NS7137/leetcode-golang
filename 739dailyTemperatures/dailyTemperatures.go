package dailytemperatures

func DailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	stack := []int{}
	res := make([]int, n)
	// 下标压栈, 将遍历到的下标与出栈的下标 温度比较，大于的，下标差放结果集
	for i := 0; i < n; i++ {
		for len(stack) != 0 && temperatures[stack[len(stack)-1]] < temperatures[i] {
			res[stack[len(stack)-1]] = i - stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	return res

}
