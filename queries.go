package main

var SENDER_INFORMATION_QUERY = `
	INSERT INTO sender_information(
		first_name,
		last_name,
		main_phone,
		email,
		address,
		city,
		post_code) values ($1,$2,$3,$4,$5,$6,$7)
		returning sender_id
`

var RECEIVER_INFORMATION_QUERY = `
	INSERT INTO RECEIVER_INFORMATION(
		sender_id
	) values ($1)
	returning receiver_id
`

var BOX_INFORMATION_QUERY = `
	INSERT INTO BOX(
		first_name,
		last_name,
		passport_serial_number,
		main_phone,
		address,
		city,
		receiver_id
		) values ($1,$2,$3,$4,$5,$6,$7)
		returning box_id
`

var LIST_OF_ITEMS_QUERY = `
	INSERT INTO LIST_OF_ITEMS(
		item_name,
		item_quantity,
		box_id
	) values ($1, $2, $3)
`

var TOTAL_VALUE_ITEM_QUERY = `
	INSERT INTO total_value_item (TOTAL_VALUE, box_id) values ($1, $2)
`

var SQL_GET_CLASSIFIEDS =`
SELECT 
   s.sender_id,
    s.first_name,
    s.last_name,
    s.main_phone,
    s.email,
    s.address,
    s.city,
    s.country,
    s.post_code,
    r.receiver_id,
    b.first_name,
    b.last_name,
    b.passport_serial_number,
    b.main_phone,
    b.address,
    b.city,
    b.country,
    array_agg(l.item_name),
    array_agg(l.item_quantity),
    t.total_value
FROM SENDER_INFORMATION as s
    join RECEIVER_INFORMATION as r using(sender_id)
    join BOX as b using(receiver_id)
    join LIST_OF_ITEMS as l on l.box_id = b.box_id
  join total_value_item as t on t.box_id = b.box_id
group by 
  b.box_id, 
  s.sender_id, 
  r.receiver_id, 
  t.total_value
;
`
