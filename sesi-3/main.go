package main

import (
	"fmt"
	"os"
	"strconv"
)

type Student struct {
	ID             int
	Name           string
	Address        string
	Job            string
	ReasonToEnroll string
}

func main() {

	arg := os.Args
	students := []Student{
		{ID: 0, Name: "Rahman", Address: "Jalan yang benar", Job: "X", ReasonToEnroll: "Pengen aja"},
		{ID: 1, Name: "Aziz", Address: "Jalan sesama", Job: "Y", ReasonToEnroll: "Bingung"},
		{ID: 2, Name: "Firmansyah", Address: "Jalan berliku", Job: "Z", ReasonToEnroll: "Gatau"},
	}

	if len(arg) <= 1 {
		fmt.Println("Tolong masukkan nama atau nomor absen")
		fmt.Println("Contoh: `go run main.go` Aziz atau `go run main.go 2`")
		return
	}

	element := elementOf(students, arg[1])
	if element == nil {
		fmt.Println("Data siswa tidak ditemukan")
	} else {
		fmt.Println(fmt.Sprintf("ID: %d", element.ID))
		fmt.Println(fmt.Sprintf("Nama: %s", element.Name))
		fmt.Println(fmt.Sprintf("Alamat: %s", element.Address))
		fmt.Println(fmt.Sprintf("Pekerjaan: %s", element.Job))
		fmt.Println(fmt.Sprintf("Alasan: %s", element.ReasonToEnroll))
	}

}

// elementOf return Student data according to an argument supplied by user
func elementOf(slc []Student, arg string) *Student {
	var element *Student
	for _, s := range slc {
		if containsArg(arg, s.ID, s.Name) {
			element = &s
			break
		}
	}
	return element
}

// containsArg checks if id or name of a student is matched with the argument
func containsArg(arg string, id int, name string) bool {
	if strconv.Itoa(id) == arg || name == arg {
		return true
	}
	return false
}
