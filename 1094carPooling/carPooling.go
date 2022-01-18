package carpooling

import "leetcode-golang/utils"

func CarPooling(trips [][]int, capacity int) bool {
	nums := make([]int, 1001)

	df := utils.Diffience(nums)

	for _, trip := range trips {
		val := trip[0]
		i := trip[1]
		j := trip[2] - 1
		df.Increment(i, j, val)
	}

	res := df.Result()

	for i := 0; i < len(res); i++ {
		if capacity < res[i] {
			return false
		}
	}

	return true

}
