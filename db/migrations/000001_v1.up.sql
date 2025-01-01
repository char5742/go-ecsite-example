BEGIN TRANSACTION;
CREATE TABLE accounts
(
    id         UUID    NOT NULL,
    is_deleted BOOLEAN NOT NULL            DEFAULT FALSE,
    created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT pk_accounts PRIMARY KEY (id)
);

COMMENT ON TABLE accounts IS 'アカウント';
COMMENT ON COLUMN accounts.id IS 'ID';
COMMENT ON COLUMN accounts.created_at IS '作成日時';

-- トリガー関数を作成
CREATE OR REPLACE FUNCTION update_updated_by()
    RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_by := current_setting('app.account_id', true);
    NEW.updated_at := current_timestamp;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TABLE breeds
(
    id         UUID                        DEFAULT gen_random_uuid(),
    name       TEXT NOT NULL UNIQUE,
    created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    created_by UUID NOT NULL               DEFAULT current_setting('app.account_id')::uuid,
    updated_by UUID NOT NULL               DEFAULT current_setting('app.account_id')::uuid,
    CONSTRAINT pk_breeds PRIMARY KEY (id)
);

COMMENT ON TABLE breeds IS '品種';
COMMENT ON COLUMN breeds.id IS 'ID';
COMMENT ON COLUMN breeds.name IS '名前';
COMMENT ON COLUMN breeds.created_at IS '作成日時';
COMMENT ON COLUMN breeds.created_by IS '作成者';
COMMENT ON COLUMN breeds.updated_at IS '更新日時';
COMMENT ON COLUMN breeds.updated_by IS '更新者';

ALTER TABLE breeds
    ADD CONSTRAINT FK_BREEDS_ON_CREATED_BY FOREIGN KEY (created_by) REFERENCES accounts (id);
ALTER TABLE breeds
    ADD CONSTRAINT FK_BREEDS_ON_UPDATED_BY FOREIGN KEY (updated_by) REFERENCES accounts (id);

CREATE TRIGGER update_breeds_trigger
    BEFORE UPDATE ON breeds
    FOR EACH ROW
EXECUTE FUNCTION update_updated_by();

CREATE TABLE colors
(
    id         UUID                        DEFAULT gen_random_uuid(),
    name       TEXT NOT NULL UNIQUE,
    created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    created_by UUID NOT NULL               DEFAULT current_setting('app.account_id')::uuid,
    updated_by UUID NOT NULL               DEFAULT current_setting('app.account_id')::uuid,
    CONSTRAINT pk_colors PRIMARY KEY (id)
);

COMMENT ON TABLE colors IS '色';
COMMENT ON COLUMN colors.id IS 'ID';
COMMENT ON COLUMN colors.name IS '名前';
COMMENT ON COLUMN colors.created_at IS '作成日時';
COMMENT ON COLUMN colors.created_by IS '作成者';
COMMENT ON COLUMN colors.updated_at IS '更新日時';
COMMENT ON COLUMN colors.updated_by IS '更新者';

ALTER TABLE colors
    ADD CONSTRAINT FK_COLORS_ON_CREATED_BY FOREIGN KEY (created_by) REFERENCES accounts (id);
ALTER TABLE colors
    ADD CONSTRAINT FK_COLORS_ON_UPDATED_BY FOREIGN KEY (updated_by) REFERENCES accounts (id);

CREATE TRIGGER update_colors_trigger
    BEFORE UPDATE ON colors
    FOR EACH ROW
EXECUTE FUNCTION update_updated_by();


CREATE TABLE genders
(
    id         UUID                        DEFAULT gen_random_uuid(),
    name       TEXT NOT NULL UNIQUE,
    created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    created_by UUID NOT NULL               DEFAULT current_setting('app.account_id')::uuid,
    updated_by UUID NOT NULL               DEFAULT current_setting('app.account_id')::uuid,
    CONSTRAINT pk_genders PRIMARY KEY (id)
);

COMMENT ON TABLE genders IS '性別';
COMMENT ON COLUMN genders.id IS 'ID';
COMMENT ON COLUMN genders.name IS '名称';
COMMENT ON COLUMN genders.created_at IS '作成日時';
COMMENT ON COLUMN genders.created_by IS '作成者';
COMMENT ON COLUMN genders.updated_at IS '更新日時';
COMMENT ON COLUMN genders.updated_by IS '更新者';

ALTER TABLE genders
    ADD CONSTRAINT FK_GENDERS_ON_CREATED_BY FOREIGN KEY (created_by) REFERENCES accounts (id);
ALTER TABLE genders
    ADD CONSTRAINT FK_GENDERS_ON_UPDATED_BY FOREIGN KEY (updated_by) REFERENCES accounts (id);

CREATE TRIGGER update_genders_trigger
    BEFORE UPDATE ON genders
    FOR EACH ROW
EXECUTE FUNCTION update_updated_by();

CREATE TABLE items
(
    id          UUID                        DEFAULT gen_random_uuid(),
    description TEXT    NOT NULL,
    price       INTEGER NOT NULL,
    image       TEXT    NOT NULL,
    birthday    date    NOT NULL,
    is_deleted  BOOLEAN NOT NULL            DEFAULT FALSE,
    gender_id   UUID    NOT NULL,
    breed_id    UUID    NOT NULL,
    color_id    UUID    NOT NULL,
    created_at  TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    created_by  UUID    NOT NULL            DEFAULT current_setting('app.account_id')::uuid,
    updated_by  UUID    NOT NULL            DEFAULT current_setting('app.account_id')::uuid,
    CONSTRAINT pk_items PRIMARY KEY (id)
);

COMMENT ON TABLE items IS 'ペット';
COMMENT ON COLUMN items.id IS 'ID';
COMMENT ON COLUMN items.description IS '説明';
COMMENT ON COLUMN items.price IS '価格';
COMMENT ON COLUMN items.image IS '画像パス';
COMMENT ON COLUMN items.birthday IS '誕生日';
COMMENT ON COLUMN items.is_deleted IS '削除フラグ';
COMMENT ON COLUMN items.gender_id IS '性別ID';
COMMENT ON COLUMN items.breed_id IS '品種ID';
COMMENT ON COLUMN items.color_id IS '色ID';
COMMENT ON COLUMN items.created_at IS '作成日時';
COMMENT ON COLUMN items.created_by IS '作成者';
COMMENT ON COLUMN items.updated_at IS '更新日時';
COMMENT ON COLUMN items.updated_by IS '更新者';

ALTER TABLE items
    ADD CONSTRAINT FK_ITEMS_ON_GENDER FOREIGN KEY (gender_id) REFERENCES genders (id);

ALTER TABLE items
    ADD CONSTRAINT FK_ITEMS_ON_BREED FOREIGN KEY (breed_id) REFERENCES breeds (id);

ALTER TABLE items
    ADD CONSTRAINT FK_ITEMS_ON_COLOR FOREIGN KEY (color_id) REFERENCES colors (id);

ALTER TABLE items
    ADD CONSTRAINT FK_ITEMS_ON_CREATED_BY FOREIGN KEY (created_by) REFERENCES accounts (id);
ALTER TABLE items
    ADD CONSTRAINT FK_ITEMS_ON_UPDATED_BY FOREIGN KEY (updated_by) REFERENCES accounts (id);

CREATE TRIGGER update_items_trigger
    BEFORE UPDATE ON items
    FOR EACH ROW
EXECUTE FUNCTION update_updated_by();

COMMIT TRANSACTION;