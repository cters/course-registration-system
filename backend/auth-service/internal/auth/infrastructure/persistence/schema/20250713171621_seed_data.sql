-- +goose Up
-- +goose StatementBegin
INSERT INTO
    go_user (
        user_account,
        user_password,
        user_email,
        user_name,
        user_phone,
        user_salt
    )
VALUES
    (
        'student1',
        '35be8821d78848c2009b7fd75cdb74538a450902f5824f75ca49492c0e02f098',
        'student1@hcmut.edu.vn',
        'Nguyen Van Student1',
        '0000000001',
        '34c7c9a65171805513a47c0eb72ac19c'
    ),
    (
        'student2',
        'a3e5d70a3bb6dc902e84a28ef8c466ac82c3cc0d23f64fb99bf1a0d20f2aae02',
        'student2@hcmut.edu.vn',
        'Nguyen Van Student2',
        '0000000002',
        '089129870e16ef98dd5d6cba558d54d5'
    ),
    (
        'student3',
        '093786c874910ecaf003a79dc70aa45102d434900e3628cfa239bf9da32613e4',
        'student3@hcmut.edu.vn',
        'Nguyen Van Student3',
        '0000000003',
        '92ce712f5df1f1be60b06b70eac073a3'
    ),
    (
        'student4',
        '121b9b1a129e1d0177159cc3000def0cfb1dcb78badaa98bcfd5350e727117a8',
        'student4@hcmut.edu.vn',
        'Nguyen Van Student4',
        '0000000004',
        '800bf18f6f987ba5ca74ce386d7e286c'
    ),
    (
        'student5',
        'fc698e2735eebab18bb1a7f3826c30c3229580618805c3858dda28ebb86f7d0e',
        'student5@hcmut.edu.vn',
        'Nguyen Van Student5',
        '0000000005',
        '222bab2b17b8d8f24df71bf09ab83177'
    ),
    (
        'instructor1',
        'c97fb70a8b5cabe93b4058bc4c9a23611c8e80fc0f86f9815d400b98f95d31bf',
        'instructor1@hcmut.edu.vn',
        'Nguyen Van Instructor1',
        '0000000011',
        '5a8186b1170d41b027127638c64b0a73'
    ),
    (
        'instructor2',
        '6c357c7511f2d7917609660eb4b0f90c2b71b91047547fcfb9ab271580c2b132',
        'instructor2@hcmut.edu.vn',
        'Nguyen Van Instructor2',
        '0000000012',
        '8db7bd46efc331b13ede255e6bea2db6'
    ),
    (
        'instructor3',
        '577c7ebeae1766bc3a460992cb1a825a89adb94e69c074d47b37a4f4fd9ceefd',
        'instructor3@hcmut.edu.vn',
        'Nguyen Van Instructor3',
        '0000000013',
        '55df1fcb49089f35474b3d1876560e07'
    ),
    (
        'instructor4',
        '63a5c666da40eca628e23616dd8402d578bd2a42f92cd579a21d8a00cded2948',
        'instructor4@hcmut.edu.vn',
        'Nguyen Van Instructor4',
        '0000000014',
        'f6d09bb810ba8d947af85b2e523ed4e1'
    ),
    (
        'instructor5',
        '6832d52807ec800479d6eea79253d7172738520176d95fdf33fa4de07f6f5e2b',
        'instructor5@hcmut.edu.vn',
        'Nguyen Van Instructor5',
        '0000000015',
        '9c4678cdef2db92910091d1286115e0c'
    ),
    (
        'operator1',
        'fba1c77bed2d1dacc14847e6444f88ef2236dbd1cc74f396f69d70f97dda90ea',
        'operator1@hcmut.edu.vn',
        'Nguyen Van Operator1',
        '0000000111',
        '7a49a4354144c23e2977ee0e47fb62bf'
    ),
    (
        'operator2',
        '04bf963b18cfa41f9157f66d0691d26e050bd1c44c09b8b40e356494a2258195',
        'operator2@hcmut.edu.vn',
        'Nguyen Van Operator2',
        '0000000112',
        '3e3160b1fa6a4da19172d7e6e963a756'
    ),
    (
        'admin',
        'ccf3f3e32fadeeed0bdd5a2d2386fbb2e884b21dbd9cbcba93fa7a33c4b3b085',
        'admin@hcmut.edu.vn',
        'Nguyen Van Admin',
        '1111111111',
        '92cc5e67212a4e3690657b140e29ad19'
    );

-- +goose StatementEnd
-- +goose StatementBegin
INSERT INTO
    go_student (
        student_id,
        student_credit,
        student_enrollment_year,
        student_gpa,
        student_department_id,
        user_id
    )
VALUES
    (2212001, 81, 2022, 2.1, 5, 1),
    (2212002, 82, 2022, 2.2, 5, 2),
    (2212003, 83, 2022, 2.3, 5, 3),
    (2212004, 84, 2022, 2.4, 5, 4),
    (2212005, 85, 2022, 2.5, 5, 5);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
-- SELECT
--     'down SQL query';
-- +goose StatementEnd