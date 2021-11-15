package ShyGinErrors

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo"
	"github.com/stretchr/testify/assert"
)

// msg tag specific the key for the corresponse error
type RegisterForm struct {
	Email    string `json:"email" binding:"required,email" msg:"error_invalid_email"`
	Username string `json:"username" binding:"required,alphanum,gte=6,lte=32" msg:"error_invalid_username"`
	Password string `json:"password" binding:"required,gte=6,lte=32" msg:"error_invalid_password"`
}

var requestErrorMessage = map[string]string{
	"error_invalid_email":    "请输入一个有效地meail地址",
	"error_invalid_username": "username仅包含大小写字母和数字，长度6-32",
	"error_invalid_password": "密码长度6-32",
}

func MockJsonPost(c *gin.Context, content interface{}) {
	c.Request.Method = "POST"
	c.Request.Header.Set("Content-Type", "application/json")

	jsonbytes, err := json.Marshal(content)
	if err != nil {
		panic(err)
	}

	c.Request.Body = io.NopCloser(bytes.NewBuffer(jsonbytes))
}

var _ = Describe("test validator", func() {
	var ctx *gin.Context
	var ge GinErrors

	BeforeEach(func() {
		ge = NewShyGinErrors(requestErrorMessage)
		w := httptest.NewRecorder()
		ctx, _ = gin.CreateTestContext(w)
		ctx.Request = &http.Request{
			Header: make(http.Header),
		}
	})

	It("test request validate username", func() {
		MockJsonPost(ctx, map[string]interface{}{"username": "shyandsy!#", "email": "shyandsy@gmail.com", "password": "123456"})

		req := RegisterForm{}
		if err := ctx.BindJSON(&req); err != nil {
			log.Println(err.Error())
			errors := ge.ListAllErrors(req, err)
			assert.True(GinkgoT(), len(errors) == 1)

			value, ok := errors["username"]
			assert.True(GinkgoT(), ok)
			assert.True(GinkgoT(), value == requestErrorMessage["error_invalid_username"])
		} else {
			assert.Fail(GinkgoT(), "error on username field has not been detected")
		}
	})

	It("test request validate email", func() {
		MockJsonPost(ctx, map[string]interface{}{"username": "shyandsy", "email": "shyandsygmail.com", "password": "123456"})

		req := RegisterForm{}
		if err := ctx.BindJSON(&req); err != nil {
			errors := ge.ListAllErrors(req, err)
			assert.True(GinkgoT(), len(errors) == 1)

			value, ok := errors["email"]
			assert.True(GinkgoT(), ok)
			assert.True(GinkgoT(), value == requestErrorMessage["error_invalid_email"])
		} else {
			assert.Fail(GinkgoT(), "error on email field has not been detected")
		}
	})

	It("test request validate password", func() {
		MockJsonPost(ctx, map[string]interface{}{"username": "shyandsy", "email": "shyandsy@gmail.com", "password": "12345"})

		req := RegisterForm{}
		if err := ctx.BindJSON(&req); err != nil {
			errors := ge.ListAllErrors(req, err)
			assert.True(GinkgoT(), len(errors) == 1)

			value, ok := errors["password"]
			assert.True(GinkgoT(), ok)
			assert.True(GinkgoT(), value == requestErrorMessage["error_invalid_password"])
		} else {
			assert.Fail(GinkgoT(), "error on email field has not been detected")
		}
	})

	It("test request validate username, email, password", func() {
		MockJsonPost(ctx, map[string]interface{}{"username": "shyandsy!", "email": "shyandsygmail.com", "password": "12345"})

		req := RegisterForm{}
		if err := ctx.BindJSON(&req); err != nil {
			errors := ge.ListAllErrors(req, err)
			assert.True(GinkgoT(), len(errors) == 3)

			value, ok := errors["username"]
			assert.True(GinkgoT(), ok)
			assert.True(GinkgoT(), value == requestErrorMessage["error_invalid_username"])

			value, ok = errors["email"]
			assert.True(GinkgoT(), ok)
			assert.True(GinkgoT(), value == requestErrorMessage["error_invalid_email"])

			value, ok = errors["password"]
			assert.True(GinkgoT(), ok)
			assert.True(GinkgoT(), value == requestErrorMessage["error_invalid_password"])

		} else {
			assert.Fail(GinkgoT(), "error on username, email, password fields has not been detected")
		}
	})
})
