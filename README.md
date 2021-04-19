# filesystem
filesystem for golang

package filesystem // import "filesystem"

##### func AppendToFile(fileName string, content []byte)
##### func Chmod(files interface{}, mode os.FileMode)
##### func ChmodWithRecur(files interface{}, mode os.FileMode)
##### func Chown(files interface{}, user, group int)
##### func ChownWithRecur(files interface{}, user, group int)
##### func Copy(originFile string, targetFile string)
##### func CopyWithOverwrite(originFile string, targetFile string)
##### func Dirname(fileName string) string
##### func Exists(paths interface{}) bool
##### func GetFilesAndDirs(dirPth string) (files []string, dirs []string)
##### func Hardlink(originFile string, targetFiles interface{})
##### func IsAbsolutePath(fileName string) bool
##### func IsDir(path string) bool
##### func IsFile(path string) bool
##### func IsReadable(fileName string) bool
##### func IsWritable(fileName string) bool
##### func MakePathRelative(endPath string, startPath string)
##### func Mkdir(paths interface{}, mode os.FileMode)
##### func Readlink(path string) string
##### func Rename(originFile string, targetFile string)
##### func Symlink(originDir string, targetDir string)
##### func Touch(files interface{})
##### func TouchFromTime(files interface{}, atime time.Time, mtime time.Time)