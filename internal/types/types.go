// Code generated by goctl. DO NOT EDIT.
package types

type InfoResponse struct {
	BlockLastUpdateTime int64
	BlockNum            int64
	CurrentTime         int64
}

type HealthRequest struct {
	Timeout int64 `form:"timeout"`
}
