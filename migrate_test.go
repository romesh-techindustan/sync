package sync

import (
    "testing"
)

func TestDatabaseConn(t *testing.T){
	//Test cases
	dsn:= "root:@tcp(127.0.0.1:3306)/cms?parseTime=true&loc=Local"
	_,err:= DatabaseConn(dsn)
	if err != nil {
		t.Fatalf("should returns error but got nil")
	}
}