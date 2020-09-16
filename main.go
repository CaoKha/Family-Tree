package main

import (
	//"example.go/go/src/Family-Tree/model"
	model "./model/src"
)

func exemple() {
	pm := model.PeopleManager{}

	p1 := pm.AddNewPerson("Phuoc Khanh", "Le", model.Male)
	p3 := pm.AddNewPerson("Kim Thai", "Le", model.Male)
	p2 := pm.AddNewPerson("Phuoc Long", "Le", model.Male)
	p4 := pm.AddNewPerson("Lien Khuong", "Tran Dinh", model.Female)
	/*
		p5 := pm.AddNewPerson("Khoi", "Tran Dinh", 'm')
		p6 := pm.AddNewPerson("Diem Chi", "Tran Dinh", 'f')
		p7 := pm.AddNewPerson("Khoi Quoc", "Tran Dinh", 'm')
		p8 := pm.AddNewPerson("Khoi Nguyen", "Tran Dinh", 'm')
		p9 := pm.AddNewPerson("Diem Phuc", "Tran Dinh", 'm')
		p10 := pm.AddNewPerson("Giang", "Tran Dinh", 'm')
		p11 := pm.AddNewPerson("Suong", "Nguyen Thi", 'f')
		p12 := pm.AddNewPerson("Ha My", "Le Phuoc", 'f')
		p13 := pm.AddNewPerson("Ba ngoai", "Nguyen", 'f')
		p14 := pm.AddNewPerson("Khuong", "Le Phuoc", 'm')

		p15 := pm.AddNewPerson("Quynh", "Tran Dinh", 'm')
		p16 := pm.AddNewPerson("Mai", "Tran Dinh", 'f')
		p17 := pm.AddNewPerson("Mien", "Tran Dinh", 'f')

		p18 := pm.AddNewPerson("Thanh Trung", "Dinh ", 'm')
		p19 := pm.AddNewPerson("Thao Nhi", "Dung", 'f')
		p20 := pm.AddNewPerson("Thanh Phuong", "Dinh", 'm')
		p21 := pm.AddNewPerson("Ong noi vis", "Dinh", 'm')
		p22 := pm.AddNewPerson("Thanh Viet", "Dinh", 'm')
		p23 := pm.AddNewPerson("Uyen My", "Dinh", 'f')

		p24 := pm.AddNewPerson("Ba co", "Nguyen", 'f')
	*/
	pm.AddNewRelation(p1.ID, p3.ID, model.ChildRole)
	pm.AddNewRelation(p3.ID, p2.ID, model.ParentRole)
	pm.AddNewRelation(p3.ID, p4.ID, model.SpouseRole)
	/*
		p1.AddDad(p3)
		p3.AddChildren(p2)
		p3.AddSpouse(p4)
		p4.AddDad(p5)
		p2.AddSpouse(p11)
		p2.AddChildren(p12)
		p5.AddChildren(p6)
		p5.AddChildren(p7)
		p5.AddChildren(p8)
		p5.AddChildren(p9)
		p5.AddChildren(p10)
		p5.AddSpouse(p13)
		p3.AddDad(p14)
		p5.AddMom(p24)
		p24.AddChildren(p15)
		p15.AddChildren(p16)
		p15.AddChildren(p17)
		p6.AddChildren(p18)
		p6.AddChildren(p19)
		p6.AddSpouse(p20)
		p20.AddDad(p21)
		p21.AddChildren(p22)
		p22.AddChildren(p23)
	*/
	pm.WriteToCSV()
}

func main() {
	//run exemple() to create file csv with people
	exemple()
	/*
		pm := &(model.PM)

		//verify dad and mom of all people
		pm.Read_CSV()
		for i := 0; i < pm.GetNbr(); i++ {
			if pm.AllPeople[i].Dad != nil {
				fmt.Println("dad of ", pm.AllPeople[i].Ten, " is ", pm.AllPeople[i].Dad.Ten)
			}
			if pm.AllPeople[i].Mom != nil {
				fmt.Println("mom of ", pm.AllPeople[i].Ten, " is ", pm.AllPeople[i].Mom.Ten)
			}
		}

		//verify all tree's root and its name
		model.BuildTrees("people.csv")
		for i := 0; i < len(model.TM); i++ {
			fmt.Println(model.TM[i].Filename_json)
		}

		//make file json
		model.TM[0].WriteToJson()

		roots_of_1 := model.GetRoot(pm.AllPeople[1])
		fmt.Println(roots_of_1)

		//check func SameRootByRank
		//
		j := 1
		for i := 1; i < 25; i++ {
			if i != j {
				fmt.Println(pm.GetPeopleByID(j).Ten, "is same root with ", pm.GetPeopleByID(i).Ten, "by rank ", model.RankOfSameRoot(pm.GetPeopleByID(j), pm.GetPeopleByID(i)))
			}
		}

		for i := 1; i < 25; i++ {
			if i != j {
				fmt.Println(pm.GetPeopleByID(j).Ten, "is distance with ", pm.GetPeopleByID(i).Ten, "by ", model.Distance(pm.GetPeopleByID(j), pm.GetPeopleByID(i)))
			}
		}
		//check func Get_people_in_view
		tmp := model.Get_people_in_view(pm.GetPeopleByID(j))
		fmt.Println("People in view of ", pm.GetPeopleByID(j).Ten)
		for i := 0; i < tmp.GetNbr(); i++ {
			fmt.Println(tmp.AllPeople[i].Ten)
		}
	*/
}
