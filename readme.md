# filecopy
Модуль для копирования файлов по маске.

>Panic: нет
>
>Потокобезопасность: да
>
>Модуль: filecopy

## Открытые функции
```go
func CopyFiles(sourceDir, fileMask, targetDir string) error
```
## Подробное описание
Модуль filecopy предназначен для копирования файлов из исходной директории в целевую директорию по заданной маске файла.
## Описание функций
```go
func CopyFiles(sourceDir, fileMask, targetDir string) error
```
### Параметры
- `sourceDir`: путь к исходной директории.
- `fileMask`: маска файла для сопоставления файлов (например, "*.txt").
- `targetDir`: путь к целевой директории.

### Возвращаемые значения
В случае возникновения ошибки, возвращается `error` отличный от `nil`.

### Пример использования
```go
err := CopyFiles("/path/to/source", "*.txt", "/path/to/target") 
if err != nil { 
	log.Fatal(err) 
}
```