# git 远程仓库的使用
[remote的使用](https://git-scm.com/book/zh/v2/Git-%E5%9F%BA%E7%A1%80-%E8%BF%9C%E7%A8%8B%E4%BB%93%E5%BA%93%E7%9A%84%E4%BD%BF%E7%94%A8)

当我们在本地执行git clone 命令时，会将远端的master分支拉取下来，拉取成功后，就会创建一个叫做origin的远端信息,可以通过git remote查看

## git remote 查看所有远程仓库的信息

### 查看远端仓库信息
> git remote
> 
> git remote -v

### 添加远程仓库的信息
```linux
git remote add <shortname> <url>
比如
git remote add origin https://github.com/xxx/xx
git remote add pd https://github.com/aa/bb
```

### 重命名和删除远端仓库信息
```linux
重命名
git remote rename <old_name> <new_name> 

删除
git remote remove origin
```

### 查看指定的远端仓库信息
```linux
git remote show origin
```

### 拉取和推送远程仓库
```linux
git fetch <remote>
比如:
git fetch origin
必须注意 git fetch 命令只会将数据下载到你的本地仓库——它并不会自动合并或修改你当前的工作。 当准备好时你必须手动将其合并入你的工作。
如果设置了本地分支和远程分支的跟踪，那么可以使用git pull，会自动合并
```

```linux
git push <remote> <branch>
比如:
git push origin master
```