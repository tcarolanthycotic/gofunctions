package gofunctions

type Mykey struct {
	clientid, clientsecret string
}

func Keyreturn(clientid string, clientsecret string) *Mykey {
	key := mykey{clientid: clientid}
	key.clientsecret = clientsecret
	return &key
}
