package sqlserver

import "fmt"

type Sqlserver struct{}

func New() *Sqlserver {
	return &Sqlserver{}
}

func (*Sqlserver) Find(id int) string {

	return "Data from Sqlserver"

}

func (*Sqlserver) Save(data string) error {
	fmt.Println("Save data to Sqlserver")
	return nil
}
