package gee

import(
	"net/http"
	"fmt"
	"encoding/json"
)

type H map[string]interface{}

//Context
type Context struct{
	Writer http.ResponseWriter
	Req *http.Request
	//request info
	Path string
	Method string
	//response info
	StatusCode int
}
//新建Context
func newContext(w http.ResponseWriter,req *http.Request) *Context {
	return &Context{
		Writer: w,
		Req: req,
		Path: req.URL.Path,
		Method: req.Method,
	}
}
//PostForm
func (c *Context) PostForm(key string) string{
	return c.Req.FormValue(key)//
}
//Query
func (c *Context) Query(key string) string{
	return c.Req.URL.Query().Get(key)//
}
//状态码
func (c *Context) Status(code int){
	c.StatusCode=code
	c.Writer.WriteHeader(code)//
}
//响应头
func (c *Context) SetHeader(key string,value string){
	c.Writer.Header().Set(key,value)//
}

func (c *Context) String(code int,format string,values ...interface{}){
	c.SetHeader("Content-Type","text/plain")
	c.Status(code)
	c.Writer.Write([]byte(fmt.Sprintf(format,values...)))
}
func (c *Context) JSON(code int,obj interface{}){
	c.SetHeader("Content-Type","text/json")
	c.Status(code)
	encoder :=json.NewEncoder(c.Writer)//
	if err :=encoder.Encode(obj);err != nil{
		http.Error(c.Writer,err.Error(),500)
	}
}
func (c *Context) Data(code int, data []byte) {
	c.Status(code)
	c.Writer.Write(data)
}

func (c *Context) HTML(code int, html string) {
	c.SetHeader("Content-Type", "text/html")
	c.Status(code)
	c.Writer.Write([]byte(html))
}




