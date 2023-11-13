package stockouts

import (
	stockouts "backend-golang/businesses/stock_outs"
	"fmt"

	"backend-golang/controllers"
	"backend-golang/controllers/stock_outs/request"
	"backend-golang/controllers/stock_outs/response"

	"net/http"

	"github.com/xuri/excelize/v2"

	"github.com/labstack/echo/v4"
)

type StockOutController struct {
	stockUseCase stockouts.Usecase
}

func NewStockOutController(authUC stockouts.Usecase) *StockOutController {
	return &StockOutController{
		stockUseCase: authUC,
	}
}

func (sc *StockOutController) GetAll(c echo.Context) error {
	ctx := c.Request().Context()

	categoriesData, err := sc.stockUseCase.GetAll(ctx)

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, http.StatusInternalServerError, true, "failed to fetch data", "")
	}

	categories := []response.StockOut{}

	for _, category := range categoriesData {
		categories = append(categories, response.FromDomain(category))
	}

	return controllers.NewResponse(c, http.StatusOK, http.StatusOK, false, "all categories", categories)
}

func (cc *StockOutController) GetByID(c echo.Context) error {
	ctx := c.Request().Context()

	stockID := c.Param("id")

	stock, err := cc.stockUseCase.GetByID(ctx, stockID)

	if err != nil {
		return controllers.NewResponse(c, http.StatusNotFound, http.StatusNotFound, true, "category not found", "")
	}

	return controllers.NewResponse(c, http.StatusOK, http.StatusOK, false, "category found", response.FromDomain(stock))
}

func (sc *StockOutController) StockOut(c echo.Context) error {
	input := request.StockOut{}
	ctx := c.Request().Context()

	if err := c.Bind(&input); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, http.StatusBadRequest, true, "invalid request", "")
	}

	err := input.Validate()

	if err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, http.StatusBadRequest, true, "invalid request", "")
	}

	stock, err := sc.stockUseCase.StockOut(ctx, input.ToDomain())

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, http.StatusInternalServerError, true, "failed to add a stock", "")
	}

	return controllers.NewResponse(c, http.StatusCreated, http.StatusCreated, false, "stock registered", response.FromDomain(stock))
}

func (sc *StockOutController) ExportToExcel(c echo.Context) error {
	ctx := c.Request().Context()
	// var stock _dbStocks.Stock

	// Buat file Excel baru
	f := excelize.NewFile()
	sheetName := "DataStok"
	f.SetSheetName("Sheet1", sheetName)

	categoriesData, err := sc.stockUseCase.ExportToExcel(ctx)
	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, http.StatusInternalServerError, true, "failed to fetch data", "")
	}

	// Contoh data stok
	stoks := []response.StockOut{
		// {Stock_Location: "Lokasi1", Stock_Code: "ABC123", Stock_Name: "Produk Anak", Stock_Unit: "pcs", Stock_Out: 10, Stock_Total: 100, StockID: 1},
		// {Stock_Location: "Lokasi2", Stock_Code: "DEF456", Stock_Name: "Produk Banyak", Stock_Unit: "pcs", Stock_Out: 5, Stock_Total: 50, StockID: 2},
		// // Tambahkan data stok lainnya sesuai kebutuhan
	}

	for _, category := range categoriesData {
		stoks = append(stoks, response.FromDomain(category))
	}

	// Tambahkan header ke lembar Excel
	headers := []string{"No", "Tanggal Keluar", "Lokasi", "Kode", "Nama Barang", "Unit", "Barang Keluar", "Total Barang", "ID"}
	err = f.SetSheetRow(sheetName, "A1", &headers)
	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("Error joining cell name: %s", err))
	}

	// Tambahkan data stok ke lembar Excel
	for i, stok := range stoks {
		rowData := []interface{}{stok.ID, stok.CreatedAt, stok.Stock_Location, stok.Stock_Code, stok.Stock_Name, stok.Stock_Unit, stok.Stock_Out, stok.Stock_Total, stok.StockID}
		startCell, err := excelize.JoinCellName("A", i+2) // Mulai dari baris kedua
		if err != nil {
			return c.String(http.StatusInternalServerError, fmt.Sprintf("Error joining cell name: %s", err))
		}
		if err := f.SetSheetRow(sheetName, startCell, &rowData); err != nil {
			return c.String(http.StatusInternalServerError, fmt.Sprintf("Error setting sheet row: %s", err))
		}
	}

	for colIndex, header := range headers {
		colName, err := excelize.ColumnNumberToName(colIndex + 1) // Kolom ke-(colIndex+1) untuk header
		if err != nil {
			return c.String(http.StatusInternalServerError, fmt.Sprintf("Error converting column number: %s", err))
		}

		maxColWidth := len(header) + 2 // Panjang maksimum awal adalah panjang header + 2 untuk memberikan ruang

		for i := 0; i < len(stoks)+1; i++ {
			cellValue, err := f.GetCellValue(sheetName, fmt.Sprintf("%s%d", colName, i+1))
			if err == nil && len(cellValue)+2 > maxColWidth {
				maxColWidth = len(cellValue) + 2
			}
		}

		f.SetColWidth(sheetName, colName, colName, float64(maxColWidth))
	}

	if err := f.SaveAs("DataStok.xlsx"); err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("Error setting headers: %s", err))
	}

	// Atur header untuk respons HTTP
	c.Response().Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.Response().Header().Set("Content-Disposition", "attachment; filename=data_stok.xlsx")

	// Salin file Excel ke response writer
	if err := f.Write(c.Response().Writer); err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("Error writing Excel file: %s", err))
	}
	return nil
}
