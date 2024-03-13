package models

import "github.com/google/uuid"

type User struct {
	Id           uuid.UUID `json:"id"`
	Name         string    `json:"name"`
	Username     string    `json:"username"`
	Password     string    `json:"password"`
	RefreshToken string    `json:"refresh_token"`
}

type OrganizerNFTs struct {
	Username string `json:"username"`
	NFTids   []int  `json:"nft_ids"`
}
