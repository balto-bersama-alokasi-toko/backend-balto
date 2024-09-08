package merchant

import (
	"backend-balto/handler/database"
	"backend-balto/models"
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os/exec"
	"strconv"
	"strings"
)

type (
	Handler interface {
		FindByCategory(category string) (response models.MerchantCategoryResponse, err error)
		GetMerchantCategories() (response models.ListCategoriesResponse, err error)
		GetPublicCategories() (response models.ListCategoriesResponse, err error)
		GetPublicPlaces(category string) (response models.ListKelurahanPublicPlaceResponse, err error)
		GetKelurahanDetails(kelurahanId int) (response models.KelurahanDetailResponse, err error)
		GetMerchantDetails(merchantId int) (response models.MerchantDetailResponse, err error)
		PredictPotentialMerchantLocation(category string) (response models.PredictionResponse, err error)
	}

	merchantUsecase struct {
		db database.Handler
	}

	ResponseModel struct {
		Predictions      []float64 `json:"predictions"`
		DeployedModelID  string    `json:"deployedModelId"`
		Model            string    `json:"model"`
		ModelDisplayName string    `json:"modelDisplayName"`
		ModelVersionID   string    `json:"modelVersionId"`
	}
)

func NewMerchantUsecase(db database.Handler) Handler {
	return &merchantUsecase{
		db: db,
	}
}

func (m *merchantUsecase) FindByCategory(category string) (response models.MerchantCategoryResponse, err error) {
	merchantsDb, err := m.db.GetMerchantPerCategory(category)
	if err != nil {
		return response, err
	}

	for _, merchant := range merchantsDb {
		response.Merchants = append(response.Merchants, models.MerchantSimpleDetail{
			ID:            merchant.ID,
			Title:         merchant.Title,
			Thumbnail:     merchant.Thumbnail,
			Category:      category,
			Phone:         merchant.Phone,
			KelurahanId:   merchant.KelurahanId,
			KelurahanName: merchant.KelurahanName,
		},
		)
	}
	response.Message = "success"
	return response, nil
}

func (m *merchantUsecase) GetMerchantCategories() (response models.ListCategoriesResponse, err error) {
	categoriesDb, err := m.db.GetMerchantCategories()
	if err != nil {
		return response, err
	}
	response.Message = "success"
	for _, category := range categoriesDb {
		response.Categories = append(response.Categories, models.CategoriesDetail{
			ID:   category.ID,
			Name: category.Name,
		})

	}
	return response, nil
}

func (m *merchantUsecase) GetPublicCategories() (response models.ListCategoriesResponse, err error) {
	categoriesDb, err := m.db.GetPublicCategories()
	if err != nil {
		return response, err
	}
	response.Message = "success"
	for _, category := range categoriesDb {
		response.Categories = append(response.Categories, models.CategoriesDetail{
			ID:   category.ID,
			Name: category.Name,
		})

	}
	return response, nil
}

func (m *merchantUsecase) GetPublicPlaces(category string) (response models.ListKelurahanPublicPlaceResponse, err error) {
	placesDb, err := m.db.GetKelurahanPublicPlaces(category)
	if err != nil {
		return response, err
	}
	response.Message = "success"
	for _, place := range placesDb {
		response.Kelurahans = append(response.Kelurahans, models.KelurahanSimpleDetail{
			ID:    place.ID,
			Name:  place.Name,
			Image: place.Image,
		},
		)
	}
	return response, nil
}

func (m *merchantUsecase) GetKelurahanDetails(kelurahanId int) (response models.KelurahanDetailResponse, err error) {

	topMerchantDb, err := m.db.GetTopMerchantsByKelurahan(kelurahanId)
	if err != nil {
		return response, err
	}

	for _, merchant := range topMerchantDb {
		response.TopMerchants = append(response.TopMerchants, models.MerchantSimpleDetailKelurahan{
			ID:        merchant.ID,
			Title:     merchant.Title,
			Thumbnail: merchant.Thumbnail,
			Category:  merchant.CategoryName,
		},
		)
	}

	merchantListDb, err := m.db.GetKelurahanDetails(kelurahanId)
	if err != nil {
		return response, err
	}
	response.KelurahanMaps = merchantListDb[0].Link
	response.KelurahanMerchantCount = merchantListDb[0].MerchantCount
	response.KelurahanPhoto = merchantListDb[0].Image
	for _, merchant := range merchantListDb {
		response.KesuluranMerchants = append(response.KesuluranMerchants, models.MerchantSimpleDetailKelurahan{
			ID:        merchant.MerchantId,
			Title:     merchant.MerchantTitle,
			Thumbnail: merchant.MerchantThumbnail,
			Category:  merchant.CategoryName,
			Phone:     merchant.MerchantPhone,
			Rating:    merchant.MerchantRating,
		},
		)
	}

	response.Message = "success"

	return response, nil
}

