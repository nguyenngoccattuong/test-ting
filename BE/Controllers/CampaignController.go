package Controllers

import (
	"net/http"
	"strconv"
	"wan-api-kol-event/Const"
	"wan-api-kol-event/Logic"
	"wan-api-kol-event/ViewModels"

	"github.com/gin-gonic/gin"
)

func GetKolsController(context *gin.Context) {
	var KolsVM ViewModels.KolViewModel

	pageIndex, err := strconv.ParseInt(context.DefaultQuery("pageIndex", "1"), 10, 64)
	if err != nil || pageIndex < 1 {
		KolsVM.Result = Const.UnSuccess
		KolsVM.ErrorMessage = "Invalid pageIndex parameter"
		context.JSON(http.StatusBadRequest, KolsVM)
		return
	}

	pageSize, err := strconv.ParseInt(context.DefaultQuery("pageSize", "100"), 10, 64)
	if err != nil || pageSize < 1 {
		KolsVM.Result = Const.UnSuccess
		KolsVM.ErrorMessage = "Invalid pageSize parameter"
		context.JSON(http.StatusBadRequest, KolsVM)
		return
	}

	kols, totalCount, error := Logic.GetKolLogic(pageIndex, pageSize)
	if error != nil {
		KolsVM.Result = Const.UnSuccess
		KolsVM.ErrorMessage = error.Error()
		KolsVM.PageIndex = pageIndex
		KolsVM.PageSize = pageSize
		context.JSON(http.StatusInternalServerError, KolsVM)
		return
	}

	// Return successful response
	KolsVM.Result = Const.Success
	KolsVM.ErrorMessage = ""
	KolsVM.PageIndex = pageIndex
	KolsVM.PageSize = pageSize
	KolsVM.TotalCount = totalCount
	KolsVM.KolInformation = kols
	context.JSON(http.StatusOK, KolsVM)
}
