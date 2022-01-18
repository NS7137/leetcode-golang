package lrucache

// 链表节点
type Node struct {
	key, val   int
	prev, next *Node
}

func initNode(key, val int) *Node {
	return &Node{key: key, val: val}
}

// 双向链表
type DoubleList struct {
	head, tail *Node // 虚拟头尾节点
	size       int   // 链表元素数
}

func initDoubleList() *DoubleList {
	dl := &DoubleList{
		head: initNode(0, 0),
		tail: initNode(0, 0),
		size: 0,
	}
	dl.head.next = dl.tail
	dl.tail.prev = dl.head
	return dl
}

// 尾部添加节点x
func (dl *DoubleList) addLast(x *Node) {
	// 新的连上再把前面的断开连接新的
	x.prev = dl.tail.prev
	x.next = dl.tail
	dl.tail.prev.next = x
	dl.tail.prev = x
	dl.size++
}

// 删除链表 x节点 x一定存在
func (dl *DoubleList) remove(x *Node) {
	x.prev.next = x.next
	x.next.prev = x.prev
	x.prev = nil
	x.next = nil
	dl.size--
}

// 删除第一个节点，并返回该节点
func (dl *DoubleList) removeFirst() *Node {
	if dl.head.next == dl.tail {
		return nil
	}
	first := dl.head.next
	dl.remove(first)
	return first
}

// 双链表从尾部插入，so 尾部是最近使用，头部是最久使用
type LRUCache struct {
	knMap    map[int]*Node // key -> Node(k,v)
	cache    *DoubleList   // Node(k1,v1) <-> Node(k2,v2)
	capacity int           // 容量
}

func Constructor(capacity int) LRUCache {
	lruCache := LRUCache{
		knMap:    map[int]*Node{},
		cache:    initDoubleList(),
		capacity: capacity,
	}
	return lruCache

}

func (lru *LRUCache) Get(key int) int {
	// 是否存在
	node, exist := lru.knMap[key]
	if !exist {
		return -1
	}
	// 提升为最近
	lru.makeRecently(key)
	return node.val
}

func (lru *LRUCache) Put(key int, value int) {
	_, exist := lru.knMap[key]
	if exist {
		// 已存在 更新key对应val 并提升为最近使用
		// 删旧
		lru.deleteKey(key)
		// 添新为最近使用
		lru.addRecently(key, value)
	} else {
		// 不存在 插入key
		if lru.cache.size == lru.capacity {
			// 容量满 删最久未使用
			lru.removeLeastRecently()
		}
		// 容量未满 插入k-v 为最近使用
		lru.addRecently(key, value)
	}

}

// 封装对get和put  同时维护cache和knMap对key操作
// key 提升为最近使用
func (lru *LRUCache) makeRecently(key int) {
	// 找节点
	node := lru.knMap[key]
	// 删节点
	lru.cache.remove(node)
	// 添加到尾部
	lru.cache.addLast(node)
}

// 添加最近使用元素
func (lru *LRUCache) addRecently(key int, val int) {
	// 初始化
	node := initNode(key, val)
	// 添加到尾部
	lru.cache.addLast(node)
	// 映射到 map
	lru.knMap[key] = node
}

// 删除一个key及对应元素
func (lru *LRUCache) deleteKey(key int) {
	// 找节点
	node := lru.knMap[key]
	// 删节点
	lru.cache.remove(node)
	// 删映射
	delete(lru.knMap, key)
}

// 删最久未使用元素
func (lru *LRUCache) removeLeastRecently() {
	// 链表头第一个就是
	deletedNode := lru.cache.removeFirst()
	// 删映射
	delete(lru.knMap, deletedNode.key)
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
