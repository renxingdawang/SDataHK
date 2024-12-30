package urban

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/renxingdawang/SDataHK/biz/dal/db"
	urban "github.com/renxingdawang/SDataHK/biz/model/urban"
)

type UrbanHeatIslandService struct {
	ctx context.Context
	c   *app.RequestContext
}

func NewUrbanHeatIslandService(ctx context.Context, c *app.RequestContext) *UrbanHeatIslandService {
	return &UrbanHeatIslandService{ctx: ctx, c: c}
}

func (s *UrbanHeatIslandService) AnalyzeUrbanHeatIsland(req *urban.UrbanHeatIslandRequest) (resp *urban.UrbanHeatIslandResponse, err error) {
	// 调用 DAO 层查询热岛效应数据
	heatIslandIndex, landUseAreas, err := db.AnalyzeUrbanHeatIsland(req.Latitude, req.Longitude, req.Radius)
	if err != nil {
		return nil, fmt.Errorf("failed to analyze urban heat island: %w", err)
	}

	// 返回响应
	resp = &urban.UrbanHeatIslandResponse{
		HeatIslandIndex: heatIslandIndex,
		LandUseAreas:    landUseAreas,
	}

	return resp, nil
}
