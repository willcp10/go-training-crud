package date_source

import (
	"errors"
	"sync"
)

type userData struct {
	id        int64
	name      string
	age       int
	docNumber string
}

func UserDataInput(id int64, name string, age int, docNumber string) userData {
	return userData{
		id:        id,
		name:      name,
		age:       age,
		docNumber: docNumber,
	}
}

func UserDataOutput(ud userData) (id int64, name string, age int, docNumber string) {
	return ud.id, ud.name, ud.age, ud.docNumber
}

var _ UserDataSource = new(UserDataSourceImpl)

type UserDataSource interface {
	Select(id int64) (userData, bool)
	SelectAll() []userData
	Insert(data userData) error
	Update(data userData) error
	Delete(id int64) error
}

type UserDataSourceImpl struct {
	users      map[int64]userData
	nextUserID int64
	mu         sync.Mutex
}

func (uds *UserDataSourceImpl) Select(id int64) (userData, bool) {
	data, ok := uds.users[id]
	return data, ok
}

func (uds *UserDataSourceImpl) SelectAll() []userData {
	var userSlice []userData
	for _, u := range uds.users {
		userSlice = append(userSlice, u)
	}

	return userSlice
}

func (uds *UserDataSourceImpl) Insert(data userData) error {
	sucess := uds.mu.TryLock()
	if !sucess {
		return errors.New("resource in use")
	}
	defer uds.mu.Unlock()

	if data.id < 0 {
		return errors.New("invalid id")
	}

	if _, found := uds.users[data.id]; found {
		return errors.New("id already in use")
	}

	if data.id == 0 {
		for {
			if _, found := uds.users[uds.nextUserID]; !found {
				break
			}
			uds.consumeID()
		}
		data.id = uds.nextUserID
		uds.consumeID()
	}

	uds.insert(data)
	return nil
}

func (uds *UserDataSourceImpl) consumeID() {
	uds.nextUserID++
}

func (uds *UserDataSourceImpl) insert(data userData) {
	uds.users[data.id] = data
}

func (uds *UserDataSourceImpl) Update(data userData) error {
	sucess := uds.mu.TryLock()
	if !sucess {
		return errors.New("resource in use")
	}
	defer uds.mu.Unlock()

	if data.id < 0 {
		return errors.New("invalid id")
	}

	if _, found := uds.users[data.id]; !found {
		return errors.New("id not found")
	}

	uds.users[data.id] = data

	return nil
}

func (uds *UserDataSourceImpl) Delete(id int64) error {
	sucess := uds.mu.TryLock()
	if !sucess {
		return errors.New("resource in use")
	}
	defer uds.mu.Unlock()

	_, found := uds.users[id]
	if !found {
		return errors.New("id not found")
	}

	delete(uds.users, id)
	return nil
}

func NewUserDataSource() *UserDataSourceImpl {
	return &UserDataSourceImpl{
		users:      make(map[int64]userData),
		nextUserID: 1,
		mu:         sync.Mutex{},
	}
}
