package live

import (
	`github.com/storezhang/gox`
)

type (
	// CreateLiveReq 创建直播请求
	CreateLiveReq struct {
		// Title 直播名称 最大长度为128个字
		Title string `json:"title" validate:"required,max=128"`
		// StartTime 直播开始时间
		StartTime gox.Timestamp `json:"startTime"`
		// EndTime 直播结束时间,结束时间必须大于开始时间
		EndTime gox.Timestamp `json:"endTime"`
	}

	// CreateLiveRsp 创建直播响应
	CreateLiveRsp struct {
		// id 直播编号
		Id int64 `json:"id,string"`
	}

	// GetLivePushReq 获取推流地址请求
	GetLivePushReq struct {
		// Id 创建直播返回编号
		Id uint64 `json:"id,string" validate:"required"`
	}

	// GetLivePushRsp 获取推流地址返回响应
	GetLivePushRsp struct {
		// cameraList 推流相机列表
		CameraList []*PushCamera `json:"cameraList"`
	}

	// PushCamera 推流相机
	PushCamera struct {
		// Status 机会状态
		// 0 未开始
		// 1 直播中
		// 2 暂停
		// 3 结束
		Status int64 `json:"status"`
		// CamIndex 机位编号
		CamIndex string `json:"camIndex"`
		// Url 推流地址
		Url string `json:"url"`
	}

	// GetLivePullReq 获取拉流信息请求
	GetLivePullReq struct {
		// Id 创建直播返回编号
		Id uint64 `json:"id,string" validate:"required"`
	}

	// GetLivePullRsp 获取拉流信息响应
	GetLivePullRsp struct {
		// 拉流相机
		CameraList []*PullCamera `json:"cameraList"`
	}

	// 拉流相机
	PullCamera struct {
		// CamIndex 机位编号
		CamIndex string `json:"camIndex"`
		// transcodeList 转码列表
		TranscodeList []*PullCameraTranscode `json:"transcodeList"`
	}

	// pullCameraTranscode 拉流相机转码
	PullCameraTranscode struct {
		// TransIndex 转码编号
		TransIndex string `json:"transIndex"`
		// TransType 转码类型
		// 与创建直播时，videoType中类型对应
		// 原画质：0
		// 流畅: 1;
		// 标清: 2;
		// 高清: 3;
		// 超清: 4;
		// 音频: 5
		TransType string `json:"transType"`
		// UrlFlv flv观看地址
		UrlFlv string `json:"urlFlv"`
		// UrlHls  hls观看地址
		UrlHls string `json:"urlHls"`
		// UrlRtmp rtmp观看地址
		UrlRtmp string `json:"urlRtmp"`
	}

	// BaseGetLiveReq 获取流基础信息
	BaseGetLiveReq struct {
		// LiveId 直播编号,如果是和直播该字段有效
		LiveId int64 `json:"liveId"`
		// ChannelId 推拉流频道编号
		ChannelId string `json:"channelId" validate:"required"`
	}
)
