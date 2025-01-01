BEGIN TRANSACTION;
-- トリガー削除
DROP TRIGGER IF EXISTS update_items_trigger ON items;
DROP TRIGGER IF EXISTS update_genders_trigger ON genders;
DROP TRIGGER IF EXISTS update_colors_trigger ON colors;
DROP TRIGGER IF EXISTS update_breeds_trigger ON breeds;

-- テーブルの外部キー制約削除
ALTER TABLE items DROP CONSTRAINT IF EXISTS FK_ITEMS_ON_UPDATED_BY;
ALTER TABLE items DROP CONSTRAINT IF EXISTS FK_ITEMS_ON_CREATED_BY;
ALTER TABLE items DROP CONSTRAINT IF EXISTS FK_ITEMS_ON_COLOR;
ALTER TABLE items DROP CONSTRAINT IF EXISTS FK_ITEMS_ON_BREED;
ALTER TABLE items DROP CONSTRAINT IF EXISTS FK_ITEMS_ON_GENDER;

ALTER TABLE genders DROP CONSTRAINT IF EXISTS FK_GENDERS_ON_UPDATED_BY;
ALTER TABLE genders DROP CONSTRAINT IF EXISTS FK_GENDERS_ON_CREATED_BY;

ALTER TABLE colors DROP CONSTRAINT IF EXISTS FK_COLORS_ON_UPDATED_BY;
ALTER TABLE colors DROP CONSTRAINT IF EXISTS FK_COLORS_ON_CREATED_BY;

ALTER TABLE breeds DROP CONSTRAINT IF EXISTS FK_BREEDS_ON_UPDATED_BY;
ALTER TABLE breeds DROP CONSTRAINT IF EXISTS FK_BREEDS_ON_CREATED_BY;

-- テーブル削除
DROP TABLE IF EXISTS items;
DROP TABLE IF EXISTS genders;
DROP TABLE IF EXISTS colors;
DROP TABLE IF EXISTS breeds;
DROP TABLE IF EXISTS accounts;

-- トリガー関数削除
DROP FUNCTION IF EXISTS update_updated_by;

COMMIT TRANSACTION;