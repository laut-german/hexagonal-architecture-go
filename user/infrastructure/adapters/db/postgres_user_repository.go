package db

import (
	"database/sql"
	"hexagonal-architecture-go/user/domain/entities"
)


type PostgresUserRepository struct {
	db *sql.DB
}


func NewPotgresUserRepository(db *sql.DB) *PostgresUserRepository {
	return &PostgresUserRepository{db: db}
}


func (repository *PostgresUserRepository) Save(user entities.User) (*entities.User, error) {
	query := `INSERT INTO users (name, email, created_at) VALUES ($1, $2, $3) RETURNING id`
	 err := repository.db.QueryRow(query, user.Name, user.Email, user.Created_At).Scan(&user.ID)
	 if err !=nil {
		return nil, err
	 }
	return &user, nil
}

func (repository *PostgresUserRepository) Update(user entities.User) error {
	query := `UPDATE users SET name = $1, email = $2 WHERE id = $3`
	_, err := repository.db.Exec(query, user.Name, user.Email)
	return err
}

func (repository *PostgresUserRepository) FindByID(id string) (*entities.User, error) {
	query := `SELECT id, name, email FROM  users  WHERE id = $1`
	row :=  repository.db.QueryRow(query, id)
	var user entities.User

	if err := row.Scan(&user.ID, &user.Name, &user.Email); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}


func (repository *PostgresUserRepository) List() ([]entities.User, error) {
    query := `SELECT id, name, email FROM users`
    rows, err := repository.db.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var users []entities.User
    for rows.Next() {
        var user entities.User
        if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
            return nil, err
        }
        users = append(users, user)
    }
    return users, nil
}