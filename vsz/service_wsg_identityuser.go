package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/identity"
)

type WSGIdentityUserService struct {
	client *Client
}

func NewWSGIdentityUserService(client *Client) *WSGIdentityUserService {
	s := new(WSGIdentityUserService)
	s.client = client
	return s
}

func (ss *WSGService) WSGIdentityUserService() *WSGIdentityUserService {
	serv := new(WSGIdentityUserService)
	serv.client = ss.client
	return serv
}

// AddIdentityUserList
//
// Use this API command to retrieve a list of identity user.
//
// Request Body:
//	 - body *identity.QueryCriteria
func (s *WSGIdentityUserService) AddIdentityUserList(ctx context.Context, body *identity.QueryCriteria) (*identity.UserList, error) {
}

// FindIdentityUsers
//
// Use this API command to retrieve a list of identity user.
//
// Query Parameters:
// - qCreatedOnFrom string
// - qCreatedOnTo string
// - qDisplayName string
// - qEmail string
// - qFirstName string
// - qIndex string
// - qIsDisabled string
// - qLastName string
// - qListSize string
// - qPhone string
// - qTimeZone string
// - qUserName string
// - qUserSource string
// - qUserType string
func (s *WSGIdentityUserService) FindIdentityUsers(ctx context.Context, qCreatedOnFrom string, qCreatedOnTo string, qDisplayName string, qEmail string, qFirstName string, qIndex string, qIsDisabled string, qLastName string, qListSize string, qPhone string, qTimeZone string, qUserName string, qUserSource string, qUserType string) (*identity.UserList, error) {
}

// FindIdentityUsersAaaserver
//
// Use this API command to retrieve a list of aaa server.
func (s *WSGIdentityUserService) FindIdentityUsersAaaserver(ctx context.Context) (*identity.AaaServerList, error) {
}

// FindIdentityUsersById
//
// Use this API command to retrieve identity user.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGIdentityUserService) FindIdentityUsersById(ctx context.Context, pId string) (*identity.UserConfiguration, error) {
}

// FindIdentityUsersCountries
//
// Use this API command to retrieve a list of countries.
func (s *WSGIdentityUserService) FindIdentityUsersCountries(ctx context.Context) (*identity.CountryList, error) {
}

// FindIdentityUsersPackages
//
// Use this API command to retrieve a list of packages.
func (s *WSGIdentityUserService) FindIdentityUsersPackages(ctx context.Context) (*identity.PackageList, error) {
}
