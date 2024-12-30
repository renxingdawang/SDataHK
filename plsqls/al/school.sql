--功能三 查询路边学校数量和列表
CREATE OR REPLACE FUNCTION find_schools_near_road(road_name TEXT, distance FLOAT)
    RETURNS TABLE (school_name VARCHAR(28), distance_to_road FLOAT) AS $$
BEGIN
    RETURN QUERY
SELECT
    s.name AS school_name,
    ST_Distance(s.geom::geography, r.geom::geography) AS distance_to_road
FROM
    gis_osm_roads_free_1 r,
    gis_osm_pois_a_free_1 s
WHERE
    r.name = road_name
  AND s.fclass = 'school'
  AND st_distancespheroid(r.geom,s.geom)<distance;
END;
$$ LANGUAGE plpgsql;

SELECT * FROM find_schools_near_road('香港仔海傍道 Aberdeen Praya Road', 500);