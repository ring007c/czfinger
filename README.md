# 项目名称
指纹扫描器
# 使用方法

```
Usage: 扫描器 [-u]|[--tf][-t]

欢迎使用扫描器

Options:
  -v, --version   Show the version and exit
  -u, --url      单个URL
  --tf, --file    加载文件
  -t ,--threads  开启线程
  ```
  

# 编译
```
go build cmd/czfinger/main.go
```
 
# 用例
```
go run .\main.go --tf 批量文件名 -t 50

go run .\mian.go -u 单个url  -t 50
