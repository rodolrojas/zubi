package enums

type HTTPMethod string

const (
	MethodGET    HTTPMethod = "GET"
    MethodPOST   HTTPMethod = "POST"
    MethodPUT    HTTPMethod = "PUT"
    MethodDELETE HTTPMethod = "DELETE"
)

func (m HTTPMethod) Valid() bool {
	switch m {
	case MethodGET, MethodPOST, MethodPUT, MethodDELETE:
		return true
	default:
		return false
	}
}
