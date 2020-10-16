package models

type Datastore interface {
	UserCred() ([]*Credentials, error)
}
