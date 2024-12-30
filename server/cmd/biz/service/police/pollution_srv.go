package service

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/renxingdawang/SDataHK/biz/dal/db"
	"github.com/renxingdawang/SDataHK/biz/model/water_pollution"
)

type PollutionService struct {
	ctx context.Context
	c   *app.RequestContext
}

func NewPollutionService(ctx context.Context, c *app.RequestContext) *PollutionService {
	return &PollutionService{ctx: ctx, c: c}
}

func (s *PollutionService) GetPollutionSources(req *water_pollution.WaterPollutionRequest) (resp *water_pollution.WaterPollutionResponse, err error) {
	// 从请求中获取经纬度和半径
	lon := req.Longitude
	lat := req.Latitude
	radius := req.Radius // 假设 Location 结构中有 Radius 字段

	// 调用 DAO 层查询污染源
	sources, err := db.GetPollutionSources(lon, lat, radius)
	if err != nil {
		return nil, fmt.Errorf("failed to get pollution sources: %w", err)
	}

	// 将查询结果转换为 WaterPollutionResponse
	resp = &water_pollution.WaterPollutionResponse{
		PollutionSources: make([]string, 0, len(sources)),
		TotalSources:     int32(len(sources)),
	}

	for _, source := range sources {
		resp.PollutionSources = append(resp.PollutionSources, source.Name)
	}

	return resp, nil
}
