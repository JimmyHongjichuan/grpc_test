package compose

type Service  interface {
	 SetData(string) error
}

type BaseService struct {
	Msg string
    impl Service
}


func (bs *BaseService) SetData(msg string) error {
	bs.Msg = msg
	return nil
}


