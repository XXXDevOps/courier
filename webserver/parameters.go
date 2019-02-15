package webserver

import "github.com/NioDevOps/courier/models"

type SendParams struct {
	Media     string           `json:"media"`
	Sender    models.Sender    `json:"sender"`
	Receivers models.Receivers `json:"receivers"`
	Message   models.Message   `json:"message"`
	Cc        models.Cc        `json:"cc"`
}
