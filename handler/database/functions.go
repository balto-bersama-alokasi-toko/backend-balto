package database

import (
	"backend-balto/utils"
	"context"
	"database/sql"
	"fmt"
	"slices"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
)

type (
	DbRepository struct {
		db       *pgxpool.Pool
		dbsecret string
		ctx      context.Context
	}

	Handler interface {
		GetMerchantPerCategory(category string) (result []MerchantSimpleDb, err error)
		GetMerchantCategories() (result []Category, err error)
		GetPublicCategories() (result []Category, err error)
		GetKelurahanPublicPlaces(category string) (result []KelurahanSimpleDetail, err error)
		GetTopMerchantsByKelurahan(kelurahanId int) (result []TopMerchantDb, err error)
		GetKelurahanDetails(kelurahanId int) (result []KelurahanDetailsDb, err error)
		GetMerchantDetailDb(merchantId int) (result MerchantDetailsDb, err error)
		GetQueryModelData(category string) (result []KelurahanDetailQueryModel, err error)
	}
)

func NewDbRepository(db *pgxpool.Pool, dbsecret string) Handler {
	dbContext := context.Background()
	return &DbRepository{
		db:       db,
		dbsecret: dbsecret,
		ctx:      dbContext,
	}
}

func (r *DbRepository) GetMerchantPerCategory(category string) (result []MerchantSimpleDb, err error) {
	var categories []Category

	rows, err := r.db.Query(r.ctx, getAllCategories)
	if err != nil {
		return result, err
	}
	defer rows.Close()
	for rows.Next() {
		var category Category
		err = rows.Scan(&category.ID, &category.Name)
		if err != nil {
			return result, err
		}
		categories = append(categories, category)
	}

	categoryIndex := slices.IndexFunc(categories, func(c Category) bool {
		return utils.StringCompare(c.Name, category)
	})
	if categoryIndex == -1 {
		return result, fiber.NewError(404, "category not found")
	}

	rows, err = r.db.Query(r.ctx, getMerchantPerCategory, categories[categoryIndex].ID)
	if err != nil {
		return result, err
	}
	defer rows.Close()
	for rows.Next() {
		var merchant MerchantSimpleDb
		err = rows.Scan(&merchant.ID, &merchant.Title, &merchant.Thumbnail, &merchant.Phone, &merchant.KelurahanId, &merchant.KelurahanName)
		if err != nil {
			return result, err
		}
		result = append(result, merchant)
	}
	return result, nil
}

func (r *DbRepository) GetMerchantCategories() (result []Category, err error) {
	rows, err := r.db.Query(r.ctx, getAllCategories)
	if err != nil {
		return result, err
	}
	defer rows.Close()
	for rows.Next() {
		var category Category
		err = rows.Scan(&category.ID, &category.Name)
		if err != nil {
			return result, err
		}
		result = append(result, category)
	}
	return result, nil
}

func (r *DbRepository) GetPublicCategories() (result []Category, err error) {
	rows, err := r.db.Query(r.ctx, getPublicCategories)
	if err != nil {
		return result, err
	}
	defer rows.Close()
	for rows.Next() {
		var category Category
		err = rows.Scan(&category.ID, &category.Name)
		if err != nil {
			return result, err
		}
		result = append(result, category)
	}
	return result, nil
}

func (r *DbRepository) GetKelurahanPublicPlaces(category string) (result []KelurahanSimpleDetail, err error) {
	var categories []Category

	rows, err := r.db.Query(r.ctx, getPublicCategories)
	if err != nil {
		return result, err
	}
	defer rows.Close()
	for rows.Next() {
		var category Category
		err = rows.Scan(&category.ID, &category.Name)
		if err != nil {
			return result, err
		}
		categories = append(categories, category)
	}

	categoryIndex := slices.IndexFunc(categories, func(c Category) bool {
		return utils.StringCompare(c.Name, category)
	})
	if categoryIndex == -1 {
		return result, fiber.NewError(404, "category not found")
	}

	rows, err = r.db.Query(r.ctx, getPublicPlaceKelurahan, categories[categoryIndex].ID)
	if err != nil {
		return result, err
	}
	defer rows.Close()
	for rows.Next() {
		var kelurahanDetail KelurahanSimpleDetail
		err = rows.Scan(&kelurahanDetail.Name, &kelurahanDetail.ID, &kelurahanDetail.Image)
		if err != nil {
			return result, err
		}
		result = append(result, kelurahanDetail)
	}
	return result, nil
}

func (r *DbRepository) GetTopMerchantsByKelurahan(kelurahanId int) (result []TopMerchantDb, err error) {
	modifiedQuery := fmt.Sprintf(getTopMerchantKelurahan, r.dbsecret)
	rows, err := r.db.Query(r.ctx, modifiedQuery, kelurahanId)
	if err != nil {
		return result, err
	}
	defer rows.Close()
	for rows.Next() {
		var topMerchant TopMerchantDb
		err = rows.Scan(
			&topMerchant.ID,
			&topMerchant.Title,
			&topMerchant.Thumbnail,
			&topMerchant.CategoryName)
		if err != nil {
			return result, err
		}
		result = append(result, topMerchant)
	}
	return result, nil
}

