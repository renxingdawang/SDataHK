package poi

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/renxingdawang/SDataHK/biz/dal/db"
	"github.com/renxingdawang/SDataHK/biz/model/poi_density"
)

type POIService struct {
	ctx context.Context
	c   *app.RequestContext
}

func NewPOIService(ctx context.Context, c *app.RequestContext) *POIService {
	return &POIService{ctx: ctx, c: c}
}

func (s *POIService) GetPOIDensity(req *poi_density.POIDensityRequest) (resp *poi_density.POIDensityResponse, err error) {
	// 调用DAO层查询POI密度
	density, err := db.GetPOIDensity(req.Longitude, req.Latitude, req.Radius)
	if err != nil {
		return nil, fmt.Errorf("failed to get POI density: %w", err)
	}

	// 返回响应
	return &poi_density.POIDensityResponse{
		DensityByType: density,
	}, nil
}
