-- accounts
INSERT INTO accounts (id)
VALUES ('00000000-0000-0000-0000-000000000000');
-- genders
INSERT INTO genders (id, name, created_by, updated_by)
VALUES (
        '11111111-1111-1111-1111-111111111111',
        'Male',
        '00000000-0000-0000-0000-000000000000',
        '00000000-0000-0000-0000-000000000000'
    );
-- breeds
INSERT INTO breeds (id, name, created_by, updated_by)
VALUES (
        '22222222-2222-2222-2222-222222222222',
        'Bulldog',
        '00000000-0000-0000-0000-000000000000',
        '00000000-0000-0000-0000-000000000000'
    );
-- colors
INSERT INTO colors (id, name, created_by, updated_by)
VALUES (
        '33333333-3333-3333-3333-333333333333',
        'Brown',
        '00000000-0000-0000-0000-000000000000',
        '00000000-0000-0000-0000-000000000000'
    );
-- items
INSERT INTO items (
        id,
        description,
        price,
        birthday,
        image,
        gender_id,
        breed_id,
        color_id,
        created_by,
        updated_by
    )
VALUES (
        '44444444-4444-4444-4444-444444444444',
        'Cute dog',
        10000,
        '2020-01-01',
        'dog.jpg',
        '11111111-1111-1111-1111-111111111111',
        '22222222-2222-2222-2222-222222222222',
        '33333333-3333-3333-3333-333333333333',
        '00000000-0000-0000-0000-000000000000',
        '00000000-0000-0000-0000-000000000000'
    );
