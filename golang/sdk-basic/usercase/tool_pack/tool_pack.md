## golang相关的工具包

### golang.org/x
```golang
import (
	"golang.org/x/sync/errgroup"
)

func Test() {

var eg errgroup.Group

eg.Go(f func() error)

eg.Wait()
}


```

### crypto
golang 的sdk下的 crypto 定义了一些hash算法，可以拿来直接使用
```go
import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
)

// MD5 _
func MD5(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

// SHA256 _
func SHA256(data []byte) string {
	hash := sha256.Sum256(data)
	return hex.EncodeToString(hash[:])
}
```
