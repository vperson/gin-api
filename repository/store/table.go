package store

import "gin-api/repository/store/tb"

func User() *tb.UserDB {
	return tb.NewUser(db)
}

func Common() *tb.CommonDB {
	return tb.NewCommon(db)
}
