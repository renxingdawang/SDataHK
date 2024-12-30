--功能二 查询特定区（行政区）内不同土地利用类型的面积占比
CREATE OR REPLACE FUNCTION calculate_landuse_area_percentage(district_name TEXT)
    RETURNS TABLE (landuse_type VARCHAR(28), percentage FLOAT) AS $$
DECLARE
    total_district_area FLOAT;
landuse_cur CURSOR FOR SELECT DISTINCT fclass FROM gis_osm_landuse_a_free_1;
current_landuse_type VARCHAR(28);
current_landuse_area FLOAT;
BEGIN
SELECT ST_Area(geom) INTO total_district_area
FROM gis_all_region
WHERE name = district_name;
OPEN landuse_cur;
LOOP
        FETCH landuse_cur INTO current_landuse_type;
EXIT WHEN NOT FOUND;
SELECT SUM(ST_Area(geom)) INTO current_landuse_area
FROM gis_osm_landuse_a_free_1
WHERE fclass = current_landuse_type
  AND ST_Intersects(
        geom,
        (SELECT geom FROM gis_all_region WHERE name = district_name)
      );
percentage := current_landuse_area / total_district_area * 100;
RETURN NEXT;
END LOOP;
CLOSE landuse_cur;
END;
$$ LANGUAGE plpgsql;

SELECT * FROM calculate_landuse_area_percentage('九龍 Kowloon');

