package main

import "fmt"

// 加权轮询的实现
// 3A  2B 1C  = AAA BB C
// 平滑加权轮询实现
// 3A  2B 1C  = A A B A B C
func main() {
	// 7A 2B 1C  平滑加权轮询
	// 初始权重: 上次的当前权重 + 权重
	// 最大值：从初始权重中取出最大值
	// total权重值: 所有的权重值
	// 返回结果
	// 当前权重: 初始权重中的选出值 - total权重值

	nodes := []Node{
		Node{Name: "A", Weight: 7},
		Node{Name: "B", Weight: 2},
		Node{Name: "C", Weight: 1},
	}
	for i:=0; i<10; i++ {
        name := smoothWeightRoundRobin(nodes)
        fmt.Println(name)
    }
}

type Node struct {
	Name string
	Weight int  // 本身的权重，一直不变
	currentWeight int // 当前权重，会变化，变化规则是 当前权重 = 上次的当前权重 + 权重，选择最大的当前权重值，然后减去总的权重值，得到选择一个节点后的当前权重值
}
// 思想：每次开始选择节点时，将各个节点的当前权重 +=权重(各个节点的)，当选中了一个节点后，该节点的当前权重值减去总的权重值
// 平滑加权轮询
func smoothWeightRoundRobin(nodes []Node)(string) {

	if len(nodes) == 0 {
        return ""
    }
	if len(nodes) == 1 {
        return nodes[0].Name
    }

	index :=0
	total :=0
	max := 0
	for i, node := range nodes {
		nodes[i].currentWeight += node.Weight
		// 计算总的权重值
		total += node.Weight
		if node.currentWeight > max {
            max = node.currentWeight
			index = i
        }
	}
	nodes[index].currentWeight -= total
	return nodes[index].Name
}