package model

import (
	"errors"
	"time"
)

// Person is special type for each person
type Person struct {
	ID        int
	FirstName string
	LastName  string
	Gender    GenderType
	Birthday  time.Time
	Deathday  time.Time
	Rank      int
	Dad       *Person
	Mom       *Person
	Spouse    []*Person
	Children  []*Person
}

// GetDadID returns Id of the father of person p
func (p Person) GetDadID() int {
	if p.Dad == nil {
		return IDNotFound
	}
	return p.Dad.ID
}

// GetMomID returns Id of the mother of person p
func (p Person) GetMomID() int {
	if p.Mom == nil {
		return IDNotFound
	}
	return p.Mom.ID
}

// GetAge returns the current age of person p
func (p Person) GetAge() int {
	return time.Now().Year() - p.Birthday.Year()
}

// UpdateRank updates the rank of person p by the rank from another person and his role to p
func (p *Person) UpdateRank(fromPerson *Person, role Role) {
	switch role {
	case ParentRole:
		p.Rank = fromPerson.Rank + 1
	case SpouseRole:
		p.Rank = fromPerson.Rank
	case ChildRole:
		p.Rank = fromPerson.Rank - 1
	}
}

// AddParent add dad/mom to the current person depending on the gender of added parent
func (p *Person) AddParent(parent *Person) error {
	if parent.Gender == Male {
		p.Dad = parent
	} else if parent.Gender == Female {
		p.Mom = parent
	} else {
		return errors.New("parent's gender is undefined")
	}
	return nil
}

// AddSpouse add new spouse to list of spouse
func (p *Person) AddSpouse(s *Person) {
	if PersonAlreadyInList(s, p.Spouse) == false {
		p.Spouse = append(p.Spouse, s)
	}
}

//AddChildren add new child to list of children
func (p *Person) AddChildren(c *Person) {
	if PersonAlreadyInList(c, p.Children) == false {
		p.Children = append(p.Children, c)
	}
}

// IsRoot returns whether a person is the root of a whole family (tree)
func (p Person) IsRoot() bool {
	return p.Dad == nil && p.Mom == nil
}

// PersonJSONForm is a convenient way to form the family trees
type PersonJSONForm struct {
	ID         int
	IDChildren []*PersonJSONForm
}

// ToJSONForm extracts useful information for PersonJSONForm from Person
func (p Person) ToJSONForm() *PersonJSONForm {
	res := PersonJSONForm{ID: p.ID}
	for _, c := range p.Children {
		tmp := c.ToJSONForm()
		res.IDChildren = append(res.IDChildren, tmp)
	}
	return &res
}
