package start

import "issuerserver/services"

func CreateDefaultUser() {
	username := "account@knn3"
	services.CreateIdentity(username)
}
