package mysql

import "fmt"

type Mysql struct {
	USER string
}

func New() *Mysql {
	return &Mysql{}
}

func (*Mysql) Find(id int) string {

	return "Data from mysql"

}

func (*Mysql) Save(data string) error {
	fmt.Println("Save data to mysql")
	return nil
}
