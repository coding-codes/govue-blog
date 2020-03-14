package model

type User struct {
	ID       int
	Name     string
	Email    string
	Password string
}

func (u *User) Create() error {
	_, err := db.Exec("INSERT INTO users (name, email, password) VALUES (?, ?, ?)",
		u.Name, u.Email, hash(u.Password))
	if err != nil {
		return err
	}
	return nil
}

func (u *User) ReadByID() error {
	return db.QueryRow("SELECT name, email, password FROM users WHERE id = ?", u.ID).
		Scan(&u.Name, &u.Email, &u.Password)
}

func (u *User) ReadByEmail() error {
	return db.QueryRow("SELECT id, name, password FROM users WHERE email = ?", u.Email).
		Scan(&u.ID, &u.Name, &u.Password)
}
