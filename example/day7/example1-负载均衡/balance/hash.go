package balance

//自定义扩展的hash算法

import (
	"fmt"
	"math/rand"
	"hash/crc64"
	"hash/crc32"
)

func init() {
	RegisterBalance("hash",&HashBalance{})
}

type HashBalance struct {

}

func (p *HashBalance) DoBalance(insts []*Instance,key ...string) (inst *Instance,err error) {
	//如果用户未传递计算hash的key值，使用默认
	defKey := fmt.Sprintf("%d",rand.Int())
	if len(key) > 0 {
		defKey = key[0]
	}

	//判断后端实例数量
	lens := len(insts)
	if lens <= 0 {
		err = fmt.Errorf("No backend instance!")
		return
	}
	//crc64.Checksum求hash值，返回uint64类型
	crcTable := crc32.MakeTable(crc32.IEEE)
	hashVal := crc32.Checksum([]byte(defKey),crcTable)
	index := int(hashVal) % lens
	inst = insts[index]
	return
}