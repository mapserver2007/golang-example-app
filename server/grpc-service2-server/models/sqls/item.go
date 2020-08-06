package sqls

func FindAllItems() string {
	return `
SELECT
  name,
  price
FROM
  items
`
}

func FindByItemId() string {
	return `
SELECT
  name,
  price
FROM
	items
WHERE
	id = ?
`
}
