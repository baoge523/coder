
## submodule
(中文文档)[https://git-scm.com/book/zh/v2/Git-%e5%b7%a5%e5%85%b7-%e5%ad%90%e6%a8%a1%e5%9d%97]


### 当前项目a将项目b作为子项目的操作
```bash
git submodule add  https://xxx
```



### 从远端clone一个项目到本地，此时本地的子项目是一个空目录

处理方式一：
```bash
git submodule init   // 用于初始化本地的配置文件
git submodule update  // 用于抓取所有数据并检出父项目中列出的合适的提交

或者

git submodule update --init
git submodule update --init --recursive  // 初始化抓取并检出任何嵌套的子模块  --recursive一般用于子项目还拥有子项目的情况
```

方式二、
在clone项目时，指定参数 --recurse-submodules,可以在clone时就把子项目拉取下来
```bash
git clone -b xxx git_clone_url --recurse-submodules
```



### 更新子项目的代码
方式一：
进入到子项目目录中
```bash
git fetch
git merge origin/master  将origin/master分支数据merge到本地的当前分支；即分支合并后，指向一个新的commit id
```

方式二、Git 将会进入子模块然后抓取并更新  -- 推荐使用
```bash
git submodule update --remote
```
当运行 git submodule update --remote 时，Git 默认会尝试更新 所有 子模块， 所以如果有很多子模块的话，你可以传递想要更新的子模块的名字
git submodule update --remote aaa


### 当.gitmodules 文件中的子模块的url发生改变后，我们执行git pull --recurse-submodules 或者git submodule update就会报错

此时我们应该使用git submodule sync 命令
--recursive 递归处理，主要用于处理子模块有子模块的场景
```bash
git submodule sync  --recursive  将新的 URL 复制到本地配置中
git submodule update --init --recursive  重新 URL 更新子模块
```


可以通过 --rebase 和 --merge 来合并子模块上的变更操作
```bash
git submodule update --remote --merge
git submodule update --remote --rebase
```