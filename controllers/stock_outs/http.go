package stockouts

import (
	stockouts "backend-golang/businesses/stock_outs"
	"fmt"
	"time"

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

	// Buat file Excel baru
	f := excelize.NewFile()
	sheetName := "DataStok"
	f.SetSheetName("Sheet1", sheetName)

	categoriesData, err := sc.stockUseCase.ExportToExcel(ctx)
	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, http.StatusInternalServerError, true, "failed to fetch data", "")
	}

	// Contoh data stok
	// {Stock_Location: "Lokasi2", Stock_Code: "DEF456", Stock_Name: "Produk Banyak", Stock_Unit: "pcs", Stock_Out: 5, Stock_Total: 50, StockID: 2},
	stoks := []response.StockOut{}

	for _, category := range categoriesData {
		stoks = append(stoks, response.FromDomain(category))
	}

	// Tambahkan judul ke lembar Excel
	companyTitle := "PT. Anugrah Hadi Electric"
	titleHeader := fmt.Sprintf("Data Stok - Rekapitulasi Barang Keluar Tahun %d", time.Now().Year())
	err = f.SetCellValue(sheetName, "A1", companyTitle)
	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, http.StatusInternalServerError, true, "Error setting title header: %s", err)

	}
	err = f.SetCellValue(sheetName, "A2", titleHeader)
	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, http.StatusInternalServerError, true, "Error setting title header: %s", err)
	}

	// Merge dan atur gaya sel untuk judul agar berada di tengah
	titleStyle, err := f.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{Horizontal: "center", Vertical: "center"},
		Font:      &excelize.Font{Bold: true, Size: 14},
	})
	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, http.StatusInternalServerError, true, "Error creating title style: %s", err)

	}
	f.SetCellStyle(sheetName, "A1", "I1", titleStyle)
	f.MergeCell(sheetName, "A1", "I1")
	f.SetCellStyle(sheetName, "A2", "I2", titleStyle)
	f.MergeCell(sheetName, "A2", "I2")

	// Tambahkan header tabel ke lembar Excel
	headers := []string{"No", "Tanggal Keluar Barang", "Lokasi Barang", "Kode Barang", "Nama Barang", "Unit", "Barang Keluar", "Total Barang", "ID"}
	err = f.SetSheetRow(sheetName, "A3", &headers)
	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, http.StatusInternalServerError, true, "Error setting sheet header: %s", err)

	}

	// Tambahkan data stok ke lembar Excel
	for i, stok := range stoks {
		//waktu dalam format "MM/DD/YYYY HH:mm:ss AM/PM".
		rowData := []interface{}{stok.ID, stok.CreatedAt.Format("01/02/2006 03:04:05 PM"), stok.Stock_Location, stok.Stock_Code, stok.Stock_Name, stok.Stock_Unit, stok.Stock_Out, stok.Stock_Total, stok.StockID}
		startCell, err := excelize.JoinCellName("A", i+4) // Mulai dari baris ke-empat
		if err != nil {
			return controllers.NewResponse(c, http.StatusInternalServerError, http.StatusInternalServerError, true, "Error joining cell name: %s", err)

		}
		if err := f.SetSheetRow(sheetName, startCell, &rowData); err != nil {
			return controllers.NewResponse(c, http.StatusInternalServerError, http.StatusInternalServerError, true, "Error setting sheet row: %s", err)
		}

		// Tambahkan border ke sel data
		endCell, err := excelize.JoinCellName("I", i+3) // Kolom I untuk data
		if err != nil {
			return controllers.NewResponse(c, http.StatusInternalServerError, http.StatusInternalServerError, true, "Error joining cell name: %s", err)
		}

		borderStyle, err := f.NewStyle(&excelize.Style{Border: []excelize.Border{
			{Type: "left", Color: "000000", Style: 1},
			{Type: "top", Color: "000000", Style: 1},
			{Type: "bottom", Color: "000000", Style: 1},
			{Type: "right", Color: "000000", Style: 1},
		}})
		if err != nil {
			return controllers.NewResponse(c, http.StatusInternalServerError, http.StatusInternalServerError, true, "Error creating border style: %s", err)

		}

		f.SetCellStyle(sheetName, startCell, endCell, borderStyle)
	}

	// Tambahkan alamat di bawah tabel
	address := "Alamat : Jl. Sriwijaya III No.9, Perumnas 3, Kec. Karawaci, Kabupaten Tangerang, Banten 15810"
	addressCell, err := excelize.JoinCellName("A", len(stoks)+5) // Gantilah "5" sesuai dengan jumlah baris data stok
	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, http.StatusInternalServerError, true, "Error joining cell name: %s", err)
	}
	err = f.SetCellValue(sheetName, addressCell, address)
	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, http.StatusInternalServerError, true, "Error setting address: %s", err)
	}

	// if err := f.SaveAs("DataStok.xlsx"); err != nil {
	// 	return controllers.NewResponse(c, http.StatusInternalServerError, http.StatusInternalServerError, true, "Error setting header: %s", err)
	// }

	// Atur header untuk respons HTTP
	c.Response().Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.Response().Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s.xlsx", titleHeader))

	// Salin file Excel ke response writer
	if err := f.Write(c.Response().Writer); err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("Error writing Excel file: %s", err))
	}
	return nil
}
