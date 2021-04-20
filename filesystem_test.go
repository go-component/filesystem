package filesystem

import (
	"testing"
)

var dirName = "filesystem/test-mkdir"
var dirs = []string{
	"filesystem/test-mkdir1",
	"filesystem/test-mkdir2",
	"filesystem/test-mkdir3",
}


func TestMkdir(t *testing.T) {


	Mkdir(dirName, 0755)

	if !Exists(dirName){
		t.Error("Mkdir test failed!")
	}

	t.Cleanup(func() {
		Remove(dirName)
	})
}

func TestMulMkdir(t *testing.T) {

	Mkdir(dirs, 0755)

	if !Exists(dirs){
		t.Error("Multi mkdir test failed!")
	}

	t.Cleanup(func() {
		Remove(dirs)
	})
}

func TestRemove(t *testing.T) {

	Mkdir(dirName, 0755)
	Remove(dirName)

	if Exists(dirName){
		t.Error("Remove test failed!")
	}
}

func TestMulRemove(t *testing.T) {

	Mkdir(dirs, 0755)
	Remove(dirs)

	if Exists(dirs){
		t.Error("Multi remove test failed!")
	}
}