package user

//go:generate go run github.com/vburenin/ifacemaker@latest --file $GOFILE --struct service --iface Service --pkg user --output service_iface.go
type service struct {
}

func CreateUser()
