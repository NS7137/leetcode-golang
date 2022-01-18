package lfucache

import "container/list"

/*
LFU 逻辑
调用 get(key) 方法，要返回该 key 对应的 val
要用 get 或 put 方法访问一次某 key，该 key 的 freq 就要加一
如果容量满时进行插入，则将frq最小的key删除，如果最小freq对应多个key，删除最久的那个

KV, KF, FK 三个映射
*/
type LFUCache struct {
	nodes    map[int]*list.Element //key-node,  keyToVal keyToFreq
	lists    map[int]*list.List    //freq-node  freq to key
	capacity int
	minFreq  int
}

type Node struct {
	key   int
	value int
	freq  int
}

func Constructor(capacity int) LFUCache {
	lfuCache := LFUCache{
		nodes:    make(map[int]*list.Element),
		lists:    make(map[int]*list.List),
		capacity: capacity,
		minFreq:  0,
	}
	return lfuCache
}

// 返回key对应val，增加key对应freq
func (lfu *LFUCache) Get(key int) int {
	el, exist := lfu.nodes[key]
	if !exist {
		return -1
	}
	node := el.Value.(*Node)
	// 增加key对应freq
	lfu.increaseFreq(key)
	return node.value
}

func (lfu *LFUCache) Put(key int, value int) {
	if lfu.capacity <= 0 {
		return
	}

	el, exist := lfu.nodes[key]
	// key 存在
	if exist {
		node := el.Value.(*Node)
		// 修改key对应val
		node.value = value
		// 增加key对应freq
		lfu.increaseFreq(key)
		return
	} else {
		// key不存在
		// 容量是否满
		if lfu.capacity <= len(lfu.nodes) {
			// 满则淘汰freq最小key
			lfu.removeMinFreqKey()
		}
		// 插入key val，freq为1
		n := &Node{key: key, value: value, freq: 1}
		// 插入FK表
		if _, ok := lfu.lists[1]; !ok {
			lfu.lists[1] = list.New()
		}
		list := lfu.lists[1]
		newnode := list.PushBack(n)
		// 更新 key node, kv kf
		lfu.nodes[key] = newnode
		// 更新最小freq
		lfu.minFreq = 1
	}
}

func (lfu *LFUCache) removeMinFreqKey() {
	// 最小freq  key表
	minFreqkeyList := lfu.lists[lfu.minFreq]
	// 最先插入的key被淘汰
	frontNode := minFreqkeyList.Front()
	deletedKey := frontNode.Value.(*Node).key
	// 更新 FK
	minFreqkeyList.Remove(frontNode)
	if minFreqkeyList.Len() == 0 {
		delete(lfu.lists, lfu.minFreq)
	}
	// 更新 KV KF
	delete(lfu.nodes, deletedKey)
}

func (lfu *LFUCache) increaseFreq(key int) {
	el := lfu.nodes[key]
	node := el.Value.(*Node)
	// 更新 FK
	// 将key从freq对应列表中删除
	lfu.lists[node.freq].Remove(el)

	// 更新 KF
	node.freq++
	// 将key加入freq + 1 列表
	if _, ok := lfu.lists[node.freq]; !ok {
		lfu.lists[node.freq] = list.New()
	}
	newlist := lfu.lists[node.freq]
	newnode := newlist.PushBack(node)
	// 更新 key node
	lfu.nodes[key] = newnode

	// 如果这个 freq 恰好是 minFreq，更新 minFreq
	if node.freq-1 == lfu.minFreq && lfu.lists[node.freq-1].Len() == 0 {
		lfu.minFreq++
	}
}
