// Code generated by goctl. DO NOT EDIT.
package types

type User struct {
	ID     int64  `json:"id"`
	Email  string `json:"email"`
	Name   string `json:"name"`
	Sex    int64  `json:"sex"`
	Avatar string `json:"avatar"`
	Info   string `json:"info"`
}

type SignInReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignInResp struct {
	AccessToken  string `json:"access_token"`
	AccessExpire int64  `json:"access_expire"`
	RefreshAfter int64  `json:"refresh_after"`
}

type SignUpReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignUpResp struct {
	AccessToken  string `json:"access_token"`
	AccessExpire int64  `json:"access_expire"`
	RefreshAfter int64  `json:"refresh_after"`
}

type WXAuthReq struct {
	Code          string `json:"code"`
	IV            string `string:"iv"`
	EncryptedData string `json:"encryptedData"`
}

type WXNAutoResp struct {
	AccessToken  string `json:"access_token"`
	AccessExpire int64  `json:"access_expire"`
	RefreshAfter int64  `json:"refresh_after"`
}

type UserInfoReq struct {
}

type UserInfoResp struct {
	UserInfo User `json:"user_info"`
}
