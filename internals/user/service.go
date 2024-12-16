// à utiliser pour générer la struct en iface ce qui permet de l'utiliser pour mocker
// exemple
//
//go:generate go run github.com/vburenin/ifacemaker@latest --file $GOFILE --struct service --iface Service --pkg user --output service_iface.go
package user

// ici les imports
// import (
//	"ApiAutoPoto/internals.models"
//)
// pour le go generate laisser les // l'outil est pris en xompte comme ceci le service_iface.go dois générer automatiquement l'interface
//go:generate go run github.com/vburenin/ifacemaker@latest --file $GOFILE --struct service --iface Service --pkg user --output service_iface.go
// type service struct {
//	db *sql.db    ici on a un pointeur vers l'outil représenté par une étoile
//}

// func NewService(db *sql.DB) Service {
//	return &service{db: db}
//}

// ici les fonctions services

// exemple func (s *service) CreateUser(user *models.User) error {
//return err  <- TOUJOURS RETOURNER UN ERROR et si vous ne voules pas vous retournez users par exemple , nil <- nil = null
//}

//run go generate en commande pour générer vos éléments
