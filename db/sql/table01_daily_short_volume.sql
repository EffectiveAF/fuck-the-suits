CREATE TABLE daily_short_volume (
    gid                 text      NOT NULL UNIQUE,
    date_orig           text      NOT NULL CHECK (LENGTH(date_orig) = 8),  -- '20210407'
    date                date      NOT NULL,
    symbol              text      NOT NULL,
    short_volume        integer   NOT NULL,
    short_exempt_volume integer   NOT NULL,
    total_volume        integer   NOT NULL,
    market              text[]    NOT NULL CHECK (ARRAY_LENGTH(market, 1) > 0),  -- ['Q', 'N']
    created             timestamp WITH time zone NOT NULL DEFAULT now()
);
ALTER TABLE daily_short_volume OWNER TO superuser;

CREATE OR REPLACE FUNCTION set_daily_short_volume_gid() RETURNS trigger AS $$
DECLARE
    correct_gid text := NEW.date_orig::text || '_' || NEW.symbol::text;
BEGIN
    IF NEW.gid IS DISTINCT FROM correct_gid THEN
        NEW.gid := correct_gid;
    end IF;
    RETURN NEW;
END;
$$ LANGUAGE PLPGSQL;
ALTER FUNCTION set_daily_short_volume_gid() OWNER TO superuser;

CREATE TRIGGER trigger_set_daily_short_volume_gid
    BEFORE INSERT OR UPDATE ON daily_short_volume
    FOR EACH ROW
    EXECUTE PROCEDURE set_daily_short_volume_gid();

CREATE INDEX daily_short_volume_gid_idx ON daily_short_volume (gid);
ALTER INDEX daily_short_volume_gid_idx OWNER TO superuser;
