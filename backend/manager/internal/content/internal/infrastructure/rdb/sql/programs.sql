-- name: CreateProgramsBatch01 :execresult
INSERT INTO programs (id, question, answer, content_id) VALUES
  (?, ?, ?, ?);

-- name: CreateProgramsBatch02 :execresult
INSERT INTO programs (id, question, answer, content_id) VALUES
  (?, ?, ?, ?),
  (?, ?, ?, ?);

-- name: CreateProgramsBatch03 :execresult
INSERT INTO programs (id, question, answer, content_id) VALUES
  (?, ?, ?, ?),
  (?, ?, ?, ?),
  (?, ?, ?, ?);

-- name: CreateProgramsBatch04 :execresult
INSERT INTO programs (id, question, answer, content_id) VALUES
  (?, ?, ?, ?),
  (?, ?, ?, ?),
  (?, ?, ?, ?),
  (?, ?, ?, ?);

-- name: CreateProgramsBatch05 :execresult
INSERT INTO programs (id, question, answer, content_id) VALUES
  (?, ?, ?, ?),
  (?, ?, ?, ?),
  (?, ?, ?, ?),
  (?, ?, ?, ?),
  (?, ?, ?, ?);
