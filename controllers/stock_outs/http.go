package stockouts

import (
	stockouts "backend-golang/businesses/stock_outs"
	"fmt"
	"sort"
	"strings"
	"time"

	"backend-golang/controllers"
	"backend-golang/controllers/stock_outs/request"
	"backend-golang/controllers/stock_outs/response"

	"net/http"

	"github.com/xuri/excelize/v2"

	// _dbStocks "backend-golang/drivers/mysql/stock_outs"

	mapset "github.com/deckarep/golang-set"

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

	var dataset [][]string

	for _, stock := range categoriesData {
		dataset = append(dataset, []string{stock.Stock_Name})
		// dataset = append(dataset, []string{stock.Stock_Out})
		// dataset = append(dataset, []string{strconv.Itoa(stock.Stock_Out)})
	}

	minimumSupport := 0.4
	minConfidence := 0.2

	result := apriori(dataset, minimumSupport, minConfidence)
	// fmt.Println("ini isi dataset", dataset)
	sort.Slice(result, func(i, j int) bool {
		return result[i].Support > result[j].Support
	})

	// Menampilkan hasil
	for _, item := range result {
		percentage := item.Support * 100
		fmt.Printf("%s - Support: %.2f%%\n", strings.Join(item.Items, ", "), percentage)
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

type ItemSet struct {
	Items   []string
	Support float64
}

func apriori(dataset [][]string, support float64, confidence float64) []ItemSet {

	elements := elements(dataset)
	freqSet := make(map[string]float64)
	largeSet := make(map[int]mapset.Set)
	oneCSet := returnItemsWithMinSupport(elements, dataset, support, &freqSet)
	currentLSet := oneCSet

	k := 2

	var result []ItemSet

	for currentLSet.Cardinality() != 0 {
		largeSet[k-1] = currentLSet
		currentLSet = joinSet(currentLSet, k)
		currentCSet := returnItemsWithMinSupport(currentLSet, dataset, support, &freqSet)
		currentLSet = currentCSet
		k = k + 1
	}

	// Mengumpulkan hasil ke dalam bentuk ItemSet
	for _, sets := range largeSet {
		for _, item := range sets.ToSlice() {
			itemSet := ItemSet{
				Items:   strings.Split(item.(string), "-"),
				Support: freqSet[item.(string)] / float64(len(dataset)),
			}
			result = append(result, itemSet)
		}
	}
	fmt.Println(largeSet)
	// Calculate stock percentages using the new function
	// percentages := calculateStockPercentages(dataset)
	// fmt.Println(percentages)
	return result
}

func returnItemsWithMinSupport(itemSet mapset.Set, dataset [][]string, minSupport float64, freqSet *map[string]float64) mapset.Set {

	localItemSet := mapset.NewSet()
	localSet := make(map[string]float64)

	for _, item := range itemSet.ToSlice() {
		dkey := strings.Split(item.(string), "-")
		sort.Strings(dkey)
		for _, line := range dataset {
			if contains(line, dkey) {
				key := strings.Join(dkey, "-")
				(*freqSet)[key] += 1.0
				localSet[key] += 1.0
			}
		}
	}

	for item, count := range localSet {
		support := count / float64(len(dataset))

		if support >= minSupport {
			localItemSet.Add(item)
		}
	}

	return localItemSet

}

func joinSet(itemSet mapset.Set, length int) mapset.Set {

	ret := mapset.NewSet()

	for _, i := range itemSet.ToSlice() {
		for _, j := range itemSet.ToSlice() {
			i := i.(string)
			j := j.(string)

			i_a := strings.Split(i, "-")
			j_a := strings.Split(j, "-")

			dkey := (union(i_a, j_a))
			if len(dkey) == length {
				sort.Strings(dkey)
				key := strings.Join(dkey, "-")
				ret.Add(key)

			}
		}
	}
	return ret
}

func union(a []string, b []string) []string {

	ret := mapset.NewSet()

	for _, v := range a {
		ret.Add(v)
	}
	for _, v := range b {
		ret.Add(v)
	}
	rets := []string{}
	for _, v := range ret.ToSlice() {
		rets = append(rets, v.(string))
	}
	return rets
}

func elements(dataset [][]string) mapset.Set {

	ret := mapset.NewSet()

	for i := 0; i < len(dataset); i++ {
		for j := 0; j < len(dataset[i]); j++ {
			if ret.Contains(dataset[i][j]) == false {
				ret.Add(dataset[i][j])
			}
		}
	}
	return ret
}

func contains_dataset(s [][]string, e []string) bool {
	ret := false
	for _, v := range s {
		ret = contains(v, e)
		if ret == true {
			break
		}
	}
	return ret
}

func contains_element(s []string, e string) bool {
	ret := false
	for _, a := range s {
		if a == e {
			ret = true
			break
		}
	}
	return ret
}

func contains(s []string, e []string) bool {
	count := 0
	if len(s) < len(e) {
		return false
	}
	mm := make(map[string]bool)
	for _, a := range e {
		mm[a] = true
	}

	for _, a := range s {
		if _, ok := mm[a]; ok {
			count += 1
		}
	}
	return count == len(e)
}

func (sc *StockOutController) GetAprioriResults(c echo.Context) error {
	ctx := c.Request().Context()

	dataset, err := sc.stockUseCase.ExportToExcel(ctx)
	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, http.StatusInternalServerError, true, "failed to fetch data", "")
	}

	var datasetStrings [][]string
	for _, stock := range dataset {
		datasetStrings = append(datasetStrings, []string{stock.Stock_Name})
	}

	minimumSupport := 0.4
	minConfidence := 0.2

	result := apriori(datasetStrings, minimumSupport, minConfidence)
	sort.Slice(result, func(i, j int) bool {
		return result[i].Support > result[j].Support
	})

	// Menampilkan hasil
	var aprioriResults []string
	for _, item := range result {
		percentage := item.Support * 100
		aprioriResult := fmt.Sprintf("%s - Support: %.2f%%", strings.Join(item.Items, ", "), percentage)
		aprioriResults = append(aprioriResults, aprioriResult)
	}

	return controllers.NewResponse(c, http.StatusOK, http.StatusOK, false, "Apriori results", aprioriResults)
}
