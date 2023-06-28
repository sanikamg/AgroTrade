package repository

import (
	"context"
	"errors"
	"fmt"
	"golang_project_ecommerce/pkg/domain"
	interfaces "golang_project_ecommerce/pkg/repository/interface"
	"golang_project_ecommerce/pkg/utils"
	"golang_project_ecommerce/pkg/utils/req"
	"golang_project_ecommerce/pkg/utils/res"

	"github.com/lib/pq"
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
func (c *productDatabase) FindAllCategory(ctx context.Context, pagination utils.Pagination) ([]domain.Category, utils.Metadata, error) {
	var categories []domain.Category
	var totalRecords int64

	db := c.DB.Model(&domain.Category{})

	// Get the total count of records
	if err := db.Count(&totalRecords).Error; err != nil {
		return categories, utils.Metadata{}, err
	}

	// Apply pagination
	db = db.Limit(pagination.Limit()).Offset(pagination.Offset())

	// Fetch categories
	if err := db.Find(&categories).Error; err != nil {
		return categories, utils.Metadata{}, err
	}

	// Compute metadata
	metadata := utils.ComputeMetadata(&totalRecords, &pagination.Page, &pagination.PageSize)

	return categories, metadata, nil
}

func (pd *productDatabase) DeleteCategory(c context.Context, categoryName string) error {
	var categories domain.Category
	err := pd.DB.Where("category_name=?", categoryName).Delete(&categories).Error
	if err != nil {
		return errors.New("failed to delete product")
	}
	return nil
}

func (pd *productDatabase) FindCategoryByName(c context.Context, categoryName string) error {
	var categories domain.Category
	err := pd.DB.Where("category_name=?", categoryName).First(&categories).Error
	if err != nil {
		return errors.New("failed find category")
	}
	return nil
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

//product management delete/update/

func (pd *productDatabase) DeleteProduct(c context.Context, productid uint) error {
	var product_details domain.ProductDetails
	err := pd.DB.Where("product_id=?", productid).Delete(&product_details).Error
	if err != nil {
		return errors.New("failed to delete product")
	}
	return nil
}

func (pd *productDatabase) UpdateProduct(c context.Context, productup req.UpdateProduct) (domain.ProductDetails, error) {
	var product domain.ProductDetails
	query := `update product_details set product_name=?, product_price=?,product_quantity=?,category_id=? where product_id=? `
	err := pd.DB.Raw(query, productup.ProductName, productup.ProductPrice, productup.ProductQuantity, productup.Categoryid, productup.ProductId).Scan(&product).Error
	if err != nil {
		return domain.ProductDetails{}, errors.New("failed to update product details")
	}

	// Retrieve the updated data
	var pro domain.ProductDetails
	err = pd.DB.Where("product_id = ?", productup.ProductId).Find(&pro).Error
	if err != nil {
		return domain.ProductDetails{}, errors.New("failed to fetch updated product_details")
	}
	category, err := pd.GetCategoryByID(c, int(pro.CategoryID))
	if err != nil {
		return domain.ProductDetails{}, errors.New("failed to find category")
	}
	pro.Category = category
	return pro, nil
}

//add image

func (pd *productDatabase) AddImage(c context.Context, pid int, filename string) (domain.Image, error) {

	// Store the image record in the database
	image := domain.Image{ProductId: uint(pid), Image: filename}
	if err := pd.DB.Create(&image).Error; err != nil {

		return domain.Image{}, errors.New("failed to store image record")
	}

	return image, nil
}

func (pd *productDatabase) GetProductByID(c context.Context, productid int) (domain.ProductDetails, error) {
	var product domain.ProductDetails
	err := pd.DB.Where("product_id=?", productid).First(&product).Error
	if err != nil {
		return domain.ProductDetails{}, errors.New("failed to find product")
	}
	return product, nil
}

func (pd *productDatabase) FindProductByName(c context.Context, productname string) error {
	var product domain.ProductDetails
	err := pd.DB.Where("product_name=?", productname).First(&product).Error
	if err != nil {
		return errors.New("failed to find product")
	}
	return nil
}

// func (pd *productDatabase) FindAllProducts(c context.Context, pagination utils.Pagination) ([]domain.ProductResponse, utils.Metadata, error) {
// 	var products []domain.ProductResponse
// 	//var images []string
// 	var totalrecords int64

// 	db := pd.DB.Model(&domain.ProductDetails{})

// 	//count all records
// 	if err := db.Count(&totalrecords).Error; err != nil {
// 		return []domain.ProductResponse{}, utils.Metadata{}, err
// 	}

// 	// Apply pagination
// 	db = db.Limit(pagination.Limit()).Offset(pagination.Offset())

// 	query := `select p.product_id,p.product_name,p.product_price,p.product_quantity c.category_name,p.discount_price from product_details as p inner join categories as c on p.category_id=c.id`

// 	// Fetch categories
// 	err := db.Raw(query).Scan(&products).Error
// 	if err != nil {
// 		return []domain.ProductResponse{}, utils.Metadata{}, errors.New("failed to select details of product")
// 	}

// 	// Compute metadata
// 	metadata := utils.ComputeMetadata(&totalrecords, &pagination.Page, &pagination.PageSize)

// 	return products, metadata, nil

// }
func (pd *productDatabase) FindAllProducts(c context.Context, pagination utils.Pagination) ([]domain.ProductResponse, utils.Metadata, error) {
	var products []domain.ProductResponse

	var totalRecords int64

	db := pd.DB.Model(&domain.ProductDetails{})

	// Count all records
	if err := db.Count(&totalRecords).Error; err != nil {
		return []domain.ProductResponse{}, utils.Metadata{}, err
	}

	// Apply pagination
	//db = db.Limit(pagination.Limit()).Offset(pagination.Offset())

	// Fetch product details and associated images
	// if err := db.
	// 	Preload("Category").
	// 	Joins("LEFT JOIN images ON product_details.product_id = images.product_id").
	// 	Select("p.product_id,p.product_name,p.product_price,p.product_quantity ,c.category_name,p.discount_price from product_details as p inner join categories as c on p.category_id=c.id, array_agg(images.image) as images").
	// 	Group("product_details.product_id").
	// 	Find(&products).
	// 	Error; err != nil {
	// 	return products, utils.Metadata{}, err
	// }

	query := `SELECT p.product_id, p.product_name, p.product_price, p.product_quantity, c.category_name, p.discount_price, array_agg(images.image) as images
	FROM product_details AS p
	INNER JOIN categories AS c ON p.category_id = c.id
	LEFT JOIN images ON p.product_id = images.product_id
	WHERE p.deleted_at IS NULL
	GROUP BY p.product_id, p.product_name, p.product_price, p.product_quantity, c.category_name, p.discount_price LIMIT $1 OFFSET $2;
	`
	rows, err := db.Raw(query, pagination.Limit(), pagination.Offset()).Rows()
	if err != nil {
		return []domain.ProductResponse{}, utils.Metadata{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var product domain.ProductResponse
		var images []string

		err := rows.Scan(
			&product.Product_Id,
			&product.ProductName,
			&product.ProductPrice,
			&product.ProductQuantity,
			&product.Category_name,
			&product.DiscountPrice,
			pq.Array(&images),
		)
		if err != nil {
			return []domain.ProductResponse{}, utils.Metadata{}, err
		}

		product.Image = images
		products = append(products, product)
	}

	// Compute metadata
	metadata := utils.ComputeMetadata(&totalRecords, &pagination.Page, &pagination.PageSize)

	return products, metadata, nil
}

func (pd *productDatabase) FindAllProductsByCategory(c context.Context, pagination utils.Pagination, category string) ([]domain.ProductResponse, utils.Metadata, error) {
	var products []domain.ProductResponse

	var totalRecords int64

	db := pd.DB.Model(&domain.ProductDetails{})

	// Count all records
	if err := db.Count(&totalRecords).Error; err != nil {
		return []domain.ProductResponse{}, utils.Metadata{}, err
	}

	// Apply pagination
	//db = db.Limit(pagination.Limit()).Offset(pagination.Offset())

	// Fetch product details and associated image

	query := `SELECT p.product_id, p.product_name, p.product_price, p.product_quantity, c.category_name, p.discount_price, array_agg(images.image) as images
	FROM product_details AS p
	INNER JOIN categories AS c ON p.category_id = c.id
	LEFT JOIN images ON p.product_id = images.product_id
	WHERE p.deleted_at IS NULL AND c.category_name = $1
	GROUP BY p.product_id, p.product_name, p.product_price, p.product_quantity, c.category_name, p.discount_price LIMIT $2 OFFSET $3;	
	`
	rows, err := db.Raw(query, category, pagination.Limit(), pagination.Offset()).Rows()
	if err != nil {
		return []domain.ProductResponse{}, utils.Metadata{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var product domain.ProductResponse
		var images []string

		err := rows.Scan(
			&product.Product_Id,
			&product.ProductName,
			&product.ProductPrice,
			&product.ProductQuantity,
			&product.Category_name,
			&product.DiscountPrice,
			pq.Array(&images),
		)
		if err != nil {
			return []domain.ProductResponse{}, utils.Metadata{}, err
		}

		product.Image = images
		products = append(products, product)
	}
	fmt.Println(products)

	// Compute metadata
	metadata := utils.ComputeMetadata(&totalRecords, &pagination.Page, &pagination.PageSize)

	return products, metadata, nil
}

func (pd *productDatabase) FindProductsById(c context.Context, pagination utils.Pagination, id int) ([]domain.ProductResponse, utils.Metadata, error) {
	var products []domain.ProductResponse

	var totalRecords int64

	db := pd.DB.Model(&domain.ProductDetails{})

	// Count all records
	if err := db.Count(&totalRecords).Error; err != nil {
		return []domain.ProductResponse{}, utils.Metadata{}, err
	}

	// Apply pagination
	//db = db.Limit(pagination.Limit()).Offset(pagination.Offset())

	// Fetch product details and associated image

	query := `SELECT p.product_id, p.product_name, p.product_price, p.product_quantity, c.category_name, p.discount_price, array_agg(images.image) as images
	FROM product_details AS p
	INNER JOIN categories AS c ON p.category_id = c.id
	LEFT JOIN images ON p.product_id = images.product_id
	WHERE p.deleted_at IS NULL AND p.product_id = $1
	GROUP BY p.product_id, p.product_name, p.product_price, p.product_quantity, c.category_name, p.discount_price LIMIT $2 OFFSET $3;	
	`
	rows, err := db.Raw(query, id, pagination.Limit(), pagination.Offset()).Rows()
	if err != nil {
		return []domain.ProductResponse{}, utils.Metadata{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var product domain.ProductResponse
		var images []string

		err := rows.Scan(
			&product.Product_Id,
			&product.ProductName,
			&product.ProductPrice,
			&product.ProductQuantity,
			&product.Category_name,
			&product.DiscountPrice,
			pq.Array(&images),
		)
		if err != nil {
			return []domain.ProductResponse{}, utils.Metadata{}, err
		}

		product.Image = images
		products = append(products, product)
	}

	// Compute metadata
	metadata := utils.ComputeMetadata(&totalRecords, &pagination.Page, &pagination.PageSize)

	return products, metadata, nil
}

//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>cart>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>

func (pd productDatabase) GetCartByUserID(c context.Context, userId uint) ([]res.CartResponse, error) {
	var cartItems []res.CartResponse
	query := `select product_id,quantity,total_price  from cart_items where user_id=? `
	err := pd.DB.Raw(query, userId).Scan(&cartItems).Error
	if err != nil {
		return []res.CartResponse{}, errors.New("cart has no items")
	}
	return cartItems, nil
}

func (pd *productDatabase) AddToCart(c context.Context, cart domain.Cart_item) (domain.Cart_item, error) {
	err := pd.DB.Create(&cart).Error
	if err != nil {
		return domain.Cart_item{}, errors.New("failed to add prduct")
	}
	return cart, nil
}

func (pd *productDatabase) UpdateCart(c context.Context, cart domain.Cart_item) ([]res.CartResponse, error) {

	query := `update cart_items set quantity=?,total_price=? where product_id=?`
	err := pd.DB.Raw(query, cart.Quantity, cart.Total_Price, cart.Product_Id).Scan(&cart).Error
	if err != nil {
		return []res.CartResponse{}, errors.New("failed to update cart")
	}
	// query1 := `select product_id,quantity,total_price,image from cart_items where user_id=? `
	// err1 := pd.DB.Raw(query1, cart.User_Id).Scan(&cartitem).Error
	// if err1 != nil {
	// 	return res.CartResponse{}, errors.New("cart has no items")
	// }
	cartitem, err := pd.GetCartByUserID(c, cart.User_Id)
	if err != nil {
		return []res.CartResponse{}, err
	}
	return cartitem, nil
}

// to view cart items
func (pd *productDatabase) ListCartItems(c context.Context, pagination utils.Pagination, userid int) ([]res.CartResponse, utils.Metadata, error) {
	var carts []res.CartResponse

	var totalRecords int64

	db := pd.DB.Model(&domain.Cart_item{})

	// Count all records
	if err := db.Count(&totalRecords).Error; err != nil {
		return []res.CartResponse{}, utils.Metadata{}, err
	}

	query := `select c.product_id,c.quantity,c.total_price, array_agg(images.image) from cart_items as c LEFT JOIN images ON c.product_id = images.product_id
	WHERE c.user_id = $1
	GROUP BY c.product_id, c.quantity, c.total_price  LIMIT $2 OFFSET $3;`

	rows, err := db.Raw(query, userid, pagination.Limit(), pagination.Offset()).Rows()
	if err != nil {
		return []res.CartResponse{}, utils.Metadata{}, errors.New("query didn't work")
	}
	defer rows.Close()

	for rows.Next() {
		var cart res.CartResponse
		var images []string

		err := rows.Scan(
			&cart.Product_Id,
			&cart.Quantity,
			&cart.Total_Price,
			pq.Array(&images),
		)
		if err != nil {
			return []res.CartResponse{}, utils.Metadata{}, errors.New("failed to scan")
		}

		cart.Image = images
		carts = append(carts, cart)
	}

	// Compute metadata
	metadata := utils.ComputeMetadata(&totalRecords, &pagination.Page, &pagination.PageSize)

	return carts, metadata, nil
}

// deleteproduct from cart item
func (pd *productDatabase) RemoveProductFromCart(c context.Context, productid uint) error {
	var cartitems domain.Cart_item
	err := pd.DB.Where("product_id=?", productid).Delete(&cartitems).Error
	if err != nil {
		return errors.New("failed to delete product")
	}
	return nil
}

//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>Coupon>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>

// add coupon
func (pd *productDatabase) AddCoupon(c context.Context, coupon domain.Coupon) (domain.Coupon, error) {
	fmt.Println(coupon)
	err := pd.DB.Create(&coupon).Error
	if err != nil {
		return domain.Coupon{}, errors.New(" failed to add coupon")
	}

	return coupon, nil

}

func (pd *productDatabase) FindCoupon(c context.Context, coupon domain.Coupon) error {
	err := pd.DB.Where("coupon=?", coupon.Coupon).First(&coupon).Error
	if err != nil {
		return errors.New("coupon already exist")
	}
	return nil
}

// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>paymentmethod>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
func (pd *productDatabase) AddPaymentMethod(c context.Context, payment domain.PaymentMethod) (domain.PaymentMethod, error) {
	err := pd.DB.Create(&payment).Error
	if err != nil {
		return domain.PaymentMethod{}, errors.New("failed to add payment method")
	}
	return payment, nil
}

func (pd *productDatabase) FindPaymentMethod(c context.Context, payment domain.PaymentMethod) error {
	var payment_methods domain.PaymentMethod
	err := pd.DB.Raw("SELECT * FROM payment_methods WHERE payment_method=?", payment.PaymentMethod).First(&payment_methods).Error
	if err != nil {

		return errors.New("failed to find payment method")
	}
	return nil
}

func (pd *productDatabase) FindPaymentMethodId(c context.Context, method_id uint) (uint, error) {
	var payment_methods domain.PaymentMethod
	err := pd.DB.Raw("SELECT * FROM payment_methods WHERE method_id=?", method_id).First(&payment_methods).Error
	if err != nil {

		return 0, errors.New("failed to find payment method")
	}
	return payment_methods.Method_id, nil
}

func (pd *productDatabase) GetPaymentMethods(c context.Context, pagination utils.Pagination) ([]res.PaymentMethodResponse, utils.Metadata, error) {
	var payments []res.PaymentMethodResponse

	var totalRecords int64

	db := pd.DB.Model(&domain.PaymentMethod{})

	// Count all records
	if err := db.Count(&totalRecords).Error; err != nil {
		return []res.PaymentMethodResponse{}, utils.Metadata{}, err
	}

	query := `select *from payment_methods limit $1 offset $2;`

	err := db.Raw(query, pagination.Limit(), pagination.Offset()).Scan(&payments).Error
	if err != nil {
		return []res.PaymentMethodResponse{}, utils.Metadata{}, errors.New("query didn't work")
	}
	// Compute metadata
	metadata := utils.ComputeMetadata(&totalRecords, &pagination.Page, &pagination.PageSize)

	return payments, metadata, nil
}

func (pd *productDatabase) UpdatePaymentMethod(c context.Context, payment domain.PaymentMethod) (domain.PaymentMethod, error) {

	query := `update payment_methods set maximum_amount=?,payment_method=? where method_id=?`
	err := pd.DB.Raw(query, payment.MaximumAmount, payment.PaymentMethod, payment.Method_id).Scan(&payment).Error
	if err != nil {
		return domain.PaymentMethod{}, errors.New("failed to update payment method details")
	}
	return payment, nil
}

func (pu *productDatabase) DeleteMethod(c context.Context, id uint) error {
	var paymentmethod domain.PaymentMethod
	query := `delete from payment_methods where method_id=?`
	err := pu.DB.Raw(query, id).Scan(&paymentmethod).Error
	if err != nil {
		return errors.New("failed to delete payment method")
	}
	return nil

}

// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>....order.....>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>......
// get total amount
func (pd *productDatabase) GetTotalAmount(c context.Context, userid int) ([]domain.Cart_item, error) {
	var cart []domain.Cart_item
	query := `select *from cart_items where user_id=?`
	err := pd.DB.Raw(query, userid).Scan(&cart).Error
	if err != nil {
		return []domain.Cart_item{}, errors.New("failed to find cart items")
	}
	return cart, nil
}

// insert into order table
func (pd *productDatabase) CreateOrder(c context.Context, order domain.Order) (res.OrderResponse, error) {
	var orderdetails res.OrderResponse
	err := pd.DB.Create(&order).Error
	if err != nil {
		return res.OrderResponse{}, errors.New("failed to place order")
	}

	query := `select o.total_amount,o.order_status,o.address_id,p.payment_method from orders as o left join payment_methods as p on o.payment_method_id=p.method_id where o.order_id=?`
	err1 := pd.DB.Raw(query, order.Order_Id).Scan(&orderdetails).Error
	if err1 != nil {
		return res.OrderResponse{}, errors.New("failed to display order details")
	}
	return orderdetails, nil
}

func (pd *productDatabase) UpdateOrderDetails(c context.Context, uporder req.UpdateOrder) (res.OrderResponse, error) {
	var order domain.Order
	var orderdetails res.OrderResponse
	query := `update orders set payment_method_id=?,address_id=? where order_id=?`
	err := pd.DB.Raw(query, uporder.PaymentMethodID, uporder.Address_Id, uporder.Order_Id).Scan(&order).Error
	if err != nil {
		return res.OrderResponse{}, errors.New("Error while updating order deatails")
	}
	query1 := `select o.total_amount,o.order_status,o.address_id,p.payment_method from orders as o left join payment_methods as p on o.payment_method_id=p.method_id where o.order_id=?`
	err1 := pd.DB.Raw(query1, uporder.Order_Id).Scan(&orderdetails).Error
	if err1 != nil {
		return res.OrderResponse{}, errors.New("failed to display order details")
	}
	return orderdetails, nil
}

func (pd *productDatabase) ListAllOrders(c context.Context, pagination utils.Pagination, usrid uint) ([]res.OrderResponse, utils.Metadata, error) {
	var orders []res.OrderResponse
	var totalRecords int64

	db := pd.DB.Model(&domain.Order{})

	// Count all records
	if err := db.Count(&totalRecords).Error; err != nil {
		return []res.OrderResponse{}, utils.Metadata{}, err
	}

	query := `select o.total_amount,o.order_status,o.address_id,p.payment_method from orders as o left join payment_methods as p on o.payment_method_id=p.method_id where user_id=$1 limit $2 offset $3;`

	err := db.Raw(query, usrid, pagination.Limit(), pagination.Offset()).Scan(&orders).Error
	if err != nil {
		return []res.OrderResponse{}, utils.Metadata{}, errors.New("query didn't work")
	}
	// Compute metadata
	metadata := utils.ComputeMetadata(&totalRecords, &pagination.Page, &pagination.PageSize)

	return orders, metadata, nil
}

func (pd *productDatabase) GetAllOrders(c context.Context, pagination utils.Pagination) ([]res.OrderResponse, utils.Metadata, error) {
	var orders []res.OrderResponse
	var totalRecords int64

	db := pd.DB.Model(&domain.Order{})

	// Count all records
	if err := db.Count(&totalRecords).Error; err != nil {
		return []res.OrderResponse{}, utils.Metadata{}, err
	}

	query := `select o.total_amount,o.order_status,o.address_id,p.payment_method from orders as o left join payment_methods as p on o.payment_method_id=p.method_id limit $1 offset $2;`

	err := db.Raw(query, pagination.Limit(), pagination.Offset()).Scan(&orders).Error
	if err != nil {
		return []res.OrderResponse{}, utils.Metadata{}, errors.New("query didn't work")
	}
	// Compute metadata
	metadata := utils.ComputeMetadata(&totalRecords, &pagination.Page, &pagination.PageSize)

	return orders, metadata, nil
}

func (pd *productDatabase) DeleteOrder(c context.Context, order_id uint) error {
	var order domain.Order
	err := pd.DB.Where("order_id=?").Delete(&order)
	if err != nil {
		return errors.New("faileed to delete orders")

	}
	return nil
}

// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>checkout>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
func (pd *productDatabase) PlaceOrder(c context.Context, order domain.Order) (res.PaymentResponse, error) {
	var paymentresp res.PaymentResponse
	query := `update orders set total_amount=? where order_id=?`
	err := pd.DB.Raw(query, order.Total_Amount, order.Order_Id).Scan(&order).Error
	if err != nil {
		return res.PaymentResponse{}, errors.New("failed to update payment")
	}
	query1 := `select total_amount,order_status,address_id,payment_method_id,payment_status from orders where order_id=?`
	err1 := pd.DB.Raw(query1, order.Order_Id).Scan(&paymentresp).Error
	if err1 != nil {
		return res.PaymentResponse{}, errors.New("failed to display order details")
	}

	return paymentresp, nil
}

func (pd *productDatabase) ValidateCoupon(c context.Context, CouponId uint) (res.CouponResponse, error) {
	var couponResp res.CouponResponse
	query := `select discount,quantity,validity from coupons where coupon_id=?`
	err := pd.DB.Raw(query, CouponId).Scan(&couponResp).Error
	if err != nil {
		return res.CouponResponse{}, errors.New("Not a valid coupon")
	}
	return couponResp, nil
}

func (pd *productDatabase) FindCouponById(c context.Context, couponId uint) error {
	var coupon domain.Coupon
	err := pd.DB.Where("coupon_id=?", couponId).First(&coupon).Error
	if err != nil {
		return errors.New("coupon already exist")
	}
	return nil
}

func (pd *productDatabase) ApplyDiscount(c context.Context, order_id uint) (domain.Order, error) {
	var order domain.Order
	query := `select *from orders where order_id=?`
	err := pd.DB.Raw(query, order_id).Scan(&order).Error
	if err != nil {
		return domain.Order{}, errors.New("failed to find order by order_id")
	}
	return order, nil
}

func (pd *productDatabase) FindPaymentMethodIdByOrderId(c context.Context, order_id uint) (uint, error) {
	var order domain.Order
	err := pd.DB.Raw("SELECT * FROM payment_methods WHERE method_id=?", order_id).First(&order).Error
	if err != nil {

		return 0, errors.New("failed to find payment method id")
	}
	return order.PaymentMethodID, nil
}

func (pd *productDatabase) UpdateOrderStatus(c context.Context, order_id uint, order_status string) (res.OrderResponse, error) {
	var order domain.Order
	var orderResp res.OrderResponse
	query := `update orders set order_status=?  where order_id=?`
	err := pd.DB.Raw(query, order_status, order_id).Scan(&order).Error
	if err != nil {
		return res.OrderResponse{}, errors.New("failed to update order status")
	}
	query1 := `select o.total_amount,o.order_status,o.address_id,p.payment_method from orders as o left join payment_methods as p on o.payment_method_id=p.method_id where o.order_id=?`
	err1 := pd.DB.Raw(query1, order_id).Scan(&orderResp).Error
	if err1 != nil {
		return res.OrderResponse{}, errors.New("failed to display order details")
	}
	return orderResp, nil
}

func (pd *productDatabase) FindTotalAmountByOrderId(c context.Context, order_id uint) (float64, error) {
	var total_amount float64
	query := `SELECT total_amount FROM orders WHERE order_id=?`
	err := pd.DB.Raw(query, order_id).Scan(&total_amount).Error
	if err != nil {
		return 0, errors.New("failed to fetch total amount")
	}
	return total_amount, nil
}

func (pd *productDatabase) FindPhnEmailByUsrId(c context.Context, usr_id int) (res.PhnEmailResp, error) {
	var phnEmail res.PhnEmailResp
	query := `SELECT phone,email FROM users WHERE user_id=?`
	err := pd.DB.Raw(query, usr_id).Scan(&phnEmail).Error
	if err != nil {
		return res.PhnEmailResp{}, errors.New("failed to fetch details")
	}
	return phnEmail, nil
}
