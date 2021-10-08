package api

import (
	"github.com/gogf/gf/net/ghttp"
	"runtime"
)

var Hello = helloApi{}

type helloApi struct {}

func (*helloApi) Index(r *ghttp.Request) {
	for i,v := range r.Request.Header {
		for _,vs := range v{
			r.Response.Header().Set(i,vs)
		}
	}
	r.Response.Writeln(runtime.Version())
}

func (*helloApi) Healthz(r *ghttp.Request){
	r.Response.Writeln("OK")
}
