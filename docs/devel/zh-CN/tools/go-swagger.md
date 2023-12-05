# go-swagger

## 安装
```bash
go install github.com/go-swagger/go-swagger/cmd/swagger

swagger version
```

## 使用

```bash
生成swagger yaml文件
swagger generate spec -o swagger.yaml

在线查看
swagger serve --no-open -F=swagger --port 36666 swagger.yaml

yaml格式转换为json
swagger generate spec -i ./swagger.yaml -o ./swagger.json
```
