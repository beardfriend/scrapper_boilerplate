package excelx

import (
	"fmt"

	"github.com/pkg/errors"

	_ "image/jpeg"

	_ "image/png"

	"github.com/xuri/excelize/v2"
)

type Excel interface {
	Open(dir string, sheetIndex *int) error
	NewFile()
	NewSheet(sheetName string) error
	Close() error
	Save() error
	SaveAs(filePath string) error
	ChangeSheet(sheetIndex int) error
	SetAlign() Excel
	SetFont(name string) error
	MergeCell(hCell, vCell string) error
	SetCellValue(cell string, value interface{}) error
	SetCellStyle(cellName, cellName2 string, styleID int) error
	SetCellFormula(cell string, formula string) error
	SetHeight(row int, height float64) error
	NewStyle(style *excelize.Style) (int, error)
	SetColStyle(col string, styleID int) error
	SetWidth(startColName, endColName string, width float64) error
	SetBackgroundColor(cell string, color string) error
	SetBackgroundColorRange(startcell, endCell string, color string) error
	SetCellImage(cell string, ext string, imageData []byte) error

	DeleteSheet(sheetName string) error
	DeletePicture(cell string) error

	CalcCellValue(cell string) (result string, err error)

	CopySheet(from int, name string) error

	GetCellValue(cellName string) (string, error)
	GetRows() ([][]string, error)
	GetFormula(cell string) (string, error)
	GetSheetsName() ([]string, error)
	GetImage(cellName string) ([]byte, string, error)

	RemoveRow(row int) error

	UpdateLinkedVal()
	GetModule() *excelize.File
}

type excel struct {
	file      *excelize.File
	sheetName string
	style     *excelize.Style
}

func NewExcel() Excel {
	return &excel{}
}

var (
	ErrCheckFile         = errors.New("please check your file")
	ErrSheetDoesNotExist = errors.New("sheet does not exist")
	ErrFileDoesNotOpen   = errors.New("please check file open")
	ErrFileDoesNotClosed = errors.New("file does not closed")
	ErrFileDoesNotSaved  = errors.New("file does not saved")
)

func (e *excel) GetModule() *excelize.File {
	return e.file
}

func (e *excel) NewFile() {
	e.file = excelize.NewFile()
	sheets, _ := e.GetSheetsName()
	e.sheetName = sheets[0]
}

func (e *excel) NewSheet(sheetName string) error {
	if !e.isOpen() {
		return ErrFileDoesNotOpen
	}
	e.file.NewSheet(sheetName)

	e.sheetName = sheetName

	return nil
}

func (e *excel) SetFont(name string) error {
	return e.file.SetDefaultFont(name)
}

func (e *excel) MergeCell(hCell, vCell string) error {
	return e.file.MergeCell(e.sheetName, hCell, vCell)
}

func (e *excel) UpdateLinkedVal() {
	e.file.UpdateLinkedValue()
}

func (e *excel) SaveAs(filePath string) error {
	return e.file.SaveAs(filePath)
}

func (e *excel) Open(dir string, sheetIndex *int) error {
	file, err := excelize.OpenFile(dir)
	if err != nil {
		return ErrCheckFile
	}

	e.file = file

	sheetNames, _ := e.GetSheetsName()

	// set SheetName default
	e.sheetName = sheetNames[0]

	// set SheetName from parameter
	if sheetIndex != nil {
		index := *sheetIndex

		if index >= len(sheetNames) {
			return ErrSheetDoesNotExist
		}

		e.sheetName = sheetNames[index]
	}

	e.style = &excelize.Style{}

	return nil
}

func (e *excel) Close() error {
	if err := e.file.Close(); err != nil {
		return ErrFileDoesNotClosed
	}

	e.file = nil

	return nil
}

func (e *excel) Save() error {
	if err := e.file.Save(); err != nil {
		return ErrFileDoesNotSaved
	}
	return nil
}

func (e *excel) SetCellValue(cell string, value interface{}) error {
	if !e.isOpen() {
		return ErrFileDoesNotOpen
	}

	if err := e.file.SetCellValue(e.sheetName, cell, value); err != nil {
		return fmt.Errorf("value not set,  detail Error : %x ", err)
	}

	return nil
}

func (e *excel) CalcCellValue(cell string) (result string, err error) {
	return e.file.CalcCellValue(e.sheetName, cell)
}

func (e *excel) GetFormula(cell string) (string, error) {
	return e.file.GetCellFormula(e.sheetName, cell)
}

func (e *excel) SetCellFormula(cell string, formula string) error {
	if !e.isOpen() {
		return ErrFileDoesNotOpen
	}

	if err := e.file.SetCellFormula(e.sheetName, cell, formula); err != nil {
		return fmt.Errorf("value not set,  detail Error : %x ", err)
	}

	return nil
}

func (e *excel) SetHeight(row int, height float64) error {
	return e.file.SetRowHeight(e.sheetName, row, height)
}

