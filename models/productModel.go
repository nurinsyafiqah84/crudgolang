// productModel.go
package models

import (
	"database/sql"
	"fmt"
	"goWeb/config"
	"goWeb/entities"
)

type ProductModel struct {
	Db *sql.DB
}

// list kan item
//1st way
//	func (productModel ProductModel) FindAll() (product []entities.Product, err error) {
//		rows, err := productModel.Db.Query("select * from product")
//		if err != nil {
//			return nil, err
//		} else {
//			var products []entities.Product
//			for rows.Next() {
//				var id int64
//				var name string
//				var price float64
//				var quantity int64
//				var description string
//				err2 := rows.Scan(&id, &name, &price, &quantity, &description)
//				if err2 != nil {
//					return nil, err2
//				} else {
//					product := entities.Product{
//						Id:          id,
//						Name:        name,
//						Price:       price,
//						Quantity:    quantity,
//						Description: description,
//					}
//					products = append(products, product)
//				}
//			}
//			return products, nil
//		}
//	}

// 2nd way
/*func (*ProductModel) FindAll() (product []entities.Product, err error) {
	db, err := config.GetDB()
	if err != nil {
		return nil, err
	} else {
		rows, err2 := db.Query("select * from product")
		if err != nil {
			return nil, err2
		} else {
			var products []entities.Product
			for rows.Next() {
				var product entities.Product
				rows.Scan(&product.Id, &product.Name, &product.Price, &product.Quantity, &product.Description)
				products = append(products, product)
			}
			return products, nil
		}
	}
}*/

// 3rd ways using gorm
func (*ProductModel) FindAll() (product []entities.Product, err error) {
	var products []entities.Product
	config.DB.Find(&products)
	return products, nil
}

func (*ProductModel) Find(id int64) (entities.Product, error) {
	/*db, err := config.GetDB()
	if err != nil {
		return entities.Product{}, err
	} else {
		rows, err2 := db.Query("select * from product where id = ?", id)
		if err != nil {
			return entities.Product{}, err2
		} else {
			var product entities.Product
			for rows.Next() {
				rows.Scan(&product.Id, &product.Name, &product.Price, &product.Quantity, &product.Description)
			}
			fmt.Println(product)
			return product, nil
		}
	}*/
	fmt.Println(id)
	var product entities.Product
	config.DB.Where("id = ?", id).Find(&product)
	return product, nil

}

func (*ProductModel) Create(product *entities.Product) bool {
	/*db, err := config.GetDB()
	if err != nil {
		return false
	} else {
		result, err2 := db.Exec("insert into product(name, price, quantity, description) values(?,?,?,?)", product.Name, product.Price, product.Quantity, product.Description)
		if err2 != nil {
			return false
		} else {
			rowsAffected, _ := result.RowsAffected()
			return rowsAffected > 0
		}
	}*/
	err := config.DB.Create(&product)
	if err != nil {
		return false
	} else {
		return true
	}

}

func (*ProductModel) Update(product entities.Product) bool {
	/*db, err := config.GetDB()
	if err != nil {
		return false
	} else {
		result, err2 := db.Exec("update product set name = ?, price = ?, quantity = ?, description = ? where id = ?", product.Name, product.Price, product.Quantity, product.Description, product.Id)
		if err2 != nil {
			return false
		} else {
			rowsAffected, _ := result.RowsAffected()
			return rowsAffected > 0
		}
	}*/

	if config.DB.Updates(&product).RowsAffected == 0 {
		return false
	} else {
		return true
	}
}

func (*ProductModel) Delete(id int64) bool {
	/*db, err := config.GetDB()
	if err != nil {
		return false
	} else {
		result, err2 := db.Exec("delete from product where id=?", id)
		if err2 != nil {
			return false
		} else {
			rowsAffected, _ := result.RowsAffected()
			return rowsAffected > 0
		}
	}*/
	fmt.Println(id)
	var product entities.Product
	if config.DB.Delete(&product, id).RowsAffected == 0 {
		return false
	} else {
		return true
	}
}
