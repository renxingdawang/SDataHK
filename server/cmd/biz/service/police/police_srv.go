package service

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/renxingdawang/SDataHK/biz/dal/db"
	police "github.com/renxingdawang/SDataHK/biz/model/police_service"
)

type PoliceService struct {
	ctx context.Context
	c   *app.RequestContext
}

func NewPoliceService(ctx context.Context, c *app.RequestContext) *PoliceService {
	return &PoliceService{ctx: ctx, c: c}
}

func (s *PoliceService) GetNearestPoliceStations(req *police.Location) (resp *police.NearestPoliceStationsResponse, err error) {
	// 检查输入参数
	if req == nil || req.Longitude == 0 || req.Latitude == 0 {
		return nil, fmt.Errorf("invalid location request")
	}

	// 查询 DAO
	stations, err := db.GetNearestPoliceStations(req.GetLongitude(), req.GetLatitude())
	if err != nil {
		return nil, fmt.Errorf("failed to get nearest police stations: %w", err)
	}

	// 构造响应
	resp = police.NewNearestPoliceStationsResponse()
	for _, station := range stations {
		// 构造 `PoliceStation` 对象并添加到响应中
		resp.Data = append(resp.Data, &police.PoliceStation{
			Name: station.Name,
		})
	}
	return resp, nil
}
