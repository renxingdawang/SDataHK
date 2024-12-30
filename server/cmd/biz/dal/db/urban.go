package db

import "fmt"

type LandUseArea struct {
	FClass string  `gorm:"column:fclass"` // 土地利用类型
	Area   float64 `gorm:"column:area"`   // 面积
}

type BuildingDensity struct {
	Count int     `gorm:"column:count"` // 建筑物数量
	Area  float64 `gorm:"column:area"`  // 区域面积
}

func AnalyzeUrbanHeatIsland(lat, lon, radius float64) (float64, map[string]float64, error) {
	// 查询土地利用类型及其面积
	var landUseAreas []LandUseArea
	landUseQuery := `
		SELECT fclass, SUM(ST_Area(geom)) AS area
		FROM gis_osm_landuse_a_free_1
		WHERE ST_DWithin(geom, ST_SetSRID(ST_MakePoint(?, ?), 4326), ?)
		GROUP BY fclass;
	`
	if err := DB.Raw(landUseQuery, lon, lat, radius).Scan(&landUseAreas).Error; err != nil {
		return 0, nil, fmt.Errorf("failed to query land use areas: %w", err)
	}

	// 查询建筑物密度
	var buildingDensity BuildingDensity
	buildingQuery := `
		SELECT COUNT(*) AS count, ST_Area(ST_Buffer(ST_SetSRID(ST_MakePoint(?, ?), 4326), ?)) AS area
		FROM gis_osm_buildings_a_free_1
		WHERE ST_DWithin(geom, ST_SetSRID(ST_MakePoint(?, ?), 4326), ?);
	`
	if err := DB.Raw(buildingQuery, lon, lat, radius, lon, lat, radius).Scan(&buildingDensity).Error; err != nil {
		return 0, nil, fmt.Errorf("failed to query building density: %w", err)
	}

	// 计算城市热岛效应指数
	heatIslandIndex := calculateHeatIslandIndex(buildingDensity, landUseAreas)

	// 将土地利用类型及其面积转换为 map
	landUseAreaMap := make(map[string]float64)
	for _, landUse := range landUseAreas {
		landUseAreaMap[landUse.FClass] = landUse.Area
	}

	return heatIslandIndex, landUseAreaMap, nil
}

// calculateHeatIslandIndex 计算城市热岛效应指数
func calculateHeatIslandIndex(buildingDensity BuildingDensity, landUseAreas []LandUseArea) float64 {
	// 简单示例：热岛效应指数 = 建筑物密度 * 0.5 + 工业区面积占比 * 0.3 + 商业区面积占比 * 0.2
	buildingDensityValue := float64(buildingDensity.Count) / buildingDensity.Area

	industrialArea := 0.0
	commercialArea := 0.0
	totalArea := 0.0

	for _, landUse := range landUseAreas {
		totalArea += landUse.Area
		if landUse.FClass == "industrial" {
			industrialArea += landUse.Area
		} else if landUse.FClass == "commercial" {
			commercialArea += landUse.Area
		}
	}

	industrialRatio := industrialArea / totalArea
	commercialRatio := commercialArea / totalArea

	heatIslandIndex := buildingDensityValue*0.5 + industrialRatio*0.3 + commercialRatio*0.2
	return heatIslandIndex
}
