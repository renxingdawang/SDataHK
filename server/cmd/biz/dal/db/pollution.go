package db

import (
	"fmt"
)

// PollutionSource 表示污染源的结构
type PollutionSource struct {
	GID      int64   `gorm:"primaryKey;column:gid"` // 主键
	OSMID    string  `gorm:"column:osm_id"`         // OSM ID
	Code     int16   `gorm:"column:code"`           // 分类代码
	FClass   string  `gorm:"column:fclass"`         // 分类
	Name     string  `gorm:"column:name"`           // 名称
	Geom     string  `gorm:"column:geom"`           // 几何信息，作为 WKT 格式
	Distance float64 `gorm:"-"`                     // 距离水域的直线距离（米），不映射到数据库
}

// GetPollutionSources 查询指定水域附近的污染源
func GetPollutionSources(lon, lat, radius float64) ([]PollutionSource, error) {
	var sources []PollutionSource
	// PostGIS 查询指定半径范围内的污染源
	query := `
		SELECT 
			gid, osm_id, code, fclass, name, 
			ST_AsText(geom) AS geom,
			ST_Distance(geom, ST_SetSRID(ST_MakePoint(?, ?), 4326)) AS distance
		FROM 
			public.gis_osm_landuse_a_free_1
		WHERE 
			fclass IN ('industrial', 'wastewater_plant') -- 假设污染源的土地利用类型为工业区或污水处理厂
			AND ST_DWithin(geom, ST_SetSRID(ST_MakePoint(?, ?), 4326), ?)
		UNION ALL
		SELECT 
			gid, osm_id, code, fclass, name, 
			ST_AsText(geom) AS geom,
			ST_Distance(geom, ST_SetSRID(ST_MakePoint(?, ?), 4326)) AS distance
		FROM 
			public.gis_osm_pois_a_free_1
		WHERE 
			fclass IN ('factory', 'chemical_plant') -- 假设污染源的POI类型为工厂或化工厂
			AND ST_DWithin(geom, ST_SetSRID(ST_MakePoint(?, ?), 4326), ?)
		ORDER BY 
			distance;
	`
	// 执行查询
	result := DB.Raw(query, lon, lat, lon, lat, radius, lon, lat, lon, lat, radius).Scan(&sources)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to query pollution sources: %w", result.Error)
	}

	if len(sources) == 0 {
		return nil, fmt.Errorf("no pollution sources found within the specified radius")
	}

	return sources, nil
}