func (m *merchantUsecase) GetMerchantDetails(merchantId int) (response models.MerchantDetailResponse, err error) {
	merchanstDb, err := m.db.GetMerchantDetailDb(merchantId)
	if err != nil {
		return response, err
	}

	response.MerchantId = strconv.Itoa(merchanstDb.ID)
	response.MerchantName = merchanstDb.Title
	response.MerchantCategory = merchanstDb.Category
	response.MerchantAddress = merchanstDb.Address
	response.MerchantMapsLink = merchanstDb.Link
	response.MerchantReviewPoint = merchanstDb.ReviewRating
	response.MerchantReviewCount = merchanstDb.ReviewCount
	response.MerchantReviewLink = merchanstDb.ReviewLink

	var MerchantReviewSlice []models.MerchantReviewDb

	err = json.Unmarshal([]byte(merchanstDb.UserReview), &MerchantReviewSlice)
	if err != nil {
		return response, err
	}
	for _, review := range MerchantReviewSlice {
		images := ""
		if len(review.Images) > 0 {
			images = review.Images[0]
		}

		response.MerchantReviews = append(response.MerchantReviews, models.MerchantReview{
			UserReviewName:    review.Name,
			UserReviewContent: review.Description,
			UserReviewPhoto:   images,
		},
		)
	}

	MerchantOpenHourSlice := make(map[string][]string)
	err = json.Unmarshal([]byte(merchanstDb.OpenHours), &MerchantOpenHourSlice)
	if err != nil {
		return response, err
	}
	for day, hourSlice := range MerchantOpenHourSlice {
		hourStr := day + ", " + hourSlice[0]
		response.MerchantOpeningHour = append(response.MerchantOpeningHour, models.MerchantOpenHour{
			TimeDesc: hourStr,
		})
	}

	response.MerchantPhone = merchanstDb.Phone
	response.MerchantSite = merchanstDb.Website.String
	response.MerchantPhoto = merchanstDb.Photo
	response.Message = "success"
	return response, nil
}

