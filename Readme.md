# 使用方法

## 安装
```
go get github.com/cboy868/cgin
```

## 简单使用
```
r := cgin.Default()

r.GET("/hello", func(c *cgin.Context) {
	c.Success(map[string]string{"name": "张三"})
})

r.Run()
```

