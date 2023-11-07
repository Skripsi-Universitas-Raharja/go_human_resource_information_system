package stocks

import (
	"backend-golang/businesses/stocks"
	// "fmt"

	// "backend-golang/drivers/mysql/stock_outs"
	// "github.com/skip2/go-qrcode"
	// "fmt"
	"context"

	// "github.com/skip2/go-qrcode"
	"gorm.io/gorm"
)

type stockRepository struct {
	conn *gorm.DB
}

func NewMySQLRepository(conn *gorm.DB) stocks.Repository {
	return &stockRepository{
		conn: conn,
	}
}

func (ur *stockRepository) GetByID(ctx context.Context, id string) (stocks.Domain, error) {
	var stock Stock

	if err := ur.conn.WithContext(ctx).First(&stock, "id = ?", id).Error; err != nil {
		return stocks.Domain{}, err
	}

	return stock.ToDomain(), nil

}

func (cr *stockRepository) Create(ctx context.Context, stockDomain *stocks.Domain) (stocks.Domain, error) {
	// var stockOuts stockouts.StockOut

	// stockOuts.Stock_Location = stockDomain.Stock_Location
	// stockOuts.Stock_Name = stockDomain.Stock_Name
	// stockOuts.Unit = stockDomain.Unit
	// stockOuts.Stock_Out = stockDomain.Stock_Out
	// stockOuts.Stock_Total = stockDomain.Stock_Total

	record := FromDomain(stockDomain)

	// qrCode, err := qrcode.New(fmt.Sprintf("ItemID: %d, Code: %s", newItem.ID, newItem.Code), qrcode.Medium)
	// if err != nil {
	// 	return c.String(http.StatusInternalServerError, "Error generating QR code")
	// }

	result := cr.conn.WithContext(ctx).Create(&record)

	if err := result.Error; err != nil {
		return stocks.Domain{}, err
	}

	if err := result.Last(&record).Error; err != nil {
		return stocks.Domain{}, err
	}

	return record.ToDomain(), nil

}

func (ur *stockRepository) DownloadBarcodeByID(ctx context.Context, id string) (stocks.Domain, error) {
	var stock Stock

	if err := ur.conn.WithContext(ctx).First(&stock, "id = ?", id).Error; err != nil {
		return stocks.Domain{}, err
	}

	// // Generate QR code
	// qrCodeData := fmt.Sprintf("ItemID: %s, Code: %s", stock.ID, stock.Stock_Code)
	// qrCode, err := qrcode.New(qrCodeData, qrcode.Medium)

	// if err != nil {
	// 	return stocks.Domain{}, err
	// }

	// // Kirim QR code sebagai respon
	// imageBytes, err := qrCode.PNG(256)
	// if err != nil {
	// 	return stocks.Domain{}, err
	// }

	return stock.ToDomain(), nil

}

func (cr *stockRepository) StockIn(ctx context.Context, stockDomain *stocks.Domain, id string) (stocks.Domain, error) {
	stock, err := cr.GetByID(ctx, id)

	if err != nil {
		return stocks.Domain{}, err
	}

	updateStock := FromDomain(&stock)

	updateStock.Stock_Total += stockDomain.Stock_In

	if err := cr.conn.WithContext(ctx).Save(&updateStock).Error; err != nil {
		return stocks.Domain{}, err
	}

	return updateStock.ToDomain(), nil
}

func (cr *stockRepository) StockOut(ctx context.Context, stockDomain *stocks.Domain, id string) (stocks.Domain, error) {
	stock, err := cr.GetByID(ctx, id)

	if err != nil {
		return stocks.Domain{}, err
	}

	updateStock := FromDomain(&stock)

	updateStock.Stock_Total -= stockDomain.Stock_Out

	if err := cr.conn.WithContext(ctx).Save(&updateStock).Error; err != nil {
		return stocks.Domain{}, err
	}

	return updateStock.ToDomain(), nil
}
