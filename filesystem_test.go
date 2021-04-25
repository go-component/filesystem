package filesystem

import (
	"os"
	"reflect"
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

	err := Mkdir(dirName, 0755)
	if err != nil {
		panic(err)
	}

	t.Cleanup(func() {
		_ = Remove(dirs)
	})

	if !Exists(dirName) {
		t.Error("Mkdir test failed!")
	}

}

func TestMultiMkdir(t *testing.T) {

	err := Mkdir(dirs, 0755)
	if err != nil {
		panic(err)
	}

	t.Cleanup(func() {
		_ = Remove(dirs)
	})

	if !Exists(dirs) {
		t.Error("Multi Mkdir test failed!")
	}

}

func TestRemove(t *testing.T) {

	err := Mkdir(dirName, 0755)
	if err != nil {
		panic(err)
	}
	if !Exists(dirName) {
		t.Error("Remove test failed!")
	}
	err = Remove(dirName)
	if err != nil {
		panic(err)
	}

	if Exists(dirName) {
		t.Error("Remove test failed!")
	}
}

func TestMultiRemove(t *testing.T) {

	err := Mkdir(dirs, 0755)
	if err != nil {
		panic(err)
	}
	if !Exists(dirs) {
		t.Error("Multi Remove test failed!")
	}
	err = Remove(dirs)
	if err != nil {
		panic(err)
	}

	if Exists(dirs) {
		t.Error("Multi Remove test failed!")
	}
}

func TestRemoveWithRecur(t *testing.T) {

	err := Mkdir(recursiveDirName, 0755)
	if err != nil {
		panic(err)
	}
	if !Exists(recursiveDirName) {
		t.Error("Multi Remove test failed!")
	}
	err = RemoveWithRecur(recursiveDirRoot)
	if err != nil {
		panic(err)
	}

	if Exists(recursiveDirName) {
		t.Error("Multi Remove test failed!")
	}
}

func TestMultiRemoveWithRecur(t *testing.T) {

	err := Mkdir(recursiveDirs, 0755)
	if err != nil {
		panic(err)
	}

	if !Exists(recursiveDirs) {
		t.Error("Multi RemoveWithRecur test failed!")
	}
	err = RemoveWithRecur(recursiveDirRoot)
	if err != nil {
		panic(err)
	}

	if Exists(recursiveDirs) {
		t.Error("Multi RemoveWithRecur test failed!")
	}
}

func TestAppendToFile(t *testing.T) {
	fileName := dirRoot + "/append-to-file.txt"

	err := AppendToFile(fileName, []byte("filesystem"))
	if err != nil {
		panic(err)
	}

	t.Cleanup(func() {
		_ = RemoveWithRecur(dirRoot)
	})

	if !Exists(fileName) {
		t.Error("AppendToFile test failed!")
	}
}

func TestCopy(t *testing.T) {
	srcFileName, dstFileName := dirRoot+"/src-file.txt", dirRoot+"/dst-file.txt"

	err := AppendToFile(srcFileName, []byte("filesystem"))
	if err != nil {
		panic(err)
	}

	t.Cleanup(func() {
		_ = RemoveWithRecur(dirRoot)
	})

	if !Exists(srcFileName) {
		t.Error("Copy test failed!")
	}

	err = Copy(srcFileName, dstFileName)
	if err != nil {
		panic(err)
	}

	if !Exists(dstFileName) {
		t.Error("Copy test failed!")
	}

}

func TestTouch(t *testing.T) {
	fileName := dirRoot + "/touch.txt"

	err := Touch(fileName)
	if err != nil {
		panic(err)
	}

	t.Cleanup(func() {
		_ = RemoveWithRecur(dirRoot)
	})

	if !Exists(fileName) {
		t.Error("Touch test failed!")
	}

}

func TestMultiTouch(t *testing.T) {
	files := []string{
		dirRoot + "/touch1.txt",
		dirRoot + "/touch2.txt",
		dirRoot + "/touch3.txt",
	}

	err := Touch(files)
	if err != nil {
		panic(err)
	}

	t.Cleanup(func() {
		_ = RemoveWithRecur(dirRoot)
	})

	if !Exists(files) {
		t.Error("Touch test failed!")
	}
}

