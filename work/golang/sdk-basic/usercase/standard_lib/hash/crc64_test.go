package hash

import (
	"fmt"
	"hash/crc64"
	"testing"
)

/*
在CRC64中，ECMA和ISO分别表示不同的多项式和计算标准：
1. **ECMA**: 表示ECMA-182标准使用的CRC64多项式，通常为`0x42F0E1EBA9EA3693`。这种标准主要用于文件系统和数据存储。
2. **ISO**: 表示ISO/IEC 3309标准使用的CRC64多项式，通常为`0xA5EB000000000000`。该标准常用于通信协议和数据传输。
*/

func TestCRC64(t *testing.T) {

	table := crc64.MakeTable(crc64.ECMA)

	cont := []byte("hello crc64") // 4608898353696507970
	a := "hello"
	b := []byte(a + " crc64")
	checksum := crc64.Checksum(cont, table)
	fmt.Printf("%v \n", checksum)
	checksum2 := crc64.Checksum(b, table)
	fmt.Printf("%v \n", checksum2)
}

func TestCRC64Hash(t *testing.T) {
	table := crc64.MakeTable(crc64.ECMA)
	hash64 := crc64.New(table)
	cont := []byte("hello crc64") // 4608898353696507970
	hash64.Write(cont)
	sum64 := hash64.Sum64()
	fmt.Printf("%v \n", sum64)
}
