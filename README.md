# filesystem
filesystem for golang

# installation

```
go get github.com/go-component/filesystem
```

# import
```
import "github.com/go-component/filesystem"
```


## Usage

#### func  AppendToFile

```go
func AppendToFile(fileName string, content []byte)
```
Appends content to a file.

#### func  Chmod

```go
func Chmod(files interface{}, mode os.FileMode)
```
Change mode for an array of files or directories.

#### func  ChmodWithRecur

```go
func ChmodWithRecur(files interface{}, mode os.FileMode)
```
Change mode for an array of files or directories with recursive mode.

#### func  Chown

```go
func Chown(files interface{}, user, group int)
```
Change the owner of an array of files or directories.

#### func  ChownWithRecur

```go
func ChownWithRecur(files interface{}, user, group int)
```
Change the owner of an array of files or directories recursive mode.

#### func  Copy

```go
func Copy(srcFileName string, dstFileName string)
```
Copies a file.

#### func  Dirname

```go
func Dirname(fileName string) string
```
Return dirname.

#### func  Exists

```go
func Exists(paths interface{}) bool
```
Checks the existence of files or directories.

#### func  GetFilesAndDirs

```go
func GetFilesAndDirs(dirPath string) (files []string, dirs []string)
```
Resolves files and directories.

#### func  Hardlink

```go
func Hardlink(srcFileName string, dstFileNames interface{})
```
Creates a hard link, or several hard links to files.

#### func  IsAbsolutePath

```go
func IsAbsolutePath(fileName string) bool
```
Return whether the file path is an absolute path.

#### func  IsDir

```go
func IsDir(path string) bool
```
Returns whether the file path is a directory.

#### func  IsFile

```go
func IsFile(path string) bool
```
Returns whether the file path is a file.

#### func  IsReadable

```go
func IsReadable(fileName string) bool
```
Tells whether a file exists and is readable.

#### func  IsWritable

```go
func IsWritable(fileName string) bool
```
Tells whether a file exists and is writable.

#### func  Mkdir

```go
func Mkdir(paths interface{}, mode os.FileMode)
```
Creates a directory of directory or directories.

#### func  Readlink

```go
func Readlink(path string) string
```
Resolves links in paths.

#### func  Remove

```go
func Remove(files interface{})
```
Remove removes the named file or (empty) directory.

#### func  RemoveWithRecur

```go
func RemoveWithRecur(files interface{})
```
Remove removes the named file or directory with recursive mode.

#### func  Rename

```go
func Rename(srcFileName string, dstFileName string)
```
Rename src to dst.

#### func  Symlink

```go
func Symlink(srcDirName string, dstDirName string)
```
Creates a symbolic link or copy a directory.

#### func  Touch

```go
func Touch(files interface{})
```
Creates new files if not exist.

#### func  TouchFromTime

```go
func TouchFromTime(files interface{}, atime time.Time, mtime time.Time)
```
Sets access and modification time of files.
