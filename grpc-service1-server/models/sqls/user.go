package sqls

func FindAllUsers() string {
	return `
SELECT
  name,
  age
FROM
  users
`
}

func FindByUserId() string {
	return `
SELECT
  name,
  age
FROM
	users
WHERE
	id = ?
`
}

func CreateUser() string {
	return `
INSERT INTO users(
	name,
	age
)
VALUES(
	?,
	?
)
`
}

func CreateUserCompensate() string {
	return `
DELETE
FROM
	users
WHERE
	id = ?
`
}

func UpdateByUserId() string {
	return `
UPDATE
	users
SET
	name = ?,
	age = ?
WHERE
	id = ?
`
}
