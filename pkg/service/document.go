package service

import (
	"context"
	"github.com/haggis-io/registry/pkg/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

type DocumentService interface {
	Get(name string) (out []*api.Document, err error)
	GetStrict(name, version string) (*api.Document, error)
	GetAll(documents ...*api.Document) (out []*api.Document, err error)
}

type documentService struct {
	client api.RegistryClient
}

func NewDocumentService(client api.RegistryClient) DocumentService {
	return &documentService{
		client: client,
	}
}

func (d *documentService) Get(name string) (out []*api.Document, err error) {

	req := api.GetDocumentsRequest{
		Name:   name,
		Status: api.Status_APPROVED,
	}

	ctx, cancelFunc := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancelFunc()

	docsRes, err := d.client.GetDocuments(ctx, &req)

	if err != nil {
		return
	}

	if docsRes != nil && len(docsRes.Documents) > 0 {
		out = docsRes.GetDocuments()
		return
	}

	return out, status.Error(codes.NotFound, "")
}

func (d *documentService) GetStrict(name, version string) (*api.Document, error) {

	req := api.GetDocumentRequest{
		Name:    name,
		Version: version,
	}

	ctx, cancelFunc := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancelFunc()

	docRes, err := d.client.GetDocument(ctx, &req)

	if err != nil {
		return nil, err
	}

	return docRes.GetDocument(), nil
}

func (d *documentService) GetAll(documents ...*api.Document) ([]*api.Document, error) {
	var out []*api.Document

	for _, document := range documents {
		if document.Version == "" {
			//TODO
			return out, nil
		}

		doc, err := d.GetStrict(document.Name, document.Version)

		if err != nil {
			return out, err
		}

		out = append(out, doc)

	}

	return out, nil
}
