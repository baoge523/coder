## prometheus 中的问题

在共计 1000 个左右的pod时，prometheus在16核独占node的总体CPU就已经长期徘徊在高位
https://github.com/prometheus/prometheus/issues/8014   issue已经解决

### 排查思路

prometheus本身也是golang程序，可以通过pprof定位到底是哪里出现了问题，然后再通过源码调试，找到其问题所在的根本原因

