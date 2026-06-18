BEGIN;

DROP TABLE IF EXISTS ical_subscription_target_tags;

ALTER TABLE ical_subscriptions
    DROP COLUMN IF EXISTS mode;

COMMIT;
