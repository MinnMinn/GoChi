package models

type Recipients struct {
	Success bool	`"json: success"`
	Recipients []string `json: "friends"`
}
