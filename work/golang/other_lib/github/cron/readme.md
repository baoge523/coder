## corn
document: https://pkg.go.dev/github.com/robfig/cron

Field name   | Mandatory? | Allowed values  | Allowed special characters
----------   | ---------- | --------------  | --------------------------
Seconds      | Yes        | 0-59            | * / , -
Minutes      | Yes        | 0-59            | * / , -
Hours        | Yes        | 0-23            | * / , -
Day of month | Yes        | 1-31            | * / , - ?
Month        | Yes        | 1-12 or JAN-DEC | * / , -
Day of week  | Yes        | 0-6 or SUN-SAT  | * / , - ?

### 特殊符号的意思
#### * 代表匹配所有
```text

```

#### / 代表range，循环处理
```text
0 0/5 * * * * // 表示每五分钟执行一次，这里 0/5 表示的是0-60/5
0 */5 * * * * // 这里的* 和0表示一样 
```
#### , 代表分别有哪些
```text
0 5,10 * * * *  // 表示每小时的第五分，第十分执行，12:05:00，12:10:00，13:05:00，13:10:00
```
#### - 代表范围
```text
0 5-8 * * * *  // 表示每小时的5-10分的0秒执行，比如 12:05:00,12:06:00,12:07:00,12:08:00

0 0 10-12 * * * // 表示每天的10-12执行 比如: 10:00:00 11:00:00
```
#### ？和*一样，但只能在 day of month and day of week 使用


Time zones：
> All interpretation and scheduling is done in the machine's local time zone
