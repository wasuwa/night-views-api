CREATE FUNCTION refresh_updated_at_step1() RETURNS trigger AS
$$
BEGIN
  IF NEW.updated_at = OLD.updated_at THEN
    NEW.updated_at := NULL;
  END IF;
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE FUNCTION refresh_updated_at_step2() RETURNS trigger AS
$$
BEGIN
  IF NEW.updated_at IS NULL THEN
    NEW.updated_at := OLD.updated_at;
  END IF;
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE FUNCTION refresh_updated_at_step3() RETURNS trigger AS
$$
BEGIN
  IF NEW.updated_at IS NULL THEN
    NEW.updated_at := CURRENT_TIMESTAMP;
  END IF;
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE SCHEMA IF NOT EXISTS extensions;
CREATE EXTENSION IF NOT EXISTS postgis WITH SCHEMA extensions;

CREATE TABLE IF NOT EXISTS public.night_views(
  id UUID NOT NULL default gen_random_uuid(),
  title TEXT NOT NULL,
  image_url TEXT NOT NULL,
  post_code TEXT NOT NULL,
  address TEXT NOT NULL,
  location extensions.GEOGRAPHY NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TRIGGER refresh_night_views_updated_at_step1
  BEFORE UPDATE ON night_views FOR EACH ROW
  EXECUTE PROCEDURE refresh_updated_at_step1();
CREATE TRIGGER refresh_night_views_updated_at_step2
  BEFORE UPDATE OF updated_at ON night_views FOR EACH ROW
  EXECUTE PROCEDURE refresh_updated_at_step2();
CREATE TRIGGER refresh_night_views_updated_at_step3
  BEFORE UPDATE ON night_views FOR EACH ROW
  EXECUTE PROCEDURE refresh_updated_at_step3();
