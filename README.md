# ShyGinErrors

[![Build Status](https://github.com/shyandsy/ShyGinErrors/workflows/Run%20Tests/badge.svg?branch=main)](https://github.com/shyandsy/ShyGinErrors/actions?query=branch%3Amain)
[![codecov](https://codecov.io/gh/shyandsy/shyginerrors/branch/main/graph/badge.svg)](https://codecov.io/gh/shyandsy/shyginerrors)
![Go Report Card](https://goreportcard.com/badge/github.com/shyandsy/shyginerrors)](https://goreportcard.com/report/github.com/shyandsy/shyginerrors)
An extension to generate key value errors for gin framework and go validator error 

#### What we want 
validate request and get error messages like below
```json
"errors" : {
  "username":"username仅包含大小写字母和数字，长度6-32"
},
```

#### How to?

sample code
```go
var requestErrorMessage = map[string]string{
    "error_invalid_email":    "请输入一个有效地meail地址",
    "error_invalid_username": "username仅包含大小写字母和数字，长度6-32",
    "error_invalid_password": "密码长度6-32",
}

type RegisterForm struct {
    Email    string `json:"email" binding:"required,email" msg:"error_invalid_email"`
    Username string `json:"username" binding:"required,alphanum,gte=6,lte=32" msg:"error_invalid_username"`
    Password string `json:"password" binding:"required,gte=6,lte=32" msg:"error_invalid_password"`
}

func (c Controller) Register(reqCtx appx.ReqContext) (interface{}, error) {

	// step 1: initialize the ge object
    ge = NewShyGinErrors(requestErrorMessage)
	
	req := model.RegisterForm{}
	if err := reqCtx.Gin().BindJSON(&req); err != nil {
		// step 2: use ge object to parse the error messages
		errors := ge.ListAllErrors(req, err)
		
		// error handling
	}
	
	return req, nil
}
```

#### How it works 

1. define error message key map, we will use the keys in msg tag for models
```go
var requestErrorMessage = map[string]string{
    "error_invalid_email":    "请输入一个有效地meail地址",
    "error_invalid_username": "username仅包含大小写字母和数字，长度6-32",
    "error_invalid_password": "密码长度6-32",
}
```

2. define models with json tag and msg tag
msg tag specific the key of the message in requestErrorMessage
```go
type RegisterForm struct {
    // 
    Email    string `json:"email" binding:"required,email" msg:"error_invalid_email"`
    Username string `json:"username" binding:"required,alphanum,gte=6,lte=32" msg:"error_invalid_username"`
    Password string `json:"password" binding:"required,gte=6,lte=32" msg:"error_invalid_password"`
}
```

3. initialize the ShyGinErrors
```go
ge = NewShyGinErrors(requestErrorMessage) // register the error message map
```

4. parse the error return by gin.BindJson()
```go
if err := reqCtx.Gin().BindJSON(&req); err != nil {

    errors := ge.ListAllErrors(req, err)
    
    // process errors
}
```

now, we get the k-v error message array to frontend
