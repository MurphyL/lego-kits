package repo

import (
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func TestName(t *testing.T) {
	ddl := `
	CREATE TABLE COMPANY(
	   ID INT PRIMARY KEY     NOT NULL,
	   NAME           TEXT    NOT NULL,
	   AGE            INT     NOT NULL,
	   ADDRESS        CHAR(50),
	   SALARY         REAL
	);
	`
	repoInst := NewRepo("sqlite3", ":memory:")
	ret := repoInst.MustExec(ddl)
	t.Log("执行结果：", ret)
}
