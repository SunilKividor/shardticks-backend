package postgresql

import (
	"bookmyshow/models"
	"log"
)

func CreateShowsQuery(show models.CreateShow) error {
	smt := `INSERT INTO shows (organizer, show, price, quantity, TicketsRem)
	VALUES ($1, $2, $3, $4, $5)`

	_, err := db.Exec(smt, show.Username, show.Show, show.Price, show.Quantity, show.TicketsRem)
	if err != nil {
		return err
	}

	return nil
}

func AddOrganizerNftsQuery(organizernts models.OrganizerNFTs) error {
	smt := `INSERT INTO organizernfts (username, nftids) VALUES($1,$2)`
	_, err := db.Exec(smt, organizernts.Username, organizernts.Username)
	if err != nil {
		return err
	}
	return nil
}

func GetAllUserShowsQuery() ([]models.CreateShow, error) {
	var shows []models.CreateShow
	smt := `SELECT organizer,show,price,quantity,ticketsrem FROM shows`
	rows, err := db.Query(smt)
	if err != nil {
		return shows, err
	}
	defer rows.Close()
	for rows.Next() {
		var show models.CreateShow
		err := rows.Scan(&show.Username, &show.Show, &show.Show, &show.TicketsRem)
		if err != nil {
			return shows, err
		}
		shows = append(shows, show)
	}
	if rows.Err() != nil {
		return shows, err
	}
	log.Println(shows)
	return shows, nil
}
