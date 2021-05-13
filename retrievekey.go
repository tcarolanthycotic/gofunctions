package gofunctions

type mykey struct {
	clientid, clientsecret string
}

func keyreturn(clientid string, clientsecret string) *mykey {
	key := mykey{clientid: clientid}
	key.clientsecret = clientsecret
	return &key
}
