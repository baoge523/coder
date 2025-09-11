## 记录数据结构与算法相关的知识，方便以后回顾使用

### 数组

#### 两个有序数组的合并
nums1 = {1,3,5,0 ,0 ,0}
nums2 = {1,4,7}
返回nums1

解决思路：插入法的思路求解

#### 搜索二维矩阵II
https://leetcode.cn/problems/search-a-2d-matrix-ii/description/

#### 螺旋矩阵
https://leetcode.cn/problems/spiral-matrix/description/

#### 旋转图像
https://leetcode.cn/problems/rotate-image/description/

### 链表


#### 相交链表
https://leetcode.cn/problems/intersection-of-two-linked-lists/

### 二叉树

#### 求一颗二叉树的最大深度

解决方式：

方式一： 深度遍历，在递归的时候，判断是叶子节点时，就比较最大深度  -- 自己想出来的，代码比较复杂
方式二： 深度遍历，递归遍历，退出条件，当节点为nil时退出，遍历条件：递归遍历左子树和右子树的高度，判断其谁更高，高的+1
方式三： 广度遍历：基于队列的方式遍历，但是需要每次将一层的数据都从queue中取出；每取出一次，就深度+1

#### 二叉树的层级遍历
这里需要一层一层的遍历二叉树，使用队列的方式，需要根据题目做进一步的查看，但思路是一样的基于queue来实现

解决方式：
   基于queue的方式处理，考虑是当个节点出队，还是一层节点出队


#### 求一颗树是一个平衡二叉树
https://leetcode.cn/problems/balanced-binary-tree/
平衡二叉树的定义： 左子树 和 右子树都是平衡二叉树，同时左子树和右子树的高度差 <= 1

解决方式：
自上而下的递归遍历：
    1、退出条件节点为nil时返回true，表示nil时，是平衡二叉树
    2、求左子树的最大深度
    3、求右子树的最大深度
    4、比较两个子树的高度差，如果<= 1,那么就是平衡二叉树，否则不是

#### 二叉树的最近公共祖先
思路：
情况1:（左右子树包含目标节点）判断当前节点的左子树和右子树是否包含目标的节点，如果都包含那么当前节点就是最近公共节点
情况2:（左右子树只有一个包含目标节点），即当前节点就是目标节点的情况  --- 这种情况，就返回目标节点即可

解决方式：

方式一： 递归求解
  退出条件： 节点等于nil，或者当前节点是目标节点，就返回当前节点
  循环条件： 递归查找左子树 和 查找右子树，通过判断左子树和右子树是否包含目标节点，如果包含，那么当前节点就是最近父节点，否则就判断哪个不为nil然后返回，如果都为nil，那么就返回nil

这种方式的效率不高 -- 思考其他方式处理

#### 二叉树路径总和
https://leetcode.cn/problems/path-sum/description/

思路： 主要就是判断递归的退出条件，该退出条件就是，当节点是叶子节点的时候，判断当前的累加值和目标值是否一致
当前是叶子节点的判断规则： root != nil && root.Left == nil && root.Right == nil

解决方式：-- 递归处理
  退出条件：当前节点是叶子节点时，判断累加值是否等于目标值
  循环条件：分别处理左子树和右子树（处理前需要判断左右子树是否不为nil），得到结果后，需要将结果做与操作，表示只要有匹配到的，就返回true

#### 二叉树右视图
https://leetcode.cn/problems/binary-tree-right-side-view/description/
思路：
按层级遍历二叉树，每一层取最后一个节点即可，时间复杂度O(N),空间复杂度O(N)

解决方式：
   基于queue的层级遍历，然后每层取最后的数据即可

#### 验证二叉搜索树
https://leetcode.cn/problems/validate-binary-search-tree/description/

给你一个二叉树的根节点 root ，判断其是否是一个有效的二叉搜索树。
```text
有效 二叉搜索树定义如下：
    节点的左子树只包含 严格小于 当前节点的数。
    节点的右子树只包含 严格大于 当前节点的数。
    所有左子树和右子树自身必须也是二叉搜索树。
```

思路：
 在递归过程中，需要判断左右子树都是二叉搜索树，同时需要拿到左右子树的最大值和最小值与当前的节点比较
 判断左子树的最大值小于当前节点的值
 判断右子树的最小值大于当前节点的值
 需要注意获取最大值最小值时：需要考虑三种情况 1️⃣：左右子树都存在的情况  2️⃣：左子树存在，右子树不存在 3️⃣：右子树存在，左子树不存在

还有一个注意点，bool判断时，需要默认是满足的

