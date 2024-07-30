package models

type (
	GeneralResponse struct {
		Message string `json:"message"`
	}

	Configuration struct {
		Db          DbConfig     `json:"db"`
		Server      ServerConfig `json:"server"`
		AccessToken string       `json:"accessToken"`
	}

	ServerConfig struct {
		Port int `json:"port"`
	}
	DbConfig struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Database string `json:"database"`
		Host     string `json:"host"`
		Port     int    `json:"port"`
	}

	MerchantDetailResponse struct {
		Message             string             `json:"message"`
		MerchantId          string             `json:"merchant_id"`
		MerchantName        string             `json:"merchant_name"`
		MerchantCategory    string             `json:"merchant_category"`
		MerchantAddress     string             `json:"merchant_address"`
		MerchantMapsLink    string             `json:"merchant_maps_link"`
		MerchantReviewPoint float64            `json:"merchant_review_point"`
		MerchantReviewCount int                `json:"merchant_review_count"`
		MerchantReviewLink  string             `json:"merchant_review_link"`
		MerchantReviews     []MerchantReview   `json:"merchant_reviews"`
		MerchantOpeningHour []MerchantOpenHour `json:"merchant_opening_hour"`
		MerchantPhone       string             `json:"merchant_phone"`
		MerchantSite        string             `json:"merchant_site"`
		MerchantPhoto       string             `json:"merchant_photo"`
	}

	MerchantReview struct {
		UserReviewName    *string `json:"user_review_name"`
		UserReviewContent *string `json:"user_review_content"`
		UserReviewPhoto   string  `json:"user_review_photo"`
	}

	MerchantReviewDb struct {
		Name           *string  `json:"Name"`
		ProfilePicture *string  `json:"ProfilePicture"`
		Rating         *float64 `json:"Rating"`
		Description    *string  `json:"Description"`
		Images         []string `json:"Images"`
		When           *string  `json:"When"`
	}

	MerchantOpenHour struct {
		TimeDesc string `json:"time_desc"`
	}

	PredictionResponse struct {
		Message             string               `json:"message"`
		LocationPredictions []LocationPrediction `json:"location_predictions"`
	}

	LocationPrediction struct {
		KelurahanId         int     `json:"kelurahan_id"`
		KelurahanName       string  `json:"kelurahan_name"`
		KelurahanRating     float64 `json:"kelurahan_rating"`
		KelurahanTransaksi  float64 `json:"kelurahan_transaksi"`
		KelurahanPopulation int     `json:"kelurahan_population"`
	}

	PayloadData struct {
		Instances [][]interface{} `json:"instances"`
	}

	Instance struct {
		KompetitorInKelurahan           int  `json:"kompetitor_in_kelurahan"`
		ApotekDanProdukKesehatanLainnya bool `json:"Apotek dan Produk Kesehatan Lainnya"`
		Clothing                        bool `json:"Clothing"`
		Entertainment                   bool `json:"Entertainment"`
		Jasa                            bool `json:"Jasa"`
		KafeDanMinuman                  bool `json:"Kafe dan Minuman"`
		Kebersihan                      bool `json:"Kebersihan"`
		Kosmetik                        bool `json:"Kosmetik"`
		Materials                       bool `json:"Materials"`
		Olahraga                        bool `json:"Olahraga"`
		Optics                          bool `json:"Optics"`
		Otomotif                        bool `json:"Otomotif"`
		PeralatanDanBarangElektronik    bool `json:"Peralatan dan Barang Elektronik"`
		Photo                           bool `json:"Photo"`
		Printing                        bool `json:"Printing"`
		Properti                        bool `json:"Properti"`
		Regional                        bool `json:"Regional"`
		RestoranUmum                    bool `json:"Restoran Umum"`
		Retail                          bool `json:"Retail"`
		RotiKueDanCemilanLainnya        bool `json:"Roti, Kue, dan Cemilan Lainnya"`
		Transportation                  bool `json:"Transportation"`
		Warung                          bool `json:"Warung"`
		JumlahPendudukAkhir2023         int  `json:"jumlah_penduduk_akhir_2023"`
		PendudukLaki2                   int  `json:"penduduk_laki2"`
		PendudukPerempuan               int  `json:"penduduk_perempuan"`
		PendudukBeragamaIslam           int  `json:"penduduk_beragama_islam"`
		PendudukBeragamaKristen         int  `json:"penduduk_beragama_kristen"`
		PendudukBeragamaKatholik        int  `json:"penduduk_beragama_katholik"`
		PendudukBeragamaHindu           int  `json:"penduduk_beragama_hindu"`
		PendudukBeragamaBuddha          int  `json:"penduduk_beragama_buddha"`
		PendudukBeragamaKonghucu        int  `json:"penduduk_beragama_konghucu"`
		PendudukBeragamaKepercayaan     int  `json:"penduduk_beragama_kepercayaan"`
		PendudukBelumSekolah            int  `json:"penduduk_belum_sekolah"`
		PendudukBelumSD                 int  `json:"penduduk_belum_SD"`
		PendudukSD                      int  `json:"penduduk_SD"`
		PendudukSMP                     int  `json:"penduduk_SMP"`
		PendudukSMA                     int  `json:"penduduk_SMA"`
		PendudukD1D2                    int  `json:"penduduk_d1_d2"`
		PendudukD3                      int  `json:"penduduk_d3"`
		PendudukS1                      int  `json:"penduduk_s1"`
		PendudukS2                      int  `json:"penduduk_s2"`
		PendudukS3                      int  `json:"penduduk_s3"`
		PendudukBelumAtauTidakBekerja   int  `json:"penduduk_belum_atau_tidak_bekerja"`
		PendudukMengurusRumahTangga     int  `json:"penduduk_mengurus_rumah_tangga"`
		PendudukPelajar                 int  `json:"penduduk_pelajar"`
		PendudukPensiunan               int  `json:"penduduk_pensiunan"`
		PendudukBekerja                 int  `json:"penduduk_bekerja"`
		Penduduk0Sampai4                int  `json:"penduduk_0_sampai_4"`
		Penduduk5Sampai9                int  `json:"penduduk_5_sampai_9"`
		Penduduk10Sampai14              int  `json:"penduduk_10_sampai_14"`
		Penduduk15Sampai19              int  `json:"penduduk_15_sampai_19"`
		Penduduk20Sampai24              int  `json:"penduduk_20_sampai_24"`
		Penduduk25Sampai29              int  `json:"penduduk_25_sampai_29"`
		Penduduk30Sampai34              int  `json:"penduduk_30_sampai_34"`
		Penduduk35Sampai39              int  `json:"penduduk_35_sampai_39"`
		Penduduk40Sampai44              int  `json:"penduduk_40_sampai_44"`
		Penduduk45Sampai49              int  `json:"penduduk_45_sampai_49"`
		Penduduk50Sampai54              int  `json:"penduduk_50_sampai_54"`
		Penduduk55Sampai59              int  `json:"penduduk_55_sampai_59"`
		Penduduk60Sampai64              int  `json:"penduduk_60_sampai_64"`
		Penduduk65Sampai69              int  `json:"penduduk_65_sampai_69"`
		Penduduk70Keatas                int  `json:"penduduk_70_keatas"`
		JumlahParksPerKelurahan         int  `json:"jumlah_Parks_per_kelurahan"`
		JumlahTemporaryAccommodations   int  `json:"jumlah_Temporary Accomodations_per_kelurahan"`
		JumlahChurches                  int  `json:"jumlah_Churches_per_kelurahan"`
		JumlahAcademicInstitutions      int  `json:"jumlah_Academic Institutions_per_kelurahan"`
		JumlahGasSPBU                   int  `json:"jumlah_Gas/SPBU_per_kelurahan"`
		JumlahMarket                    int  `json:"jumlah_Market_per_kelurahan"`
		JumlahOffices                   int  `json:"jumlah_Offices_per_kelurahan"`
		JumlahResidences                int  `json:"jumlah_Residences_per_kelurahan"`
		JumlahTourist                   int  `json:"jumlah_Tourist_per_kelurahan"`
		JumlahUnclassified              int  `json:"jumlah_Unclassified_per_kelurahan"`
		JumlahMall                      int  `json:"jumlah_Mall_per_kelurahan"`
		JumlahMosques                   int  `json:"jumlah_Mosques_per_kelurahan"`
		JumlahTransportationHub         int  `json:"jumlah_Transportation Hub_per_kelurahan"`
		JumlahMedicalServices           int  `json:"jumlah_Medical Services_per_kelurahan"`
	}
)
