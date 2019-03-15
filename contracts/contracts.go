package contracts

type Finder interface {
	FindByNumber(number string) (Lyr, error)
}

type Lyr interface {
}