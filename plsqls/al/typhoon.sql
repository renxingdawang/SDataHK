--功能四 台风影响区域分析
CREATE OR REPLACE FUNCTION analyze_typhoon_impact(typhoon_buffer geometry)
    RETURNS TABLE(place_name VARCHAR(28),road_name VARCHAR(28)) AS $$
BEGIN
    RETURN QUERY
SELECT
    c.name AS community_id,
    r.name AS road_id
FROM
    gis_osm_places_a_free_1 c,
    gis_osm_roads_free_1 r
WHERE
    st_intersects(c.geom,typhoon_buffer)
   OR st_intersects(r.geom,typhoon_buffer);
END;
$$ LANGUAGE plpgsql;

-- 创建一个缓冲区并直接调用函数
WITH typhoon_buffer AS (
    SELECT ST_Buffer(ST_SetSRID(ST_MakePoint(114.15, 22.29), 4326)::geography, 5000)::geometry AS geom
)
SELECT *
FROM analyze_typhoon_impact((SELECT geom FROM typhoon_buffer));