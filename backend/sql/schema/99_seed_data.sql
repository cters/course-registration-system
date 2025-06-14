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
    go_role (role_name)
VALUES
    ('guest'),
    ('student'),
    ('instructor'),
    ('operator'),
    ('admin');

-- +goose StatementEnd
-- +goose StatementBegin
INSERT INTO
    go_user_role (user_id, role_id)
VALUES
    (1, 2),
    (2, 2),
    (3, 2),
    (4, 2),
    (5, 2),
    (6, 3),
    (7, 3),
    (8, 3),
    (9, 3),
    (10, 3),
    (11, 4),
    (12, 4),
    (13, 5);

-- +goose StatementEnd
-- +goose StatementBegin
INSERT INTO
    go_permission (permission_name)
VALUES
    ('read:catalog'),
    ('read:course'),
    ('read:course_detail'),
    ('register:course'),
    ('drop:course'),
    ('read:own_registrations'),
    ('read:own_grades'),
    ('read:own_courses_taught'),
    ('read:students_in_own_courses'),
    ('update:grades_own_courses'),
    ('create:course'),
    ('update:course'),
    ('delete:course'),
    ('manage:schedule'),
    ('manage:enrollment'),
    ('manage:user'),
    ('manage:role'),
    ('manage:system');

-- +goose StatementEnd
-- +goose StatementBegin
INSERT INTO
    go_subject (subject_id, subject_title, subject_credit)
VALUES
    ('MT1003', 'Calculus 1', 4),
    ('MT1005', 'Calculus 2', 4),
    ('MT1007', 'Linear Algebra', 3),
    ('MT2013', 'Probability and Statistics', 4),
    ('CH1003', 'General Chemistry', 3),
    ('PH1003', 'General Physics 1', 4),
    ('PH1007', 'General Physics Labs', 1),
    ('CO1007', 'Introduction to Computing', 4),
    ('CO2011', 'Digital Systems', 3),
    ('SP1007', 'Professional Skills for Engineers', 2),
    ('SP1031', 'Programming Fundamentals', 3),
    ('SP1033', 'Advanced Programming', 2),
    (
        'SP1035',
        'Principles of Programming Languages',
        2
    ),
    ('SP1037', 'Computer Architecture', 2),
    ('SP1039', 'Data Structures and Algorithms', 2),
    ('CO1005', 'Mathematical Modeling', 3),
    (
        'IM1013',
        'Introduction to Artificial Intelligence',
        3
    ),
    ('IM1023', 'Machine Learning', 3),
    ('IM1025', 'Natural Language Processing', 3),
    ('IM1027', 'Data Mining', 3),
    ('IM3001', 'Capstone Project', 3),
    ('CO2001', 'Programming Integration Project', 3),
    ('LA1003', 'English 1', 2),
    ('LA1005', 'English 2', 2),
    ('LA1007', 'English 3', 2),
    ('LA1009', 'English 4', 2),
    ('CO1023', 'Operating Systems', 3),
    ('CO1027', 'Software Engineering', 3),
    ('CO2003', 'Computer Networks', 4),
    ('CO2007', 'Database Systems', 4),
    ('CO2013', 'Advanced Programming', 4),
    ('CO2039', 'Computer Architecture', 3),
    ('CO2017', 'Data Structures and Algorithms', 3),
    ('CO3001', 'Software Engineering', 3),
    ('CO3005', 'Programming Integration Project', 4),
    ('CO3093', 'Capstone Project', 3),
    ('CO3101', 'Internship', 1),
    ('CO3103', 'Specialized Project', 1),
    ('CO3105', 'Japanese 1', 1),
    ('CO3127', 'Japanese 2', 1),
    ('CO3107', 'Japanese 3', 1),
    ('CO3109', 'Japanese 4', 1),
    ('CO3111', 'Japanese 5', 1),
    (
        'CO3011',
        'Distributed and Object-Oriented Databases',
        3
    ),
    ('CO3013', 'Electronic Commerce', 3),
    (
        'CO3015',
        'Data Warehouses and Decision Support Systems',
        3
    ),
    ('CO3017', 'Information System Security', 3),
    ('CO3021', 'Systems Analysis and Design', 3),
    (
        'CO3023',
        'Big Data Analytics and Business Intelligence',
        3
    ),
    (
        'CO3027',
        'Enterprise Resource Planning Systems',
        3
    ),
    ('CO3029', 'Management Information Systems', 3),
    ('CO3031', 'Biometric Security', 3),
    ('CO3033', 'Software Project Management', 3),
    ('CO3035', 'Compiler Construction', 3),
    ('CO3037', 'Software Testing', 3),
    ('CO3041', 'Software Architecture', 3),
    ('CO3043', 'Advanced Software Engineering', 3),
    (
        'CO3045',
        'Selected Topics in High Performance Computing',
        3
    ),
    ('CO3047', 'Mobile Application Development', 3),
    ('CO3049', 'Game Programming', 3),
    ('CO3051', 'Web Programming', 3),
    ('CO3057', 'Advanced Computer Networks', 3),
    ('CO3059', 'Cryptography and Network Security', 3),
    ('CO3061', 'Distributed Systems', 3),
    (
        'CO3065',
        'Advance Cryptography and Coding Theory',
        3
    ),
    ('CO3067', 'Information and Social Networks', 3),
    (
        'CO3069',
        'Data Warehouses and Decision Support Systems',
        3
    ),
    (
        'CO3071',
        'Big Data Analytics and Business Intelligence',
        3
    ),
    (
        'CO3083',
        'Enterprise Resource Planning Systems',
        3
    ),
    ('CO3085', 'Management Information Systems', 3),
    ('CO3089', 'Biometric Security', 3),
    ('CO3115', 'Internship', 3),
    ('CO3117', 'Capstone Project', 3),
    ('CO3335', 'Japanese Culture', 2),
    ('CO4025', 'Military Training', 3),
    ('CO4031', 'Physical Education', 3),
    ('CO4033', 'Free Electives', 3),
    ('CO4035', 'Graduation Project', 3),
    ('CO4037', 'Specialized Project', 3),
    ('CO4039', 'Capstone Project', 3);

-- +goose StatementEnd
-- +goose StatementBegin
INSERT INTO
    go_course (
        subject_id,
        course_term,
        course_max_slot,
        course_current_slot
    )
SELECT
    subject_id,
    251, -- Set term to 251
    2000, -- Default max slot
    0 -- Default current slot
FROM
    go_subject;

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
-- +goose StatementEnd