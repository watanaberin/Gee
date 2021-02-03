# Day1

**见注解**

# Day2

## Gee.go

*Engine*封装了一个router，主要的业务逻辑：

**GET()/POST()**:将其和pattern（/、/hello、/login）和handler一一对应并注册到路由表中。

**Run()**:封装了http的listenAndServe()。

## Router.go

handle封装了一个map。

**addRoute()**:添加router项。把POST/GET + pattern作为key，如（GET - /hello），handler为value。

**handle()**:读取Context中的key，查询路由表是否存在该key，存在则执行对应的handler，否则显示404。

## Context.go

**上下文**

*Context*封装了

```
Response、Request

Path、Method（Request info）

StatusCode（Response info）
```

**newContext()**:新建Context。

**PostForm()**:表单提交，输入key返回value。

**Query()**

**Status()**: 状态码更新（Context直接封装的StatusCode，Response头的更新）

**SetHeader()**:响应头（key，value）

**String()**:响应头/状态码/写入

**JSON()**:响应头/状态码/写入（存疑）

**Data()**:状态码/写入

**HTML()**:响应头/状态码/写入

写入 **c.Writer.Write([]byte(Text))**