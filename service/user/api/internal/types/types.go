// Code generated by goctl. DO NOT EDIT.
package types

type LoginReq struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type ListReq struct {
	Name string `form:"name,optional"`
	Id   int64  `form:"id,optional"`
}

type CreateReq struct {
	Name     string `form:"name"`
	Password string `form:"password"`
}

type UpdateReq struct {
	Id       int64  `form:"id"`
	Name     string `form:"name"`
	Password string `form:"password"`
}

type LoginResp struct {
	Id           int64  `json:"id"`
	Name         string `json:"name"`
	AccessToken  string `json:"accessToken"`
	AccessExpire int64  `json:"accessExpire"`
	RefreshAfter int64  `json:"refreshAfter"`
}

type ListResp struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}
