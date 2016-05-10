# log-server
日志服务器

## 客户端上传日志

发送请求`POST /log/stream`

`request body`内容为每条日志一行，格式如下：

```json
{
	"id": "消息在当前进程的唯一ID",
	"path": "日志路径，如 server/api",
	"data": "日志内容"
}
```

服务器成功写入日志后，会在`response body`中返回一行数据（`chunked`）：

```
消息的ID
```

## 启动服务器

