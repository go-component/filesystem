package filesystem

import (
	"testing"
	"time"
)

var dirRoot = "filesystem/test-mkdir"

var dirName = dirRoot + "/0"
var dirs = []string{
	dirRoot + "/1",
	dirRoot + "/2",
	dirRoot + "/3",
}

var recursiveDirRoot = dirRoot + "/sub-mkdir"

var recursiveDirName = recursiveDirRoot + "/0"

var recursiveDirs = []string{
	recursiveDirRoot + "/1",
	recursiveDirRoot + "/2",
	recursiveDirRoot + "/3",
}

func TestMkdir(t *testing.T) {

	Mkdir(dirName, 0755)

	if !Exists(dirName) {
		t.Error("Mkdir test failed!")
	}

	Remove(dirName)

}

func TestMultiMkdir(t *testing.T) {

	Mkdir(dirs, 0755)

	if !Exists(dirs) {
		t.Error("Multi mkdir test failed!")
	}

	Remove(dirs)

}

func TestRemove(t *testing.T) {

	Mkdir(dirName, 0755)
	if !Exists(dirName) {
		t.Error("Remove test failed!")
	}
	Remove(dirName)

	if Exists(dirName) {
		t.Error("Remove test failed!")
	}
}

func TestMultiRemove(t *testing.T) {

	Mkdir(dirs, 0755)
	if !Exists(dirs) {
		t.Error("Multi remove test failed!")
	}
	Remove(dirs)

	if Exists(dirs) {
		t.Error("Multi remove test failed!")
	}
}

func TestRemoveWithRecur(t *testing.T) {

	Mkdir(recursiveDirName, 0755)
	if !Exists(recursiveDirName) {
		t.Error("Multi remove test failed!")
	}
	RemoveWithRecur(recursiveDirRoot)

	if Exists(recursiveDirName) {
		t.Error("Multi remove test failed!")
	}
}

func TestMultiRemoveWithRecur(t *testing.T) {

	Mkdir(recursiveDirs, 0755)
	if !Exists(recursiveDirs) {
		t.Error("Multi remove test failed!")
	}
	RemoveWithRecur(recursiveDirRoot)

	if Exists(recursiveDirs) {
		t.Error("Multi remove test failed!")
	}
}

func TestAppendToFile(t *testing.T) {
	fileName := dirRoot+"/append-to-file.txt"

	AppendToFile(fileName, []byte("filesystem"))

	if !Exists(fileName){
		t.Error("AppendToFile test failed!")
	}

	Remove(fileName)

}

func TestCopy(t *testing.T) {
	srcFileName, dstFileName := dirRoot+"/src-file.txt" , dirRoot+"/dst-file.txt"

	AppendToFile(srcFileName, []byte("filesystem"))

	if !Exists(srcFileName){
		t.Error("Copy test failed!")
	}

	Copy(srcFileName, dstFileName)

	if !Exists(dstFileName){
		t.Error("Copy test failed!")
	}

	Remove([]string{srcFileName, dstFileName})

}

func TestTouch(t *testing.T) {
	fileName := dirRoot+"/touch.txt"

	Touch(fileName)

	if !Exists(fileName){
		t.Error("Touch test failed!")
	}

	Remove(fileName)

}

func TestMultiTouch(t *testing.T) {
	files := []string{
		dirRoot+"/touch1.txt",
		dirRoot+"/touch2.txt",
		dirRoot+"/touch3.txt",
	}

	Touch(files)

	if !Exists(files){
		t.Error("Touch test failed!")
	}
	Remove(files)
}

func TestTouchFromTime(t *testing.T) {
	fileName := dirRoot+"/touch.txt"

	now := time.Now().Local()

	TouchFromTime(fileName, now, now)

	if !Exists(fileName){
		t.Error("TouchFromTime test failed!")
	}

	Remove(fileName)

	Touch(fileName)

	TouchFromTime(fileName, now, now)
	Remove(fileName)
}

func TestMultiTouchFromTime(t *testing.T) {
	files := []string{
		dirRoot+"/touch1.txt",
		dirRoot+"/touch2.txt",
		dirRoot+"/touch3.txt",
	}

	now := time.Now().Local()

	TouchFromTime(files, now, now)

	if !Exists(files){
		t.Error("Multi touchFromTime test failed!")
	}

	Remove(files)

	Touch(files)

	TouchFromTime(files, now, now)
	Remove(files)
}

func TestChmod(t *testing.T) {
	fileName := dirRoot+"/chmod.txt"

	Touch(fileName)

	if !Exists(fileName){
		t.Error("Chmod test failed!")
	}

	Chmod(fileName, 0755)

	Remove(fileName)
}

func TestMultiChmod(t *testing.T) {
	files := []string{
		dirRoot+"/chmod1.txt",
		dirRoot+"/chmod2.txt",
		dirRoot+"/chmod3.txt",
	}

	Touch(files)

	if !Exists(files){
		t.Error("Multi chmod test failed!")
	}

	Chmod(files,755)

	Remove(files)
}

