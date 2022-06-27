package mysql

const (
	getAll       = "SELECT id, cid, company_name, address, telephone FROM sellers"
	getOne       = "SELECT id, cid, company_name, address, telephone FROM sellers WHERE id = ?"
	create       = "INSERT INTO sellers(id, cid, company_name, address, telephone) VALUES (?, ?, ?, ?, ?)"
	deleteSeller = "DELETE FROM sellers WHERE id = ?"
	update       = "UPDATE sellers SET cid = ?, company_name = ?, address = ?, telephone = ? WHERE id = ?"
)
