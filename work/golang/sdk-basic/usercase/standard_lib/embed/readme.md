## embed
支持内嵌文件信息

支持三种类型来接收嵌入的文件
 - string
 - []byte
 - embed.FS

其中的string、[]byte 只支持单个文件

embed.FS 可以支持单个或者多个(文件、目录)，指定方式也可以使用path.Match patterns
> 因为 embed.FS实现了io/fs，所以embed.FS可以使用在net/http,text/template,html/template等地方



