package filesystem

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"syscall"
	"time"
)

// creates a directory recursively
func Mkdir(paths interface{}, mode os.FileMode) {

	for _, pathName := range toIterable(paths) {
		err := os.MkdirAll(pathName, mode)
		if err != nil {
			panic(err)
		}
	}

}

func Copy(originFile string, targetFile string) {
	
}

func CopyWithOverwrite(originFile string, targetFile string) {

}

//checks the existence of files or directories
func Exists(paths interface{}) bool {
	for _, pathName := range toIterable(paths) {
		_, err := os.Stat(pathName) // For read access.
		if err != nil {
			return false
		}
	}

	return true
}

func Touch(files interface{}) {
	for _, fileName := range toIterable(files) {
		if !Exists(fileName) {
			file, err := os.Create(fileName)
			if err != nil {
				panic(file)
			}
			err = file.Close()
			if err != nil {
				panic(file)
			}
		}
	}
}

func TouchFromTime(files interface{}, atime time.Time, mtime time.Time) {
	for _, fileName := range toIterable(files) {

		if Exists(fileName) {
			err := os.Chtimes(fileName, atime, mtime)
			if err != nil {
				panic(err)
			}
		} else {
			file, err := os.Create(fileName)
			if err != nil {
				panic(file)
			}
			err = file.Close()
			if err != nil {
				panic(file)
			}
		}
	}
}

func Chmod(files interface{}, mode os.FileMode) {
	for _, fileName := range toIterable(files) {

		if !Exists(fileName) {
			continue
		}
		err := os.Chmod(fileName, mode)
		if err != nil {
			panic(err)
		}
	}
}

func ChmodWithRecur(files interface{}, mode os.FileMode) {

	for _, fileName := range toIterable(files) {

		if !Exists(fileName) {
			continue
		}
		if IsDir(fileName) {

			fs, ds := GetFilesAndDirs(fileName)

			ChmodWithRecur(append(fs, ds...), mode)

		} else {
			err := os.Chmod(fileName, mode)

			if err != nil {
				panic(err)
			}
		}
	}
}

func Chown(files interface{}, user, group int) {

	for _, fileName := range toIterable(files) {

		if !Exists(fileName) {
			continue
		}
		err := os.Chown(fileName, user, group)

		if err != nil {
			panic(err)
		}
	}
}

func ChownWithRecur(files interface{}, user, group int) {

	for _, fileName := range toIterable(files) {

		if !Exists(fileName) {
			continue
		}

		if IsDir(fileName) {

			fs, ds := GetFilesAndDirs(fileName)

			ChownWithRecur(append(fs, ds...), user, group)

		} else {
			err := os.Chown(fileName, user, group)

			if err != nil {
				panic(err)
			}
		}
	}
}

func GetFilesAndDirs(dirPth string) (files []string, dirs []string) {
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		panic(err)
	}

	PthSep := string(os.PathSeparator)

	for _, fi := range dir {
		if fi.IsDir() { // 目录, 递归遍历
			dirs = append(dirs, dirPth+PthSep+fi.Name())
			_, _ = GetFilesAndDirs(dirPth + PthSep + fi.Name())

		} else {
			files = append(files, dirPth+PthSep+fi.Name())
		}
	}

	return files, dirs
}

func Rename(originFile string, targetFile string) {
	err := os.Rename(originFile, targetFile)

	if err != nil{
		panic(err)
	}
}

func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// 判断所给路径是否为文件
func IsFile(path string) bool {
	return !IsDir(path)
}

func IsReadable(fileName string) bool {
	err := syscall.Access(fileName, syscall.O_RDONLY)
	if err != nil {
		return false
	}

	return true
}

func IsWritable(fileName string) bool {

	err := syscall.Access(fileName, syscall.O_RDWR)
	if err != nil {
		return false
	}

	return true
}

func Symlink(originDir string, targetDir string) {
	err := os.Symlink(originDir, targetDir)

	if err != nil{
		panic(err)
	}
}

func Hardlink(originFile string, targetFiles interface{}) {

}

func Readlink(path string) string{
	link,err := os.Readlink(path)

	if err != nil{
		panic(err)
	}
	return link
}

func MakePathRelative(endPath string, startPath string) {

}

func IsAbsolutePath(fileName string) bool {
	return filepath.IsAbs(fileName)
}

func Dirname(fileName string) string {
	return filepath.Dir(fileName)
}

func AppendToFile(fileName string, content []byte) {

	dir := Dirname(fileName)

	if !Exists(dir) {
		Mkdir(dir, 0755)
	}

	if !IsWritable(fileName) {
		panic(fmt.Sprintf("Unable to write to the \"%s\" directory.", fileName))
	}

	file, _ := os.Open(fileName)

	_, err := file.Write(content)

	if err != nil {
		panic(err)
	}
}

func toIterable(files interface{}) []string {
	switch files.(type) {
	case string:
		return []string{files.(string)}

	case []string:
		return files.([]string)
	default:
		panic("invalid interface type")
	}
}
