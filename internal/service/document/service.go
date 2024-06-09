package document

type DocumentStorage interface{
}

type DocumentService struct {
	docStorage DocumentStorage
}

func NewDocumentService(docStorage DocumentStorage) *DocumentService {
	return &DocumentService{
		docStorage: docStorage,
	}
}