```go
func isValidBST(root *TreeNode) bool {
    if root == nil {
        return true
    }
    b,_,_ := isValidBSTReturn(root)
    return b
}

func isValidBSTReturn(root *TreeNode) (bool,*TreeNode,*TreeNode) {
	// 如果是叶子节点 那么就返回true，且最大值，最小值都是自己
    if root != nil && root.Left == nil && root.Right == nil {
        return true,root,root
    }
    // 左子树或者右子树不存在时，默认是满足的
    left := true
    right := true
    var lmin,lmax,rmin,rmax *TreeNode
    if  root.Left != nil {
		// 判断左子树是否是二叉搜索树，并且返回其最大的节点和最小的节点
        left,lmin,lmax = isValidBSTReturn(root.Left)
    }
    if root.Right != nil {
		// 判断右子树是否是二叉搜索树，并且返回其最大的节点和最小的节点
         right,rmin,rmax = isValidBSTReturn(root.Right)
    }
    // 与当前节点比较，判断当前节点树，是不是二叉搜索树
    leftValid := true
    rightValid := true
    if lmax != nil {
        leftValid = root.Val > lmax.Val
    }
    if rmin != nil {
       rightValid = root.Val < rmin.Val
    }
	// 返回当前树的最大值和最小值，这里需要考虑的情况是：如果左子树或者右子树不存在时，需要用当前节点代替
    min := root
    max := root
    if lmin != nil {
        min = lmin
    }
    if rmax != nil {
        max = rmax
    }
    return left && right && leftValid && rightValid,min,max
}
```

#### 寻找二叉搜索树中的目标节点
https://leetcode.cn/problems/er-cha-sou-suo-shu-de-di-kda-jie-dian-lcof/description/

二叉搜索树：左子树的任意值 < 根 < 右子树的任意值

中序遍历： 左根右 可以顺序遍历二叉搜索树的所有节点 ----》逆遍历  --〉 逆中序   右根左 --》按大小逆遍历

```go
// 基于中序遍历的方式处理
 // 中序 左根右   中逆序 右根左

func findTargetNode(root *TreeNode, cnt int) int {
    i := 0 // i从0到k
    ans := 0
    var dfs func(n *TreeNode) // 声明递归函数
    dfs = func(n *TreeNode) {
        if n == nil {return} 
        dfs(n.Right)  // 顺序遍历： 中序： 左根右  逆中序 右根左
        i++
        if i == cnt {ans = n.Val; return}
        dfs(n.Left)
    }
    dfs(root)
    return ans
}

```

####  二叉树的直径
https://leetcode.cn/problems/diameter-of-binary-tree/description/
给你一棵二叉树的根节点，返回该树的 直径 。
二叉树的 直径 是指树中任意两个节点之间最长路径的 长度 。这条路径可能经过也可能不经过根节点 root 。
两节点之间路径的 长度 由它们之间边数表示。

思路：
1、需要一个额外的全局变量来记录最长路径
2、深度遍历，递归遍历每一个节点，自下而上的遍历，遍历每一个节点时，需要分两步判断
3、 第一步：左子树的路径 + 右子树的路径 == 当前节点的路径，需要和最长路径做判断，如果是最长的，需要记录一下
4、 第二步：判断左右子树哪个路径更长，将其返回给上层



#### 对称二叉树  -- 没有一次性做出来
https://leetcode.cn/problems/symmetric-tree/description/
给你一个二叉树的根节点 root ， 检查它是否轴对称。是基于root根节点进行对称的

思路是：递归遍历，左子树的左，右和右子树的右，左 进行比较

退出条件：
  1、左右子节点同时为nil时，返回true
  2、左右子节点一个为nil，一个不为nil时，返回false

循环条件：
  循环判断左节点的左和右节点的右
   左节点的右和右节点的左
  同时当前的左右节点的值需要是相同的


迭代的思路：基于队列的方式层级遍历，但是层级遍历需要将相互比较的节点同时存放到queue中

1、将root两次加入到queue中（主要的目的就是方便取出root的左右进行比较）
2、循环处理queue，每次从队列中取出两个节点出来，包含nil节点，然后比较则两个节点，判断是否相同（是否为nil，是否值相同）
3、如果不相同就返回false退出，如果相同，就将其left.Left和right.Right 、left.Right和right.Left 放入队列中
4、一直循环直到遍历结束返回


#### 翻转二叉树  -- 第一时间没有做出来
https://leetcode.cn/problems/invert-binary-tree/description/
给你一棵二叉树的根节点 root ，翻转这棵二叉树，并返回其根节点。

思路：递归的方式，自下而上的将所有的节点都翻转了，就完成了二叉树的翻转

```go
 // 思路： 递归翻转每一个节点的左右节点，叶子节点的时候不翻转
func invertTree(root *TreeNode) *TreeNode {
    if root == nil {
        return root
    }
    left := invertTree(root.Left) // 翻转左子树
    right := invertTree(root.Right) // 翻转右子树
	// 翻转当前节点，并返回
    root.Left = right 
    root.Right = left
    return root
}
```


#### 处理二叉树的思维方式
1、看整体需要达到的效果，但不要陷入到整体中去，学会化整为零
2、将二叉树拆分成一个节点（叶子节点）、两个节点（一个父节点，一个子节点）、三个节点（一个父节点、两个子节点）
3、分别将上面的三种场景拿来整理成公式处理






































