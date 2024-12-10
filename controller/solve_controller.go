package controller

/*
#cgo LDFLAGS: -L../ -ltest_ffi_func
#include <stdlib.h>

extern char* solve(int numChamps, int highTier, char* augment, double tierCoefficient, char* champions, char* traits);
extern void free_rust_string(char* str);
*/
import "C"
import (
	"encoding/json"
	"net/http"
	"strconv"
	"tft_su_bd_backend/model"
	"tft_su_bd_backend/usecase"
	"unsafe"

	"github.com/gin-gonic/gin"
)

type SolveController struct {
	usecase usecase.SolveUseCase
}

func NewSolveController(usecase usecase.SolveUseCase) SolveController {
	return SolveController{
		usecase,
	}
}

func (pc *SolveController) GetSolution(ctx *gin.Context) {
	numChampsParam := ctx.Param("numChamps")
	highTierQuery := ctx.DefaultQuery("high_tier", "false")
	augment := ctx.DefaultQuery("augment", "standUnited")
	tierCoefficientQuery := ctx.DefaultQuery("tier_coefficient", "1.0")

	numChamps, err := strconv.Atoi(numChampsParam)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response[*struct{}]{
			Status:  http.StatusBadRequest,
			Message: "Invalid Number of champions",
			Data:    nil,
		})
		return
	}

	highTier, err := strconv.ParseBool(highTierQuery)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response[*struct{}]{
			Status:  http.StatusBadRequest,
			Message: "Invalid 'High Tier' value",
			Data:    nil,
		})
		return
	}

	tierCoefficient, err := strconv.ParseFloat(tierCoefficientQuery, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response[*struct{}]{
			Status:  http.StatusBadRequest,
			Message: "Invalid tier coefficient",
			Data:    nil,
		})
		return
	}

	if augment != "standUnited" && augment != "builtDifferent" {
		ctx.JSON(http.StatusBadRequest, model.Response[*struct{}]{
			Status:  http.StatusBadRequest,
			Message: "Invalid augment",
			Data:    nil,
		})
		return
	}

	dataToSolve, err := pc.usecase.GetDataToSolve(augment)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.Response[*struct{}]{
			Status: http.StatusInternalServerError,
			Data:   nil,
		})
		return
	}

	championsJson, err := json.Marshal(dataToSolve.Champions)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.Response[*struct{}]{
			Status: http.StatusInternalServerError,
			Data:   nil,
		})
		return
	}

	traitsJson, err := json.Marshal(dataToSolve.Traits)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.Response[*struct{}]{
			Status: http.StatusInternalServerError,
			Data:   nil,
		})
		return
	}

	resultString := solve(numChamps, highTier, augment, tierCoefficient, string(championsJson), string(traitsJson))

	var solution model.Solution
	json.Unmarshal([]byte(resultString), &solution)
	ctx.JSON(http.StatusOK, model.Response[model.Solution]{
		Status: http.StatusOK,
		Data:   solution,
	})
}

func solve(numChamps int, highTier bool, augment string, tierCoefficient float64, champions string, traits string) string {
	cAugment := C.CString(augment)
	cChampions := C.CString(champions)
	cTraits := C.CString(traits)

	result := C.solve(
		C.int(numChamps),
		C.int(boolToInt(highTier)),
		cAugment,
		C.double(tierCoefficient),
		cChampions,
		cTraits,
	)

	defer C.free_rust_string(result)
	goResult := C.GoString(result)

	C.free(unsafe.Pointer(cAugment))
	C.free(unsafe.Pointer(cChampions))
	C.free(unsafe.Pointer(cTraits))

	return goResult
}

func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}
