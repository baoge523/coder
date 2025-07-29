## linxu安装golang sdk
https://go.dev/doc/install

### 下载某一个版本
```text
wget https://go.dev/dl/go1.23.11.linux-amd64.tar.gz
```

```bash
 rm -rf /usr/local/go && tar -C /usr/local -xzf go1.24.5.linux-amd64.tar.gz
 export PATH=$PATH:/usr/local/go/bin
 go version
```