package main

import (
	"errors"
	"fmt"

	uuidv4 "github.com/google/uuid"
)

var (
	ErrUserNotFound = errors.New("user not found")
)

type User struct {
	UUID     string
	FistName string
	LastName string
}

type Users map[string]*User

func (us *Users) Add(u *User) error {
	if u == nil {
		return fmt.Errorf("no user to add")
	}
	u.UUID = uuidv4.NewString()

	(*us)[u.UUID] = u
	// implémenter l'ajout dans us
	return nil
}

func (us *Users) GetByUUID(uuid string) (*User, error) {
	err := uuidv4.Validate(uuid)
	if err != nil {
		return nil, err
	}
	// implémenter le get by UUID.
	u, ok := (*us)[uuid]
	if !ok {
		return nil, ErrUserNotFound
	}
	return u, nil
}

func (us *Users) UpdateByUUID(uuid string, u *User) error {
	err := uuidv4.Validate(uuid)
	if err != nil {
		return err
	}
	// implémenter l'update
	uOld, err := us.GetByUUID(uuid)
	if err != nil {
		return err
	}
	uOld.FistName = u.FistName
	uOld.LastName = u.LastName
	return nil

}

func (us *Users) DeleteByUUID(uuid string) error {
	err := uuidv4.Validate(uuid)
	if err != nil {
		return err
	}
	// implémenter le delete
	_, ok := (*us)[uuid]
	if !ok {
		return ErrUserNotFound
	}
	delete((*us), uuid)
	return nil
}

func main() {
	us := make(Users)

	u := User{
		FistName: "Rob",
		LastName: "Pike",
	}

	u2 := User{
		FistName: "Yannick",
		LastName: "Rinnebach",
	}
	err := us.Add(&u)
	if err != nil {
		panic(err)
	}

	err = us.Add(&u2)
	if err != nil {
		panic(err)
	}
	fmt.Println(us)

	err = us.UpdateByUUID(u.UUID, &u2)
	if err != nil {
		panic(err)
	}

	err = us.DeleteByUUID(u2.UUID)
	if err != nil {
		panic(err)
	}

	fmt.Println(us)
}
