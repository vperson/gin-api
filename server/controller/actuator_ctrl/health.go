package actuator_ctrl

type HealthReq struct {
}

type HealthResp struct {
	Msg string `yaml:"msg"`
}

func (r *HealthReq) Exec() (*HealthResp, error) {
	return &HealthResp{Msg: "hello world"}, nil
}
