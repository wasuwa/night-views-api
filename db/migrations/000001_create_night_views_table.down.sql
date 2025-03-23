DROP TRIGGER IF EXISTS refresh_night_views_updated_at_step1 ON night_views;
DROP TRIGGER IF EXISTS refresh_night_views_updated_at_step2 ON night_views;
DROP TRIGGER IF EXISTS refresh_night_views_updated_at_step3 ON night_views;

DROP FUNCTION IF EXISTS refresh_updated_at_step1();
DROP FUNCTION IF EXISTS refresh_updated_at_step2();
DROP FUNCTION IF EXISTS refresh_updated_at_step3();

DROP TABLE IF EXISTS night_views;

DROP EXTENSION IF EXISTS postgis;
DROP SCHEMA IF EXISTS extensions;
