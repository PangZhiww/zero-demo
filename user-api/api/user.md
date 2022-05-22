### 1. "获取用户信息"

1. route definition

- Url: /home/pangzhi/go/src/zero-demo/user-api/api/foo
- Method: POST
- Request: `UserInfoReq`
- Response: `UserInfoResp`

2. request definition



```golang
type UserInfoReq struct {
	UserId int64 `json:"UserId"`
}
```


3. response definition



```golang
type UserInfoResp struct {
	UserId int64 `json:"UserId"`
	Nickname string `json:"Nickname"`
}
```

