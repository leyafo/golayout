package path

import "testing"

func TestMakeParentDir(t *testing.T) {
	MakeParentDir("./test/a.log")
	if !PathIsExist("test"){
		t.Fatal("MakeParentDir failed")
	}
	if err := RemoveDir("test"); err!=nil{
		t.Fatal(err)
	}
}
