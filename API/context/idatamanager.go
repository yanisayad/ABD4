/*
 * File: idatamanager.go
 * Project: ABD4/VMD Escape Game
 * File Created: Thursday, 11th October 2018 4:41:33 pm
 * Author: ayad_y billaud_j castel_a masera_m
 * Contact: (ayad_y@etna-alternance.net billaud_j@etna-alternance.net castel_a@etna-alternance.net masera_m@etna-alternance.net)
 * -----
 * Last Modified: Monday, 29th October 2018 10:41:30 pm
 * Modified By: Aurélien Castellarnau
 * -----
 * Copyright © 2018 - 2018 ayad_y billaud_j castel_a masera_m, ETNA - VDM EscapeGame API
 */

package context

import (
	"ABD4/API/model"
)

// IDataManager is the common part for all Data Manager
type IDataManager interface {
	Init(map[string]string) error
	GetDB() interface{}
	SetDB(dbObject interface{}) error
	GetEntity() string
	SetEntity(entity string)
	GetDBName() string
	SetDBName(dbName string)
}

// IUserManager define what is needed to access
// user data whatever database is used
type IUserManager interface {
	IDataManager
	FindAll() ([]*model.User, error)
	FindBy(map[string]string) ([]*model.User, error)
	FindOneBy(map[string]string) (*model.User, error)
	RemoveBy(map[string]string) (int, error)
	RemoveAll() (int, error)
	Create(user *model.User) (*model.User, error)
	// For dev
	FakePersist(n int) error
}

// ITransactionManager define what is needed to access
// transaction data whatever database is used
type ITransactionManager interface {
	IDataManager
	FindAll() ([]*model.Transaction, error)
	FindBy(map[string]string) ([]*model.Transaction, error)
	FindOneBy(map[string]string) (*model.Transaction, error)
	RemoveBy(map[string]string) (int, error)
	RemoveAll() (int, error)
	Create(tx *model.Transaction) (*model.Transaction, error)
	// For dev
	FakePersist(n int) error
}

// ITarifManager define what is needed to access
// tarif data whatever database is used
type ITarifManager interface {
	IDataManager
	FindAll() ([]*model.Tarif, error)
	FindBy(map[string]string) ([]*model.Tarif, error)
	FindOneBy(map[string]string) (*model.Tarif, error)
	RemoveBy(map[string]string) (int, error)
	RemoveAll() (int, error)
	Create(tx *model.Tarif) (*model.Tarif, error)
}
