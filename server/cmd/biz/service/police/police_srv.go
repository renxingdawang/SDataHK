package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
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
	return nil, nil
}
