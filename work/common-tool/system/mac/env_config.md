## macos 中的环境变量配置文件

### 全局配置文件，对所有的用户生效

/etc/profile
系统级配置文件，所有用户登录时（Login Shell）会读取，但通常不建议直接修改。

/etc/paths
系统默认的 PATH 配置，每行一个路径，对所有用户生效。

/etc/paths.d/
存放追加路径的目录，可通过在此目录下添加文件来扩展 PATH。

/etc/zprofile、/etc/zshrc
Zsh 的系统级配置文件（若使用 Zsh）。

### 用户级配置文件（对当前用户生效）

### bash shell
~/.bash_profile
> 用户登录时（Login Shell）读取，优先级高于 ~/.bashrc。适合配置环境变量。

~/.bashrc
> 非登录交互式 Shell（如新开终端窗口）读取。需在 ~/.bash_profile 中显式调用

### zsh shell (macOS 默认 Shell)

~/.zprofile
> 类似 ~/.bash_profile，登录时执行一次。

~/.zshrc    用户shell窗口级别 -- 常用
> 交互式 Shell 启动时读取，适合日常变量和别名

~/.zshenv    # 哈哈哈，我本机没有这个问题
> 所有 Zsh 实例（包括脚本）启动时读取，优先级最高。

### 修改后，生效操作
修改后需重启终端或执行 source <文件>（如 source ~/.zshrc）立即生效


### 当发现用户级别的~/.zshrc文件是root所有的时，无法编辑该文件
```bash
cd ~
ls -a -l 
# 发现.zshrc 是root用户的，其他用户没有w权限，这里该文件应该是属于当前用户的

sudo chown $USER ./zshrc

ls -a -l
vim .zshrc
```