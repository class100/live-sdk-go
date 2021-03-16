package live

import (
	`github.com/class100/core`
)

const (
	ApiPathLiveCreate ApiPath = "live"
	ApiPathLivePush   ApiPath = "live/pull"
	ApiPathLivePull   ApiPath = "live/push"
)

type liver interface {
	// CreateLive 创建直播
	CreateLive(req *CreateLiveReq) (live *CreateLiveRsp, err error)
	// GetLivePush 获取直播推流
	GetLivePush(req *GetLivePushReq) (rsp *GetLivePushRsp, err error)
	// GetLivePull 获取直播拉流
	GetLivePull(req *GetLivePullReq) (rsp *GetLivePullRsp, err error)
}

func (client *httpClient) CreateLive(req *CreateLiveReq) (live *CreateLiveRsp, err error) {
	live = new(CreateLiveRsp)
	err = client.requestApi(
		ApiPathLiveCreate,
		core.HttpMethodPost,
		req,
		live,
	)

	return
}

func (client *httpClient) GetLivePush(req *GetLivePushReq) (rsp *GetLivePushRsp, err error) {
	rsp = new(GetLivePushRsp)
	err = client.requestApi(
		ApiPathLivePush,
		core.HttpMethodGet,
		req,
		rsp,
	)

	return
}

func (client *httpClient) GetLivePull(req *GetLivePullReq) (rsp *GetLivePullRsp, err error) {
	rsp = new(GetLivePullRsp)
	err = client.requestApi(
		ApiPathLivePull,
		core.HttpMethodGet,
		req,
		rsp,
	)

	return
}
