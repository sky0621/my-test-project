-- name: ListContentsWithPrograms :many
SELECT
    c.id AS content_id,
    c.name AS content_name,
    p.id AS program_id,
    p.question AS question,
    p.answer AS answer
FROM
    contents AS c
        LEFT JOIN programs AS p ON p.content_id = c.id;

-- name: SearchContentsWithPrograms :many
SELECT
    c.id AS content_id,
    c.name AS content_name,
    p.id AS program_id,
    p.question AS question,
    p.answer AS answer
FROM
    contents AS c
        lEFT JOIN programs AS p ON p.content_id = c.id
WHERE
    c.name LIKE ?;

-- name: GetContentWithProgramsById :many
SELECT
    c.id AS content_id,
    c.name AS content_name,
    p.id AS program_id,
    p.question AS question,
    p.answer AS answer
FROM
    contents AS c
        lEFT JOIN programs AS p ON p.content_id = c.id
WHERE
    c.id = UUID_TO_BIN(?);
