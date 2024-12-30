namespace go poi_density

// 请求结构体
struct POIDensityRequest {
    1: double latitude,   // 纬度
    2: double longitude,  // 经度
    3: double radius,     // 半径（米）
}

// 响应结构体
struct POIDensityResponse {
    1: map<string, double> densityByType, // POI类型 -> 密度
}

// 服务接口
service POIDensityService {
    POIDensityResponse GetPOIDensity(1: POIDensityRequest request)(api.get="/poi");
}