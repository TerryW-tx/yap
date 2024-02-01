package main

import (
	"github.com/goplus/yap/ytest"
	"github.com/goplus/yap/ytest/auth/jwt"
	"testing"
	"time"
)

type case_jwtdemo struct {
	ytest.Case
}
//line ytest/demo/jwtdemo/jwtdemo_ytest.gox:7
func (this *case_jwtdemo) Main() {
//line ytest/demo/jwtdemo/jwtdemo_ytest.gox:7:1
	this.Mock("foo.com", new(jwtdemo))
//line ytest/demo/jwtdemo/jwtdemo_ytest.gox:9:1
	testuser := jwt.HS256("<secret key>").Audience("testuser").Expiration(time.Now().Add(time.Hour))
//line ytest/demo/jwtdemo/jwtdemo_ytest.gox:11:1
	this.Run("test get /p/$id", func() {
//line ytest/demo/jwtdemo/jwtdemo_ytest.gox:12:1
		id := "123"
//line ytest/demo/jwtdemo/jwtdemo_ytest.gox:13:1
		this.Get("http://foo.com/p/" + id)
//line ytest/demo/jwtdemo/jwtdemo_ytest.gox:14:1
		this.Auth(testuser)
//line ytest/demo/jwtdemo/jwtdemo_ytest.gox:15:1
		this.RetWith(200)
//line ytest/demo/jwtdemo/jwtdemo_ytest.gox:16:1
		this.Json(map[string]string{"id": id})
	})
}
func Test_jwtdemo(t *testing.T) {
	ytest.Gopt_Case_TestMain(new(case_jwtdemo), t)
}
