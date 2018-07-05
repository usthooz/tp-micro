// Code generated by 'micro gen' command.
// DO NOT EDIT!

package model

import (
	"database/sql"
	"unsafe"

	"github.com/henrylee2cn/goutil/coarsetime"
	tp "github.com/henrylee2cn/teleport"
	"github.com/xiaoenai/tp-micro/model/mysql"
	"github.com/xiaoenai/tp-micro/model/sqlx"

	"github.com/xiaoenai/tp-micro/examples/project/args"
)

// Log comment...
type Log args.Log

// ToLog converts to *Log type.
func ToLog(_l *args.Log) *Log {
	return (*Log)(unsafe.Pointer(_l))
}

// ToArgsLog converts to *args.Log type.
func ToArgsLog(_l *Log) *args.Log {
	return (*args.Log)(unsafe.Pointer(_l))
}

// ToLogSlice converts to []*Log type.
func ToLogSlice(a []*args.Log) []*Log {
	return *(*[]*Log)(unsafe.Pointer(&a))
}

// ToArgsLogSlice converts to []*args.Log type.
func ToArgsLogSlice(a []*Log) []*args.Log {
	return *(*[]*args.Log)(unsafe.Pointer(&a))
}

// TableName implements 'github.com/xiaoenai/tp-micro/model'.Cacheable
func (*Log) TableName() string {
	return "log"
}

func (_l *Log) isZeroPrimaryKey() bool {
	var _id int64
	if _l.Id != _id {
		return false
	}
	return true
}

var logDB, _ = mysqlHandler.RegCacheableDB(new(Log), cacheExpire, args.LogSql)

// GetLogDB returns the Log DB handler.
func GetLogDB() *mysql.CacheableDB {
	return logDB
}

// InsertLog insert a Log data into database.
// NOTE:
//  Primary key: 'id';
//  Without cache layer.
func InsertLog(_l *Log, tx ...*sqlx.Tx) (int64, error) {
	_l.UpdatedAt = coarsetime.FloorTimeNow().Unix()
	if _l.CreatedAt == 0 {
		_l.CreatedAt = _l.UpdatedAt
	}
	return _l.Id, logDB.Callback(func(tx sqlx.DbOrTx) error {
		var (
			query            string
			isZeroPrimaryKey = _l.isZeroPrimaryKey()
		)
		if isZeroPrimaryKey {
			query = "INSERT INTO `log` (`text`,`updated_at`,`created_at`)VALUES(:text,:updated_at,:created_at);"
		} else {
			query = "INSERT INTO `log` (`id`,`text`,`updated_at`,`created_at`)VALUES(:id,:text,:updated_at,:created_at);"
		}
		r, err := tx.NamedExec(query, _l)
		if isZeroPrimaryKey {
			_l.Id, err = r.LastInsertId()
		}
		return err
	}, tx...)
}

// UpsertLog insert or update the Log data by primary key.
// NOTE:
//  Primary key: 'id';
//  With cache layer;
//  Insert data if the primary key is specified;
//  Update data based on _updateFields if no primary key is specified;
//  _updateFields' members must be db field style (snake format);
//  Automatic update 'updated_at' field;
//  Don't update the primary keys, 'created_at' key and 'deleted_ts' key;
//  Update all fields except the primary keys, 'created_at' key and 'deleted_ts' key, if _updateFields is empty.
func UpsertLog(_l *Log, _updateFields []string, tx ...*sqlx.Tx) (int64, error) {
	if _l.UpdatedAt == 0 {
		_l.UpdatedAt = coarsetime.FloorTimeNow().Unix()
	}
	if _l.CreatedAt == 0 {
		_l.CreatedAt = _l.UpdatedAt
	}
	return _l.Id, logDB.Callback(func(tx sqlx.DbOrTx) error {
		var (
			query            string
			isZeroPrimaryKey = _l.isZeroPrimaryKey()
		)
		if isZeroPrimaryKey {
			query = "INSERT INTO `log` (`text`,`updated_at`,`created_at`)VALUES(:text,:updated_at,:created_at)"
		} else {
			query = "INSERT INTO `log` (`id`,`text`,`updated_at`,`created_at`)VALUES(:id,:text,:updated_at,:created_at)"
		}
		query += " ON DUPLICATE KEY UPDATE "
		if len(_updateFields) == 0 {
			query += "`text`=VALUES(`text`),`updated_at`=VALUES(`updated_at`);"
		} else {
			for _, s := range _updateFields {
				if s == "updated_at" || s == "created_at" || s == "deleted_ts" || s == "id" {
					continue
				}
				query += "`" + s + "`=VALUES(`" + s + "`),"
			}
			if query[len(query)-1] != ',' {
				return nil
			}
			query += "`updated_at`=VALUES(`updated_at`),`deleted_ts`=0;"
		}
		r, err := tx.NamedExec(query, _l)
		if isZeroPrimaryKey {
			rowsAffected, err := r.RowsAffected()
			if err == nil && rowsAffected == 1 {
				_l.Id, err = r.LastInsertId()
			}
		}
		return err
	}, tx...)
}

