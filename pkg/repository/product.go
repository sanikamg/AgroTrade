package repository

import (
	"context"
	"errors"
	"fmt"
	"golang_project_ecommerce/pkg/domain"
	interfaces "golang_project_ecommerce/pkg/repository/interface"
	"golang_project_ecommerce/pkg/utils/res"

	"gorm.io/gorm"
)

type productDatabase struct {
	DB *gorm.DB
}

func NewProductRepository(DB *gorm.DB) interfaces.ProductRepository {
	return &productDatabase{DB}
}

// category
func (pd *productDatabase) FindCategory(c context.Context, category domain.Category) (domain.Category, error) {
	var tempCategory domain.Category
	err := pd.DB.Where("category_name=?", category.CategoryName).First(&tempCategory).Error
	if err != nil {
		return domain.Category{}, errors.New("failed find category")
	}
	return tempCategory, nil
}
func (pd *productDatabase) AddCategory(c context.Context, category domain.Category) (domain.Category, error) {
	err := pd.DB.Create(&category).Error

	if err != nil {
		return domain.Category{}, errors.New("failed to add category")
	}
	return category, nil
}
func (pd *productDatabase) FindAllCategory(c context.Context) ([]res.AllCategories, error) {
	var categories []res.AllCategories
	query := `select *from categories`
	if pd.DB.Raw(query).Scan(&categories).Error != nil {
		return []res.AllCategories{}, errors.New("failed to get categories ")
	}
	fmt.Println(categories)
	return categories, nil
}

//.........................................................................................//

// product
func (pd *productDatabase) AddProduct(c context.Context, product domain.ProductDetails) (domain.ProductDetails, error) {
	err := pd.DB.Create(&product).Error
	if err != nil {
		return domain.ProductDetails{}, errors.New("failed to add product")
	}
	return product, nil
}

// find
func (pd *productDatabase) FindProductById(c context.Context, productid uint) error {
	var product domain.ProductDetails
	err := pd.DB.Where("product_id=?", productid).First(&product).Error
	if err != nil {
		return errors.New("failed to find product")
	}
	return nil
}

func (pd *productDatabase) FindProduct(c context.Context, product domain.ProductDetails) (domain.ProductDetails, error) {
	err := pd.DB.Where("product_id=? OR product_name=?", product.Product_Id, product.ProductName).First(&product).Error
	if err != nil {
		return domain.ProductDetails{}, errors.New("failed to find product")
	}

	return product, nil
}

func (pd *productDatabase) AddQuantity(c context.Context, product domain.ProductDetails) (domain.ProductDetails, error) {
	// Execute the update query
	fmt.Println("repo", product)
	query := `UPDATE product_details SET product_quantity = product_quantity + ? WHERE product_id = ?`
	err := pd.DB.Exec(query, product.ProductQuantity, product.Product_Id).Error
	if err != nil {
		return domain.ProductDetails{}, errors.New("failed to update product_details")
	}

	// Retrieve the updated data
	var pro domain.ProductDetails
	err = pd.DB.Where("product_id = ?", product.Product_Id).Find(&pro).Error
	if err != nil {
		return domain.ProductDetails{}, errors.New("failed to fetch updated product_details")
	}
	pro.Category = product.Category

	return pro, nil
}

func (pd *productDatabase) GetCategoryByID(c context.Context, categoryId int) (domain.Category, error) {
	var category domain.Category
	query := `select * from categories where id=?`
	err := pd.DB.Raw(query, categoryId).Scan(&category).Error
	if err != nil {
		return domain.Category{}, errors.New("failed to find category name")
	}

	return category, nil
}

//product management delete/update/

func (pd *productDatabase) DeleteProduct(c context.Context, productid uint) error {
	var product_details domain.ProductDetails
	err := pd.DB.Where("product_id=?", productid).Delete(&product_details)
	if err != nil {
		return errors.New("failed to delete product")
	}
	return nil
}
