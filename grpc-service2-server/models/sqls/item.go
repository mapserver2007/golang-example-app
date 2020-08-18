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

func CreateItem() string {
	return `
  INSERT INTO items(
    name,
    price
  )
  VALUES(
    ?,
    ?
  )
`
}

func CreateItemCompensate() string {
	return `
DELETE
FROM
  items
WHERE
  id = ?
`
}
