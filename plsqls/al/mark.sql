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