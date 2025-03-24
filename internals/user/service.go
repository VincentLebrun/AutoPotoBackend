package user

import "github.com/go-playground/validator"

//go:generate go run github.com/vburenin/ifacemaker@latest --file $GOFILE --struct service --iface Service --pkg user --output service_iface.go
var validate = validator.New()
