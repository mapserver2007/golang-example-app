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
