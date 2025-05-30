package postgres

import (
	"amantana/internal/domain"

	"gorm.io/gorm"
)

func SeedPlants(db *gorm.DB) error {
	// Hapus data yang ada
	if err := db.Exec("TRUNCATE TABLE plants RESTART IDENTITY CASCADE").Error; err != nil {
		return err
	}

	plants := []domain.Plant{
		{
			Name:        "Philodendron Pink Princess",
			Image:       "https://www.beardsanddaisies.co.uk/cdn/shop/products/Beards-Daises-27.9.220162copy_900x.jpg?v=1664960440",
			Category:    "Tanaman Hias",
			Difficulty:  "Sedang",
			Description: "Philodendron Pink Princess terkenal dengan daunnya yang memiliki variegasi merah muda yang cantik.",
			Benefits:    domain.StringArray{"Menambah nilai estetika ruangan", "Pembersih udara alami", "Meningkatkan produktivitas", "Menciptakan suasana tropis"},
			Care: domain.PlantCare{
				Watering:    "Siram saat lapisan atas tanah mulai kering",
				Sunlight:    "Cahaya tidak langsung sedang",
				Temperature: "18-29°C",
				Soil:        "Campuran tanah organik yang gembur",
			},
		},
		{
			Name:        "Alocasia Black Velvet",
			Image:       "https://www.plantvault.com/cdn/shop/articles/Alocasia_Black_Velvet_Houseplant.jpg?v=1697146921",
			Category:    "Tanaman Hias",
			Difficulty:  "Sedang",
			Description: "Tanaman hias dengan daun beludru hitam yang eksotis",
			Benefits:    domain.StringArray{"Mempercantik ruangan", "Pembersih udara", "Meningkatkan kelembaban"},
			Care: domain.PlantCare{
				Watering:    "Siram 2-3 kali seminggu",
				Sunlight:    "Cahaya tidak langsung",
				Temperature: "20-28°C",
				Soil:        "Campuran tanah yang gembur dan well-draining",
			},
		},
		{
			Name:        "Lidah Buaya",
			Image:       "https://asset.kompas.com/crops/dQTf5ZO60uyxGigGM7H1JScn3sE=/0x83:1000x750/1200x800/data/photo/2021/08/18/611cb4014c9a3.jpg",
			Category:    "Tanaman Obat",
			Difficulty:  "Mudah",
			Description: "Tanaman sukulen yang dikenal karena khasiatnya untuk kulit dan rambut.",
			Benefits:    domain.StringArray{"Mengobati luka bakar", "Pelembap alami", "Meningkatkan kualitas udara"},
			Care: domain.PlantCare{
				Watering:    "Siram seminggu sekali",
				Sunlight:    "Cahaya terang tidak langsung",
				Temperature: "18-30°C",
				Soil:        "Tanah berpasir dengan drainase baik",
			},
		},
		{
			Name:        "Bird of Paradise",
			Image:       "https://www.thespruce.com/thmb/-rU-vGiOQ-8jlT8RkHKeN88_lik=/4000x0/filters:no_upscale():max_bytes(150000):strip_icc()/SPR-bird-of-paradise-plants-2132859-hero-83ba0a370f284175a229ce271790e133.jpg",
			Category:    "Tanaman Hias",
			Difficulty:  "Sedang",
			Description: "Bird of Paradise adalah tanaman dengan daun besar dan bunga eksotis yang menawan.",
			Benefits:    domain.StringArray{"Memberikan nuansa tropis", "Focal point yang dramatis", "Pembersih udara alami", "Tahan lama"},
			Care: domain.PlantCare{
				Watering:    "Siram secara teratur, jaga kelembaban",
				Sunlight:    "Cahaya terang langsung hingga tidak langsung",
				Temperature: "18-30°C",
				Soil:        "Tanah yang kaya nutrisi dan well-draining",
			},
		},
		{
			Name:        "Monstera Deliciosa",
			Image:       "https://images.unsplash.com/photo-1614594975525-e45190c55d0b?w=800",
			Category:    "Tanaman Hias",
			Difficulty:  "Sedang",
			Description: "Tanaman dengan daun besar berlubang yang sangat dekoratif.",
			Benefits:    domain.StringArray{"Mempercantik ruangan", "Pembersih udara", "Meningkatkan kelembaban"},
			Care: domain.PlantCare{
				Watering:    "Siram 1-2 kali seminggu",
				Sunlight:    "Cahaya terang tidak langsung",
				Temperature: "20-30°C",
				Soil:        "Tanah gambut atau humus",
			},
		},
		{
			Name:        "Sirih Gading",
			Image:       "https://images.unsplash.com/photo-1596724878582-76f1a8fdc24f?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=1024&h=890&q=80",
			Category:    "Tanaman Hias",
			Difficulty:  "Mudah",
			Description: "Tanaman rambat dengan daun hijau bercorak kuning yang populer sebagai tanaman indoor.",
			Benefits:    domain.StringArray{"Pembersih udara", "Mengurangi stres", "Mudah dirawat"},
			Care: domain.PlantCare{
				Watering:    "Siram 2-3 kali seminggu",
				Sunlight:    "Cahaya redup hingga sedang",
				Temperature: "18-27°C",
				Soil:        "Tanah pot serbaguna",
			},
		},
		{
			Name:        "Kangkung",
			Image:       "https://i0.wp.com/raisa.aeonstore.id/wp-content/uploads/2024/02/806244.jpg?fit=2700%2C2700&ssl=1",
			Category:    "Tanaman Sayur",
			Difficulty:  "Mudah",
			Description: "Sayuran hijau cepat tumbuh yang populer dalam masakan Indonesia.",
			Benefits:    domain.StringArray{"Sumber zat besi", "Menyehatkan mata", "Antioksidan alami"},
			Care: domain.PlantCare{
				Watering:    "Siram setiap hari",
				Sunlight:    "Sinar matahari penuh",
				Temperature: "24-32°C",
				Soil:        "Tanah berlumpur atau basah",
			},
		},
		{
			Name:        "Tomat",
			Image:       "https://awsimages.detik.net.id/community/media/visual/2020/05/22/9c2387dd-eb74-4c21-8c57-0f01a3d75e19.jpeg?w=600&q=90",
			Category:    "Tanaman Buah",
			Difficulty:  "Sedang",
			Description: "Tanaman buah kecil yang kaya akan vitamin C dan likopen.",
			Benefits:    domain.StringArray{"Menjaga kesehatan kulit", "Menurunkan risiko kanker", "Kaya vitamin dan mineral"},
			Care: domain.PlantCare{
				Watering:    "Siram 2 kali sehari saat berbunga",
				Sunlight:    "Sinar matahari penuh",
				Temperature: "22-30°C",
				Soil:        "Tanah subur, gembur, dan kaya organik",
			},
		},
		{
			Name:        "Kemangi",
			Image:       "https://mysiloam-api.siloamhospitals.com/storage-down/file-web-cms/17237989961225591.webp",
			Category:    "Tanaman Obat",
			Difficulty:  "Mudah",
			Description: "Tanaman daun aromatik yang digunakan dalam banyak masakan dan pengobatan tradisional.",
			Benefits:    domain.StringArray{"Anti bakteri", "Melancarkan pencernaan", "Menambah cita rasa makanan"},
			Care: domain.PlantCare{
				Watering:    "Siram setiap pagi",
				Sunlight:    "Matahari penuh",
				Temperature: "20-35°C",
				Soil:        "Tanah lempung berpasir",
			},
		},
		{
			Name:        "Chinese Evergreen",
			Image:       "https://i.pinimg.com/736x/0f/b2/8b/0fb28bd72dcb55a4c0bae59d4ed9a125.jpg",
			Category:    "Tanaman Hias",
			Difficulty:  "Mudah",
			Description: "Chinese Evergreen adalah tanaman yang mudah dirawat dengan daun berpola menarik.",
			Benefits:    domain.StringArray{"Tahan berbagai kondisi", "Pembersih udara efektif", "Banyak variasi warna", "Cocok untuk pemula"},
			Care: domain.PlantCare{
				Watering:    "Siram saat lapisan atas tanah kering",
				Sunlight:    "Toleran cahaya rendah hingga sedang",
				Temperature: "16-27°C",
				Soil:        "Tanah yang well-draining",
			},
		},
		{
			Name:        "Kaktus Mini",
			Image:       "https://media-public.dekoruma.com/article/2025/03/kaktus-mini.jpg",
			Category:    "Tanaman Hias",
			Difficulty:  "Mudah",
			Description: "Tanaman berduri kecil yang cocok sebagai dekorasi meja.",
			Benefits:    domain.StringArray{"Dekorasi ruangan", "Mudah dirawat", "Hemat air"},
			Care: domain.PlantCare{
				Watering:    "Siram 1 minggu sekali",
				Sunlight:    "Matahari langsung",
				Temperature: "20-35°C",
				Soil:        "Campuran pasir dan tanah kaktus",
			},
		},
		{
			Name:        "Jahe",
			Image:       "https://sumeksradio.disway.id/upload/c94515eb4101a3f86c00f6d34a7653e2.jpg",
			Category:    "Tanaman Obat",
			Difficulty:  "Sedang",
			Description: "Rimpang yang digunakan sebagai rempah dan obat tradisional.",
			Benefits:    domain.StringArray{"Menghangatkan tubuh", "Mengatasi mual", "Anti-inflamasi"},
			Care: domain.PlantCare{
				Watering:    "Siram 2-3 kali seminggu",
				Sunlight:    "Cahaya tidak langsung",
				Temperature: "22-28°C",
				Soil:        "Tanah gembur dan lembap",
			},
		},
		{
			Name:        "Lavender",
			Image:       "https://cdn.mos.cms.futurecdn.net/vMVNCBkEjtnZE8efv7foyX.jpg",
			Category:    "Tanaman Hias",
			Difficulty:  "Mudah",
			Description: "Tanaman dengan bunga ungu harum yang sering digunakan untuk aromaterapi.",
			Benefits:    domain.StringArray{"Aromaterapi", "Mengusir nyamuk", "Menenangkan pikiran"},
			Care: domain.PlantCare{
				Watering:    "Siram 1-2 kali seminggu",
				Sunlight:    "Sinar matahari penuh",
				Temperature: "15-30°C",
				Soil:        "Tanah gembur dan berdrainase baik",
			},
		},
		{
			Name:        "Bayam",
			Image:       "https://image.astronauts.cloud/product-images/2024/5/1Cover14_bb1e11a6-1398-482f-9cac-7ec4e97de7a9_900x900.png",
			Category:    "Tanaman Sayur",
			Difficulty:  "Mudah",
			Description: "Sayuran hijau kaya nutrisi yang mudah ditanam.",
			Benefits:    domain.StringArray{"Sumber zat besi", "Meningkatkan energi", "Baik untuk kesehatan jantung"},
			Care: domain.PlantCare{
				Watering:    "Siram setiap hari",
				Sunlight:    "Sinar matahari penuh hingga teduh",
				Temperature: "18-24°C",
				Soil:        "Tanah subur dan berdrainase baik",
			},
		},
		{
			Name:        "Basil",
			Image:       "https://images.tokopedia.net/img/cache/1500/VqbcmM/2022/4/14/f73e4c9d-01d4-4ca0-bff2-a8fa03a985a3.jpg",
			Category:    "Tanaman Obat",
			Difficulty:  "Mudah",
			Description: "Tanaman aromatik yang sering dipakai untuk masakan dan obat tradisional.",
			Benefits:    domain.StringArray{"Meningkatkan pencernaan", "Antioksidan", "Menenangkan saraf"},
			Care: domain.PlantCare{
				Watering:    "Siram 2 kali seminggu",
				Sunlight:    "Cahaya penuh atau tidak langsung",
				Temperature: "20-30°C",
				Soil:        "Tanah gembur dan berdrainase baik",
			},
		},
		{
			Name:        "Cabe Rawit",
			Image:       "https://cdn0-production-images-kly.akamaized.net/g5cMiDqczLvbo1Tlk0M-eQi-hMQ=/1200x675/smart/filters:quality(75):strip_icc():format(jpeg)/kly-media-production/medias/2742391/original/099306700_1551679752-foto_5_perawatan_cabe.jpg",
			Category:    "Tanaman Buah",
			Difficulty:  "Sedang",
			Description: "Tanaman buah pedas yang banyak dipakai dalam masakan Indonesia.",
			Benefits:    domain.StringArray{"Meningkatkan metabolisme", "Mengandung vitamin C tinggi", "Membantu pencernaan"},
			Care: domain.PlantCare{
				Watering:    "Siram 2-3 kali seminggu",
				Sunlight:    "Sinar matahari penuh",
				Temperature: "22-30°C",
				Soil:        "Tanah subur dan gembur",
			},
		},
		{
			Name:        "Fiddle Leaf Fig",
			Image:       "https://www.palasa.co.in/cdn/shop/articles/IMG_20220226_173034_1.jpg?crop=center&height=2048&v=1694161186&width=2048",
			Category:    "Tanaman Hias",
			Difficulty:  "Sulit",
			Description: "Fiddle Leaf Fig atau Ficus Lyrata adalah tanaman populer dengan daun besar berbentuk biola.",
			Benefits:    domain.StringArray{"Focal point ruangan yang sempurna", "Menyaring polutan udara", "Meningkatkan konsentrasi", "Memberikan kesan mewah"},
			Care: domain.PlantCare{
				Watering:    "Siram saat 1-2 inci lapisan atas tanah kering",
				Sunlight:    "Cahaya terang tidak langsung",
				Temperature: "18-30°C",
				Soil:        "Tanah yang kaya nutrisi dengan drainase baik",
			},
		},
		{
			Name:        "Jasmine",
			Image:       "https://thursd.com/storage/media/79907/jasmine-flowers-blossoming-in-green-foliage.jpg",
			Category:    "Tanaman Hias",
			Difficulty:  "Sedang",
			Description: "Tanaman bunga harum yang populer sebagai tanaman hias dan pengharum ruangan.",
			Benefits:    domain.StringArray{"Aromaterapi", "Menenangkan pikiran", "Mengusir serangga"},
			Care: domain.PlantCare{
				Watering:    "Siram 2-3 kali seminggu",
				Sunlight:    "Cahaya terang tidak langsung",
				Temperature: "15-25°C",
				Soil:        "Tanah gembur dan kaya nutrisi",
			},
		},
		{
			Name:        "Pachira Aquatica",
			Image:       "https://florastore.com/cdn/shop/files/3381701-9826001_Atmosphere_03_SQ_MJ.jpg?v=1742554350&width=1080",
			Category:    "Tanaman Hias",
			Difficulty:  "Mudah",
			Description: "Tanaman uang yang dipercaya membawa keberuntungan.",
			Benefits:    domain.StringArray{"Membawa keberuntungan", "Pembersih udara", "Dekorasi rumah"},
			Care: domain.PlantCare{
				Watering:    "Siram 1-2 kali seminggu",
				Sunlight:    "Cahaya tidak langsung",
				Temperature: "18-24°C",
				Soil:        "Tanah pot serbaguna",
			},
		},
		{
			Name:        "Mint",
			Image:       "https://imgx.sonora.id/crop/0x0:0x0/700x465/filters:format(webp):quality(50)/photo/2024/01/17/1-ilustrasi-tanaman-mintjpg-20240117124623.jpg",
			Category:    "Tanaman Obat",
			Difficulty:  "Mudah",
			Description: "Tanaman herbal dengan aroma segar yang banyak dipakai dalam minuman dan obat tradisional.",
			Benefits:    domain.StringArray{"Meningkatkan pencernaan", "Mengurangi sakit kepala", "Menyegarkan nafas"},
			Care: domain.PlantCare{
				Watering:    "Siram 2-3 kali seminggu",
				Sunlight:    "Cahaya terang tidak langsung",
				Temperature: "15-25°C",
				Soil:        "Tanah lembab dan subur",
			},
		},
		{
			Name:        "Bonsai Serut",
			Image:       "https://down-id.img.susercontent.com/file/b62a45fdaf92d7258b476a523ec6549d",
			Category:    "Tanaman Hias",
			Difficulty:  "Sulit",
			Description: "Tanaman hias miniatur dengan teknik khusus.",
			Benefits:    domain.StringArray{"Dekorasi unik", "Menenangkan pikiran", "Melatih kesabaran"},
			Care: domain.PlantCare{
				Watering:    "Siram setiap hari",
				Sunlight:    "Cahaya terang tidak langsung",
				Temperature: "20-28°C",
				Soil:        "Tanah khusus bonsai",
			},
		},
		{
			Name:        "Anggrek Bulan",
			Image:       "https://asset.kompas.com/crops/ukrKA5VpT8BXQeIgO6Fk0kcgTuA=/0x0:1920x1280/1200x800/data/photo/2022/02/02/61fa8689dc71f.jpg",
			Category:    "Tanaman Hias",
			Difficulty:  "Sulit",
			Description: "Anggrek dengan bunga putih cantik yang tahan lama.",
			Benefits:    domain.StringArray{"Dekorasi mewah", "Menenangkan ruangan", "Simbol keindahan"},
			Care: domain.PlantCare{
				Watering:    "Siram 1-2 kali seminggu",
				Sunlight:    "Cahaya redup",
				Temperature: "18-25°C",
				Soil:        "Media tanam anggrek khusus",
			},
		},
		{
			Name:        "Bayam Merah",
			Image:       "https://d1vbn70lmn1nqe.cloudfront.net/prod/wp-content/uploads/2021/10/15060945/manfaat-bayam-merah-yang-jarang-diketahui-halodoc.jpg",
			Category:    "Tanaman Sayur",
			Difficulty:  "Mudah",
			Description: "Varietas bayam dengan daun berwarna merah kaya antioksidan.",
			Benefits:    domain.StringArray{"Kaya antioksidan", "Menyehatkan kulit", "Meningkatkan daya tahan tubuh"},
			Care: domain.PlantCare{
				Watering:    "Siram setiap hari",
				Sunlight:    "Cahaya matahari penuh",
				Temperature: "18-25°C",
				Soil:        "Tanah subur dan berdrainase baik",
			},
		},
		{
			Name:        "Bunga Matahari",
			Image:       "https://cdn1-production-images-kly.akamaized.net/xbgYRrudNnMau80_IvNe9Vb1Z30=/1200x900/smart/filters:quality(75):strip_icc():format(webp)/kly-media-production/medias/2368223/original/039312400_1537962596-sunflower-1627193_1920.jpg",
			Category:    "Tanaman Hias",
			Difficulty:  "Mudah",
			Description: "Tanaman bunga berwarna cerah yang mengikuti arah matahari.",
			Benefits:    domain.StringArray{"Dekorasi taman", "Meningkatkan mood", "Penghasil biji"},
			Care: domain.PlantCare{
				Watering:    "Siram 2 kali seminggu",
				Sunlight:    "Sinar matahari penuh",
				Temperature: "20-30°C",
				Soil:        "Tanah subur dan gembur",
			},
		},
	}

	// Batch insert semua data
	result := db.Create(&plants)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
