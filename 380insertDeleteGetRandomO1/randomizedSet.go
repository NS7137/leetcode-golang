package insertdeletegetrandomo1

import "math/rand"

/*
getRandom 要 O1，所以底层用数组实现且必须是紧凑的
插入，删除 O1，对数组尾部进行插入删除不会涉及数据搬移
删除时交换val到尾部，再pop，数据交换通过hash表来记录valToIndex
*/
type RandomizedSet struct {
	vals       []int
	valToIndex map[int]int
}

func Constructor() RandomizedSet {
	set := RandomizedSet{
		vals:       []int{},
		valToIndex: make(map[int]int),
	}
	return set
}

// val不存在时，向集合插入该项
func (rs *RandomizedSet) Insert(val int) bool {
	// val 已存在，不用再插入
	_, exist := rs.valToIndex[val]
	if exist {
		return false
	}
	// 不存在，插入到vals尾部
	rs.vals = append(rs.vals, val)
	// 并记录val索引
	rs.valToIndex[val] = len(rs.vals) - 1
	return true

}

// val存在时，从集合移除该项
func (rs *RandomizedSet) Remove(val int) bool {
	// val 不存在，不用删除
	index, exist := rs.valToIndex[val]
	if !exist {
		return false
	}
	lastIndex := len(rs.vals) - 1
	// 存在
	// swap val 和 last
	rs.vals[index], rs.vals[lastIndex] = rs.vals[lastIndex], rs.vals[index]
	// 修改元素对应的索引
	rs.valToIndex[rs.vals[index]] = index
	// pop back
	rs.vals = rs.vals[:lastIndex]
	// 删 val 对应索引
	delete(rs.valToIndex, val)
	return true
}

// 随机返回现有集合中的一项，每个元素概率相同
func (rs *RandomizedSet) GetRandom() int {
	return rs.vals[rand.Intn(len(rs.vals))]
}
