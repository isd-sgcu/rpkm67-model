package constant

type Role string

const (
	USER  Role = "user"
	STAFF Role = "staff"
)

func (r Role) String() string {
	return string(r)
}
