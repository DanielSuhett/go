package client

type Owner struct {
	name string
	SSN  string
	profession string
}

func Create(name string, SSN string, profession string) Owner {
	return Owner{name, SSN, profession}
}