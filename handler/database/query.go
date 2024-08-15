package database

var (
	getMerchantPerCategory = `
		SELECT 
		    m.id,
		    title,
		    thumbnail,
		    phone,
		    kelurahan_id,
		    dk.nama_kelurahan
		FROM 
		    merchants m
		INNER JOIN 
		    daftar_kelurahan dk
		ON
			m.kelurahan_id = dk.id
		WHERE kategori_id = $1`

	getAllCategories = `
		SELECT 
		    id,
		    name
		FROM 
		    categories`

	getPublicPlaceKelurahan = `
		SELECT 
		    DISTINCT(dk.nama_kelurahan),
		    dk.id AS "kelurahan_id",
		    dk.image
		FROM
		    popular_places pp
		JOIN 
			popular_place_categories ppc
		ON pp.kategori_id = ppc.id
		JOIN 
			daftar_kelurahan dk
		ON pp.kelurahan_id = dk.id
		WHERE
			ppc.id = $1
	`

	getPublicCategories = `
		SELECT 
		    id,
		    name
		FROM 
		    popular_place_categories`

	getTopMerchantKelurahan = `
		SELECT
		    m.id as "merchant_id",
			m.title,
			m.thumbnail,
			c.name AS "category_name"
		FROM
			daftar_kelurahan dk
		JOIN 
			merchants m
		ON
		    dk.id = m.kelurahan_id
		JOIN
			categories c
		ON 
			c.id = m.kategori_id
		where
		    m.kelurahan_id = $1
		order by
		    pgp_sym_decrypt(m.last_month_qr_total_nominal_transactions,%s)::INT desc
		LIMIT 3;
`

	getDetailKelurahan = `
		SELECT
			dk.id AS "kelurahan_id",
			dk.nama_kelurahan,
			dk.image,
			dk.maps_link,
			m.id AS "merchant_id",
			m.title,
			m.thumbnail,
			c.name AS "category_name",
			COUNT(m.id) OVER (PARTITION BY dk.id) AS "merchant_count",
            m.phone,
            m.review_rating
		FROM
			daftar_kelurahan dk
		JOIN
			merchants m ON dk.id = m.kelurahan_id
		JOIN
			categories c ON c.id = m.kategori_id
		WHERE
			dk.id = $1;
		`

	getMerchantDetailDb = `
	    SELECT 
			m.id as "merchant_id",
			m.title,
			c.name as "category_name",
			m.address,
			m.link,
			m.review_rating,
			m.review_count,
			m.reviews_link,
			m.user_reviews,
			m.open_hours,
			m.phone,
			m.website,
			m.thumbnail
		FROM 
		    merchants m
		JOIN
		    categories c
		ON 
		    m.kategori_id = c.id
	    WHERE
	         m.id = $1;
`

	getCompetitorCount = `
        SELECT
            count(*) as total_competitor
        FROM
            merchants m
        WHERE
            m.kategori_id = $1 AND
            m.kelurahan_id = $2;`

	getModelQueryDetail = `
        SELECT
            dk.id AS "kelurahan_id",
            dk.nama_kelurahan,
            detk.jumlah_penduduk_akhir_2023,
            detk.penduduk_laki2,
            detk.penduduk_perempuan,
            detk.penduduk_beragama_islam,
            detk.penduduk_beragama_kristen,
            detk.penduduk_beragama_katholik,
            detk.penduduk_beragama_hindu,
            detk.penduduk_beragama_buddha,
            detk.penduduk_beragama_konghucu,
            detk.penduduk_beragama_kepercayaan,
            detk.penduduk_belum_sekolah,
            detk.penduduk_belum_sd,
            detk.penduduk_sd,
            detk.penduduk_smp,
            detk.penduduk_sma,
            detk.penduduk_d1_d2,
            detk.penduduk_d3,
            detk.penduduk_s1,
            detk.penduduk_s2,
            detk.penduduk_s3,
            detk.penduduk_belum_atau_tidak_bekerja,
            detk.penduduk_mengurus_rumah_tangga,
            detk.penduduk_pelajar,
            detk.penduduk_pensiunan,
            detk.penduduk_bekerja,
            detk.penduduk_0_sampai_4,
            detk.penduduk_5_sampai_9,
            detk.penduduk_10_sampai_14,
            detk.penduduk_15_sampai_19,
            detk.penduduk_20_sampai_24,
            detk.penduduk_25_sampai_29,
            detk.penduduk_30_sampai_34,
            detk.penduduk_35_sampai_39,
            detk.penduduk_40_sampai_44,
            detk.penduduk_45_sampai_49,
            detk.penduduk_50_sampai_54,
            detk.penduduk_55_sampai_59,
            detk.penduduk_60_sampai_64,
            detk.penduduk_65_sampai_69,
            detk.penduduk_70_keatas,
            detk.jumlah_parks_per_kelurahan,
            detk."jumlah_Temporary Accomodations_per_kelurahan",
            detk.jumlah_churches_per_kelurahan,
            detk."jumlah_Academic Institutions_per_kelurahan",
            detk."jumlah_Gas/SPBU_per_kelurahan",
            detk.jumlah_market_per_kelurahan,
            detk.jumlah_offices_per_kelurahan,
            detk.jumlah_residences_per_kelurahan,
            detk.jumlah_tourist_per_kelurahan,
            detk.jumlah_unclassified_per_kelurahan,
            detk.jumlah_mall_per_kelurahan,
            detk.jumlah_mosques_per_kelurahan,
            detk."jumlah_Transportation Hub_per_kelurahan",
            detk."jumlah_Medical Services_per_kelurahan"
        FROM 
            daftar_kelurahan dk
        JOIN
            detail_kelurahan detk
        ON
            dk.nama_kelurahan  = detk.kelurahan;`
)
