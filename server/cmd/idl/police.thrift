namespace go police_service

struct Location {
  1: double latitude;
  2: double longitude;
}

struct PoliceStation {
  1: string name;
  2: Location location;
  3: double distance; // 距离单位：米
}

struct NearestPoliceStationsResponse {
  1: string status;
  2: list<PoliceStation> data;
}

service PoliceService {
  NearestPoliceStationsResponse getNearestPoliceStations(1: Location userLocation)(api.get="/police");
}
