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
			Image:       "uploads/1d4c237f-215f-4c4f-8b05-92d1045fd76a.jpg",
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
			Image:       "uploads/alocasia-black-velvet.jpg",
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
			Image:       "uploads/lidah-buaya.jpg",
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
			Image:       "uploads/bird-of-paradise.jpg",
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
			Image:       "uploads/monstera-deliciosa.jpg",
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
			Image:       "uploads/sirih-gading.jpg",
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
			Image:       "uploads/kangkung.jpeg",
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
			Image:       "uploads/tomat.jpg",
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
			Image:       "uploads/kemangi.jpg",
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
			Image:       "uploads/chinese-evergreen.jpg",
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
			Image:       "uploads/kaktus-mini.jpeg",
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
			Image:       "uploads/jahe.jpg",
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
			Image:       "uploads/lavender.jpg",
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
			Image:       "uploads/bayam.jpg",
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
			Image:       "uploads/basil.jpg",
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
			Image:       "uploads/cabe-rawit.jpg",
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
			Image:       "uploads/fiddle-leaf-fig.jpg",
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
			Image:       "uploads/jasmine.jpg",
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
			Image:       "uploads/pachira-aquatica.jpg",
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
			Image:       "uploads/mint.jpg",
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
			Image:       "uploads/bonsai-serut.jpg",
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
			Image:       "uploads/anggrek-bulan.jpg",
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
			Image:       "uploads/bayam-merah.jpg",
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
			Image:       "uploads/bunga-matahari.jpg",
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
