package gofunctions

type mykey struct {
	clientid, clientsecret string
}

func Keyreturn(clientid string, clientsecret string) *mykey {
	key := mykey{clientid: clientid}
	key.clientsecret = clientsecret
	return &key
}
