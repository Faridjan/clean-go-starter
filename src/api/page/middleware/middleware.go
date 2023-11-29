package middleware

import (
	"tiny-template/src/api/page"
)

type Middleware func(service page.ServiceInterface) page.ServiceInterface
