package usecasesinterface

import (
	"context"
	helperstructs "ecommerce/web/helpers/helper_structs"
	"ecommerce/web/helpers/responce"
)

type ProductUsecaseInterface interface {
	AddProduct(ctx context.Context, productreq helperstructs.ProductReq) (responce.ProuctData, error)
	GetProducts(ctx context.Context, email, count, page string) ([]responce.ProuctData, error)
	UpdateProduct(ctx context.Context, productreq helperstructs.ProductReq) (responce.ProuctData, error)
	DeleteProduct(ctx context.Context, product_id uint) error
	UpdateStock(ctx context.Context, id uint, stock int) error
	GetProductByID(ctx context.Context, id string, email string) (responce.ProuctData, error)
	AddProductImages(ctx context.Context, id uint, path string) error
}
