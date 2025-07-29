## wc linux的使用方式
wc主要用于统计数量并将数量输出

统计源：（取决于参数指定）
 - 可以是标准输入
 - 可以是文件

输出：
 - 可以是标准输出(控制台)
 - 可以是文件

### usage
通过man 查看使用方式
```bash
man wc
```

wc - print newline, word, and byte counts for each file
```text
wc [OPTION]... [FILE]...
wc [OPTION]... ‐‐files0‐from=F
```

options
```text
-c, --bytes
      print the byte counts

-m, --chars
      print the character counts

-l, --lines
      print the newline counts

--files0-from=F
      read input from the files specified by NUL-terminated names in file F; If F is - then read names from standard input

-L, --max-line-length
      print the maximum display width

-w, --words
      print the word counts

--total=WHEN
      when to print a line with total counts; WHEN can be: auto, always, only, never

--help display this help and exit

--version
      output version information and exit
```

demo
```text
wc -l 
wc -l a.txt 

```