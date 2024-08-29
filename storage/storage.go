package storage

import (
	"database/sql"
	"sale/storage/mongosh"
	"sale/storage/postgres"
	"sale/storage/repo"

	"go.mongodb.org/mongo-driver/mongo"
)

type Istorage interface {
	Product() repo.Product
	Bought() repo.Bought
	Feedback() repo.Feedback
	Wishlist() repo.Wishlist
	Processes() repo.Processes
}

type StoragePro struct {
	Mdb *mongo.Database
	PDB *sql.DB
}

func NewStoragePro(mdb *mongo.Database, pdb *sql.DB) Istorage {
	return &StoragePro{
		Mdb: mdb,
		PDB: pdb,
	}
}

func (pro *StoragePro) Product() repo.Product {
	return mongosh.NewProductsRepository(pro.Mdb)
}

func (pro *StoragePro) Bought() repo.Bought {
	return postgres.NewBougthRepository(pro.PDB)
}

func (pro *StoragePro) Feedback() repo.Feedback {
	return postgres.NewFeedbackRepository(pro.PDB)
}

func (pro *StoragePro) Wishlist() repo.Wishlist {
	return postgres.NewWishlistRepository(pro.PDB)
}

func (pro *StoragePro) Processes() repo.Processes {
	return postgres.NewProcessRepository(pro.PDB)
}
