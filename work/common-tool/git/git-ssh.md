## ssh秘钥生产，并支持多账号

```TEXT
1. **生成 SSH 密钥**：
ssh-keygen -t rsa -f ~/.ssh/id_rsa_account1 -C "XXXX1.qq.com"
ssh-keygen -t rsa -f ~/.ssh/id_rsa_account2 -C "XXXX2.qq.com"


2. **编辑 SSH 配置文件**：
   打开或创建 `~/.ssh/config` 文件，添加以下内容：
   
   Host account1
   HostName example.com
   User your_username1
   IdentityFile ~/.ssh/id_rsa_account1
   
   Host account2
   HostName example.com
   User your_username2
   IdentityFile ~/.ssh/id_rsa_account2

3. **设置权限**：
   确保 SSH 文件夹和密钥的权限正确：

   chmod 700 ~/.ssh
   chmod 600 ~/.ssh/id_rsa_account1
   chmod 600 ~/.ssh/id_rsa_account2

```

~/.ssh/config 文件信息如下:
```text
Host github.com
HostName github.com
User 191xxxx898@qq.com
IdentityFile ~/.ssh/id_rsa_qq
```

### 使用git ssh
将对应账号生产的ssh公钥添加到github中，然后就可以使用git ssh来访问项目了

### 将本地的https访问改成ssh方式访问github
git config -l
```text
remote.origin.url=https://github.com/baoge523/coder.git
```
git config --unset-all  remote.origin.url

git config --global remote.origin.url "git@github.com:baoge523/coder.git"

然后再通过 git remote -v 查看origin的地址是否修改过来了，如果修改过来了，那么就可以直接使用了