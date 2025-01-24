CREATE TABLE IF NOT EXISTS sequences
(
    sequence_id                     SERIAL PRIMARY KEY,
    account_id                      BIGINT       NOT NULL,
    created_at                      BIGINT       NOT NULL,
    updated_at                      BIGINT DEFAULT NULL,
    sequence_name                   VARCHAR(255) NOT NULL,
    sequence_open_tracking_enabled  BOOLEAN      NOT NULL,
    sequence_click_tracking_enabled BOOLEAN      NOT NULL
);

CREATE TABLE IF NOT EXISTS steps
(
    account_id          BIGINT       NOT NULL,
    step_id             SERIAL PRIMARY KEY,
    sequence_id         BIGINT       NOT NULL,
    created_at          BIGINT       NOT NULL,
    updated_at          BIGINT DEFAULT NULL,
    step_email_subject  VARCHAR(255) NOT NULL,
    step_email_body     TEXT         NOT NULL,
    wait_days           INT          NOT NULL,
    eligible_start_time BIGINT       NOT NULL,
    eligible_end_time   BIGINT       NOT NULL,
    FOREIGN KEY (sequence_id) REFERENCES sequences (sequence_id)
);

-- Insert sample data into sequences table
INSERT INTO sequences (account_id, created_at, sequence_name, sequence_open_tracking_enabled,
                       sequence_click_tracking_enabled)
VALUES (1, 1633036800, 'Welcome Sequence', true, true),
       (2, 1633123200, 'Onboarding Sequence', false, true);

-- Insert sample data into steps table
INSERT INTO steps (account_id, sequence_id, created_at, step_email_subject, step_email_body, wait_days,
                   eligible_start_time, eligible_end_time)
VALUES (1, 1, 1633036800, 'Welcome to our service', 'Thank you for joining us!', 1, 1633036800, 1633123200),
       (1, 1, 1633123200, 'Getting Started', 'Here are some tips to get started.', 2, 1633123200,
        1633209600),
       (2, 2, 1633123200, 'Welcome to the team', 'We are excited to have you!', 1, 1633123200, 1633209600),
       (2, 2, 1633209600, 'Next Steps', 'Here is what you need to do next.', 3, 1633209600, 1633296000);
