# oconv

Office Converter

### 格式转换接口

**url**

`/conv`

**method**

`POST`

**参数**

|名|类型|说明|
|-|-|-|
|file|文件|大小不超过32m, 类型限制为`.doc`,`.docx`,`.pdf`|
|out|string|输出文件格式, 限制为`pdf`,`jpg`|

**返回示例**

```
成功
status code = 200

[
    "https://oconv.oss-cn-beijing.aliyuncs.com/res/fc97cb88eebe4fa83b9a005b99f805d0ee5b407a674c4a88000ea54731e142b7.jpg"
]
```

```
失败
status code <> 200

eg.
status: 400
output file format invalid
```

**注意**

1. Content-Type: multipart/form-data
2. 请求body最大不超过32M

### gofmt

`gofmt -w -l *.go utils/*.go config/*.go model/*.go`

