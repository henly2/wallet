package method

import(
	"log"
	"errors"
)

type Module struct{
	Handle func(*string, *string)error
}

func (m *Module) Dispath(req *string, res * string) error {
	log.Println("client call module : ", *req)

	if m.Handle==nil {
		return errors.New("Module Handle is nil")
	}
	m.Handle(req, res)
	return nil;
}
