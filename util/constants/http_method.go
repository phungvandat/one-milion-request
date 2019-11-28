package constants

// HTTPMethod struct
type HTTPMethod string

const (
	// POSTMethod is post method
	POSTMethod HTTPMethod = "POST"
	// GETMethod is get method
	GETMethod HTTPMethod = "GET"
	// PUTMethod is put method
	PUTMethod HTTPMethod = "PUT"
)

func (m HTTPMethod) String() string {
	return string(m)
}
