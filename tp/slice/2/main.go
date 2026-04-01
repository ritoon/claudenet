package main

import (
	"errors"
	"fmt"
)

var (
	ErrUserNotFound = errors.New("user not found")
	ErrIDOutOfRange = errors.New("id is out of range")
)

var id uint64

type User struct {
	ID       uint64
	FistName string
	LastName string
}

type Users []User

type UserList struct {
	us []User
}

func (us *Users) Add(u *User) error {
	if u == nil {
		return ErrUserNotFound
	}
	id++
	if id == 0 {
		return ErrIDOutOfRange
	}
	// implémenter l'ajout dans us
	u.ID = id
	*us = append(*us, *u)
	return nil
}

func (us *Users) GetByID(id uint64) (*User, error) {
	// implémenter le get by id.
	// for _, v := range *us {
	// 	if v.ID == id {
	// 		return &v, nil
	// 	}
	// }

	for k := range *us {
		if []User(*us)[k].ID == id {
			return &(*us)[k], nil
		}
	}
	return nil, ErrUserNotFound
}

func (us *Users) UpdateByID(id uint64, u *User) error {
	// implémenter l'update
	for k, v := range *us {
		if v.ID == id {
			u.ID = id
			(*us)[k] = *u
			return nil
		}
	}
	return ErrUserNotFound
}

func (us *Users) DeleteByID(id uint64) error {
	// implémenter le delete
	for k, v := range *us {
		if v.ID == id {
			//*us = slices.Delete(*us, k, k+1)
			*us = append((*us)[:k], (*us)[k+1:]...)
			return nil
		}
	}
	return ErrUserNotFound
}

func main() {

	u2 := User{FistName: "Mopheus", LastName: "Pull"}

	var us Users

	err := us.Add(&User{FistName: "Rob", LastName: "Pike"})
	if err != nil {
		panic(err)
	}
	err = us.Add(&u2)
	if err != nil {
		panic(err)
	}

	fmt.Println(us)

	err = us.DeleteByID(2)
	if err != nil {
		panic(err)
	}
	us.UpdateByID(1, &u2)

	fmt.Println(us)
}
