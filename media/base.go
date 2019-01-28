package media

import (
	"fmt"
	"github.com/NioDevOps/courier/models"
)

var MediaCenter = &Center{Map: map[string]models.Media{}}

type Center struct {
	Map map[string]models.Media
}
type BaseMedia struct {
	Name string
}

func GetMediaCenter() *Center {
	return MediaCenter
}

func (cs *Center) Register(m models.Media) error {
	_, ok := cs.Map[m.GetName()]
	var err error = nil
	if ok {
		err = fmt.Errorf("media named '%s' is already registed", m.GetName())
	} else {
		cs.Map[m.GetName()] = m
	}
	fmt.Println(cs.Map)
	return err
}

func (cs *Center) Get(name string) (models.Media, error) {
	m, ok := cs.Map[name]
	var err error = nil
	if !ok {
		err = fmt.Errorf("no media named '%s'", name)
	}
	return m, err
}

func (bm *BaseMedia) GetName() string {
	return bm.Name
}
