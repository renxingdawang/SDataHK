namespace go water_pollution

struct WaterPollutionRequest {
    1: double latitude,
    2: double longitude,
    3: double radius, // in meters
}

struct WaterPollutionResponse {
    1: list<string> pollutionSources,
    2: i32 totalSources,
}

service WaterPollutionService {
    WaterPollutionResponse GetPollutionSources(1: WaterPollutionRequest request)(api.get="/pollution");
}