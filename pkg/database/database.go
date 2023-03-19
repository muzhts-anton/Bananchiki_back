package database

import (
	"banana/pkg/utils/log"
	
	"context"

	pgx "github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type DBbyterow [][]byte

type ConnectionPool interface {
	Begin(context.Context) (pgx.Tx, error)
	Close()
}

type DBManager struct {
	Pool ConnectionPool
}

func InitDatabase() *DBManager {
	return &DBManager{
		Pool: nil,
	}
}

func (dbm *DBManager) Connect() {
	connString := "user=bananchiki" +
		" password=1234" +
		" host=localhost" +
		" port=5432" +
		" dbname=tpfinal"

	pool, err := pgxpool.Connect(context.Background(), connString)
	if err != nil {
		log.Warn("{Connect} Postgres error")
		log.Error(err)
		return
	}

	err = pool.Ping(context.Background())
	if err != nil {
		log.Warn("{Connect} Ping error")
		log.Error(err)
		return
	}

	log.Info("Successful connection to postgres")
	log.Info("Connection params: " + connString)
	dbm.Pool = pool
}

func (dbm *DBManager) Disconnect() {
	dbm.Pool.Close()
	log.Info("Postgres disconnected")
}

func (dbm *DBManager) Query(queryString string, params ...interface{}) ([]DBbyterow, error) {
	transactionContext := context.Background()
	tx, err := dbm.Pool.Begin(transactionContext)
	if err != nil {
		log.Warn("{Query} Error connecting to a pool")
		log.Error(err)
		return nil, err
	}

	defer func() {
		err := tx.Rollback(transactionContext)
		if err != nil {
			log.Error(err)
		}
	}()

	rows, err := tx.Query(transactionContext, queryString, params...)
	if err != nil {
		log.Warn("{Query} Error in query: " + queryString)
		log.Error(err)
		return nil, err
	}
	defer rows.Close()

	result := make([]DBbyterow, 0)
	for rows.Next() {
		rowBuffer := make(DBbyterow, 0)
		rowBuffer = append(rowBuffer, rows.RawValues()...)
		result = append(result, rowBuffer)
	}

	err = tx.Commit(transactionContext)
	if err != nil {
		log.Warn("{Query} Error committing")
		log.Error(err)
		return nil, err
	}

	return result, nil
}

func (dbm *DBManager) Execute(queryString string, params ...interface{}) error {
	transactionContext := context.Background()
	tx, err := dbm.Pool.Begin(transactionContext)
	if err != nil {
		log.Warn("{Execute} Error connecting to a pool")
		log.Error(err)
		return err
	}

	defer func() {
		err := tx.Rollback(transactionContext)
		if err != nil {
			log.Error(err)
		}
	}()

	_, err = tx.Exec(transactionContext, queryString, params...)
	if err != nil {
		log.Warn("{Execute} Error in query: " + queryString)
		log.Error(err)
		return err
	}

	err = tx.Commit(transactionContext)
	if err != nil {
		log.Warn("{Execute} Error committing")
		log.Error(err)
		return err
	}

	return nil
}
