# protobuf
[protobuf github地址](https://github.com/protocolbuffers/protobuf)

[protobuf-go](https://github.com/protocolbuffers/protobuf-go)

[文档](https://protobuf.dev/overview/)
### protobuf-go 使用方式
[指南](https://protobuf.dev/reference/go/go-generated/)

1、安装protobuf
```text
macos
brew install protobuf
```
2、go application 中下载protoc-gen-go插件
注意：这个需要 Go 1.16 or higher，如果是1.16之前的，需要使用相关版本(踩过坑)
```linux
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
```

3、编辑proto文件

4、生成文件xx.pb.go文件
```text
protoc --proto_path=src --go_out=out --go_opt=paths=source_relative foo.proto bar/baz.proto

--proto_path  proto的文件路径

--go_out 输出路径

--go_opt 参数信息

# source_relative 表示相对路径
```

```text
protoc --proto_path=src \
  --go_opt=Mprotos/buzz.proto=example.com/project/protos/fizz \
  --go_opt=Mprotos/bar.proto=example.com/project/protos/foo \
  protos/buzz.proto protos/bar.proto

```

### proto文件的消息定义

#### Messages 定义结构体
```protobuf
message Artist {}
```
#### 内嵌类型
```protobuf
message Artist {
  message Name {
  }
}
```

#### fields 定义成员信息
proto3
```protobuf
int32 birth_year = 1;
optional int32 first_active_year = 2;

```
生成go代码的字段名称为：
BirthYear
FirstActiveYear
#### Singular Message Fields 单个消息体字段
```protobuf
message Band {}

// proto2
message Concert {
  optional Band headliner = 1;
  // The generated code is the same result if required instead of optional.
}

// proto3
messadge Concert {
  Band headliner = 1;
}
```
生成的go对象
```go
type Concert struct {
    Headliner *Band   // 注意Message结果体的成员是指针，默认值为nil
}
```

#### Repeated Fields 可重复成员，对应go中的切片
```protobuf
message Band {}

message Concert {
  // Best practice: use pluralized names for repeated fields:
  // /programming-guides/style#repeated-fields
  repeated Band support_acts = 1;
}

```
生成的go对象
```go
type Concert struct {
    SupportActs []*Band  // 注意这里是Band的切片指针
}
```

#### Map Fields map类型的字段
```text
message MerchItem {}

message MerchBooth {
  // items maps from merchandise item name ("Signed T-Shirt") to
  // a MerchItem message with more details about the item.
  map<string, MerchItem> items = 1;
}
```
生成的go对象
```go
type MerchBooth struct {
    Items map[string]*MerchItem  // 注意这里的value是对象指针
}
```

#### Oneof Fields

```protobuf
package account;
message Profile {
  oneof avatar {
    string image_url = 1;
    bytes image_data = 2;
  }
}
```
生成的go代码
```go
type Profile struct {
    // Types that are valid to be assigned to Avatar:
    //  *Profile_ImageUrl
    //  *Profile_ImageData
    Avatar isProfile_Avatar `protobuf_oneof:"avatar"`
}

type Avatar interface {
	xxx()
}

// Profile_ImageUrl、Profile_ImageData都实现了Avatar中的方法

type Profile_ImageUrl struct {
        ImageUrl string
}
type Profile_ImageData struct {
        ImageData []byte
}
```

使用方式
```go
p1 := &account.Profile{
  Avatar: &account.Profile_ImageUrl{ImageUrl: "http://example.com/image.png"},
}

// imageData is []byte
imageData := getImageData()
p2 := &account.Profile{
  Avatar: &account.Profile_ImageData{ImageData: imageData},
}
```
可以通过switch的方式访问值类型
```go
switch x := m.Avatar.(type) {
case *account.Profile_ImageUrl:
    // Load profile image based on URL
    // using x.ImageUrl
case *account.Profile_ImageData:
    // Load profile image based on bytes
    // using x.ImageData
case nil:
    // The field is not set.
default:
    return fmt.Errorf("Profile.Avatar has unexpected type %T", x)
}
```

#### Enumerations 枚举类型
注意： enum下标是从0开始的、message下标是从1开始的

```protobuf
message Venue {
  
  enum Kind {
    KIND_UNSPECIFIED = 0;
    KIND_CONCERT_HALL = 1;
    KIND_STADIUM = 2;
    KIND_BAR = 3;
    KIND_OPEN_AIR_FESTIVAL = 4;
  }
  Kind kind = 1;
  // ...
}
```

生成的go代码
```go
type Venue_Kind int32  // 定义的枚举类型 如果枚举定义在message内，命名为message_enum,如果定义在外面，命名为enum

// 枚举对应的常量值
const (
    Venue_KIND_UNSPECIFIED       Venue_Kind = 0
    Venue_KIND_CONCERT_HALL      Venue_Kind = 1
    Venue_KIND_STADIUM           Venue_Kind = 2
    Venue_KIND_BAR               Venue_Kind = 3
    Venue_KIND_OPEN_AIR_FESTIVAL Venue_Kind = 4
)

```