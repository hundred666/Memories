package dao

import "model"

func AddMCR(m model.MCR) (error) {
	insert := "INSERT INTO MCR(Mid,Cid,Time) VALUES(?,?,?)"
	_, err := Modify(insert, m.Mid, m.Cid, m.Time)
	return err
}

func GetMCRS(m model.Move) []model.MCR {
	query := "SELECT * FROM MCR WHERE Mid=?"
	results := Get(query, &model.MCR{}, m.Id)
	if len(results) == 0 {
		return nil
	}
	mcrs := make([]model.MCR, 0)
	for _, res := range results {
		v, ok := res.(model.MCR)
		if ok {
			mcrs = append(mcrs, v)
		}
	}
	return mcrs
}

func DelMCR(m model.MCR) (int, error) {
	del := "DELETE FROM MCR WHERE  Cid=?"
	return Modify(del, m.Cid)
}
