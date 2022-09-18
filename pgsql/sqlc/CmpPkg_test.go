package pgsql

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"

	"github.com/joshmalouf/fleetmanager/utils"
	"github.com/stretchr/testify/require"
)

func createRndCmpPkg(t *testing.T) AssetsCmppkg {

	args := CreateCmpPkgParams{
		UnitNumber:      rndUnitNum(),
		Stages:          rndStages(),
		OpStatus:        rndOpStatus(),
		ComStatus:       rndComStatus(),
		CurrentLocation: rndLocation(),
		DriverID:        rndSqlId(true),
		CompressorID:    rndSqlId(true),
		CoolerID:        rndSqlId(true),
		VesselID:        rndSqlId(true),
		DrawingRef:      rndPath(),
		Bom:             rndSqlId(true),
		CreatedAt:       time.Now().Add(-time.Hour * 15 * 24),
		ModifiedAt:      time.Now().Add(time.Hour * 15 * 24),
	}

	cmppkg, err := testQueries.CreateCmpPkg(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, cmppkg)

	require.NotEmpty(t, cmppkg.ID)
	require.Equal(t, args.UnitNumber, cmppkg.UnitNumber)
	require.Equal(t, args.Stages, cmppkg.Stages)
	require.Equal(t, args.OpStatus, cmppkg.OpStatus)
	require.Equal(t, args.ComStatus, cmppkg.ComStatus)
	require.Equal(t, args.CurrentLocation, cmppkg.CurrentLocation)
	require.Equal(t, args.DriverID, cmppkg.DriverID)
	require.Equal(t, args.CoolerID, cmppkg.CoolerID)
	require.Equal(t, args.VesselID, cmppkg.VesselID)
	require.Equal(t, args.DrawingRef, cmppkg.DrawingRef)
	require.Equal(t, args.Bom, cmppkg.Bom)
	require.NotEmpty(t, cmppkg.CreatedAt)
	require.NotEmpty(t, cmppkg.ModifiedAt)
	require.NotEqual(t, cmppkg.CreatedAt, cmppkg.ModifiedAt)

	return cmppkg
}

func TestCreateAccount(t *testing.T) {
	createRndCmpPkg(t)
}

func TestGetCmpPkgByID(t *testing.T) {
	cmppkg1 := createRndCmpPkg(t)
	cmppkg2, err := testQueries.GetCmpPkgByID(context.Background(), cmppkg1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, cmppkg2)

	require.Equal(t, cmppkg1.ID, cmppkg2.ID)
}

func TestGetCmpPkgByUnitNumber(t *testing.T) {
	cmppkg1 := createRndCmpPkg(t)
	cmppkg2, err := testQueries.GetCmpPkgByUnitNumber(context.Background(), cmppkg1.UnitNumber)
	require.NoError(t, err)
	require.NotEmpty(t, cmppkg2)

	require.Equal(t, cmppkg1.UnitNumber, cmppkg2.UnitNumber)
}

func TestGetCmpPkgs(t *testing.T) {
	for i := 0; i <= 10; i++ {
		createRndCmpPkg(t)
	}

	cmppkgs, err := testQueries.GetCmpPkgs(context.Background())
	require.NoError(t, err)
	require.NotEmpty(t, cmppkgs)
}

func TestGetCmpPkgsByStages(t *testing.T) {
	for i := 0; i <= 10; i++ {
		createRndCmpPkg(t)
	}

	stageFilter := rndStages()

	cmppkgs, err := testQueries.GetCmpPkgsByStages(context.Background(), stageFilter)
	require.NoError(t, err)
	require.NotEmpty(t, cmppkgs)

	for pkg := range cmppkgs {
		require.Equal(t, cmppkgs[pkg].Stages, stageFilter)
	}
}

func TestGetCmpPksByEngines(t *testing.T) {
	for i := 0; i <= 10; i++ {
		createRndCmpPkg(t)
	}

	args := GetCmpPkgsByEngineParams{
		Make:  "",
		Model: "",
	}

	cmppkgs, err := testQueries.GetCmpPkgsByEngine(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, cmppkgs)

	for pkg := range cmppkgs {
		require.Equal(t, cmppkgs[pkg].Make, args.Make)
		require.Equal(t, cmppkgs[pkg].Model, args.Model)
	}
}

