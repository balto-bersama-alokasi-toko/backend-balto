package database

import "database/sql"

type (
	Category struct {
		ID   int
		Name string
	}

	MerchantSimpleDb struct {
		ID            int
		Title         string
		Thumbnail     string
		Phone         string
		KelurahanId   int
		KelurahanName string
	}

	KelurahanSimpleDetail struct {
		Name  string
		ID    int
		Image string
	}

	KelurahanDetailsDb struct {
		ID                int
		Name              string
		Image             string
		Link              string
		MerchantId        int
		MerchantTitle     string
		MerchantThumbnail string
		CategoryName      string
		MerchantCount     int
		MerchantPhone     string
		MerchantRating    float64
	}

	TopMerchantDb struct {
		ID           int
		Title        string
		Thumbnail    string
		CategoryName string
	}

	MerchantDetailsDb struct {
		ID           int
		Title        string
		Category     string
		Address      string
		Link         string
		ReviewRating float64
		ReviewCount  int
		ReviewLink   string
		UserReview   string
		OpenHours    string
		Phone        string
		Website      sql.NullString
		Photo        string
	}

	KelurahanDetailQueryModel struct {
		CategoryIndex                             int
		TotalCompetitor                           int    `json:"total_competitor"`
		KelurahanID                               int    `json:"kelurahan_id"`
		NamaKelurahan                             string `json:"nama_kelurahan"`
		JumlahPendudukAkhir2023                   int    `json:"jumlah_penduduk_akhir_2023"`
		PendudukLaki2                             int    `json:"penduduk_laki2"`
		PendudukPerempuan                         int    `json:"penduduk_perempuan"`
		PendudukBeragamaIslam                     int    `json:"penduduk_beragama_islam"`
		PendudukBeragamaKristen                   int    `json:"penduduk_beragama_kristen"`
		PendudukBeragamaKatholik                  int    `json:"penduduk_beragama_katholik"`
		PendudukBeragamaHindu                     int    `json:"penduduk_beragama_hindu"`
		PendudukBeragamaBuddha                    int    `json:"penduduk_beragama_buddha"`
		PendudukBeragamaKonghucu                  int    `json:"penduduk_beragama_konghucu"`
		PendudukBeragamaKepercayaan               int    `json:"penduduk_beragama_kepercayaan"`
		PendudukBelumSekolah                      int    `json:"penduduk_belum_sekolah"`
		PendudukBelumSD                           int    `json:"penduduk_belum_sd"`
		PendudukSD                                int    `json:"penduduk_sd"`
		PendudukSMP                               int    `json:"penduduk_smp"`
		PendudukSMA                               int    `json:"penduduk_sma"`
		PendudukD1D2                              int    `json:"penduduk_d1_d2"`
		PendudukD3                                int    `json:"penduduk_d3"`
		PendudukS1                                int    `json:"penduduk_s1"`
		PendudukS2                                int    `json:"penduduk_s2"`
		PendudukS3                                int    `json:"penduduk_s3"`
		PendudukBelumAtauTidakBekerja             int    `json:"penduduk_belum_atau_tidak_bekerja"`
		PendudukMengurusRumahTangga               int    `json:"penduduk_mengurus_rumah_tangga"`
		PendudukPelajar                           int    `json:"penduduk_pelajar"`
		PendudukPensiunan                         int    `json:"penduduk_pensiunan"`
		PendudukBekerja                           int    `json:"penduduk_bekerja"`
		Penduduk0Sampai4                          int    `json:"penduduk_0_sampai_4"`
		Penduduk5Sampai9                          int    `json:"penduduk_5_sampai_9"`
		Penduduk10Sampai14                        int    `json:"penduduk_10_sampai_14"`
		Penduduk15Sampai19                        int    `json:"penduduk_15_sampai_19"`
		Penduduk20Sampai24                        int    `json:"penduduk_20_sampai_24"`
		Penduduk25Sampai29                        int    `json:"penduduk_25_sampai_29"`
		Penduduk30Sampai34                        int    `json:"penduduk_30_sampai_34"`
		Penduduk35Sampai39                        int    `json:"penduduk_35_sampai_39"`
		Penduduk40Sampai44                        int    `json:"penduduk_40_sampai_44"`
		Penduduk45Sampai49                        int    `json:"penduduk_45_sampai_49"`
		Penduduk50Sampai54                        int    `json:"penduduk_50_sampai_54"`
		Penduduk55Sampai59                        int    `json:"penduduk_55_sampai_59"`
		Penduduk60Sampai64                        int    `json:"penduduk_60_sampai_64"`
		Penduduk65Sampai69                        int    `json:"penduduk_65_sampai_69"`
		Penduduk70Keatas                          int    `json:"penduduk_70_keatas"`
		JumlahParksPerKelurahan                   int    `json:"jumlah_parks_per_kelurahan"`
		JumlahTemporaryAccommodationsPerKelurahan int    `json:"jumlah_temporary_accommodations_per_kelurahan"`
		JumlahChurchesPerKelurahan                int    `json:"jumlah_churches_per_kelurahan"`
		JumlahAcademicInstitutionsPerKelurahan    int    `json:"jumlah_academic_institutions_per_kelurahan"`
		JumlahGasSPBUPerKelurahan                 int    `json:"jumlah_gas_spbu_per_kelurahan"`
		JumlahMarketPerKelurahan                  int    `json:"jumlah_market_per_kelurahan"`
		JumlahOfficesPerKelurahan                 int    `json:"jumlah_offices_per_kelurahan"`
		JumlahResidencesPerKelurahan              int    `json:"jumlah_residences_per_kelurahan"`
		JumlahTouristPerKelurahan                 int    `json:"jumlah_tourist_per_kelurahan"`
		JumlahUnclassifiedPerKelurahan            int    `json:"jumlah_unclassified_per_kelurahan"`
		JumlahMallPerKelurahan                    int    `json:"jumlah_mall_per_kelurahan"`
		JumlahMosquesPerKelurahan                 int    `json:"jumlah_mosques_per_kelurahan"`
		JumlahTransportationHubPerKelurahan       int    `json:"jumlah_transportation_hub_per_kelurahan"`
		JumlahMedicalServicesPerKelurahan         int    `json:"jumlah_medical_services_per_kelurahan"`
	}
)
