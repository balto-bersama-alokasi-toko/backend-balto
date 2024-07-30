package models

type (
	MerchantCategoryResponse struct {
		Message   string                 `json:"message"`
		Merchants []MerchantSimpleDetail `json:"merchants"`
	}

	MerchantSimpleDetail struct {
		ID            int    `json:"merchant_id"`
		Title         string `json:"merchant_name"`
		Thumbnail     string `json:"merchant_photo"`
		Category      string `json:"merchant_category"`
		Phone         string `json:"merchant_phone"`
		KelurahanId   int    `json:"merchant_kelurahan_id"`
		KelurahanName string `json:"merchant_kelurahan_name"`
	}

	ListCategoriesResponse struct {
		Message    string             `json:"message"`
		Categories []CategoriesDetail `json:"categories"`
	}

	CategoriesDetail struct {
		ID   int    `json:"category_id"`
		Name string `json:"category_name"`
	}

	ListKelurahanPublicPlaceResponse struct {
		Message    string                  `json:"message"`
		Kelurahans []KelurahanSimpleDetail `json:"kelurahans"`
	}

	KelurahanSimpleDetail struct {
		Name  string `json:"kelurahan_name"`
		ID    int    `json:"kelurahan_id"`
		Image string `json:"kelurahan_url_photo"`
	}

	KelurahanDetailResponse struct {
		Message                string                          `json:"message"`
		KelurahanMaps          string                          `json:"kelurahan_maps"`
		KelurahanPhoto         string                          `json:"kelurahan_photo"`
		KelurahanMerchantCount int                             `json:"kelurahan_merchant_count"`
		TopMerchants           []MerchantSimpleDetailKelurahan `json:"top_merchants"`
		KesuluranMerchants     []MerchantSimpleDetailKelurahan `json:"kesuluran_merchants"`
	}

	MerchantSimpleDetailKelurahan struct {
		ID        int     `json:"merchant_id"`
		Title     string  `json:"merchant_name"`
		Thumbnail string  `json:"merchant_photo"`
		Category  string  `json:"merchant_category"`
		Phone     string  `json:"merchant_phone"`
		Rating    float64 `json:"merchant_rating"`
	}
)
