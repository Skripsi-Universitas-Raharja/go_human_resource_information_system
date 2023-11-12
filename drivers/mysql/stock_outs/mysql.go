package stockouts

import (
	stockouts "backend-golang/businesses/stock_outs"
	"context"
	"fmt"

	_dbStocks "backend-golang/drivers/mysql/stocks"

	"gorm.io/gorm"
)

type stockOutRepository struct {
	conn *gorm.DB
}

func NewMySQLRepository(conn *gorm.DB) stockouts.Repository {
	return &stockOutRepository{
		conn: conn,
	}
}

func (ur *stockOutRepository) GetByID(ctx context.Context, id string) (stockouts.Domain, error) {
	var stockout StockOut

	if err := ur.conn.WithContext(ctx).First(&stockout, "id = ?", id).Error; err != nil {
		return stockouts.Domain{}, err
	}

	return stockout.ToDomain(), nil

}

func (sr *stockOutRepository) StockOut(ctx context.Context, stockOutDomain *stockouts.Domain) (stockouts.Domain, error) {
	record := FromDomain(stockOutDomain)

	// Buat rekam stok pada riwayat
	result := sr.conn.WithContext(ctx).Create(&record)
	if err := result.Error; err != nil {
		return stockouts.Domain{}, err
	}

	// Perbarui total stok di tabel stok
	var stock _dbStocks.Stock

	if err := sr.conn.WithContext(ctx).First(&stock, "id = ?", record.StockID).Error; err != nil {
		return stockouts.Domain{}, err
	}

	stock.Stock_Total -= record.Stock_Out
	fmt.Println("StockTotal after addition:", stock.Stock_Total)

	// Simpan total stok yang diperbarui di tabel stok
	if err := sr.conn.WithContext(ctx).Save(&stock).Error; err != nil {
		return stockouts.Domain{}, err
	}

	// Perbarui total stok pada rekam riwayat stok
	record.Stock_Location = stock.Stock_Location
	record.Stock_Code = stock.Stock_Code
	record.Stock_Name = stock.Stock_Name
	record.Stock_Unit = stock.Unit
	// record.StockInTransactions.Stock_Total = stock.Stock_Total
	record.Stock_Total = stock.Stock_Total

	fmt.Println("Record StockTotal:", record.Stock_Total)

	// Simpan total stok yang diperbarui di tabel riwayat stok in
	if err := sr.conn.WithContext(ctx).Save(&record).Error; err != nil {
		return stockouts.Domain{}, err
	}

	return record.ToDomain(), nil

}

func (sr *stockOutRepository) ExportToExcel(ctx context.Context) ([]stockouts.Domain, error) {
	var records []StockOut
	if err := sr.conn.WithContext(ctx).Find(&records).Error; err != nil {
		return nil, err
	}

	categories := []stockouts.Domain{}

	for _, category := range records {
		categories = append(categories, category.ToDomain())
	}

	return categories, nil
}

// func (sr *stockOutRepository) ExportToExcel(ctx context.Context, filename string) error {
// 	var records []StockOut

// 	if err := sr.conn.WithContext(ctx).Find(&records).Error; err != nil {
// 		return nil
// 	}

// 	// Buat file Excel baru
// 	f := excelize.NewFile()

// 	// Tambahkan header
// 	for col, header := range []string{"ID", "Created At", "Deleted At", "Stock Location", "Stock Code", "Stock Name", "Stock Unit", "Stock Out", "Stock Total", "Stock ID"} {
// 		cell := string('A'+rune(col)) + "1"
// 		f.SetCellValue("Sheet1", cell, header)
// 	}

// 	// Tambahkan data
// 	for row, stockOut := range records {
// 		f.SetCellValue("Sheet1", fmt.Sprintf("A%d", row+2), stockOut.ID)
// 		f.SetCellValue("Sheet1", fmt.Sprintf("B%d", row+2), stockOut.CreatedAt)
// 		f.SetCellValue("Sheet1", fmt.Sprintf("C%d", row+2), stockOut.DeletedAt)
// 		f.SetCellValue("Sheet1", fmt.Sprintf("D%d", row+2), stockOut.Stock_Location)
// 		f.SetCellValue("Sheet1", fmt.Sprintf("E%d", row+2), stockOut.Stock_Code)
// 		f.SetCellValue("Sheet1", fmt.Sprintf("F%d", row+2), stockOut.Stock_Name)
// 		f.SetCellValue("Sheet1", fmt.Sprintf("G%d", row+2), stockOut.Stock_Unit)
// 		f.SetCellValue("Sheet1", fmt.Sprintf("H%d", row+2), stockOut.Stock_Out)
// 		f.SetCellValue("Sheet1", fmt.Sprintf("I%d", row+2), stockOut.Stock_Total)
// 		f.SetCellValue("Sheet1", fmt.Sprintf("J%d", row+2), stockOut.StockID)
// 	}

// 	// Simpan ke file Excel
// 	if err := f.SaveAs(filename); err != nil {
// 		return nil
// 	}

// 	return nil
// }
