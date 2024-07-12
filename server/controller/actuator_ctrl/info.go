package actuator_ctrl

import (
	"gin-api/repository/store"
	"gin-api/repository/store/tb"
)

type InfoReq struct {
}
type InfoResp struct {
	Name string `json:"name"`
}

func (r *InfoReq) Exec() (*InfoResp, error) {
	record, err := store.Common().GetByTypeAndName(tb.CommonTypeSystem, tb.CommonKeySystemName)
	if err != nil {
		return nil, err
	}

	var (
		name string
	)

	if record.ID == 0 {
		name = "the app name is not set"
	} else {
		name = record.Value
	}

	return &InfoResp{Name: name}, nil
}
