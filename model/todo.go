package model

import (
	"fmt"
	"github.com/rs/xid"
)
type Todo struct {
	ID string
	Description string
}

func CreateTodo(params *Todo)  *Todo{
	td:=new(Todo)
	td.Description=params.Description
	td.ID=xid.New().String()
	fmt.Println(td)
	return td
}