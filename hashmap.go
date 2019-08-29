package hashmap

import (
	"unsafe"
	"fmt"
	"hash/adler32"
)

type listNode struct {
	key string
	val string
	next *listNode
}

type HashMap struct {
	listArray  []*listNode
	size uint32
}

const (
	c1_32 uint32 = 0xcc9e2d51
	c2_32 uint32 = 0x1b873593

)

func hashFunc2(data []byte) uint32{
	return adler32.Checksum(data)
}

// GetHash returns a murmur32 hash for the data slice.
func hashFunc(data []byte) uint32 {
	// Seed is set to 37, same as C# version of emitter
	var h1 uint32 = 37
	nblocks := len(data) / 4
	var p uintptr
	if len(data) > 0 {
		p = uintptr(unsafe.Pointer(&data[0]))
	}

	p1 := p + uintptr(4*nblocks)
	for ; p < p1; p += 4 {
		k1 := *(*uint32)(unsafe.Pointer(p))

		k1 *= c1_32
		k1 = (k1 << 15) | (k1 >> 17) // rotl32(k1, 15)
		k1 *= c2_32

		h1 ^= k1
		h1 = (h1 << 13) | (h1 >> 19) // rotl32(h1, 13)
		h1 = h1*5 + 0xe6546b64
	}

	tail := data[nblocks*4:]

	var k1 uint32
	switch len(tail) & 3 {
	case 3:
		k1 ^= uint32(tail[2]) << 16
		fallthrough
	case 2:
		k1 ^= uint32(tail[1]) << 8
		fallthrough
	case 1:
		k1 ^= uint32(tail[0])
		k1 *= c1_32
		k1 = (k1 << 15) | (k1 >> 17) // rotl32(k1, 15)
		k1 *= c2_32
		h1 ^= k1
	}

	h1 ^= uint32(len(data))

	h1 ^= h1 >> 16
	h1 *= 0x85ebca6b
	h1 ^= h1 >> 13
	h1 *= 0xc2b2ae35
	h1 ^= h1 >> 16

	return (h1 << 24) | (((h1 >> 8) << 16) & 0xFF0000) | (((h1 >> 16) << 8) & 0xFF00) | (h1 >> 24)
}

func find(head *listNode, key string)(*listNode){
	var p *listNode
	p = head
	for p != nil{
		if p.key == key{
			// 找到
			break
		}
		// 没找到则继续找
		p = p.next
	}
	return p
}

func Init(size uint32) *HashMap{
	if size < 0 {
		return nil
	}

	hm := &HashMap{}
	hm.listArray = make([]*listNode, size)
	hm.size = size
	return hm
}

func (hm *HashMap)CalculateHashIndex(key string)(uint32){
	return hashFunc2([]byte(key)) % (hm.size)
}

func (hm *HashMap) Set(key, val string){
	var head *listNode
	var node *listNode
	k := hm.CalculateHashIndex(key)

	// 创建listNode
	node = &listNode{
		key: key,
		val: val,
		next: nil,
	}

	head = hm.listArray[k]
	if head == nil{
		// 没有代表当前hash链上没有数据
		hm.listArray[k] = node
	}else{
		// 代表链上有数据，则需要进行判断，当前key是否在链表中
		p := find(head, key)
		if p != nil{
			// 代表找到，则更新
			p.val = val
		}else{
			// 代表没找到，则新增到链表头部
			node.next = head
			hm.listArray[k] = node
		}
	}
}

func (hm *HashMap) Get(key string)(string, bool){
	var head *listNode
	var val string
	var exist bool
	k := hm.CalculateHashIndex(key)

	head = hm.listArray[k]
	if head == nil{
		// 代表没有
		goto DONE
	}else{
		// 代表该key有对应的链存在
		p := find(head, key)
		if p == nil{
			// 则代表不存在
			goto DONE
		}else{
			// 存在
			val = p.val
			exist = true
			goto DONE
		}
	}

DONE:
	return val, exist
}

func (hm *HashMap) Print(){
	for idx, head := range hm.listArray{
		p := head
		fmt.Print("index:", idx)
		if head != nil{
			for p != nil{
				fmt.Print(*p)
				p = p.next
			}
		}
		fmt.Println()
	}
}

