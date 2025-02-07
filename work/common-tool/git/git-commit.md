# git commit 相关操作
[文档](https://git-scm.com/book/zh/v2/Git-%E5%9F%BA%E7%A1%80-%E6%92%A4%E6%B6%88%E6%93%8D%E4%BD%9C)

## commit 的一个常见操作
```linux
git add aaa
git commit -m 'add aaa'
// 然后发现文件bbb漏提了，可以执行如下操作
git add bbb
git commit --amend 
```