func (r *DbRepository) GetKelurahanDetails(kelurahanId int) (result []KelurahanDetailsDb, err error) {
	rows, err := r.db.Query(r.ctx, getDetailKelurahan, kelurahanId)
	if err != nil {
		return result, err
	}
	defer rows.Close()
	for rows.Next() {
		var kelurahanDetail KelurahanDetailsDb
		err = rows.Scan(
			&kelurahanDetail.ID,
			&kelurahanDetail.Name,
			&kelurahanDetail.Image,
			&kelurahanDetail.Link,
			&kelurahanDetail.MerchantId,
			&kelurahanDetail.MerchantTitle,
			&kelurahanDetail.MerchantThumbnail,
			&kelurahanDetail.CategoryName,
			&kelurahanDetail.MerchantCount,
			&kelurahanDetail.MerchantPhone,
			&kelurahanDetail.MerchantRating)
		if err != nil {
			return result, err
		}
		result = append(result, kelurahanDetail)
	}
	return result, nil
}

func (r *DbRepository) GetMerchantDetailDb(merchantId int) (result MerchantDetailsDb, err error) {
	var merchantDetail MerchantDetailsDb
	err = r.db.QueryRow(r.ctx, getMerchantDetailDb, merchantId).Scan(&merchantDetail.ID,
		&merchantDetail.Title,
		&merchantDetail.Category,
		&merchantDetail.Address,
		&merchantDetail.Link,
		&merchantDetail.ReviewRating,
		&merchantDetail.ReviewCount,
		&merchantDetail.ReviewLink,
		&merchantDetail.UserReview,
		&merchantDetail.OpenHours,
		&merchantDetail.Phone,
		&merchantDetail.Website,
		&merchantDetail.Photo)
	if err != nil {
		return merchantDetail, err
	}

	return merchantDetail, nil
}

func (r *DbRepository) GetQueryModelData(category string) (result []KelurahanDetailQueryModel, err error) {

	rows, err := r.db.Query(r.ctx, getModelQueryDetail)
	if err != nil {
		return result, err
	}
	defer rows.Close()

	for rows.Next() {
		var detail KelurahanDetailQueryModel
		err := rows.Scan(
			&detail.KelurahanID,
			&detail.NamaKelurahan,
			&detail.JumlahPendudukAkhir2023,
			&detail.PendudukLaki2,
			&detail.PendudukPerempuan,
			&detail.PendudukBeragamaIslam,
			&detail.PendudukBeragamaKristen,
			&detail.PendudukBeragamaKatholik,
			&detail.PendudukBeragamaHindu,
			&detail.PendudukBeragamaBuddha,
			&detail.PendudukBeragamaKonghucu,
			&detail.PendudukBeragamaKepercayaan,
			&detail.PendudukBelumSekolah,
			&detail.PendudukBelumSD,
			&detail.PendudukSD,
			&detail.PendudukSMP,
			&detail.PendudukSMA,
			&detail.PendudukD1D2,
			&detail.PendudukD3,
			&detail.PendudukS1,
			&detail.PendudukS2,
			&detail.PendudukS3,
			&detail.PendudukBelumAtauTidakBekerja,
			&detail.PendudukMengurusRumahTangga,
			&detail.PendudukPelajar,
			&detail.PendudukPensiunan,
			&detail.PendudukBekerja,
			&detail.Penduduk0Sampai4,
			&detail.Penduduk5Sampai9,
			&detail.Penduduk10Sampai14,
			&detail.Penduduk15Sampai19,
			&detail.Penduduk20Sampai24,
			&detail.Penduduk25Sampai29,
			&detail.Penduduk30Sampai34,
			&detail.Penduduk35Sampai39,
			&detail.Penduduk40Sampai44,
			&detail.Penduduk45Sampai49,
			&detail.Penduduk50Sampai54,
			&detail.Penduduk55Sampai59,
			&detail.Penduduk60Sampai64,
			&detail.Penduduk65Sampai69,
			&detail.Penduduk70Keatas,
			&detail.JumlahParksPerKelurahan,
			&detail.JumlahTemporaryAccommodationsPerKelurahan,
			&detail.JumlahChurchesPerKelurahan,
			&detail.JumlahAcademicInstitutionsPerKelurahan,
			&detail.JumlahGasSPBUPerKelurahan,
			&detail.JumlahMarketPerKelurahan,
			&detail.JumlahOfficesPerKelurahan,
			&detail.JumlahResidencesPerKelurahan,
			&detail.JumlahTouristPerKelurahan,
			&detail.JumlahUnclassifiedPerKelurahan,
			&detail.JumlahMallPerKelurahan,
			&detail.JumlahMosquesPerKelurahan,
			&detail.JumlahTransportationHubPerKelurahan,
			&detail.JumlahMedicalServicesPerKelurahan,
		)
		if err != nil {
			return nil, err
		}
		result = append(result, detail)
	}
	type CategoryQuery struct {
		ID       int
		Name     string
		CatIndex sql.NullInt64
	}

	var categories []CategoryQuery

	rows, err = r.db.Query(r.ctx, `SELECT id, name, array_index FROM categories`)
	if err != nil {
		return result, err
	}
	defer rows.Close()
	for rows.Next() {
		var category CategoryQuery
		err = rows.Scan(&category.ID, &category.Name, &category.CatIndex)
		if err != nil {
			return result, err
		}
		categories = append(categories, category)
	}

	categoryIndex := slices.IndexFunc(categories, func(c CategoryQuery) bool {
		return utils.StringCompare(c.Name, category)
	})
	if categoryIndex == -1 {
		return result, fiber.NewError(404, "category not found")
	}

	for i, detail := range result {
		var totalCompetitor int
		err := r.db.QueryRow(r.ctx, getCompetitorCount, categories[categoryIndex].ID, detail.KelurahanID).Scan(&totalCompetitor)
		if err != nil {
			return result, err
		}
		result[i].CategoryIndex = int(categories[categoryIndex].CatIndex.Int64)
		result[i].TotalCompetitor = totalCompetitor
	}

	return result, nil
}
