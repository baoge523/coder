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
