# preview

document preview online

### B站视频演示

[基于golang开发的一款高性能文档在线预览服务]()


### 文档格式转换接口

将文档格式转换为`jpg`或者`pdf`, 这样就可以前端渲染预览.

如果选择的输出格式为`jpg`, 为了提高预览速度, 会把文档按照页顺序切分成多张图片

**url**

`/oconv`

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
    "https://preview.mysite.test/oconv/res/fc97cb88eebe4fa83b9a005b99f805d0ee5b407a674c4a88000ea54731e142b7.jpg"
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
3. **重要:** 为了预览方便, 服务器只会缓存文档3天, 而且**保证不会用作其他用途**, 请提前做好备份

### gofmt

`gofmt -w -l *.go utils/*.go config/*.go model/*.go`

