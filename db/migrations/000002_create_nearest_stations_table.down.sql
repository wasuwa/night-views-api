DROP TRIGGER IF EXISTS refresh_nearest_stations_updated_at_step1 ON nearest_stations;
DROP TRIGGER IF EXISTS refresh_nearest_stations_updated_at_step2 ON nearest_stations;
DROP TRIGGER IF EXISTS refresh_nearest_stations_updated_at_step3 ON nearest_stations;

DROP INDEX IF EXISTS idx_night_view_id;
DROP TABLE IF EXISTS nearest_stations;
