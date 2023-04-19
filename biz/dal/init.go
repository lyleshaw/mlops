package dal

import (
	"github.com/lyleshaw/mlops/biz/dal/mysql"
	"github.com/lyleshaw/mlops/biz/model/query"
)

func Init() {
	mysql.Init()
	query.SetDefault(mysql.DB)
}
