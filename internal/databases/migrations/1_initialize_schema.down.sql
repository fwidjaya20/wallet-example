DROP TABLE IF EXISTS "wallets";
DROP TABLE IF EXISTS "wallet_balance_events";
DROP TABLE IF EXISTS "wallet_balance_snapshots";

DROP INDEX IF EXISTS "ix_wallets_id";
DROP INDEX IF EXISTS "ix_wallets_account_holder";
DROP INDEX IF EXISTS "ix_wallets_account_number";
DROP INDEX IF EXISTS "ix_wallet_balance_events_id";
DROP INDEX IF EXISTS "ix_wallet_balance_events_wallet_id";
DROP INDEX IF EXISTS "ix_wallet_balance_events_balance_type";
DROP INDEX IF EXISTS "ix_wallet_balance_snapshots_id";
DROP INDEX IF EXISTS "ix_wallet_balance_snapshots_wallet_id";
DROP INDEX IF EXISTS "ix_wallet_balance_snapshots_last_event_id";