// UpdateLogByPrimary update the Log data in database by primary key.
// NOTE:
//  Primary key: 'id';
//  With cache layer;
//  _updateFields' members must be db field style (snake format);
//  Automatic update 'updated_at' field;
//  Don't update the primary keys, 'created_at' key and 'deleted_ts' key;
//  Update all fields except the primary keys, 'created_at' key and 'deleted_ts' key, if _updateFields is empty.
func UpdateLogByPrimary(_l *Log, _updateFields []string, tx ...*sqlx.Tx) error {
	_l.UpdatedAt = coarsetime.FloorTimeNow().Unix()
	err := logDB.Callback(func(tx sqlx.DbOrTx) error {
		query := "UPDATE `log` SET "
		if len(_updateFields) == 0 {
			query += "`text`=:text,`updated_at`=:updated_at WHERE `id`=:id AND `deleted_ts`=0 LIMIT 1;"
		} else {
			for _, s := range _updateFields {
				if s == "updated_at" || s == "created_at" || s == "deleted_ts" || s == "id" {
					continue
				}
				query += "`" + s + "`=:" + s + ","
			}
			if query[len(query)-1] != ',' {
				return nil
			}
			query += "`updated_at`=:updated_at WHERE `id`=:id AND `deleted_ts`=0 LIMIT 1;"
		}
		_, err := tx.NamedExec(query, _l)
		return err
	}, tx...)
	if err != nil {
		return err
	}
	err = logDB.DeleteCache(_l)
	if err != nil {
		tp.Errorf("%s", err.Error())
	}
	return nil
}

// DeleteLogByPrimary delete a Log data in database by primary key.
// NOTE:
//  Primary key: 'id';
//  With cache layer.
func DeleteLogByPrimary(_id int64, deleteHard bool, tx ...*sqlx.Tx) error {
	var err error
	if deleteHard {
		// Immediately delete from the hard disk.
		err = logDB.Callback(func(tx sqlx.DbOrTx) error {
			_, err := tx.Exec("DELETE FROM `log` WHERE `id`=? AND `deleted_ts`=0;", _id)
			return err
		}, tx...)

	} else {
		// Delay delete from the hard disk.
		ts := coarsetime.FloorTimeNow().Unix()
		err = logDB.Callback(func(tx sqlx.DbOrTx) error {
			_, err := tx.Exec("UPDATE `log` SET `updated_at`=?, `deleted_ts`=? WHERE `id`=? AND `deleted_ts`=0;", ts, ts, _id)
			return err
		}, tx...)
	}

	if err != nil {
		return err
	}
	err = logDB.DeleteCache(&Log{
		Id: _id,
	})
	if err != nil {
		tp.Errorf("%s", err.Error())
	}
	return nil
}

// GetLogByPrimary query a Log data from database by primary key.
// NOTE:
//  Primary key: 'id';
//  With cache layer;
//  If @return bool=false error=nil, means the data is not exist.
func GetLogByPrimary(_id int64) (*Log, bool, error) {
	var _l = &Log{
		Id: _id,
	}
	err := logDB.CacheGet(_l)
	switch err {
	case nil:
		if _l.CreatedAt == 0 {
			return nil, false, nil
		}
		return _l, true, nil
	case sql.ErrNoRows:
		err2 := logDB.PutCache(_l)
		if err2 != nil {
			tp.Errorf("%s", err2.Error())
		}
		return nil, false, nil
	default:
		return nil, false, err
	}
}

// GetLogByWhere query a Log data from database by WHERE condition.
// NOTE:
//  Without cache layer;
//  If @return bool=false error=nil, means the data is not exist.
func GetLogByWhere(whereCond string, arg ...interface{}) (*Log, bool, error) {
	var _l = new(Log)
	err := logDB.Get(_l, "SELECT `id`,`text`,`updated_at`,`created_at` FROM `log` WHERE "+insertZeroDeletedTsField(whereCond)+" LIMIT 1;", arg...)
	switch err {
	case nil:
		return _l, true, nil
	case sql.ErrNoRows:
		return nil, false, nil
	default:
		return nil, false, err
	}
}

// SelectLogByWhere query some Log data from database by WHERE condition.
// NOTE:
//  Without cache layer.
func SelectLogByWhere(whereCond string, arg ...interface{}) ([]*Log, error) {
	var objs = new([]*Log)
	err := logDB.Select(objs, "SELECT `id`,`text`,`updated_at`,`created_at` FROM `log` WHERE "+insertZeroDeletedTsField(whereCond), arg...)
	return *objs, err
}

// CountLogByWhere count Log data number from database by WHERE condition.
// NOTE:
//  Without cache layer.
func CountLogByWhere(whereCond string, arg ...interface{}) (int64, error) {
	var count int64
	err := logDB.Get(&count, "SELECT count(1) FROM `log` WHERE "+insertZeroDeletedTsField(whereCond), arg...)
	return count, err
}
