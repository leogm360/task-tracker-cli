package data

type DataSourcer interface {
	Read(any) error
	Write(any) (bool, error)
}
