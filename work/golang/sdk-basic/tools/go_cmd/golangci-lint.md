## golangci-lint

```bash
golangci-lint
```
https://golangci-lint.run/welcome/install/

## 安装golangci-lint
使用wget的方式，可以按照指定的版本
```bash
# 会将命令放在当前目录的./bin/golangci-lint
wget -O- -nv https://raw.githubusercontent.com/golangci/golangci-lint/HEAD/install.sh | sh -s v1.62.2


# 通过修改~/.zshrc 将命令放在$PATH下， source ~/.zshrc
```

也可以在macos 中通过以下命令安装
```bash
brew install golangci-lint
brew upgrade golangci-lint
```

## config file
https://golangci-lint.run/usage/configuration/

.golangci.yml

也可以通过  golangci-lint config path 查看配置的使用位置


### 在golangci-lint中遇到的问题
```text
当安装了v1.62.2版本的golangci-lint 后执行，发现存在error错误；需要检测一下当前golang的版本
1、通过 go version 查看golang版本
2、通过golangci-lint version 查看 golangci-lint的版本
```