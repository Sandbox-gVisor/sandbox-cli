package cli

func sendChange(address string, dtos []CallbackDto) {
	SendToSocket(address, "change-callbacks", dtos)
}

func sendState(address string, dtos []CallbackDto) {
	SendToSocket(address, "change-state", dtos)
}

func sendInfo(address string) {
	SendToSocket(address, "hooks-info", []CallbackDto{})
}
