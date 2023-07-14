package generic

type Dao[T int | string] interface {
	Process(T) (T, error)

	Save(data T) error
}

func DataOperate(p Dao[string]) {

}

type JsonData struct {
}

func (j JsonData) Process(t string) (string, error) {
	// TODO implement me
	panic("implement me")
}

func (j JsonData) Save(data string) error {
	// TODO implement me
	panic("implement me")
}
