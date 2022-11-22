package models

type Contacto struct {
	ID             int64
	PrimerNombre   string
	SegundoNombre  string
	PrimerApellido string
	SegundoApelido string
	Email          string
	Birthday       string
	Telefono       string
}
