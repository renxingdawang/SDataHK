package db

import (
	"fmt"
)

type POI struct {
	GID    int64  `gorm:"primaryKey;column:gid"` // 主键
	OSMID  string `gorm:"column:osm_id"`         // OSM ID
	Code   int16  `gorm:"column:code"`           // 分类代码
	FClass string `gorm:"column:fclass"`         // 分类
	Name   string `gorm:"column:name"`           // 名称
	Geom   string `gorm:"column:geom"`           // 几何信息，作为 WKT 格式
}

// GetPOIDensity 统计POI密度
func GetPOIDensity(lon, lat, radius float64) (map[string]float64, error) {
	var results []struct {
		FClass string `gorm:"column:fclass"`
		Count  int    `gorm:"column:count"`
	}

	// 统计不同类型POI的数量
	query := `
		SELECT fclass, COUNT(*) AS count
		FROM gis_osm_pois_free_1
		WHERE ST_DWithin(geom, ST_SetSRID(ST_MakePoint(?, ?), 4326), ?)
		GROUP BY fclass;
	`

	result := DB.Raw(query, lon, lat, radius).Scan(&results)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to query POI density: %w", result.Error)
	}

	// 计算区域面积
	areaQuery := `
		SELECT ST_Area(ST_Buffer(ST_SetSRID(ST_MakePoint(?, ?), 4326), ?)) AS area;
	`
	var area float64
	if err := DB.Raw(areaQuery, lon, lat, radius).Scan(&area).Error; err != nil {
		return nil, fmt.Errorf("failed to calculate area: %w", err)
	}

	// 计算POI密度（数量/区域面积）
	poiDensity := make(map[string]float64)
	for _, res := range results {
		poiDensity[res.FClass] = float64(res.Count) / area
	}

	return poiDensity, nil
}
