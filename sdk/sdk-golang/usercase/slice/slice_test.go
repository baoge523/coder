package slice

import (
	"fmt"
	"testing"
)

func TestSliceZeroLen(t *testing.T){
	var fruits []string
	fmt.Printf("len = %d, cap = %d \n",len(fruits),cap(fruits))
	fruits[0] = "aa"  // 当前的fruits没有空间，所以通过索引方式赋值，会报索引越界 index out of range [0] with length 0
}



// go test -run TestSliceAppend
func TestSliceAppend(t *testing.T) {

	var fruits []string
	fruits = append(fruits,"apple")
	fmt.Printf("len = %d, cap = %d \n",len(fruits),cap(fruits))
	fmt.Printf("heap address %p\n", &fruits[0])
	fruits = append(fruits,"banana")
	fmt.Printf("len = %d, cap = %d \n",len(fruits),cap(fruits))
	fmt.Printf("heap address %p\n",&fruits[0])
	fruits = append(fruits,"pear")
	fmt.Printf("len = %d, cap = %d \n",len(fruits),cap(fruits))
	fmt.Printf("heap address %p\n",&fruits[0])
	fruits = append(fruits,"grape")
	fmt.Printf("len = %d, cap = %d \n",len(fruits),cap(fruits))
	fmt.Printf("heap address %p\n",&fruits[0])
	fruits = append(fruits,"litchi")
	fmt.Printf("len = %d, cap = %d \n",len(fruits),cap(fruits))
	fmt.Printf("heap address %p\n",&fruits[0])
	fruits = append(fruits,"peach")
	fmt.Printf("len = %d, cap = %d \n",len(fruits),cap(fruits))
	fmt.Printf("heap address %p\n",&fruits[0])  // 数组第一个元素的地址
	fmt.Printf("heap address %p\n",&fruits)  // 数组指针的地址

	// len = 1, cap = 1
	// len = 2, cap = 2
	// len = 3, cap = 4
	// len = 4, cap = 4
	// len = 5, cap = 8
	// len = 6, cap = 8
}

func TestSliceMake(t *testing.T) {
	fruits := make([]string,1,3)
	fmt.Printf("init len = %d, cap = %d \n",len(fruits),cap(fruits))
	fruits = append(fruits,"apple")
	fmt.Printf(" len = %d, cap = %d \n",len(fruits),cap(fruits))
	fruits = append(fruits,"peach")
	fmt.Printf(" len = %d, cap = %d \n",len(fruits),cap(fruits))
	fruits = append(fruits,"litchi")
	fmt.Printf(" len = %d, cap = %d \n",len(fruits),cap(fruits))
	fruits = append(fruits,"banana")
	fmt.Printf("len = %d, cap = %d \n",len(fruits),cap(fruits))
	fruits = append(fruits,"pear")
	fmt.Printf("len = %d, cap = %d \n",len(fruits),cap(fruits))
	fruits = append(fruits,"grape")
	fmt.Printf("len = %d, cap = %d \n",len(fruits),cap(fruits))

	// 结论: make 可以指定slice的len 和cap,如果在之后使用append，
	//      会从len之后开始存储，当cap不够用时，会扩容到原来的两倍
}
// 切片的[:]方式操作切片，存在内存共享问题 - 需要注意
func TestSliceShareMemory(t *testing.T) {
	fruits := make([]string,0,5)
	fruits = append(fruits,"apple")
	fruits = append(fruits,"peach")
	fruits = append(fruits,"banana")
	fruits = append(fruits,"grape")
	fruits = append(fruits,"litchi")
	fmt.Printf("fruits = %v \n",fruits)
	fru1 := fruits[:3]  // 0 1 2   这创建出来的slice len = 3 , cap = 5, 其中就会有2个位置会共享内存，在未扩容前,使用索引3 4 会变更到fru2的内容
	fru2 := fruits[3:]  // 3 4     这里创建出来的slice len = 2 cap = 2
	fmt.Printf("fru1 = %v len = %d, cap = %d\n",fru1,len(fru1),cap(fru1))
	fmt.Printf("fru2 = %v len = %d, cap = %d\n",fru2,len(fru2),cap(fru2))
	fru1 = append(fru1,"pear")
	fmt.Printf("fru1 = %v len = %d, cap = %d\n",fru1,len(fru1),cap(fru1))
	fmt.Printf("fru2 = %v len = %d, cap = %d\n",fru2,len(fru2),cap(fru2))
	fru1 = append(fru1,"hahaha")
	fmt.Printf("fru1 = %v len = %d, cap = %d\n",fru1,len(fru1),cap(fru1))
	fmt.Printf("fru2 = %v len = %d, cap = %d\n",fru2,len(fru2),cap(fru2))
	fru1 = append(fru1,"heiheihei")
	fmt.Printf("fru1 = %v len = %d, cap = %d\n",fru1,len(fru1),cap(fru1))
	fmt.Printf("fru2 = %v len = %d, cap = %d\n",fru2,len(fru2),cap(fru2))
	fru2 = append(fru2,"hello")
	fmt.Printf("fru1 = %v len = %d, cap = %d\n",fru1,len(fru1),cap(fru1))
	fmt.Printf("fru2 = %v len = %d, cap = %d\n",fru2,len(fru2),cap(fru2))

	fru1[3] = "xxxxx"
	fmt.Printf("fru1 = %v len = %d, cap = %d\n",fru1,len(fru1),cap(fru1))
	fmt.Printf("fru2 = %v len = %d, cap = %d\n",fru2,len(fru2),cap(fru2))

	// 结论: 在使用截断数组时，应该使用 fru := fruits[:3:3] 表示 len = 3 cap = 3 就不会存在共享内存了

}