package productos

import (
	"errors"
	"fmt"

	"strconv"
)

type Producto struct {
	ID          int
	Name        string
	Quantity    int
	CodeValue   string
	IsPublished bool
	Expiration  string
	Price       float64
}

type RequestProduct struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

// vamos a practicar ya hacer lo que es similar dto
type Data struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

type ResponseBody struct {
	Mensjage string `json:"mensaje"`
	Data     *Data  `json:"product"`
}

var Productos []Producto

func ValidateProduct(data *Producto) (err error) {
	AddID(data)
	err = nil
	_, err1 := ValidateEmpty(data)
	if err1 != nil {

		err = err1
	}

	return
}

func ValidateEmpty(data *Producto) (ok bool, err error) {

	switch {

	case (data).Price == 0.0:
		err = errors.New("Price empty")
		ok = false
	case (data).Expiration == "":
		err = errors.New("Expiration empty")
		ok = false
	case (data).Name == "":
		err = errors.New("Name empty")
		ok = false
	case (data).CodeValue == "":
		err = errors.New("Codevalue empty")
		ok = false
	case (data).Quantity <= 0:
		err = errors.New("Quantity empty o negative number")
		ok = false
	default:
		err = nil
		ok = false
		return

	}

	return
}

func AddID(data *Producto) {
	for {
		ind := 1
		c := len(Productos) + ind
		fmt.Println("el valor de len", c)
		flag := false
		for _, v := range Productos {
			if c == v.ID {
				flag = true

			}
		}
		if flag == false {
			(*data).ID = c
			break

		}
		ind++
		fmt.Println(ind, flag)
		fmt.Println(data)
	}

}

func ValidateCodeValue(data *Producto) (err error) {
	//tiene que ser todo mayuscula almenos 1 letra y un numero
	var (
		cletra = 0
		cnum   = 0
		flag   = false
	)
	for i, v := range (*data).CodeValue {
		fmt.Println(v, i)
		switch {
		case i == 0 && v == 48:
			err = errors.New("El primer caracter no puede ser cero")
			return
		case v >= 65 && v <= 90:
			cletra++
			flag = true
			err = nil
		case i != 0 && v >= 48 && v <= 57:
			flag = true
			cnum++
			err = nil

		default:

		}

	}
	if cletra == 0 {
		err = errors.New("Debe haber almenos una letra")
	}
	if flag == false && cnum == 0 {
		err = errors.New("Debe haber almenos un numero")
	}
	return
}

func ValidateDate(data *Producto) (err error) {

	if len((*data).Expiration) != 10 {
		err = errors.New("DATE NO FORMAT dd/mm/yyyy")
		return
	}
	v := (*data).Expiration
	fmt.Println(v)
	fmt.Println(v[0], v[1], v[2], v[3], v[4], v[5], v[6], v[7], v[8], v[9])

	if v[2] != 47 && v[5] != 47 {
		err = errors.New("separate parameters invalid please use /")
	}

	dd, err2 := strconv.Atoi(v[0:2])
	if err2 != nil {
		err = errors.New("Days no valid")
	}
	fmt.Println(dd)
	mm, err3 := strconv.Atoi(v[3:5])
	if err3 != nil {
		err = errors.New("Mouth no valid")
		return
	}
	fmt.Println(mm)
	yyyy, err4 := strconv.Atoi(v[6:])
	if err4 != nil {
		err = errors.New("Year no valid")
		return
	}
	fmt.Println(yyyy)
	switch {
	case dd > 31 || dd <= 0:
		err = errors.New("days out of range ")
		return
	case mm > 12 || mm <= 0:
		err = errors.New("mouth out of range ")
		return
	case yyyy < 2023:
		err = errors.New("year out of range ")
		return
	default:
	}

	return
}

func SearchById(id int) (val Producto, err error) {
	flag := false
	for _, v := range Productos {
		if id == v.ID {
			val = v
			flag = true
			err = nil
		}
	}
	if flag == false {
		err = errors.New(fmt.Sprintf("Not found with id %d", id))
	}

	return

}
func SearchMayPrice(id float64) (out []Producto, err error) {

	flag := false
	for _, v := range Productos {
		if v.Price >= id {
			out = append(out, v)
			err = nil
			flag = true
		}
	}
	if flag == false {
		err = errors.New("Not found...")
		return
	}
	return

}
