package Information

var current_client_connected int

func Initialize() {
	current_client_connected = 0
}

func GetClients() int {
	return current_client_connected
}

func AddClient() {
	current_client_connected++
}

func RemoveClient() {
	if(current_client_connected > 0) {
		current_client_connected--
	}
}
