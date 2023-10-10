package repositories

import (
	"control-api/model"
	"database/sql"
	"fmt"
	"log"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (repo *UserRepository) CheckTable() {
	_, err := repo.db.Exec("create table if not exists test_go_sql(id int primary key, number char(9), name char(10)")
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("Table created !")
}

func (repo *UserRepository) Create(id int, number string, name string) int {
	statement, err := repo.db.Prepare("insert into test_go_sql(id, number, name) values (?, ?, ?)")
	if err != nil {
		log.Fatal(err.Error())
	}
	res, err := statement.Exec(id, number, name)
	if err != nil {
		log.Fatal(err.Error())
	}
	cur_id, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err.Error())
	}
	return int(cur_id)
}

func (repo *UserRepository) Retrieve(id int) (*model.UserRecord, error) {

	rows, err := repo.db.Query("select * from test_go_sql where id = ? ", id)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	var rec model.UserRecord

	for rows.Next() {
		err = rows.Scan(&rec.Id, &rec.Number, &rec.Name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(rec.Id, rec.Number, rec.Name)
		return &rec, nil
	}
	if err != nil {
		log.Fatal(err)
	}

	return nil, err
}

func (repo *UserRepository) Update(id int, number string, name string) int {

	stmt, err := repo.db.Prepare("update test_go_sql set number = ?, name = ? where id = ?")
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec(number, name, id)
	if err != nil {
		log.Fatal(err)
	}
	num, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	return int(num)

}

func (repo *UserRepository) Delete(id int) int {
	stmt, err := repo.db.Prepare("delete from test_go_sql where id = ?")
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec(id)
	if err != nil {
		log.Fatal(err)
	}
	num, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	return int(num)
}
