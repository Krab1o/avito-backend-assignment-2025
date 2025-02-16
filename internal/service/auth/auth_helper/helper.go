package authhelper

type helper struct {}

//This interface was created basically for ease of testing

type AuthHelper interface {
	GenerateJWT(userID int64, jwtSecret []byte, jwtTimeout int) (string, error)
	VerifyPassword(hashedPassword string, candidatePassword string) bool
}

func NewHelper() AuthHelper {
	return &helper{}
}