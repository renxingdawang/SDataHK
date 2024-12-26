SELECT COUNT(*) FROM gis_osm_buildings_a_free_1;


SELECT gis_osm_pois_a_free_1.name FROM gis_osm_pois_a_free_1 where fclass='park';

SELECT COUNT(*) FROM gis_osm_pofw_a_free_1;

SELECT SUM(gis_osm_places_free_1.population) FROM gis_osm_places_free_1;

SELECT gis_osm_places_free_1.name FROM gis_osm_places_free_1 WHERE population>0;

--select police
SELECT *
FROM gis_osm_pois_free_1
WHERE fclass = 'police';

SELECT *
FROM gis_osm_buildings_a_free_1
WHERE name LIKE '%警%';
--从traffic_a_free_1找到所有的码头并放到新表里：pier
CREATE TABLE gis_all_pier
(
    gid    serial PRIMARY KEY,
    osm_id varchar(12),
    code   smallint,
    fclass varchar(28),
    name   varchar(100),
    geom   geometry(MultiPolygon, 4326)
);
alter table gis_all_pier
    owner to postgres;
create index gis_all_pier_geom_idx
    on gis_all_pier using gist (geom);

INSERT INTO gis_all_pier (gid, osm_id, code, fclass, name, geom)
SELECT gid, osm_id, code, fclass, name, geom
FROM gis_osm_traffic_a_free_1
WHERE fclass = 'pier';

--找到所有警察局 gis_osm_buildings_a_free_1 插入新表gis_all_police中
create table gis_all_police
(
    gid    serial
        primary key,
    osm_id varchar(12),
    code   smallint,
    fclass varchar(28),
    name   varchar(100),
    type   varchar(20),
    geom   geometry(MultiPolygon, 4326)
);
alter table gis_all_police
    owner to postgres;
create index gis_all_police_geom_idx
    on gis_all_police using gist (geom);

INSERT INTO gis_all_police (gid, osm_id, code, fclass, name, type,geom)
SELECT gid, osm_id, code, fclass, name, type, geom
FROM gis_osm_buildings_a_free_1
WHERE name LIKE '%警%';
--找出所有的region并插入到新表里
create table gis_all_region
(
    gid        serial
        primary key,
    osm_id     varchar(12),
    code       smallint,
    fclass     varchar(28),
    population double precision,
    name       varchar(100),
    geom       geometry(MultiPolygon, 4326)
);
alter table gis_all_region
    owner to postgres;
create index gis_all_region_geom_idx
    on gis_all_region using gist (geom);
INSERT INTO gis_all_region (gid, osm_id, code, fclass, population, name,geom)
SELECT gid, osm_id, code, fclass, population, name, geom
FROM gis_osm_places_a_free_1
WHERE fclass='region';

--
DROP FUNCTION find_schools_near_road(TEXT, FLOAT);


--功能一 地铁站推荐商城
CREATE OR REPLACE FUNCTION find_supermarkets_near_subway(subway_name TEXT)
    RETURNS refcursor AS $$
DECLARE
    result_cursor refcursor;
BEGIN
    OPEN result_cursor FOR
SELECT p.name AS supermarket_name, s.name AS subway_name,
       ST_Distance(s.geom, p.geom) AS distance
FROM gis_osm_pois_a_free_1 p, gis_osm_transport_a_free_1 s
WHERE p.fclass = 'mall' AND s.fclass ='railway_station'
  AND s.name = subway_name
  AND st_distancespheroid(s.geom,p.geom)<=1000;
RETURN result_cursor;
END;
$$ LANGUAGE plpgsql;
-- 声明一个游标变量
DO $$
    DECLARE
        cur refcursor;
supermarket_name TEXT;
subway_name TEXT;
distance FLOAT;
BEGIN
    -- 调用函数并获取游标
    cur := find_supermarkets_near_subway('中環 Central');
-- 从游标中获取数据
LOOP
            FETCH cur INTO supermarket_name, subway_name, distance;
EXIT WHEN NOT FOUND;
RAISE NOTICE 'Supermarket: %, Subway: %, Distance: %', supermarket_name, subway_name, distance;
END LOOP;
        -- 关闭游标
CLOSE cur;
END$$;

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

--功能四
