package live

import (
	`github.com/storezhang/gox`
)

const (
	// original  0 原画
	// 360p  1: 流畅
	// 480p  2: 标清
	// 720p  3: 高清
	// 1080p 4: 超清
	// audio 5: 音频
	// VideoTypeAudio audio 音频
	VideoTypeAudio VideoType = "audio"
	// VideoType360p 360p
	VideoType360p VideoType = "360p"
	// VideoType480p 480p
	VideoType480p VideoType = "480p"
	// VideoType720p 720p
	VideoType720p VideoType = "720p"
	// VideoType1080p  1080p
	VideoType1080p VideoType = "1080p"
	// VideoTypeOriginal  Original 原画
	VideoTypeOriginal VideoType = "original"

	// 视频格式类型 VideoFlowingTypeHls
	VideoFlowingTypeHls VideoFlowingType = "hls"
	// 视频格式类型 VideoFlowingTypeFlv
	VideoFlowingTypeFlv VideoFlowingType = "flv"
	// 视频格式类型 VideoFlowingTypeMp4
	VideoFlowingTypeMp4 VideoFlowingType = "mp4"
	// 视频格式类型 VideoFlowingTypeRtmp
	VideoFlowingTypeRtmp VideoFlowingType = "rtmp"
)

var (
	// VideoTypeMap 视频类型转化
	VideoTypeMap = map[int]string{
		0: "original",
		1: "360p",
		2: "480p",
		3: "720p",
		4: "1080p",
		5: "audio",
	}
)

type (
	// VideoType 视频类型
	VideoType string
	// original  0 原画
	// 360p  1: 流畅
	// 480p  2: 标清
	// 720p  3: 高清
	// 1080p 4: 超清
	// audio 5: 音频

	// VideoFlowingType 视频格式类型
	VideoFlowingType string

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
		// cameras 推流相机列表
		Cameras []*PushCamera `json:"cameras"`
	}

	// PushCamera 推流相机
	PushCamera struct {
		// Index 相机号位
		Index int `json:"index"`
		// Urls 相机视频流信息
		Urls []*VideoFlowing `json:"urls"`
	}

	VideoFlowing struct {
		// Type 视频流类型
		// hls flv mp4 rtmp
		Type VideoFlowingType `json:"type"`
		// Url 流地址
		Url string `json:"url"`
	}

	// GetLivePullReq 获取拉流信息请求
	GetLivePullReq struct {
		// Id 创建直播返回编号
		Id uint64 `json:"id,string" validate:"required"`
	}

	// GetLivePullRsp 获取拉流信息响应
	GetLivePullRsp struct {
		// Cameras 拉流相机
		Cameras []*PullCamera `json:"cameras"`
	}

	// 拉流相机
	PullCamera struct {
		// Index 机位编号
		Index int `json:"index"`
		// Videos 视频信息
		Videos []*PullVideo `json:"videos"`
	}

	PullVideo struct {
		// Type 类型
		// 与创建直播时，videoType中类型对应
		// 原画质：0
		// 流畅: 1;
		// 标清: 2;
		// 高清: 3;
		// 超清: 4;
		// 音频: 5
		Type VideoType `json:"type"`
		// Urls 相机视频流信息
		Urls []*VideoFlowing `json:"urls"`
	}

	// BaseGetLiveReq 获取流基础信息
	BaseGetLiveReq struct {
		// LiveId 直播编号,如果是和直播该字段有效
		LiveId int64 `json:"liveId"`
		// ChannelId 推拉流频道编号
		ChannelId string `json:"channelId" validate:"required"`
	}
)
