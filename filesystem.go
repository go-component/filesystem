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

// Creates a directory of directory or directories.
func Mkdir(paths interface{}, mode os.FileMode) error {

	for _, pathName := range toIterable(paths) {

		if Exists(pathName) {
			continue
		}

		err := os.MkdirAll(pathName, mode)
		if err != nil {
			return err
		}
	}

	return nil
}

// Remove removes the named file or (empty) directory.
func Remove(files interface{}) error {

	for _, fileName := range toIterable(files) {

		if fileName == "" {
			continue
		}

		if !Exists(fileName) {
			continue
		}

		err := os.Remove(fileName)

		if err != nil {
			return err
		}
	}
	return nil
}

// Remove removes the named file or directory with recursive mode.
func RemoveWithRecur(files interface{}) error {

	for _, fileName := range toIterable(files) {

		if fileName == "" {
			continue
		}

		if !Exists(fileName) {
			continue
		}

		if IsDir(fileName) {
			fs, ds, err := ResolveFilesAndDirs(fileName)
			if err != nil {
				return err
			}

			moreFiles := append(fs, ds...)
			if len(moreFiles) > 0 {
				err := RemoveWithRecur(moreFiles)
				if err != nil {
					return err
				}
			}
		}

		err := os.Remove(fileName)

		if err != nil {
			return err
		}
	}

	return nil
}

// Copies a file.
func Copy(srcFileName string, dstFileName string) error {

	srcFile, err := os.Open(srcFileName)
	if err != nil {
		return err
	}

	dstFile, err := os.OpenFile(dstFileName, os.O_CREATE|os.O_WRONLY, 0666)

	if err != nil {
		return err
	}

	_, err = io.Copy(dstFile, srcFile)

	if err != nil {
		return err
	}

	return nil
}

// Checks the existence of files or directories.
func Exists(paths interface{}) bool {

	for _, pathName := range toIterable(paths) {
		_, err := os.Stat(pathName) // For read access.
		if err != nil {
			return false
		}
	}

	return true
}

// Creates new files if not exist.
func Touch(files interface{}) error {

	for _, fileName := range toIterable(files) {

		dir := Dirname(fileName)
		if !Exists(dir) {
			err := Mkdir(dir, 0755)
			if err != nil {
				return err
			}
		}

		if !Exists(fileName) {
			file, err := os.Create(fileName)
			if err != nil {
				return err
			}
			err = file.Close()
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// Sets access and modification time of files.
func TouchFromTime(files interface{}, atime time.Time, mtime time.Time) error {

	for _, fileName := range toIterable(files) {
		dir := Dirname(fileName)
		if !Exists(dir) {
			err := Mkdir(dir, 0755)
			if err != nil {
				return err
			}
		}
		if Exists(fileName) {
			err := os.Chtimes(fileName, atime, mtime)
			if err != nil {
				return err
			}
		} else {
			file, err := os.Create(fileName)
			if err != nil {
				return err
			}
			err = file.Close()
			if err != nil {
				return err
			}
			err = TouchFromTime(fileName, atime, mtime)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// Change mode for an array of files or directories.
func Chmod(files interface{}, mode os.FileMode) error {

	for _, fileName := range toIterable(files) {

		if !Exists(fileName) {
			continue
		}
		err := os.Chmod(fileName, mode)
		if err != nil {
			return err
		}
	}
	return nil
}

// Change mode for an array of files or directories with recursive mode.
func ChmodWithRecur(files interface{}, mode os.FileMode) error {

	for _, fileName := range toIterable(files) {

		if !Exists(fileName) {
			continue
		}
		if IsDir(fileName) {

			fs, ds, err := ResolveFilesAndDirs(fileName)
			if err != nil {
				return err
			}

			err = ChmodWithRecur(append(fs, ds...), mode)
			if err != nil {
				return err
			}

		} else {
			err := os.Chmod(fileName, mode)

			if err != nil {
				return err
			}
		}
	}

	return nil
}

// Change the owner of an array of files or directories.
func Chown(files interface{}, user, group int) error {

	for _, fileName := range toIterable(files) {

		if !Exists(fileName) {
			continue
		}
		err := os.Chown(fileName, user, group)

		if err != nil {
			return err
		}
	}

	return nil
}

// Change the owner of an array of files or directories recursive mode.
func ChownWithRecur(files interface{}, user, group int) error {

	for _, fileName := range toIterable(files) {

		if !Exists(fileName) {
			continue
		}

		if IsDir(fileName) {

			fs, ds, err := ResolveFilesAndDirs(fileName)
			if err != nil {
				return err
			}

			err = ChownWithRecur(append(fs, ds...), user, group)
			if err != nil {
				return err
			}

		} else {
			err := os.Chown(fileName, user, group)

			if err != nil {
				return err
			}
		}
	}

	return nil
}

// Resolves files and directories.
func ResolveFilesAndDirs(dirPath string) (files []string, dirs []string, err error) {

	dir, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return []string{}, []string{}, err
	}

	PathSep := string(os.PathSeparator)

	for _, fi := range dir {
		if fi.IsDir() {
			dirs = append(dirs, dirPath+PathSep+fi.Name())
			_, _, _ = ResolveFilesAndDirs(dirPath + PathSep + fi.Name())

		} else {
			files = append(files, dirPath+PathSep+fi.Name())
		}
	}

	return files, dirs, nil
}

// Rename src to dst.
func Rename(srcFileName string, dstFileName string) error {

	err := os.Rename(srcFileName, dstFileName)

	if err != nil {
		return err
	}

	return nil
}

// Returns whether the file path is a directory.
func IsDir(path string) bool {

	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// Returns whether the file path is a file.
func IsFile(path string) bool {

	return !IsDir(path)
}

// Tells whether a file exists and is readable.
func IsReadable(fileName string) bool {

	err := syscall.Access(fileName, syscall.O_RDONLY)
	if err != nil {
		return false
	}

	return true
}

// Tells whether a file exists and is writable.
func IsWritable(fileName string) bool {

	err := syscall.Access(fileName, syscall.O_RDWR)
	if err != nil {
		return false
	}

	return true
}

// Creates a symbolic link or copy a directory.
func Symlink(srcDirName string, dstDirName string) error {
	err := os.Symlink(srcDirName, dstDirName)

	if err != nil {
		return err
	}

	return nil
}

// Creates a hard link, or several hard links to files.
func Hardlink(srcFileName string, dstFileName string) error {
	err := os.Link(srcFileName, dstFileName)
	if err != nil {
		return err
	}
	return nil
}

// Resolves links in paths.
func Readlink(path string) (string, error) {
	link, err := os.Readlink(path)

	if err != nil {
		return "", err
	}
	return link, nil
}

// Return whether the file path is an absolute path.
func IsAbsolutePath(fileName string) bool {
	return filepath.IsAbs(fileName)
}

// Return dirname.
func Dirname(fileName string) string {
	return filepath.Dir(fileName)
}

// Appends content to a file.
func AppendToFile(fileName string, content []byte) error {

	dir := Dirname(fileName)
	if !Exists(dir) {
		err := Mkdir(dir, 0755)
		if err != nil {
			return err
		}
	}

	if !IsWritable(dir) {
		panic(fmt.Sprintf("Unable to write to the \"%s\" directory.", fileName))
	}

	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}

	_, err = file.Write(content)

	if err != nil {
		return err
	}

	return nil
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
