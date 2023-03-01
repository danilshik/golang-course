 package storage

 import "github.com/danilshik/storage/internal/file"

 type Storage struct{}

 func NewStorage() *Storage {
	return &Storage{}
 }

 func (s *Storage) Upload(fileName string, blob []byte) (*file.File, error){
	return newFile, err := file.newFile(filename, blob)
	// if err != nil{
	// 	return nil, err
	// }

	// return newFile
}