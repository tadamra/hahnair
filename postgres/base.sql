CREATE TABLE IF NOT EXISTS cards_info (
   card_id                NUMERIC       NOT NULL,
   card_holder_name       VARCHAR(64)   NOT NULL,
   car_type               VARCHAR(64) NOT NULL,
   card_expiry_date_month NUMERIC NOT NULL,
   card_expiry_date_year  NUMERIC NOT NULL,
   address1               VARCHAR(256) NOT NULL,
   address2               VARCHAR(256) NOT NULL,
   address3               VARCHAR(256) NOT NULL,
   postal_code            VARCHAR(128) NOT NULL,
   city                   VARCHAR(64) NOT NULL,
   holder_state           VARCHAR(64) NOT NULL,
   country_code           VARCHAR(16)         ,
   CONSTRAINT card_pk PRIMARY KEY (card_id)
);