package model

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

//------------------------ from model to database --------------------

func GetIdByName(FirstName string, LastName string) (int, error) {
	id_row, err := selectIDByName.Query(FirstName, LastName)
	if err != nil {
		log.Println("query error", err)
		return 0, err
	}

	defer id_row.Close()

	if !id_row.Next() {
		return 0, nil
	}

	var id int
	err = id_row.Scan(&id)
	if err != nil {
		log.Println("scan error", err)
		return 0, err
	}
	return id, nil
}

func GetAllRoot() ([]int, error) {
	rows, err := selectAllRoot.Query()
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	if !rows.Next() {
		return nil, nil
	}

	var res []int
	for {
		var tmp int
		err = rows.Scan(&tmp)
		res = append(res, tmp)
		if !rows.Next() {
			break
		}
	}
	return res, nil
}

func GetPersonById(ID_person int) (*Person, error) {
	rows, err := selectPersonById.Query(ID_person)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	if !rows.Next() {
		return nil, nil
	}
	var res *Person

	err = rows.Scan(res.ID, res.FirstName, res.LastName,
		res.NickName, res.Gender, res.Rank,
		res.Birthday, res.Deathday)

	return res, nil
}
func GetNumberPerson() int {
	rows, _ := selectNumberPerson.Query()

	defer rows.Close()

	var id int
	_ = rows.Scan(&id)
	return id
}

func isPersonEmpty() bool {
	if GetNumberPerson() == 0 {
		return true
	} else {
		return false
	}
}

func SetFatherTree(ID_person int, ID_tree int) error {
	return nil
}

func SetMotherTree(ID_person int, ID_tree int) error {
	return nil
}

func createFirstTree() error {
	return nil
}

func GetIdFatherTree(id_person int) (int, error) {
	var id_tree int
	return id_tree, nil
}

func GetIdMotherTree(id_person int) (int, error) {
	var id_tree int
	return id_tree, nil
}

func UpdateTreeRoot(id_tree int, id_root int) error {
	return nil
}

func InsertTree(id_root int) (int, error) {
	var id_newest_tree int
	return id_newest_tree, nil
}

func Clear_tables() {
	var clearTablePerson *sql.Stmt
	var clearTableRelation *sql.Stmt

	clearTablePerson, err := db.Prepare("DELETE FROM Person")
	if err != nil {
		log.Fatal(err)
	}

	clearTableRelation, err = db.Prepare("DELETE FROM Relation")
	if err != nil {
		log.Fatal(err)
	}

	clearTableRelation.Exec()
	clearTablePerson.Exec()
}