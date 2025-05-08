## grep search command

grep 是可以强大的文件检索工具，可以通过man查看其具体用法
比如：
   可以对输出的数据标记颜色、求count数等等

```bash
man grep  // show grep usage
```
### 使用方式
如果没有指定文件，那么就查询当前目录中的文件，-r 表示递归查询
文件也支持pattern
```text
   grep [OPTIONS] PATTERN [FILE...]
```

### 常使用的参数
```text
   - `-i`：忽略大小写。
   - `-v`：反向匹配，输出不包含匹配规则的行。
   - `-n`：显示行号。
   - `-r`：递归搜索目录中的文件。
   - `-l`：只输出匹配文件的名称。
```

### 使用例子
查询匹配套件的所有文件
```text
grep "aaa" *.log.20250428
```
