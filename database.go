package main

import (
	"database/sql"
	_ "github.com/lib/pq"
)


func PostClassifieds(body PostInformation) PostedInformation {

	db, err := sql.Open("postgres", DB_CONFIG)
	
	if err != nil {
		panic (err)
	}

	defer db.Close()

	var posted PostedInformation

	err = db.QueryRow(SENDER_INFORMATION_QUERY, 
		body.Sender.FirstName,
		body.Sender.LastName,
		body.Sender.MainPhone,
		body.Sender.Email,
		body.Sender.Address,
		body.Sender.City,
		body.Sender.PostCode,
	).Scan(&posted.SenderId)

	if err != nil {
		panic(err)
	}

	

	err = db.QueryRow(RECEIVER_INFORMATION_QUERY, 
		posted.SenderId,
	).Scan(&posted.ReceiverId)

	if err != nil {
		panic(err)
	}

	

	err = db.QueryRow(BOX_INFORMATION_QUERY, 
		body.Box.FirstName,
		body.Box.LastName,
		body.Box.PassportSerialNumber,
		body.Box.MainPhone,
		body.Box.Address,
		body.Box.City,
		posted.ReceiverId,
	).Scan(&posted.BoxId)


	if err != nil {
		panic(err)
	}


	for _, i := range body.Items{
		db.QueryRow(LIST_OF_ITEMS_QUERY, i.ItemName, i.Quantity, posted.BoxId)
	}

	db.QueryRow(TOTAL_VALUE_ITEM_QUERY, body.TotalValue.TotalValueItem, posted.BoxId)

	if err != nil {
		panic(err)
	}

	return posted
}

func GetClassifieds () []PostInformation {

	db, err := sql.Open("postgres", DB_CONFIG)

	defer db.Close()

	if err != nil {
		panic(err)
	}

	rows, err := db.Query(SQL_GET_CLASSIFIEDS)

	defer rows.Close()

	if err != nil {
		panic(err)
	}

	var classifieds []PostInformation
	var it = Item{}
	for rows.Next() {

		var classified PostInformation

		rows.Scan(
			&classified.Sender.SenderId,
			&classified.Sender.FirstName,
			&classified.Sender.LastName,
			&classified.Sender.MainPhone,
			&classified.Sender.Email,
			&classified.Sender.Address,
			&classified.Sender.City,
			&classified.Sender.Country,
			&classified.Sender.PostCode,
			&classified.Receiver.ReceiverId,
			&classified.Box.FirstName,
			&classified.Box.LastName,
			&classified.Box.PassportSerialNumber,
			&classified.Box.MainPhone,
			&classified.Box.Address,
			&classified.Box.City,
			&classified.Box.Country,
			&it.ItemName,
			&it.Quantity,
			&classified.TotalValue.TotalValueItem,
		)

		classified.Items = append(classified.Items, it)
		classifieds = append(classifieds, classified)

	}


	return classifieds
}