namespace go urban

// 请求结构体
struct UrbanHeatIslandRequest {
    1: double latitude,  // 中心点纬度
    2: double longitude, // 中心点经度
    3: double radius,    // 半径（米）
}

// 响应结构体
struct UrbanHeatIslandResponse {
    1: double heatIslandIndex, // 城市热岛效应指数
    2: map<string, double> landUseAreas, // 土地利用类型 -> 面积
}

// 服务接口
service UrbanHeatIslandService {
    UrbanHeatIslandResponse AnalyzeUrbanHeatIsland(1: UrbanHeatIslandRequest request)(api.get="/urban");
}