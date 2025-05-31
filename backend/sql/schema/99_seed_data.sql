-- +goose Up
-- +goose StatementBegin
INSERT INTO
    go_role (role_name)
VALUES
    ('guest'),
    ('student'),
    ('instructor'),
    ('operator');

-- +goose StatementEnd
-- +goose StatementBegin
INSERT INTO
    go_permission (permission_name)
VALUES
    ('readOwn'),
    ('readAny'),
    ('createOwn'),
    ('createAny'),
    ('deleteOwn'),
    ('deleteAny'),
    ('updateOwn'),
    ('updateAny');

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
-- +goose Down
-- +goose StatementBegin
-- +goose StatementEnd