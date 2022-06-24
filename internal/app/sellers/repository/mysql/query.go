package mysql

const (
	getAll = "SELECT id, cid, company_name, address, telephone FROM sellers"
	getOne = "SELECT id, cid, company_name, address, telephone FROM sellers WHERE id = ?"
	create = "INSERT INTO sellers(id, cid, company_name, address, telephone) VALUES (?, ?, ?, ?, ?)"
	delete = "DELETE FROM sellers WHERE id = ?"
	update = "UPDATE FROM sellers SET cid = ?, company_name = ?, address = ?, telephone = ? WHERE id = ?"
)
