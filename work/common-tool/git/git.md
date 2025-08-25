## GIT 使用学习

Git 是一个开源的分布式版本控制系统，用于敏捷高效地处理任何或小或大的项目

GIT官网命令手册地址：https://git-scm.com/docs

### 一、Git 工作区、暂存区和版本库

基本概念

我们先来理解下 Git 工作区、暂存区和版本库概念：

- **工作区：**就是你在电脑里能看到的目录。
- **暂存区：**英文叫 stage 或 index。一般存放在 <u>**.git** 目录下的 index 文件</u>（.git/index）中，所以我们把暂存区有时也叫作索引（index）。
- **版本库：**工作区有一个隐藏目录 **.git**，这个不算工作区，而是 Git 的版本库。

![img](https://www.runoob.com/wp-content/uploads/2015/02/1352126739_7909.jpg)

### 二、Git创建仓库

#### git init

Git 使用 **git init** 命令来初始化一个 Git 仓库，Git 的很多命令都需要在 Git 的仓库中运行，所以 **git init** 是使用 Git 的第一个命令。

在执行完成 **git init** 命令后，Git 仓库会生成一个 .git 目录，该目录包含了资源的所有元数据，其他的项目目录保持不变。

git init                  初始化当前目录

git init newrepo   初始化指定目录

#### git clone

我们使用 **git clone** 从现有 Git 仓库中拷贝项目。

git clone <repo>

git clone -b 分支名  git地址

#### 配置

git 的设置使用 **git config** 命令。

显示当前的 git 配置信息：

```linux
$ git config --list
credential.helper=osxkeychain
core.repositoryformatversion=0
core.filemode=true
core.bare=false
core.logallrefupdates=true
core.ignorecase=true
core.precomposeunicode=true
```

添加 --global 表示针对系统上所有仓库

设置提交代码时的用户信息：

> $ git config --global user.name "runoob"
> $ git config --global user.email test@runoob.com

### 三、Git基本操作

Git 的工作就是创建和保存你项目的快照及与之后的快照进行对比。

![img](https://www.runoob.com/wp-content/uploads/2015/02/git-command.jpg)

**说明：**

- workspace：工作区
- staging area：暂存区/缓存区
- local repository：版本库或本地仓库
- remote repository：远程仓库



**HEAD 说明：**

- HEAD 表示当前版本

- HEAD^ 上一个版本

- HEAD^^ 上上一个版本

- HEAD^^^ 上上上一个版本

- 以此类推...

  

可以使用 ～数字表示

- HEAD~0 表示当前版本
- HEAD~1 上一个版本
- HEAD^2 上上一个版本
- HEAD^3 上上上一个版本
- 以此类推...



#### 3.1 Git reset 命令

git reset 命令用于回退版本，可以指定退回某一次提交的版本。

> git reset [--soft | --mixed | --hard] [HEAD]

**--mixed** 为默认，可以不用带该参数，用于重置暂存区的文件与上一次的提交(commit)保持一致，工作区文件内容保持不变。

```linux
$ git reset HEAD^            # 回退所有内容到上一个版本  
$ git reset HEAD^ hello.php  # 回退 hello.php 文件的版本到上一个版本  
$ git  reset  052e           # 回退到指定版本
```

**--soft** 参数用于回退到某个版本：

```linux
$ git reset --soft HEAD~3   # 回退上上上一个版本 
```

**--hard** 参数撤销工作区中所有未提交的修改内容，将暂存区与工作区都回到上一次版本，并删除之前的所有信息提交：

```linux
$ git reset --hard HEAD~3  # 回退上上上一个版本  
$ git reset –hard bae128  # 回退到某个版本回退点之前的所有信息。 
$ git reset --hard origin/master    # 将本地的状态回退到和远程的一样 
```

**注意：**谨慎使用 **–-hard** 参数，它会删除回退点之前的所有信息。

案例：

```
自己的私有仓库领先于远端的公有仓库，现在需要将自己的提交回滚掉，然后再拉取远程公有仓库的代码

upstream: 远端公有仓库
origin: 远端私有仓库
local: 本地仓库和远端私有仓库代码一致，领先于 upstream的提交

git reset --hard  commit_id 这个提交号是需要回滚的提交号的前一个提交号

git log 查看提交信息

git push -f 将回滚后的本地信息，提交到远端私有仓库

git pull upstream master --rebase  将upstream master分支 rebase到本地

git push -f  提交拉取到的分支

此时 upsteram、origin、local代码一致，无多余的提交



```







#### 3.2 Git rm 命令

git rm 命令用于删除文件。

1、将文件从暂存区和工作区中删除：

> git rm <file>

如果删除之前修改过并且已经放到暂存区域的话，则必须要用强制删除选项 **-f**。

> git rm -f runoob.txt 

2、如果想把文件从暂存区域移除，但仍然希望保留在当前工作目录中

> git rm --cached <file>



### git clone 制定commitId，并修改版本提交

```linux
1、克隆分支
git clone <repository_url>
cd <repository_name>
git checkout <commit_id>

2、为切换的commitId分支，创建一个名称
git switch -c 

dev/ins_mod_regionMap_380


```

```linux
在克隆的指定 commit ID 后，要提交代码到仓库，您可以按照以下步骤操作：
1. 在克隆的存储库中进行所需的更改。
2. 使用以下命令将更改添加到暂存区：
   ```bash
   git add <file_name>
```
   或者，如果要添加所有更改，可以使用：
   ```bash
   git add .
   ```
3. 提交更改到本地存储库：
   ```bash
   git commit -m "提交说明"
   ```
   替换 `"提交说明"` 为对提交的简要描述。
4. 最后，将更改推送到远程仓库：
   ```bash
   git push origin <branch_name>
   ```
   替换 `<branch_name>` 为您要推送到的远程分支名称。
   这些步骤将允许您将更改提交到指定的克隆的 commit ID 后的仓库。

如果远程分支dev/ins_mod_require不存在
git push -u origin dev/ins_mod_require
```

### 分支相关的操作

1、查看本地分支关联的远程分支

```linux
git branch -vv
```

2、删除本地分支

```linux
git branch -d branch_name

// 如果要强制删除未合并的分支，可以使用以下命令
git branch -D branch_name
```

3、要基于某个远程分支创建本地分支

```linux
git checkout -b local_branch_name origin/remote_branch_name
例如：
git checkout -b release/tce3.10.0 origin/release/tce3.10.0
```

4、要基于本地分支创建新的本地分支

```linux
git checkout -b new_branch_name local_branch_name

```

### 如何将一个地方的分支拉取到本地，并推送到另一个远程上 

```tet
1、拉取分支代码到本地
git clone -b xx  git地址

2、重新设置origin
git remote -v 查看现有remote信息
git remote remove origin 删除现有的remote origin
git remote add origin git地址

3、push代码
git push -u origin local_branch:remote_branch
其中 -u 表示远端不存在分支时，就创建



```

