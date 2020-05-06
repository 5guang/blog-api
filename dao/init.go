package dao

import (
	"blog/pkg/setting"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"time"
)

var (
	DB, TX *gorm.DB
)

func Init() {
	var (
		err error
		// 配置mysql数据库信息
		MysqlDsn = fmt.Sprintf(fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
			setting.DataBaseSetting.User,
			setting.DataBaseSetting.Password,
			setting.DataBaseSetting.Host,
			setting.DataBaseSetting.Name))
	)
	fmt.Println("-----------------------",MysqlDsn)
	// 链接mysql数据库
	DB, err = gorm.Open(setting.DataBaseSetting.Type, MysqlDsn)
	if err != nil {
		log.Println("----",err)
	}
	TX = DB.Begin()
	// 设置表名前缀
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return setting.DataBaseSetting.TablePrefix + defaultTableName
	}
	// gorm表名默认就是结构体名称的复数
	// 禁用默认表名的复数形式 详情查看 https://gorm.io/zh_CN/docs/conventions.html
	DB.SingularTable(true)
	DB.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	DB.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)
	DB.Callback().Delete().Replace("gorm:delete", deleteCallback)
	DB.LogMode(true)
	//设置连接池
	//空闲
	DB.DB().SetMaxIdleConns(10)
	//打开
	DB.DB().SetMaxOpenConns(100)
	//超时
	DB.DB().SetConnMaxLifetime(time.Second * 30)
	migration()
}


func CloseDB() {
	defer DB.Close()
}

func Rollback() {
	TX.Rollback()
}

func Commit() error {
	return TX.Commit().Error
}
func migration() {
	// 自动迁移模式
	DB.AutoMigrate(&Tag{}, &Article{}, &User{}, &Role{}, &Permission{})
}

func updateTimeStampForCreateCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		nowTime := time.Now().Unix()
		// scope.FieldByName 通过 scope.Fields() 获取所有字段，判断当前是否包含所需字段
		if createdTimeField, ok := scope.FieldByName("CreatedOn"); ok {
			// field.IsBlank 可判断该字段的值是否为空
			if createdTimeField.IsBlank {
				// 若为空则 field.Set 用于给该字段设置值
				createdTimeField.Set(nowTime)
			}
		}

		//if modifyTimeField, ok := scope.FieldByName("UpdatedOn"); ok {
		//	if modifyTimeField.IsBlank {
		//		modifyTimeField.Set(nowTime)
		//	}
		//}
	}
}

// updateTimeStampForUpdateCallback will set `ModifyTime` when updating
func updateTimeStampForUpdateCallback(scope *gorm.Scope) {
	// scope.Get(...) 根据入参获取设置了字面值的参数， gorm:update_column ，它会去查找含这个字面值的字段属性
	if _, ok := scope.Get("gorm:update_column"); !ok {
		// scope.SetColumn(...) 假设没有指定 update_column 的字段，我们默认在更新回调设置 UpdatedOn 的值
		scope.SetColumn("UpdatedOn", time.Now().Unix())
	}
}

func deleteCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		var extraOption string
		if str, ok := scope.Get("gorm:delete_option"); ok {
			extraOption = fmt.Sprint(str)
		}

		deletedOnField, hasDeletedOnField := scope.FieldByName("DeletedOn")

		if !scope.Search.Unscoped && hasDeletedOnField {
			scope.Raw(fmt.Sprintf(
				"UPDATE %v SET %v=%v%v%v",
				scope.QuotedTableName(),
				scope.Quote(deletedOnField.DBName),
				scope.AddToVars(time.Now().Unix()),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		} else {
			scope.Raw(fmt.Sprintf(
				"DELETE FROM %v%v%v",
				scope.QuotedTableName(),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		}
	}
}

func addExtraSpaceIfExist(str string) string {
	if str != "" {
		return " " + str
	}
	return ""
}
