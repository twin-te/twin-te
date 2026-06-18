BEGIN;

ALTER TABLE ical_subscriptions
    ADD COLUMN mode text NOT NULL DEFAULT 'sync';

CREATE TABLE ical_subscription_target_tags (
    ical_subscription_id uuid NOT NULL,
    tag_id uuid NOT NULL,
    PRIMARY KEY (ical_subscription_id, tag_id),
    FOREIGN KEY (ical_subscription_id) REFERENCES ical_subscriptions(id) ON DELETE CASCADE,
    FOREIGN KEY (tag_id) REFERENCES tags(id) ON DELETE CASCADE
);

COMMIT;
