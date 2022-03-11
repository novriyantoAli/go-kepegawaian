package repository

import (
	"context"
	"database/sql"

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
type pegawaiRepositoryImpl struct {
	DB *sql.DB
}

func NewPegawaiRepository(database *sql.DB) PegawaiRepository {
	return &pegawaiRepositoryImpl{DB: database}
}

func (repository *pegawaiRepositoryImpl) query(id *string) (pegawai []entity.Pegawai) {
	sql := "SELECT * FROM pegawai "
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
		en := entity.Pegawai{}
		err2 := rows.Scan(
			&en.Id,
			&en.NamaLengkap,
			&en.Tmt,
			&en.CreatedAt,
			&en.UpdatedAt,
		)
		// Exit if we get an error
		if err2 != nil {
			logrus.Panic(err2)
		}
		pegawai = append(pegawai, en)
	}

	return
}

func (repository *pegawaiRepositoryImpl) Insert(pegawai entity.Pegawai) (err error) {
	// Create a new context, and begin a transaction
	ctx := context.Background()
	tx, err := repository.DB.BeginTx(ctx, nil)
	if err != nil {
		logrus.Panic(err)
	}

	// Here, the query is executed on the transaction instance, and not applied to the database yet
	sql := "INSERT INTO users(id, nama_lengkap, tmt, created_at, updated_at) VALUES(?,?,?,?,?)"
	_, err = tx.ExecContext(ctx, sql,
		pegawai.Id, pegawai.NamaLengkap, pegawai.Tmt, pegawai.CreatedAt, pegawai.UpdatedAt,
	)
	if err != nil {
		// Incase we find any error in the query execution, rollback the transaction
		tx.Rollback()

		logrus.Panic(err)
	}

	// Finally, if no errors are recieved from the queries, commit the transaction
	// this applies the above changes to our database
	err = tx.Commit()
	if err != nil {
		logrus.Panic(err)
	}

	logrus.Info("commit all data ...")

	return
}

func (repository *pegawaiRepositoryImpl) FindAll() (pegawais []entity.Pegawai) {
	pegawais = repository.query(nil)

	return
}

func (repository *pegawaiRepositoryImpl) Delete(id string) {
	sql := "DELETE FROM pegawai WHERE id = ?"

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