func TestTouchFromTime(t *testing.T) {
	fileName := dirRoot + "/touch.txt"

	now := time.Now().Local()

	err := TouchFromTime(fileName, now, now)
	if err != nil {
		panic(err)
	}

	t.Cleanup(func() {
		_ = RemoveWithRecur(dirRoot)
	})

	if !Exists(fileName) {
		t.Error("TouchFromTime test failed!")
	}

	err = RemoveWithRecur(dirRoot)
	if err != nil {
		panic(err)
	}

	err = Touch(fileName)
	if err != nil {
		panic(err)
	}

	err = TouchFromTime(fileName, now, now)
	if err != nil {
		panic(err)
	}
}

func TestMultiTouchFromTime(t *testing.T) {
	files := []string{
		dirRoot + "/touch1.txt",
		dirRoot + "/touch2.txt",
		dirRoot + "/touch3.txt",
	}

	now := time.Now().Local()

	err := TouchFromTime(files, now, now)
	if err != nil {
		panic(err)
	}

	t.Cleanup(func() {
		_ = RemoveWithRecur(dirRoot)
	})

	if !Exists(files) {
		t.Error("Multi TouchFromTime test failed!")
	}

	err = RemoveWithRecur(dirRoot)
	if err != nil {
		panic(err)
	}

	err = Touch(files)
	if err != nil {
		panic(err)
	}

	err = TouchFromTime(files, now, now)
	if err != nil {
		panic(err)
	}
}

func TestChmod(t *testing.T) {
	fileName := dirRoot + "/chmod.txt"

	err := Touch(fileName)
	if err != nil {
		panic(err)
	}

	t.Cleanup(func() {
		_ = RemoveWithRecur(dirRoot)
	})

	if !Exists(fileName) {
		t.Error("Chmod test failed!")
	}

	err = Chmod(fileName, 0755)
	if err != nil {
		panic(err)
	}
}

func TestMultiChmod(t *testing.T) {
	files := []string{
		dirRoot + "/chmod1.txt",
		dirRoot + "/chmod2.txt",
		dirRoot + "/chmod3.txt",
	}

	err := Touch(files)
	if err != nil {
		panic(err)
	}

	t.Cleanup(func() {
		_ = RemoveWithRecur(dirRoot)
	})

	if !Exists(files) {
		t.Error("Multi Chmod test failed!")
	}

	err = Chmod(files, 755)
	if err != nil {
		panic(err)
	}
}

func TestChmodWithRecur(t *testing.T) {
	fileName := recursiveDirRoot + "/chmod.txt"

	err := Touch(fileName)
	if err != nil {
		panic(err)
	}

	t.Cleanup(func() {
		_ = RemoveWithRecur(dirRoot)
	})

	if !Exists(fileName) {
		t.Error("ChmodWithRecur test failed!")
	}

	err = ChmodWithRecur(dirRoot, 0755)
	if err != nil {
		panic(err)
	}
}

func TestMultiChmodWithRecur(t *testing.T) {
	fileName := []string{
		recursiveDirRoot + "/chmod1.txt",
		recursiveDirRoot + "/chmod2.txt",
		recursiveDirRoot + "/chmod3.txt",
	}

	err := Touch(fileName)
	if err != nil {
		panic(err)
	}

	t.Cleanup(func() {
		_ = RemoveWithRecur(dirRoot)
	})

	if !Exists(fileName) {
		t.Error("Multi MultiChmodWithRecur test failed!")
	}

	err = ChmodWithRecur(dirRoot, 0755)
	if err != nil {
		panic(err)
	}
}

func TestChown(t *testing.T) {
	fileName := dirRoot + "/chown.txt"

	err := Touch(fileName)
	if err != nil {
		panic(err)
	}

	t.Cleanup(func() {
		_ = RemoveWithRecur(dirRoot)
	})

	if !Exists(fileName) {
		t.Error("Chown test failed!")
	}

	err = Chown(fileName, 501, 20)
	if err != nil {
		panic(err)
	}
}

func TestMultiChown(t *testing.T) {
	files := []string{
		dirRoot + "/chown1.txt",
		dirRoot + "/chown2.txt",
		dirRoot + "/chown3.txt",
	}

	err := Touch(files)
	if err != nil {
		panic(err)
	}

	t.Cleanup(func() {
		_ = RemoveWithRecur(dirRoot)
	})

	if !Exists(files) {
		t.Error("Multi Chown test failed!")
	}

	err = Chown(files, 501, 20)
	if err != nil {
		panic(err)
	}
}

