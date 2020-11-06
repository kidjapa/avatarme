package user

import (
    "crypto/md5"
    "fmt"
    "io"
)

// Create user struct for hash e-mail/id/etc.. for user string
type User struct {
    Str string
    Hash string
}

func New(str string) *User {
    u := &User{Str: str}
    u.Hash = u._getStringMD5()
    return u
}

func(u *User) _getStringMD5() string{
    Hash := md5.New()
    _, _ = io.WriteString(Hash, u.Str)
    return fmt.Sprintf("%x", Hash.Sum(nil))
}

func (u *User) String() string  {
    return fmt.Sprintf("string used is: %s\ngenerated an image with: %s\n",u.Str, u.Hash)
}