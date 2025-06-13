## tcpdump
用于在linux系统中抓取数据包

### 使用的场景有

1、当业务无法判断是否有数据进入时，可以通过抓包查看，是否有访问指定的端口请求

2、当需要确认请求的参数信息时，可以通过抓包查看请求体


### 命令解释
```bash
man tcpdump
```
常用的参数有:
```text
-A  print each package use ASCII

-d human readable to standard output

-i interface:网络接口  可以填写，比如: "eth0" "any" 表示任何一个网卡

-c 获取指定数据包后退出

-n don't convert host address to names

-nn don't convert protocol and port numbers etc. to names

--print print parsed packet output

-v when parsing and printing, produce verbose output

-vv even more verbose output

-0
--no-optimize  don't run the package-matching code optimizer

-s 指定截取package的长度，默认是262144bytes，sets 0 表示不限制  （指定每个数据包的捕获长度（字节）。）

```
### tcpdump 表达式
可以通过man tcpdump查看
```text
在 `tcpdump` 中，表达式用于过滤捕获的数据包。它们由以下几个部分组成：
1. **协议**：如 `tcp`、`udp`、`icmp` 等。
2. **地址**：可以是源地址 (`src`)、目标地址 (`dst`)，例如 `src 192.168.1.1`。
3. **端口**：指定源或目标端口，如 `port 80` 或 `src port 22`。
4. **逻辑运算符**：使用 `and`、`or`、`not` 进行组合。
示例表达式：
- 捕获 TCP 数据包：`tcp`
- 捕获来自特定 IP 的数据包：`src 192.168.1.1`
- 捕获特定端口的数据包：`port 80`
- 组合条件：`tcp and src 192.168.1.1 and (dst port 80 or dst port 443)`
可以根据需要组合这些元素，形成复杂的过滤条件。
```
例子
```text
tcpdump tcp src 10.1.24.100 and src port 80

tcpdump udp dst port 8080

```



### man提供的例子
```text
To print all packets arriving at or departing from sundown:
          tcpdump host sundown

To print traffic between helios and either hot or ace:
      tcpdump host helios and \( hot or ace \)

To print all IP packets between ace and any host except helios:
      tcpdump ip host ace and not helios

To print all traffic between local hosts and hosts at Berkeley:
      tcpdump net ucb-ether

To print all ftp traffic through internet gateway snup: (note that the expression is quoted to prevent the shell from (mis-)interpreting the parentheses):
      tcpdump 'gateway snup and (port ftp or ftp-data)'

To print traffic neither sourced from nor destined for local hosts (if you gateway to one other net, this stuff should never make it onto your local net).
      tcpdump ip and not net localnet

To print the start and end packets (the SYN and FIN packets) of each TCP conversation that involves a non-local host.
      tcpdump 'tcp[tcpflags] & (tcp-syn|tcp-fin) != 0 and not src and dst net localnet'

To print the TCP packets with flags RST and ACK both set.  (i.e. select only the RST and ACK flags in the flags field, and if the result is "RST and ACK both set", match)
      tcpdump 'tcp[tcpflags] & (tcp-rst|tcp-ack) == (tcp-rst|tcp-ack)'

To print all IPv4 HTTP packets to and from port 80, i.e. print only packets that contain data, not, for example, SYN and FIN packets and ACK-only packets.  (IPv6 is left as an exercise for the reader.)
      tcpdump 'tcp port 80 and (((ip[2:2] - ((ip[0]&0xf)<<2)) - ((tcp[12]&0xf0)>>2)) != 0)'

To print IP packets longer than 576 bytes sent through gateway snup:
      tcpdump 'gateway snup and ip[2:2] > 576'

To print IP broadcast or multicast packets that were not sent via Ethernet broadcast or multicast:
      tcpdump 'ether[0] & 1 = 0 and ip[16] >= 224'

To print all ICMP packets that are not echo requests/replies (i.e., not ping packets):
      tcpdump 'icmp[icmptype] != icmp-echo and icmp[icmptype] != icmp-echoreply'

```

### 常用例子
tcpdump -nn -i any port 80 -vvvAs0
```text
- `-nn`: 不解析主机名和端口名，显示数字格式。
- `-i any`: 监听所有网络接口。
- `port 80`: 过滤只捕获目标端口为80的流量（通常是HTTP流量）。
- `-vvv`: 提供详细输出。
- `-A`: 以ASCII格式显示数据包的内容。
- `-s0`: 捕获整个数据包，而不仅仅是头部（0表示无限制）。
```


