package models

type Request struct {
	Requestor string `"json: requestor"`
	Target string `"json: target"`
}