func (m *merchantUsecase) PredictPotentialMerchantLocation(category string) (response models.PredictionResponse, err error) {
	detailKelurahan, err := m.db.GetQueryModelData(category)
	if err != nil {
		return response, err
	}
	var payloadData models.PayloadData
	for _, v := range detailKelurahan {
		categorySlices := make([]bool, 21)
		categorySlices[v.CategoryIndex] = true
		var interfaceArray []interface{}

		interfaceArray = append(interfaceArray, v.TotalCompetitor)
		for _, i := range categorySlices {
			interfaceArray = append(interfaceArray, i)
		}
		interfaceArray = append(interfaceArray, v.JumlahPendudukAkhir2023)
		interfaceArray = append(interfaceArray, v.PendudukLaki2)
		interfaceArray = append(interfaceArray, v.PendudukPerempuan)
		interfaceArray = append(interfaceArray, v.PendudukBeragamaIslam)
		interfaceArray = append(interfaceArray, v.PendudukBeragamaKristen)
		interfaceArray = append(interfaceArray, v.PendudukBeragamaKatholik)
		interfaceArray = append(interfaceArray, v.PendudukBeragamaHindu)
		interfaceArray = append(interfaceArray, v.PendudukBeragamaBuddha)
		interfaceArray = append(interfaceArray, v.PendudukBeragamaKonghucu)
		interfaceArray = append(interfaceArray, v.PendudukBeragamaKepercayaan)
		interfaceArray = append(interfaceArray, v.PendudukBelumSekolah)
		interfaceArray = append(interfaceArray, v.PendudukBelumSD)
		interfaceArray = append(interfaceArray, v.PendudukSD)
		interfaceArray = append(interfaceArray, v.PendudukSMP)
		interfaceArray = append(interfaceArray, v.PendudukSMA)
		interfaceArray = append(interfaceArray, v.PendudukD1D2)
		interfaceArray = append(interfaceArray, v.PendudukD3)
		interfaceArray = append(interfaceArray, v.PendudukS1)
		interfaceArray = append(interfaceArray, v.PendudukS2)
		interfaceArray = append(interfaceArray, v.PendudukS3)
		interfaceArray = append(interfaceArray, v.PendudukBelumAtauTidakBekerja)
		interfaceArray = append(interfaceArray, v.PendudukMengurusRumahTangga)
		interfaceArray = append(interfaceArray, v.PendudukPelajar)
		interfaceArray = append(interfaceArray, v.PendudukPensiunan)
		interfaceArray = append(interfaceArray, v.PendudukBekerja)
		interfaceArray = append(interfaceArray, v.Penduduk0Sampai4)
		interfaceArray = append(interfaceArray, v.Penduduk5Sampai9)
		interfaceArray = append(interfaceArray, v.Penduduk10Sampai14)
		interfaceArray = append(interfaceArray, v.Penduduk15Sampai19)
		interfaceArray = append(interfaceArray, v.Penduduk20Sampai24)
		interfaceArray = append(interfaceArray, v.Penduduk25Sampai29)
		interfaceArray = append(interfaceArray, v.Penduduk30Sampai34)
		interfaceArray = append(interfaceArray, v.Penduduk35Sampai39)
		interfaceArray = append(interfaceArray, v.Penduduk40Sampai44)
		interfaceArray = append(interfaceArray, v.Penduduk45Sampai49)
		interfaceArray = append(interfaceArray, v.Penduduk50Sampai54)
		interfaceArray = append(interfaceArray, v.Penduduk55Sampai59)
		interfaceArray = append(interfaceArray, v.Penduduk60Sampai64)
		interfaceArray = append(interfaceArray, v.Penduduk65Sampai69)
		interfaceArray = append(interfaceArray, v.Penduduk70Keatas)
		interfaceArray = append(interfaceArray, v.JumlahParksPerKelurahan)
		interfaceArray = append(interfaceArray, v.JumlahTemporaryAccommodationsPerKelurahan)
		interfaceArray = append(interfaceArray, v.JumlahChurchesPerKelurahan)
		interfaceArray = append(interfaceArray, v.JumlahAcademicInstitutionsPerKelurahan)
		interfaceArray = append(interfaceArray, v.JumlahGasSPBUPerKelurahan)
		interfaceArray = append(interfaceArray, v.JumlahMarketPerKelurahan)
		interfaceArray = append(interfaceArray, v.JumlahOfficesPerKelurahan)
		interfaceArray = append(interfaceArray, v.JumlahResidencesPerKelurahan)
		interfaceArray = append(interfaceArray, v.JumlahTouristPerKelurahan)
		interfaceArray = append(interfaceArray, v.JumlahMallPerKelurahan)
		interfaceArray = append(interfaceArray, v.JumlahMosquesPerKelurahan)
		interfaceArray = append(interfaceArray, v.JumlahTransportationHubPerKelurahan)
		interfaceArray = append(interfaceArray, v.JumlahMedicalServicesPerKelurahan)

		payloadData.Instances = append(payloadData.Instances, interfaceArray)

		response.LocationPredictions = append(response.LocationPredictions, models.LocationPrediction{
			KelurahanId:         v.KelurahanID,
			KelurahanName:       v.NamaKelurahan,
			KelurahanPopulation: v.JumlahPendudukAkhir2023,
		})
	}

	// Define the URL for the POST request
	urlRatingModel := "https://asia-southeast2-aiplatform.googleapis.com/v1/projects/555399740612/locations/asia-southeast2/endpoints/224397129089548288:predict" // Replace with the actual URL
	urlQrModel := "https://asia-southeast2-aiplatform.googleapis.com/v1/projects/555399740612/locations/asia-southeast2/endpoints/4237104397076660224:predict"

	ratingModelResponse, err := sendRequestToModel(payloadData, urlRatingModel)
	if err != nil {
		return response, err
	}

	for i, v := range ratingModelResponse.Predictions {
		response.LocationPredictions[i].KelurahanRating = v
	}

	//for i, _ := range response.LocationPredictions {
	//	response.LocationPredictions[i].KelurahanRating = randomFloat64InRange(0, 5)
	//	response.LocationPredictions[i].KelurahanTransaksi = randomFloat64InRange(0, 5000000)
	//}

	qrModelResponse, err := sendRequestToModel(payloadData, urlQrModel)
	if err != nil {
		return response, err
	}
	for i, v := range qrModelResponse.Predictions {
		response.LocationPredictions[i].KelurahanTransaksi = v
	}

	response.Message = "success"

	return response, nil
}

func sendRequestToModel(payload models.PayloadData, url string) (response ResponseModel, err error) {
	jsonData, err := json.Marshal(payload)
	if err != nil {
		log.Fatalf("Error marshaling JSON: %v", err)
	}

	// Create a new POST request with the JSON data
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}
	gcloudToken, err := executeCommand("gcloud", "auth", "print-access-token")
	if err != nil {
		log.Fatalf("Error executing command: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+gcloudToken)

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error sending request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return response, err
	}

	var result ResponseModel
	if err := json.Unmarshal(body, &result); err != nil {
		return response, err
	}
	return result, nil
}

func executeCommand(command string, args ...string) (string, error) {
	cmd := exec.Command(command, args...)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	// Trim the output to remove any leading/trailing whitespace or newlines
	return strings.TrimSpace(out.String()), nil
}

//func randomFloat64InRange(min, max float64) float64 {
//	r := rand.New(rand.NewSource(time.Now().UnixNano()))
//	return min + r.Float64()*(max-min)
//}