func TestNewCmpPkg(t *testing.T) {

	args := NewCmpPkgParams{
		UnitNumber: rndUnitNum(),
		Stages:     rndStages(),
		DrawingRef: nil,
	}

	cmppkg, err := testQueries.NewCmpPkg(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, cmppkg)

	require.Equal(t, args.UnitNumber, cmppkg.UnitNumber)
	require.Equal(t, args.Stages, cmppkg.Stages)
	require.Equal(t, args.DrawingRef, cmppkg.DrawingRef)
	require.WithinDuration(t, time.Now(), cmppkg.CreatedAt, time.Second)
}

func TestCmpPkgChgDriver(t *testing.T) {
	testpkg := createRndCmpPkg(t)

	args := CmpPkgChgDriverParams{
		ID:       testpkg.ID,
		DriverID: rndSqlId(true),
	}

	cmppkg, err := testQueries.CmpPkgChgDriver(context.Background(), args)
	require.NoError(t, err)
	require.Equal(t, cmppkg.ID, args.ID)
	require.Equal(t, cmppkg.DriverID, args.DriverID)
}

func TestCmpPkgChgComp(t *testing.T) {
	testpkg := createRndCmpPkg(t)

	args := CmpPkgChgCompParams{
		ID:           testpkg.ID,
		CompressorID: rndSqlId(true),
	}

	cmppkg, err := testQueries.CmpPkgChgComp(context.Background(), args)
	require.NoError(t, err)
	require.Equal(t, cmppkg.ID, args.ID)
	require.Equal(t, cmppkg.CompressorID, args.CompressorID)
}

func TestDeactivateCmpPkg(t *testing.T) {
	testpkg := createRndCmpPkg(t)

	cmppkg, err := testQueries.DeactivateCmpPkg(context.Background(), testpkg.ID)
	require.NoError(t, err)
	require.Equal(t, cmppkg.ID, testpkg.ID)
	require.Equal(t, cmppkg.OpStatus, "inactive")
}

func TestDisposeCmpPkg(t *testing.T) {
	testpkg := createRndCmpPkg(t)

	cmppkg, err := testQueries.DeactivateCmpPkg(context.Background(), testpkg.ID)
	require.NoError(t, err)
	require.Equal(t, cmppkg.ID, testpkg.ID)
	require.Equal(t, cmppkg.OpStatus, "disposed")
}

func TestDeleteCmpPkg(t *testing.T) {
	testpkg := createRndCmpPkg(t)

	err := testQueries.DeleteCmpPkg(context.Background(), testpkg.ID)
	require.NoError(t, err)

	cmppkg,getserr := testQueries.GetCmpPkgByID(context.Background(), testpkg.ID)
	require.Error(t, getserr)
	require.Empty(t, cmppkg)
}

//  ---------------Random value generators----------------------------------------
func rndUnitNum() string {
	return fmt.Sprintf("V%d", utils.RandomInt(1000, 9999))
}

func rndStages() string {
	stages := []string{"1", "1-2", "2", "2-3", "3", "4"}
	return stages[utils.RandomInt(0, int64(len(stages)))]
}

func rndSqlId(state bool) sql.NullInt32 {
	if state {
		return sql.NullInt32{
			Int32: int32(utils.RandomInt(1, 99999999999)),
			Valid: state,
		}
	}
	return sql.NullInt32{
		Int32: 0,
		Valid: !state,
	}
}

func rndOpStatus() string {
	opStatus := []string{"active", "inactive", "disposed"}
	return opStatus[utils.RandomInt(0, int64(len(opStatus)))]
}

func rndComStatus() string {
	comStatus := []string{"available", "under contract", "released", "in service"}
	return comStatus[utils.RandomInt(0, int64(len(comStatus)))]
}

func rndLocation() string {
	return fmt.Sprintf("%d, %d", utils.RandomInt(0, 90), utils.RandomInt(0, 90))
}

func rndPath() string {
	return fmt.Sprintf("/%s/%s/%d", utils.RandomString(5), utils.RandomString(8), utils.RandomInt(0, 999999))
}
