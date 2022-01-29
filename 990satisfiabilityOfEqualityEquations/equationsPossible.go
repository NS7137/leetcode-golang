package satisfiabilityofequalityequations

import "leetcode-golang/utils"

func EquationsPossible(equations []string) bool {
	// 26 字母
	uf := utils.NewUF(26)
	// 相等的字母形成连通分量
	for _, eq := range equations {
		if eq[1] == '=' {
			x := int(eq[0] - 'a')
			y := int(eq[3] - 'a')
			uf.Union(x, y)
		}
	}
	// 检查不等关系是否打破相等关系的连通性
	for _, eq := range equations {
		if eq[1] == '!' {
			x := int(eq[0] - 'a')
			y := int(eq[3] - 'a')
			// 如果相等关系成立，就逻辑冲突
			if uf.Connected(x, y) {
				return false
			}
		}
	}
	return true
}