func (e *excel) SetWidth(startColName, endColName string, width float64) error {
	return e.file.SetColWidth(e.sheetName, startColName, endColName, width)
}

func (e *excel) NewStyle(style *excelize.Style) (int, error) {
	return e.file.NewStyle(style)
}

func (e *excel) SetColStyle(col string, styleID int) error {
	return e.file.SetColStyle(e.sheetName, col, styleID)
}

func (e *excel) SetCellStyle(cellName, cellName2 string, styleID int) error {
	return e.file.SetCellStyle(e.sheetName, cellName, cellName2, styleID)
}

func (e *excel) SetCellImage(cell string, ext string, imageData []byte) error {
	if !e.isOpen() {
		return ErrFileDoesNotOpen
	}

	if err := e.file.AddPictureFromBytes(e.sheetName, cell, &excelize.Picture{
		Extension: ext,
		File:      imageData,
		Format: &excelize.GraphicOptions{
			OffsetX: 0,
			AutoFit: true,
		},
	}); err != nil {
		return fmt.Errorf("image not added,  detail Error : %x ", err)
	}

	return nil
}

func (e *excel) SetBackgroundColor(cell string, color string) error {
	styleID, err := e.file.NewStyle(&excelize.Style{
		Fill: excelize.Fill{Type: "pattern", Color: []string{color}, Pattern: 1},
		Border: []excelize.Border{
			{Type: "left", Color: "000000", Style: 1},
			{Type: "top", Color: "000000", Style: 1},
			{Type: "bottom", Color: "000000", Style: 1},
			{Type: "right", Color: "000000", Style: 1},
		},
		Font: &excelize.Font{
			Bold: true,
		},
		Alignment: &excelize.Alignment{Vertical: "center", Horizontal: "center"},
	})
	if err != nil {
		return err
	}

	return e.file.SetCellStyle(e.sheetName, cell, cell, styleID)
}

func (e *excel) SetBackgroundColorRange(startcell, endCell string, color string) error {
	styleID, err := e.file.NewStyle(&excelize.Style{
		Fill: excelize.Fill{Type: "pattern", Color: []string{color}, Pattern: 1},
		Border: []excelize.Border{
			{Type: "left", Color: "000000", Style: 1},
			{Type: "top", Color: "000000", Style: 1},
			{Type: "bottom", Color: "000000", Style: 1},
			{Type: "right", Color: "000000", Style: 1},
		},
		Font: &excelize.Font{
			Bold: true,
		},
		Alignment: &excelize.Alignment{Vertical: "center", Horizontal: "center"},
	})
	if err != nil {
		return err
	}

	return e.file.SetCellStyle(e.sheetName, startcell, endCell, styleID)
}

func (e *excel) SetAlign() Excel {
	align := &excelize.Alignment{
		Horizontal: "center",
	}
	e.style.Alignment = align

	return &excel{}
}

func (e *excel) GetRows() ([][]string, error) {
	return e.file.GetRows(e.sheetName, excelize.Options{RawCellValue: true})
}

func (e *excel) GetCellValue(cellName string) (string, error) {
	return e.file.GetCellValue(e.sheetName, cellName)
}

func (e *excel) GetSheetsName() ([]string, error) {
	if !e.isOpen() {
		return nil, ErrFileDoesNotOpen
	}
	return e.file.GetSheetList(), nil
}

func (e *excel) GetImage(cellName string) ([]byte, string, error) {
	if !e.isOpen() {
		return nil, "", ErrFileDoesNotOpen
	}
	pics, err := e.file.GetPictures(e.sheetName, cellName)
	if err != nil {
		return nil, "", err
	}
	fmt.Println(cellName)
	return pics[0].File, pics[0].Extension, nil
}

func (e *excel) RemoveRow(row int) error {
	return e.file.RemoveRow(e.sheetName, row)
}

func (e *excel) DeleteSheet(sheetName string) error {
	if !e.isOpen() {
		return ErrFileDoesNotOpen
	}

	if err := e.file.DeleteSheet(sheetName); err != nil {
		return fmt.Errorf("failed to delete sheet: %v", err)
	}

	return nil
}

func (e *excel) CopySheet(from int, name string) error {
	if !e.isOpen() {
		return ErrFileDoesNotOpen
	}

	index, err := e.file.NewSheet(name)
	if err != nil {
		return err
	}

	if err := e.file.CopySheet(from, index); err != nil {
		return fmt.Errorf("failed to duplicate sheet: %v", err)
	}

	return nil
}

func (e *excel) DeletePicture(cell string) error {
	return e.file.DeletePicture(e.sheetName, cell)
}

func (e *excel) ChangeSheet(sheetIndex int) error {
	if !e.isOpen() {
		return ErrFileDoesNotOpen
	}

	sheetNames, _ := e.GetSheetsName()

	if sheetIndex >= len(sheetNames) {
		return ErrSheetDoesNotExist
	}

	e.sheetName = sheetNames[sheetIndex]

	return nil
}

func (e *excel) isOpen() bool {
	return e.file != nil
}
