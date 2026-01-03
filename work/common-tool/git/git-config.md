## git config

git config --global --list
git config --local --list

git config --global --unset config_key

git config --local --add key value





### git clone 总是下载相同的项目的问题排查

根据下面的命令，总是下载到相同的仓库代码文件
```bash
git clone git@github.com:baoge523/trpc-go.git

git clone https://github.com/baoge523/opentelemetry-go.git
```

排查原因是，git的config配置文件导致的

查询全局的配置文件
```bash
git config --global --list
```
发现设置了remote.origin.url 导致通过git clone 总是下载到coder的仓库文件
```text
remote.origin.url=git@github.com:baoge523/coder.git
```

查询项目的local配置信息
```bash
git config --local --list
```

解决方式： 移出掉global作用域的remote.origin.url，并将其设置到对应的项目的git config --local下
```text

git config --global --unset remote.origin.url

git config --local --add remote.origin.url git@github.com:baoge523/coder.git 
```