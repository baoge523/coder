
## proto在传输过程中，会将具有默认值的字段隐藏
比如：pb定义如下
```protobuf
message UserResponse {
  string name = 1;
  int32 isPreset = 2; // 0表示 非预定义  1表示预定义
}
```
当使用该pb定义的响应时，如果isPreset的值为0，那么在传输过程中，就不会有isPreset这个字段

但我们又期望有isPreset的字段是需要有值的

解决方式：
使用将isPreset定义为string类型