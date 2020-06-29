CREATE TABLE IF NOT EXISTS "wallets"
(
    "id"             UUID        NOT NULL,
    "account_holder" TEXT        NOT NULL,
    "account_number" TEXT        NOT NULL,
    "created_at"     TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "created_by"     VARCHAR(255),
    "updated_at"     TIMESTAMPTZ,
    "updated_by"     VARCHAR(255),
    CONSTRAINT "uq_wallets_id" UNIQUE ("id"),
    CONSTRAINT "pk_wallets" PRIMARY KEY ("id")
);
CREATE INDEX IF NOT EXISTS "ix_wallets_id" ON "wallets" USING btree("id");
CREATE INDEX IF NOT EXISTS "ix_wallets_account_holder" ON "wallets" USING btree("account_holder");
CREATE INDEX IF NOT EXISTS "ix_wallets_account_number" ON "wallets" USING btree("account_number");

CREATE TABLE IF NOT EXISTS "wallet_balance_events"
(
    "id"           UUID         NOT NULL,
    "wallet_id"    UUID         NOT NULL,
    "amount"       NUMERIC      NOT NULL,
    "balance_type" VARCHAR(100) NOT NULL,
    "notes"        JSONB        NOT NULL,
    "created_at"   TIMESTAMPTZ  NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "created_by"   VARCHAR(255),
    CONSTRAINT "uq_wallet_balance_events_id" UNIQUE ("id"),
    CONSTRAINT "pk_wallet_balance_events" PRIMARY KEY ("id")
);
CREATE INDEX IF NOT EXISTS "ix_wallet_balance_events_id" ON "wallet_balance_events" USING btree("id");
CREATE INDEX IF NOT EXISTS "ix_wallet_balance_events_wallet_id" ON "wallet_balance_events" USING btree("wallet_id");
CREATE INDEX IF NOT EXISTS "ix_wallet_balance_events_balance_type" ON "wallet_balance_events" USING btree("balance_type");

CREATE TABLE IF NOT EXISTS "wallet_balance_snapshots"
(
    "id"            UUID        NOT NULL,
    "wallet_id"     UUID        NOT NULL,
    "balance"       NUMERIC     NOT NULL,
    "last_event_id" UUID        NOT NULL,
    "created_at"    TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "created_by"    VARCHAR(255),
    CONSTRAINT "uq_wallet_balance_snapshots_id" UNIQUE ("id"),
    CONSTRAINT "pk_wallet_balance_snapshots_id" PRIMARY KEY ("id")
);
CREATE INDEX IF NOT EXISTS "ix_wallet_balance_snapshots_id" ON "wallet_balance_snapshots" USING btree("id");
CREATE INDEX IF NOT EXISTS "ix_wallet_balance_snapshots_wallet_id" ON "wallet_balance_snapshots" USING btree("wallet_id");
CREATE INDEX IF NOT EXISTS "ix_wallet_balance_snapshots_last_event_id" ON "wallet_balance_snapshots" USING btree("last_event_id");