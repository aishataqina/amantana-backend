package postgres

import (
	"amantana/internal/domain"
	"time"

	"gorm.io/gorm"
)

func SeedPlants(db *gorm.DB) error {
	// Hapus data yang ada
	if err := db.Exec("TRUNCATE TABLE plants RESTART IDENTITY CASCADE").Error; err != nil {
		return err
	}

	plants := []domain.Plant{
		{
			Name:        "Monstera Deliciosa",
			Image:       "https://images.unsplash.com/photo-1614594975525-e45190c55d0b?w=800",
			Category:    "Indoor",
			Difficulty:  "Mudah",
			Description: "Monstera deliciosa adalah tanaman hias tropis yang populer dengan daunnya yang besar dan berlubang unik.",
			Benefits:    domain.StringArray{"Menyaring udara dalam ruangan", "Meningkatkan kelembaban udara", "Mengurangi stres dengan keindahannya", "Mudah dirawat"},
			Care: domain.PlantCare{
				Watering:    "Siram sekali seminggu, biarkan tanah mengering di antara penyiraman",
				Sunlight:    "Cahaya tidak langsung yang terang",
				Temperature: "20-30°C",
				Soil:        "Campuran tanah yang kaya nutrisi dengan drainase baik",
			},
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			Name:        "Snake Plant",
			Image:       "https://images.squarespace-cdn.com/content/v1/54fbb611e4b0d7c1e151d22a/1610074066643-OP8HDJUWUH8T5MHN879K/Snake+Plant.jpg?format=1000w",
			Category:    "Indoor",
			Difficulty:  "Mudah",
			Description: "Snake Plant atau Lidah Mertua adalah tanaman yang sangat tangguh dan pembersih udara alami.",
			Benefits:    domain.StringArray{"Menyerap racun dari udara", "Melepaskan oksigen di malam hari", "Tahan berbagai kondisi", "Membantu menyaring formaldehida"},
			Care: domain.PlantCare{
				Watering:    "Siram setiap 2-3 minggu",
				Sunlight:    "Toleran terhadap berbagai kondisi cahaya",
				Temperature: "18-27°C",
				Soil:        "Tanah berpasir dengan drainase baik",
			},
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			Name:        "Philodendron Pink Princess",
			Image:       "https://www.beardsanddaisies.co.uk/cdn/shop/products/Beards-Daises-27.9.220162copy_900x.jpg?v=1664960440",
			Category:    "Indoor",
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
			Name:        "Calathea Orbifolia",
			Image:       "https://www.thespruce.com/thmb/DizsBW_GJ0-8NE8FI-qg2UPYfiw=/1500x0/filters:no_upscale():max_bytes(150000):strip_icc()/calathea-orbifolia-growing-guide-5270824-hero-2a3b8667f05b40a49b27da573d2486fb.jpg",
			Category:    "Indoor",
			Difficulty:  "Sedang",
			Description: "Calathea Orbifolia memiliki daun bundar dengan motif bergaris yang elegan.",
			Benefits:    domain.StringArray{"Memperbaiki kualitas udara", "Dekorasi alami yang menenangkan", "Indikator kelembaban ruangan", "Cocok untuk ruang meditasi"},
			Care: domain.PlantCare{
				Watering:    "Jaga kelembaban tanah secara konsisten",
				Sunlight:    "Cahaya tidak langsung rendah hingga sedang",
				Temperature: "18-27°C",
				Soil:        "Tanah yang kaya humus dan dapat menahan air",
			},
		},
		{
			Name:        "Fiddle Leaf Fig",
			Image:       "https://www.palasa.co.in/cdn/shop/articles/IMG_20220226_173034_1.jpg?crop=center&height=2048&v=1694161186&width=2048",
			Category:    "Indoor",
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
			Name:        "ZZ Plant",
			Image:       "https://images.unsplash.com/photo-1572688484438-313a6e50c333?w=800",
			Category:    "Indoor",
			Difficulty:  "Mudah",
			Description: "ZZ Plant adalah tanaman yang sangat mudah dirawat dan tahan dalam kondisi minim cahaya.",
			Benefits:    domain.StringArray{"Sangat mudah dirawat", "Tahan dalam ruangan gelap", "Membersihkan udara", "Tahan kekeringan"},
			Care: domain.PlantCare{
				Watering:    "Siram setiap 2-3 minggu",
				Sunlight:    "Toleran terhadap cahaya rendah",
				Temperature: "15-30°C",
				Soil:        "Campuran tanah yang well-draining",
			},
		},
		{
			Name:        "String of Pearls",
			Image:       "https://cdn.mos.cms.futurecdn.net/aetijUZ6LLsrZA8EvavrVU.jpg",
			Category:    "Indoor",
			Difficulty:  "Sedang",
			Description: "String of Pearls adalah tanaman sukulen merambat dengan daun berbentuk mutiara.",
			Benefits:    domain.StringArray{"Cocok untuk hanging basket", "Hemat tempat", "Unik dan menarik", "Mudah diperbanyak"},
			Care: domain.PlantCare{
				Watering:    "Siram saat tanah benar-benar kering",
				Sunlight:    "Cahaya terang tidak langsung",
				Temperature: "20-25°C",
				Soil:        "Tanah khusus sukulen dengan drainase sangat baik",
			},
		},
		{
			Name:        "Peace Lily",
			Image:       "https://asset.kompas.com/crops/w0pNnO9h0MiiDoLrg3E9GxmFAsk=/0x83:1000x750/1200x800/data/photo/2024/01/18/65a9152b26a63.jpg",
			Category:    "Indoor",
			Difficulty:  "Mudah",
			Description: "Peace Lily adalah tanaman dengan bunga putih elegan yang efektif membersihkan udara.",
			Benefits:    domain.StringArray{"Menghilangkan racun dari udara", "Meningkatkan kelembaban", "Mudah merawatnya", "Bunga yang cantik"},
			Care: domain.PlantCare{
				Watering:    "Jaga kelembaban tanah, siram saat mulai layu",
				Sunlight:    "Cahaya tidak langsung rendah hingga sedang",
				Temperature: "18-30°C",
				Soil:        "Tanah yang dapat menahan air dengan baik",
			},
		},
		{
			Name:        "Chinese Evergreen",
			Image:       "https://images.unsplash.com/photo-1597305877032-0668b3c6413a?w=800",
			Category:    "Indoor",
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
			Name:        "Bird of Paradise",
			Image:       "https://www.thespruce.com/thmb/-rU-vGiOQ-8jlT8RkHKeN88_lik=/4000x0/filters:no_upscale():max_bytes(150000):strip_icc()/SPR-bird-of-paradise-plants-2132859-hero-83ba0a370f284175a229ce271790e133.jpg",
			Category:    "Outdoor",
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
	}

	// Batch insert semua data
	result := db.Create(&plants)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
