package middleware

import (
	"tiny-template/src/api/category"
)

type Middleware func(service category.ServiceInterface) category.ServiceInterface
