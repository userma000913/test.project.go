package dao

func CreateTestDemo() error {
	sql := "insert into user(id,name) values(1,'tom')"
	return Db.Exec(sql).Error
}
