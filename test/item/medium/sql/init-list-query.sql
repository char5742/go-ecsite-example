-- accounts
INSERT INTO accounts (id)
VALUES ('00000000-0000-0000-0000-000000000000');
SET app.account_id = '00000000-0000-0000-0000-000000000000';
-- genders
INSERT INTO genders (id, name) VALUES
    ('11111111-1111-1111-1111-111111111111', 'Male'),
    ('22222222-2222-2222-2222-222222222222', 'Female'),
    ('99999999-9999-9999-9999-999999999999', 'Unknown'),
    ('AAAAAAAA-AAAA-AAAA-AAAA-AAAAAAAAAAAA', 'Undefined');

-- breeds
INSERT INTO breeds (id, name) VALUES
    ('33333333-3333-3333-3333-333333333333', 'Bulldog'),
    ('44444444-4444-4444-4444-444444444444', 'Poodle'),
    ('BBBBBBBB-BBBB-BBBB-BBBB-BBBBBBBBBBBB', 'Golden Retriever'),
    ('CCCCCCCC-CCCC-CCCC-CCCC-CCCCCCCCCCCC', 'Shiba Inu');

-- colors
INSERT INTO colors (id, name) VALUES
    ('55555555-5555-5555-5555-555555555555', 'Brown'),
    ('66666666-6666-6666-6666-666666666666', 'White'),
    ('DDDDDDDD-DDDD-DDDD-DDDD-DDDDDDDDDDDD', 'Black'),
    ('EEEEEEEE-EEEE-EEEE-EEEE-EEEEEEEEEEEE', 'Golden');

-- items
INSERT INTO items (
    id,
    description,
    price,
    birthday,
    image,
    gender_id,
    breed_id,
    color_id
)
VALUES
    -- 既存データ
    ('77777777-7777-7777-7777-777777777777',
     'Cute dog',
     10000,
     '2020-01-01',
     'dog.jpg',
     '11111111-1111-1111-1111-111111111111',
     '33333333-3333-3333-3333-333333333333',
     '55555555-5555-5555-5555-555555555555'
    ),
    ('88888888-8888-8888-8888-888888888888',
     'Another dog',
     8000,
     '2021-05-10',
     'another.jpg',
     '22222222-2222-2222-2222-222222222222',
     '44444444-4444-4444-4444-444444444444',
     '66666666-6666-6666-6666-666666666666'
    ),
    -- 追加データ1: Male + Golden Retriever + White
    ('99999999-9999-9999-9999-111111111111',
     'Golden retriever puppy',
     12000,
     '2022-02-10',
     'golden_puppy.jpg',
     '11111111-1111-1111-1111-111111111111',
     'BBBBBBBB-BBBB-BBBB-BBBB-BBBBBBBBBBBB',
     '66666666-6666-6666-6666-666666666666'
    ),
    -- 追加データ2: Female + Shiba Inu + Black
    ('99999999-9999-9999-9999-222222222222',
     'Black Shiba Inu',
     25000,
     '2021-03-15',
     'black_shiba.jpg',
     '22222222-2222-2222-2222-222222222222',
     'CCCCCCCC-CCCC-CCCC-CCCC-CCCCCCCCCCCC',
     'DDDDDDDD-DDDD-DDDD-DDDD-DDDDDDDDDDDD'
    ),
    -- 追加データ3: Unknown + Bulldog + Brown
    ('99999999-9999-9999-9999-333333333333',
     'Cute unknown dog',
     5000,
     '2023-01-01',
     'unknown.jpg',
     '99999999-9999-9999-9999-999999999999',
     '33333333-3333-3333-3333-333333333333',
     '55555555-5555-5555-5555-555555555555'
    ),
    -- 追加データ4: Undefined + Poodle + White
    ('99999999-9999-9999-9999-444444444444',
     'Undefined Poodle',
     15000,
     '2019-06-20',
     'undefined_poodle.jpg',
     'AAAAAAAA-AAAA-AAAA-AAAA-AAAAAAAAAAAA',
     '44444444-4444-4444-4444-444444444444',
     '66666666-6666-6666-6666-666666666666'
    ),
    -- 追加データ5: Male + Shiba Inu + Golden
    ('99999999-9999-9999-9999-555555555555',
     'Golden Shiba Inu',
     30000,
     '2020-11-11',
     'golden_shiba.jpg',
     '11111111-1111-1111-1111-111111111111',
     'CCCCCCCC-CCCC-CCCC-CCCC-CCCCCCCCCCCC',
     'EEEEEEEE-EEEE-EEEE-EEEE-EEEEEEEEEEEE'
    ),
    -- 追加データ6: Female + Golden Retriever + Golden
    ('99999999-9999-9999-9999-666666666666',
     'Female golden retriever older',
     35000,
     '2018-02-28',
     'female_golden.jpg',
     '22222222-2222-2222-2222-222222222222',
     'BBBBBBBB-BBBB-BBBB-BBBB-BBBBBBBBBBBB',
     'EEEEEEEE-EEEE-EEEE-EEEE-EEEEEEEEEEEE'
    );
