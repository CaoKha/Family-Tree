package model

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

//------------------------ from model to database --------------------

// GetIdByInfo return id of a person by Lastname, FirstName and Birthday which are unique by person
func GetIdByInfo(FirstName string, LastName string, Birthday time.Time) int {
	id_row, err := selectIDByInfo.Query(FirstName, LastName, TimeToString(Birthday)[0:10])
	if err != nil {
		log.Println("query error", err)
		return 0
	}

	defer id_row.Close()

	if !id_row.Next() {
		return 0
	}

	var id int
	err = id_row.Scan(&id)
	if err != nil {
		log.Println("scan error", err)
		return 0
	}
	return id
}

func GetIdByInfo_(p Person) int {
	return GetIdByInfo(p.FirstName, p.LastName, p.Birthday)
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

func GetNumberPerson() int {
	rows, _ := selectNumberPerson.Query()

	defer rows.Close()

	if !rows.Next() {
		return -1
	}

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
	_, err := updateFatherTree.Exec(ID_tree, ID_person)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func SetMotherTree(ID_person int, ID_tree int) error {
	_, err := updateMotherTree.Exec(ID_tree, ID_person)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

//func createFirstTree() error {
//	return nil
//}

func GetIdFatherTree(id_person int) (int, error) {
	rows, err := selectFatherTreePerson.Query(id_person)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	if !rows.Next() {
		return -1, nil
	}

	var res int
	err = rows.Scan(&res)

	return res, nil
}

func GetIdMotherTree(id_person int) (int, error) {
	rows, err := selectMotherTreePerson.Query(id_person)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	if !rows.Next() {
		return -1, nil
	}

	var res int
	err = rows.Scan(&res)

	return res, nil
}

func UpdateTreeRoot(id_tree int, id_root int) error {
	_, err := updateTreeRootID.Exec(id_root, id_tree)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func GetIdTreeByRoot(id_root int) int {
	rows, err := selectIdTreeByRoot.Query(id_root)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	if !rows.Next() {
		return -1
	}

	var res int
	err = rows.Scan(&res)

	return res
}

func GetRootByIdTree(id_tree int) int {
	rows, err := selectRootByIdTree.Query(id_tree)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	if !rows.Next() {
		return -1
	}

	var res int
	err = rows.Scan(&res)

	return res
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
