namespace go police_service

struct Location {
  1: double latitude;
  2: double longitude;
}

struct PoliceStation {
  1: string name;
}

struct NearestPoliceStationsResponse {
  1: string status;
  2: list<PoliceStation> data;
}

service PoliceService {
  NearestPoliceStationsResponse getNearestPoliceStations(1: Location userLocation)(api.get="/police");
}
