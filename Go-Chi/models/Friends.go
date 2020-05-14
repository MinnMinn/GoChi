package models

type Friends struct {
	Success bool	`"json: success"`
	Friends []string `json: "friends"`
	Count int		`"json: count"`
}
