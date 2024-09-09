# git 标签
[git 标签](https://git-scm.com/book/zh/v2/Git-%E5%9F%BA%E7%A1%80-%E6%89%93%E6%A0%87%E7%AD%BE)

标签主要用来记录版本、划时代意义
标签主要分为两种标签： 轻量级标签、注释标签
轻量级标签：基于某一次提交取别名，不会记录标签信息，及通过git show查询不到tag信息，没有存放到git数据库
注释标签: 基于某一个提交的tag，可以添加msg信息，同时会记录到git数据库，可以通过git show 查看到打标签的人和注释信息

## 标签列表
```linux
git tag
git tag -l
git tag -list

模糊查询
git tag -l 'v1.0*'
git tag -list 'v1.0*'
```

## 创建标签
注释标签
```linux
// 以当前最后一次提交创建tag
git tag -a v1.1 -m '这是我的v1.1 tag'

git tag -a v1.0  <commit_id> -m '基于指定commit id打标签'
```

轻量级标签:
```linux
git tag v1.1-lw

git tag v1.0-lw <commit_id>
```

## 查看标签存储情况
```linux
git show v1.1
git show v1.1-lw  // 这个是没有的
```

## 删除标签
```linux
git tag -d v1.1-lw

// 推送删除的标签
git push origin :refs/tags/v1.1-lw

git push origin --delete <tagname>
```


## 共享标签
```linux
git push origin v1.1

git push origin --tags 推送所有的标签
```