package filesystem

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"syscall"
	"time"
)

// creates a directory of directory or directories.
func Mkdir(paths interface{}, mode os.FileMode) {

	for _, pathName := range toIterable(paths) {
		err := os.MkdirAll(pathName, mode)
		if err != nil {
			panic(err)
		}
	}

}

// copies a file
func Copy(srcFileName string, dstFileName string) {

	srcFile, err := os.Open(srcFileName)
	if err != nil{
		panic(err)
	}

	dstFile, err := os.OpenFile(dstFileName, os.O_CREATE|os.O_WRONLY, 0666)

	if err != nil{
		panic(err)
	}

	_, err = io.Copy(dstFile, srcFile)

	if err != nil{
		panic(err)
	}
}

// checks the existence of files or directories
func Exists(paths interface{}) bool {
	for _, pathName := range toIterable(paths) {
		_, err := os.Stat(pathName) // For read access.
		if err != nil {
			return false
		}
	}

	return true
}

// creates new files
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

// sets access and modification time of files
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

// change mode for an array of files or directories.
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

// change mode for an array of files or directories with recursive
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

// change the owner of an array of files or directories
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

// change the owner of an array of files or directories recursive
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

// resolves files and directories
func GetFilesAndDirs(dirPath string) (files []string, dirs []string) {
	dir, err := ioutil.ReadDir(dirPath)
	if err != nil {
		panic(err)
	}

	PthSep := string(os.PathSeparator)

	for _, fi := range dir {
		if fi.IsDir() { // 目录, 递归遍历
			dirs = append(dirs, dirPath+PthSep+fi.Name())
			_, _ = GetFilesAndDirs(dirPath + PthSep + fi.Name())

		} else {
			files = append(files, dirPath+PthSep+fi.Name())
		}
	}

	return files, dirs
}

// rename src to dst
func Rename(srcFileName string, dstFileName string) {
	err := os.Rename(srcFileName, dstFileName)

	if err != nil{
		panic(err)
	}
}

// returns whether the file path is a directory
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// returns whether the file path is a file
func IsFile(path string) bool {
	return !IsDir(path)
}

// tells whether a file exists and is readable
func IsReadable(fileName string) bool {
	err := syscall.Access(fileName, syscall.O_RDONLY)
	if err != nil {
		return false
	}

	return true
}

// tells whether a file exists and is writable
func IsWritable(fileName string) bool {

	err := syscall.Access(fileName, syscall.O_RDWR)
	if err != nil {
		return false
	}

	return true
}

// creates a symbolic link or copy a directory
func Symlink(srcDirName string, dstDirName string) {
	err := os.Symlink(srcDirName, dstDirName)

	if err != nil{
		panic(err)
	}
}

// creates a hard link, or several hard links to a file
func Hardlink(srcFileName string, dstFileNames interface{}) {

}
// resolves links in paths.
func Readlink(path string) string{
	link,err := os.Readlink(path)

	if err != nil{
		panic(err)
	}
	return link
}

// returns whether the file path is an absolute path
func IsAbsolutePath(fileName string) bool {
	return filepath.IsAbs(fileName)
}

// get dirname
func Dirname(fileName string) string {
	return filepath.Dir(fileName)
}

// appends content to a file
func AppendToFile(fileName string, content []byte) {

	dir := Dirname(fileName)

	if !Exists(dir) {
		Mkdir(dir, 0666)
	}

	if !IsWritable(fileName) {
		panic(fmt.Sprintf("Unable to write to the \"%s\" directory.", fileName))
	}

	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE, 0666)

	if err != nil{
		panic(err)
	}

	_, err = file.Write(content)

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
