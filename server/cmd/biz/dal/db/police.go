package db

import (
	"fmt"
	"github.com/renxingdawang/SDataHK/pkg/errno"
)

type PoliceStation struct {
	GID    int64  `gorm:"primaryKey;column:gid"` // 主键
	OSMID  string `gorm:"column:osm_id"`         // OSM ID
	Code   int16  `gorm:"column:code"`           // 分类代码
	FClass string `gorm:"column:fclass"`         // 分类
	Name   string `gorm:"column:name"`           // 名称
	Type   string `gorm:"column:type"`           // 类型
	Geom   string `gorm:"column:geom"`           // 几何信息，作为 WKT 格式
}

// GetNearestPoliceStations 查询最近的三个警察局
func GetNearestPoliceStations(lon, lat float64) ([]PoliceStation, error) {
	var stations []PoliceStation
	// PostGIS 查询最近的警察局
	query := `
		SELECT gid, osm_id, code, fclass, name, type, ST_AsText(geom) AS geom
		FROM gis_all_police
		ORDER BY ST_Distance(geom, ST_SetSRID(ST_MakePoint(?, ?), 4326))
		LIMIT 3;
	`
	// 执行查询
	result := DB.Raw(query, lon, lat).Scan(&stations)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to query nearest police stations: %w", result.Error)
	}

	if len(stations) == 0 {
		return nil, errno.PoliceSrvErr.WithMessage("len station =0")
	}

	return stations, nil
}
