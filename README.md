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
appends content to a file

#### func  Chmod

```go
func Chmod(files interface{}, mode os.FileMode)
```
change mode for an array of files or directories.

#### func  ChmodWithRecur

```go
func ChmodWithRecur(files interface{}, mode os.FileMode)
```
change mode for an array of files or directories with recursive

#### func  Chown

```go
func Chown(files interface{}, user, group int)
```
change the owner of an array of files or directories

#### func  ChownWithRecur

```go
func ChownWithRecur(files interface{}, user, group int)
```
change the owner of an array of files or directories recursive

#### func  Copy

```go
func Copy(srcFileName string, dstFileName string)
```
copies a file

#### func  Dirname

```go
func Dirname(fileName string) string
```
get dirname

#### func  Exists

```go
func Exists(paths interface{}) bool
```
checks the existence of files or directories

#### func  GetFilesAndDirs

```go
func GetFilesAndDirs(dirPath string) (files []string, dirs []string)
```
resolves files and directories

#### func  Hardlink

```go
func Hardlink(srcFileName string, dstFileNames interface{})
```
creates a hard link, or several hard links to a file

#### func  IsAbsolutePath

```go
func IsAbsolutePath(fileName string) bool
```
returns whether the file path is an absolute path

#### func  IsDir

```go
func IsDir(path string) bool
```
returns whether the file path is a directory

#### func  IsFile

```go
func IsFile(path string) bool
```
returns whether the file path is a file

#### func  IsReadable

```go
func IsReadable(fileName string) bool
```
tells whether a file exists and is readable

#### func  IsWritable

```go
func IsWritable(fileName string) bool
```
tells whether a file exists and is writable

#### func  Mkdir

```go
func Mkdir(paths interface{}, mode os.FileMode)
```
creates a directory of directory or directories.

#### func  Readlink

```go
func Readlink(path string) string
```
resolves links in paths.

#### func  Rename

```go
func Rename(srcFileName string, dstFileName string)
```
rename src to dst

#### func  Symlink

```go
func Symlink(srcDirName string, dstDirName string)
```
creates a symbolic link or copy a directory

#### func  Touch

```go
func Touch(files interface{})
```
creates new files

#### func  TouchFromTime

```go
func TouchFromTime(files interface{}, atime time.Time, mtime time.Time)
```
sets access and modification time of files