package service

import (
	"context"
	"github.com/haggis-io/jenerate/cmd/errors"
	"github.com/haggis-io/registry/pkg/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

type DocumentService interface {
	GetAll(name string) (out []string, err error)
	Get(name, version string) (*api.Document, error)
}

type documentService struct {
	client api.RegistryClient
}

func NewDocumentService(client api.RegistryClient) DocumentService {
	return &documentService{
		client: client,
	}
}

func (d *documentService) GetAll(name string) (out []string, err error) {

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
		out = ExtractVersionsFromDocumentSlice(docsRes.Documents)
		return
	}

	return out, status.Error(codes.NotFound, errors.DocumentNotFoundErr.Error())
}

func (d *documentService) Get(name, version string) (*api.Document, error) {

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
