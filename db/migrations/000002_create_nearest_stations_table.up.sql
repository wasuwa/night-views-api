CREATE TABLE IF NOT EXISTS public.nearest_stations (
  id UUID NOT NULL DEFAULT gen_random_uuid() PRIMARY KEY,
  night_view_id UUID NOT NULL,
  route_name TEXT NOT NULL,
  station_name TEXT NOT NULL,
  walking_time_from_station integer NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT fk_night_view
    FOREIGN KEY (night_view_id)
    REFERENCES night_views (id)
    ON DELETE RESTRICT
);

CREATE TRIGGER refresh_nearest_stations_updated_at_step1
  BEFORE UPDATE ON nearest_stations FOR EACH ROW
  EXECUTE PROCEDURE refresh_updated_at_step1();
CREATE TRIGGER refresh_nearest_stations_updated_at_step2
  BEFORE UPDATE OF updated_at ON nearest_stations FOR EACH ROW
  EXECUTE PROCEDURE refresh_updated_at_step2();
CREATE TRIGGER refresh_nearest_stations_updated_at_step3
  BEFORE UPDATE ON nearest_stations FOR EACH ROW
  EXECUTE PROCEDURE refresh_updated_at_step3();

CREATE INDEX idx_night_view_id ON nearest_stations(night_view_id);
