package user

type UserTypeID uint

const (
	Admin  UserTypeID = 1
	Client UserTypeID = 2
)

type User struct {
	ID       int64      `json:"id,omitempty"`
	Email    string     `json:"email"`
	Name     string     `json:"name"`
	Type     UserTypeID `json:"type"`
	LdapUUID string     `json:"-"`
	Password string     `json:"-"`
}

type UserType struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name"`
}
