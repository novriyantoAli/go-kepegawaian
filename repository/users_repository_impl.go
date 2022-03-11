package repository

import (
	"context"
	"database/sql"

	"github.com/casbin/casbin/v2"
	"github.com/novriyantoAli/go-kepegawaian/entity"
	"github.com/sirupsen/logrus"
)

/**
* Insert(manager entity.Manager)
*
* FindAll() (managers []entity.Manager)
*
* DeleteAll()
 */
type usersRepositoryImpl struct {
	DB      *sql.DB
	Enforce *casbin.Enforcer
}

func NewUsersRepository(database *sql.DB, casen *casbin.Enforcer) UsersRepository {
	return &usersRepositoryImpl{DB: database, Enforce: casen}
}

func (repository *usersRepositoryImpl) query(id *string) (users []entity.Users) {
	sql := "SELECT * FROM users "
	if id != nil {
		sql += " WHERE id = ? "
	}

	sql += " ORDER BY nama_lengkap DESC "

	rows, err := repository.DB.Query(sql)
	// Exit if the SQL doesn't work for some reason
	if err != nil {
		logrus.Panic(err)
	}

	// make sure to cleanup when the program exits
	defer rows.Close()

	for rows.Next() {
		en := entity.Users{}
		err2 := rows.Scan(
			&en.Id,
			&en.Username,
			&en.Password,
			&en.Role,
			&en.NamaLengkap,
			&en.NoTelp,
			&en.Email,
			&en.CreatedAt,
			&en.UpdatedAt,
		)
		// Exit if we get an error
		if err2 != nil {
			logrus.Panic(err2)
		}
		users = append(users, en)
	}

	return
}

func (repository *usersRepositoryImpl) Insert(users entity.Users) (err error) {
	// Create a new context, and begin a transaction
	ctx := context.Background()
	tx, err := repository.DB.BeginTx(ctx, nil)
	if err != nil {
		logrus.Panic(err)
	}

	// Here, the query is executed on the transaction instance, and not applied to the database yet
	sql := "INSERT INTO users(id, username, password, nama_lengkap, email, no_telp, role, created_at, updated_at) VALUES(?,?,?,?,?,?,?,?,?)"
	_, err = tx.ExecContext(ctx, sql,
		users.Id, users.Username, users.Password, users.NamaLengkap, users.Email, users.NoTelp, users.Role, users.CreatedAt, users.UpdatedAt,
	)
	if err != nil {
		// Incase we find any error in the query execution, rollback the transaction
		tx.Rollback()

		logrus.Panic(err)
	}

	ok, err := repository.Enforce.AddGroupingPolicy(users.Id, users.Role)
	if err != nil {
		tx.Rollback()

		logrus.Panic(err)
	}

	if !ok {
		logrus.Warning("add policy to group fail")
	}

	// Finally, if no errors are recieved from the queries, commit the transaction
	// this applies the above changes to our database
	err = tx.Commit()
	if err != nil {
		repository.Enforce.RemoveGroupingPolicy(users.Id, users.Role)

		logrus.Panic(err)
	}

	logrus.Info("commit all data ...")

	return
}

func (repository *usersRepositoryImpl) FindAll() (users []entity.Users) {
	users = repository.query(nil)

	return
}

func (repository *usersRepositoryImpl) Delete(id string) {
	sql := "DELETE FROM users WHERE id = ?"

	// Create a prepared SQL statement
	stmt, err := repository.DB.Prepare(sql)
	// Exit if we get an error
	if err != nil {
		logrus.Panic(err)
	}

	// Make sure to cleanup after the program exits
	defer stmt.Close()

	// Replace the '?' in our prepared statement with 'name'
	_, err = stmt.Exec(id)
	// Exit if we get an error
	if err != nil {
		logrus.Panic(err)
	}
}
