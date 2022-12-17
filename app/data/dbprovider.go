package data

type DbProvider[T any] interface {
	Create(*T) error
	Get() ([]T, error)
	GetById(id string) (T, error)
	Update(*T) error
}

// type DbProvider interface {
// 	Create(*entities.CustomerEntity) error
// 	Get() ([]entities.CustomerEntity, error)
// 	GetById(id string) (entities.CustomerEntity, error)
// 	Update(*entities.CustomerEntity) error
// }
