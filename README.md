# safecustody-request

### go mod 安装
```go
import (
    req "github.com/chainlife-doc/safecustody-request"
)

```
#### POST
```go
body,err:=req.Http().Post(url, param)
fmt.Println(string(body))
```
#### GET
```go
body,err:=req.Http().Get(url)
fmt.Println(string(body))
```
#### 设置请求头
```go
header := make(req.Header)
header["Content-Type"] = "application/x-www-form-urlencoded"
req.Http().SetHeader(header).Post(url, nil)
```
#### 设置超时时间
```go
header := make(req.Header)
header["Content-Type"] = "application/x-www-form-urlencoded"
req.Http().SetHeader(header).SetTimeout(1 * time.Second).Post(url, nil)
```

#### 更多设置
```go
type Request struct {
    //参考源码设置
    HttpC *http.Client

    t time.Duration

    header Header
}
```

```go
func Http() *Request {
	header := make(Header)
	header["Content-Type"] = "application/x-www-form-urlencoded"

	return &Request{
		HttpC: &http.Client{
			Timeout:   timeOutDefault,
			Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: InsecureSkipVerify}},
		},
		header: header,
	}
}
```
