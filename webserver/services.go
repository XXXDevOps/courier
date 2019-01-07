package webserver

import (
	"github.com/NioDevOps/courier/media"
	"github.com/NioDevOps/courier/models"
)

func SendService(media2 string, sender models.Sender, receivers models.Receivers, message models.Message, cc models.Cc)error{
	m,err :=media.GetMediaCenter().Get(media2)
	if err !=nil{
		return err
	}
	t := models.NewTransaction(sender,receivers,message,cc)
	return m.Send(t)
}
