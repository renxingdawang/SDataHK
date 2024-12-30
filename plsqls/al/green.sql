--功能五 计算每个区域的绿色覆盖率
CREATE OR REPLACE FUNCTION calculate_green_coverage()
    RETURNS TABLE(community_id integer, community_name varchar, green_area double precision, total_area double precision, green_coverage_percentage double precision) AS $$
BEGIN
    RETURN QUERY
SELECT
    c.gid AS community_id,
    c.name AS community_name,
    COALESCE(SUM(ST_Area(ST_Intersection(c.geom, l.geom))), 0) AS green_area,
    ST_Area(c.geom) AS total_area,
    COALESCE(SUM(ST_Area(ST_Intersection(c.geom, l.geom))) / ST_Area(c.geom) * 100, 0) AS green_coverage_percentage
FROM
    public.gis_osm_places_a_free_1 c
        LEFT JOIN
    public.gis_osm_landuse_a_free_1 l
    ON
        l.fclass = 'park' AND ST_Intersects(c.geom, l.geom)
GROUP BY
    c.gid, c.name;
END;
$$ LANGUAGE plpgsql;
SELECT *
FROM calculate_green_coverage();