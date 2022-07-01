package mysql

const (
	getAllSellers = "SELECT id, cid, company_name, address, telephone FROM sellers"
	getOneSeller  = "SELECT id, cid, company_name, address, telephone FROM sellers WHERE id = ?"
	createSeller  = "INSERT INTO sellers(id, cid, company_name, address, telephone) VALUES (?, ?, ?, ?, ?)"
	deleteSeller  = "DELETE FROM sellers WHERE id = ?"
	updateSeller  = "UPDATE sellers SET cid = ?, company_name = ?, address = ?, telephone = ? WHERE id = ?"

	getCountryID  = "SELECT id FROM countries WHERE country_name = ?"
	createCountry = "INSERT INTO countries(id, country_name) VALUES (?, ?)"

	getProvinceID  = "SELECT id FROM provinces WHERE province_name = ? AND country_id = ?"
	createProvince = "INSERT INTO provinces(id, province_name, country_id) VALUES (?, ?, ?)"

	checkLocalityExists = "SELECT id FROM localities WHERE id = ?"
	createLocality      = "INSERT INTO localities(id, locality_name, province_id) VALUES(?, ?, ?)"
)
