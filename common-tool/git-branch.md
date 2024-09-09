## git branch相关的操作

[分支管理](https://git-scm.com/book/zh/v2/Git-%E5%88%86%E6%94%AF-%E5%88%86%E6%94%AF%E7%AE%A1%E7%90%86)


查看git命令的相关帮助
git branch --help
git checkout --help

### git branch
git branch 可以查看分支列表，创建分支、删除分支、查看合并状态等操作

#### 查看分支信息
```linux
git branch
git branch -v
git branch -vv
```

#### 创建分支
```linux
git branch testing
```
#### 删除分支
```linux
git branch -d aaa
git branch -D bbb  // -D表示强制删除，因为在删除的时候可能会报错说，删除的分支没有合并到HEAD分支上
```

#### 查看合并状态
查看与当前HEAD分支已经合并的分支,有包含没有带*的，都是可以删除的分支
```linux
git branch --merged
```

查看与当前HEAD分支没有合并的分支
```linux
git branch --no-merged
```

查看尚未合并到指定分支的其他分支
```linux
git branch --merged master   // 查看没有合并到master的所有分支
```


### 实用操作
1、基于远程分支创建本地分支
> git checkout -b local_branch origin/remote_branch </br>
> 比如基于远程分支master创建本地分支dev  </br>
> git checkout -b dev origin/master

2、基于本地分支创建本地分支
> git checkout -b new_branch old_branch  </br>
> 比如基于本地分支dev创建本地分支master   </br>
> git checkout -b master dev  

3、将本地分支推送到远程(但是远程分支不存在)
> git push -u origin local_branch:remote_branch </br>
> 比如将本地分支dev推送到远端去(远端没有dev分支)  </br>
> git push -u origin dev:dev


```linux
git checkout -b aaa
==
git branch aaa
git checkout aaa

```