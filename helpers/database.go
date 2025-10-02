package helpers

import (
	"backend-test/internal/models"
	"fmt"
	"log"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func SetupPostgreSQL() {
	var err error
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		GetEnv("DB_HOST", "127.0.0.1"),
		GetEnv("DB_PORT", "5432"),
		GetEnv("DB_USER", "postgres"),
		GetEnv("DB_PASSWORD", "postgres"),
		GetEnv("DB_NAME", "mlm_db"),
	)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to database: ", err)
	}

	logrus.Info("Successfully connected to database..")

	// AutoMigrate
	err = DB.AutoMigrate(
		&models.Member{},
		&models.Manager{},
		&models.Location{},
		&models.Registration{},
		&models.Paket{},
	)
	if err != nil {
		log.Fatal("failed to migrate database: ", err)
	}

	seedDatabase()
}

func seedDatabase() {
	// Seeder Paket
	var paketCount int64
	DB.Model(&models.Paket{}).Count(&paketCount)
	if paketCount == 0 {
		pakets := []models.Paket{
			{NamaPaket: "Fast Track", JenisPaket: "Silver", Wilayah: "A", Price: 1000000},
			{NamaPaket: "Smart Reward", JenisPaket: "Gold", Wilayah: "B", Price: 2000000},
			{NamaPaket: "Fast Track & Smart Reward", JenisPaket: "Platinum", Wilayah: "A", Price: 5000000},
		}
		for _, p := range pakets {
			DB.FirstOrCreate(&p, models.Paket{NamaPaket: p.NamaPaket, JenisPaket: p.JenisPaket})
		}
	}

	// Seeder Location
	var locationCount int64
	DB.Model(&models.Location{}).Count(&locationCount)
	if locationCount == 0 {
		locations := []models.Location{
			{Kelurahan: "Kelurahan 1", Kecamatan: "Kecamatan 1", Kabupaten: "Kabupaten 1", KodePos: "11111", Detail: "Jl. Mawar No.1"},
			{Kelurahan: "Kelurahan 2", Kecamatan: "Kecamatan 2", Kabupaten: "Kabupaten 2", KodePos: "22222", Detail: "Jl. Melati No.2"},
			{Kelurahan: "Kelurahan 3", Kecamatan: "Kecamatan 3", Kabupaten: "Kabupaten 3", KodePos: "33333", Detail: "Jl. Anggrek No.3"},
			{Kelurahan: "Kelurahan 4", Kecamatan: "Kecamatan 4", Kabupaten: "Kabupaten 4", KodePos: "44444", Detail: "Jl. Kenanga No.4"},
		}
		for _, l := range locations {
			DB.FirstOrCreate(&l, models.Location{Kelurahan: l.Kelurahan, Kecamatan: l.Kecamatan})
		}
	}

	// Seeder Manager
	var managerCount int64
	DB.Model(&models.Manager{}).Count(&managerCount)
	if managerCount == 0 {
		var locations []models.Location
		if err := DB.Find(&locations).Error; err != nil || len(locations) == 0 {
			log.Fatal("Seeder gagal: Location belum ada")
		}

		managers := []models.Manager{
			{Nama: "Manager A", LocationID: locations[0].ID},
		}
		if len(locations) > 1 {
			managers = append(managers, models.Manager{Nama: "Manager B", LocationID: locations[1].ID})
		}
		if len(locations) > 2 {
			managers = append(managers, models.Manager{Nama: "Manager C", LocationID: locations[2].ID})
		}
		if len(locations) > 3 {
			managers = append(managers, models.Manager{Nama: "Manager D", LocationID: locations[3].ID})
		}

		for _, m := range managers {
			DB.FirstOrCreate(&m, models.Manager{Nama: m.Nama})
		}
	}

	// Seeder Member + Registration
	var memberCount int64
	DB.Model(&models.Member{}).Count(&memberCount)
	if memberCount == 0 {
		var manager models.Manager
		if err := DB.First(&manager).Error; err != nil {
			log.Fatal("Seeder gagal: Manager belum ada")
		}

		var paket models.Paket
		if err := DB.First(&paket).Error; err != nil {
			log.Fatal("Seeder gagal: Paket belum ada")
		}

		// Buat Member
		member := models.Member{
			ID:             "MEM-INIT-0001",
			Nama:           "Member Test",
			JenisKelamin:   "Laki-laki",
			NoKtp:          "1234567890123456",
			TempatLahir:    "Jakarta",
			TanggalLahir:   time.Now().Format("2006-01-02"),
			NoHp:           "08123456789",
			Email:          "member@test.com",
			NoRekening:     "1234567890",
			ManagerID:      manager.ID,
			UplineMemberID: "MEM-INIT-0001",
		}
		if err := DB.Create(&member).Error; err != nil {
			log.Fatal("Seeder gagal insert Member:", err)
		}

		// Buat Registration setelah Member ada
		registration := models.Registration{
			MemberID: member.ID,
			PaketID:  paket.ID,
		}
		if err := DB.Create(&registration).Error; err != nil {
			log.Fatal("Seeder gagal insert Registration:", err)
		}

		// Update Member dengan RegistrationID
		member.RegistartionID = registration.ID
		DB.Save(&member)
	}

	logrus.Info("Seeder executed successfully")
}
