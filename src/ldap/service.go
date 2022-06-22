package ldap

import (
	"fmt"

	goldap "github.com/go-ldap/ldap/v3"
)

const (
	rootDN = "dc=example,dc=com"
	userDN = "ou=users"
)

type Service interface {
	Get(cn string) (*goldap.Entry, error)
}

type ldapService struct{}

func NewLDAPService() Service {
	return &ldapService{}
}

func (s *ldapService) connect() (*goldap.Conn, error) {
	conn, err := goldap.DialURL("ldap://localhost:1389")
	if err != nil {
		return nil, err
	}

	err = conn.Bind("cn=admin,"+rootDN, "admin")
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func (s *ldapService) Get(cn string) (*goldap.Entry, error) {
	result, err := s.search("ou=users,dc=example,dc=com", 1, "(&(cn=%s))", cn, []string{"dn", "cn", "email", "entryUUID"})
	if err != nil {
		return nil, err
	}

	return result.Entries[0], nil
}

func (s *ldapService) search(base string, scope int, filter, value string, attributes []string) (*goldap.SearchResult, error) {
	conn, err := s.connect()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	searchRequest := goldap.NewSearchRequest(base, scope, 0, 0, 0, false, fmt.Sprintf(filter, value), attributes, nil)

	sr, err := conn.Search(searchRequest)
	if err != nil {
		return nil, err
	}

	return sr, nil
}