func TestChownWithRecur(t *testing.T) {
	fileName := recursiveDirRoot + "/chown.txt"

	err := Touch(fileName)
	if err != nil {
		panic(err)
	}

	t.Cleanup(func() {
		_ = RemoveWithRecur(dirRoot)
	})

	if !Exists(fileName) {
		t.Error("ChownWithRecur test failed!")
	}

	err = ChownWithRecur(dirRoot, 501, 20)
	if err != nil {
		panic(err)
	}
}

func TestMultiChownWithRecur(t *testing.T) {
	fileName := []string{
		recursiveDirRoot + "/chown1.txt",
		recursiveDirRoot + "/chown2.txt",
		recursiveDirRoot + "/chown3.txt",
	}

	err := Touch(fileName)
	if err != nil {
		panic(err)
	}

	t.Cleanup(func() {
		_ = RemoveWithRecur(dirRoot)
	})

	if !Exists(fileName) {
		t.Error("Multi ChownWithRecur test failed!")
	}

	err = ChownWithRecur(dirRoot, 501, 20)
	if err != nil {
		panic(err)
	}
}

func TestRename(t *testing.T) {
	srcFileName, dstFileName := dirRoot+"/src-file.txt", dirRoot+"/dst-file.txt"

	err := AppendToFile(srcFileName, []byte("filesystem"))
	if err != nil {
		panic(err)
	}

	t.Cleanup(func() {
		_ = RemoveWithRecur(dirRoot)
	})

	if !Exists(srcFileName) {
		t.Error("Rename test failed!")
	}

	err = Rename(srcFileName, dstFileName)
	if err != nil {
		panic(err)
	}

	if !Exists(dstFileName) {
		t.Error("Rename test failed!")
	}
}

func TestIsDir(t *testing.T) {
	err := Mkdir(dirRoot, 0755)
	if err != nil {
		panic(err)
	}
	fileName := dirRoot + "/test.txt"
	err = Touch(fileName)
	if err != nil {
		panic(err)
	}
	t.Cleanup(func() {
		_ = RemoveWithRecur(dirRoot)
	})

	if !Exists(dirRoot) {
		t.Error("IsDir test failed!")
	}

	if IsDir(dirRoot) != true {
		t.Error("IsDir test failed!")
	}
	if IsDir(fileName) == true {
		t.Error("IsDir test failed!")
	}
}

func TestIsFile(t *testing.T) {
	err := Mkdir(dirRoot, 0755)
	if err != nil {
		panic(err)
	}
	fileName := dirRoot + "/test.txt"
	err = Touch(fileName)
	if err != nil {
		panic(err)
	}

	t.Cleanup(func() {
		_ = RemoveWithRecur(dirRoot)
	})

	if !Exists(dirRoot) {
		t.Error("IsFile test failed!")
	}

	if IsFile(dirRoot) == true {
		t.Error("IsFile test failed!")
	}
	if IsFile(fileName) != true {
		t.Error("IsFile test failed!")
	}

}

func TestIsReadable(t *testing.T) {
	writeDir := "privilege" + "/read"
	err := Mkdir(writeDir, 0755)
	if err != nil {
		panic(err)
	}

	t.Cleanup(func() {
		_ = RemoveWithRecur("privilege")
	})

	if !IsReadable(writeDir) {
		t.Error("IsReadable test failed!")
	}
}

func TestIsWriteable(t *testing.T) {

	writeDir := "privilege" + "/write"
	err := Mkdir(writeDir, 0755)
	if err != nil {
		panic(err)
	}

	t.Cleanup(func() {
		_ = RemoveWithRecur("privilege")
	})

	if !IsWritable(writeDir) {
		t.Error("IsWriteable test failed!")
	}
}


func TestHardlink(t *testing.T) {
	srcFileName := dirRoot + "/src.txt"
	dstFileName := dirRoot + "/dst.txt"
	err := Touch(srcFileName)
	if err != nil{
		panic(err)
	}

	err = Hardlink(srcFileName, dstFileName)
	if err != nil{
		panic(err)
	}

	srcFileInfo, err := os.Stat(srcFileName)

	if err != nil{
		panic(err)
	}

	dstFileInfo, err := os.Stat(dstFileName)

	if err != nil{
		panic(err)
	}


	srcNode := reflect.ValueOf(srcFileInfo.Sys()).Elem().Field(3).Uint()
	dstNode := reflect.ValueOf(dstFileInfo.Sys()).Elem().Field(3).Uint()


	t.Cleanup(func() {
		_ = RemoveWithRecur(dirRoot)
	})

	if srcNode != dstNode{
		t.Error("Hardlink test failed!")
	}
}
