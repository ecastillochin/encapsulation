package course

import "fmt"

//Course plplp
type Course struct {
	Name    string
	Price   float64
	IsFree  bool
	UserIDs []uint
	Moduls  map[uint]string
}

//tener en cuenta que al trabajar con punteros en metodos se recomienda usar todos de tipos punteros
//en este ejemplo ChangePrice es de tipo puntero por lo cual actualizamos PrintCourse
func (c *Course) ChangePrice(price float64) {
	c.Price = price
}

//se agrega funcion a la clase, se indica el tipo que recibe, no se usa this, self
func (c *Course) PrintCourse() {
	text := " Las clases son: "
	for _, class := range c.Moduls {
		text += class + ", "
	}

	fmt.Println(text[:len(text)-2])
}
