package buku_test

import (
	"errors"
	"testing"

	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/app/business/buku"
	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/app/business/buku/mocks"
	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/app/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var AuthorbukRepoInterfaceMock = &mocks.IBukuRepository{Mock: mock.Mock{}}
var AuthorbukServices = buku.BukuService{Repository: AuthorbukRepoInterfaceMock}

func TestGetDataAuthorBook(t *testing.T) {
	t.Run("Get All data", func(t *testing.T) {
		t.Run("Success Get All data", func(t *testing.T) {
			PenulisBukDummy := []models.Penulis_Buku{
				{
					IDPenulis:     "01H019WHW4A45KH65AH9XW89PC",
					PenulisBuku:   "ini penulis",
					AlamatPenulis: "ini alamat penulis",
					EmailPenulis:  "a@a.com",
					Deskripsi:     "ini deskripsi buku",
				},
			}
			AuthorbukRepoInterfaceMock.On("GetAllPenulisBuku").Return(PenulisBukDummy, nil).Once()

			data, err := AuthorbukServices.GetAllPenulisBuku()
			assert.Equal(t, nil, err)
			assert.Equal(t, PenulisBukDummy, data)
		})

		t.Run("Failed get all no data", func(t *testing.T) {
			PenulisBukDummy := []models.Penulis_Buku{}
			AuthorbukRepoInterfaceMock.On("GetAllPenulisBuku").Return(PenulisBukDummy, nil).Once()

			data, err := AuthorbukServices.GetAllPenulisBuku()
			assert.Equal(t, errors.New("data kosong"), err)
			assert.Equal(t, PenulisBukDummy, data)
		})
	})

	t.Run("Single data", func(t *testing.T) {
		t.Run("Success Single data by id", func(t *testing.T) {
			PenulisBukDummy := models.Penulis_Buku{
				IDPenulis:     "01H019WHW4A45KH65AH9XW89PC",
				PenulisBuku:   "ini penulis",
				AlamatPenulis: "ini alamat penulis",
				EmailPenulis:  "a@a.com",
				Deskripsi:     "ini deskripsi buku",
			}
			AuthorbukRepoInterfaceMock.On("GetPenulisBukuById", PenulisBukDummy.IDPenulis).Return(&PenulisBukDummy, nil).Once()
			data, err := AuthorbukServices.GetPenulisBukuById(PenulisBukDummy.IDPenulis)
			assert.Equal(t, nil, err)
			assert.Equal(t, &PenulisBukDummy, data)
		})

		t.Run("Failed Single no data by id", func(t *testing.T) {
			PenulisBukDummy := models.Penulis_Buku{}
			PenbukRepoInterfaceMock.On("GetPenulisBukuById", "01H019WHW4A45KH65AH9XW89PC").Return(&PenulisBukDummy, nil).Once()

			data, err := PenbukServices.GetPenulisBukuById("01H019WHW4A45KH65AH9XW89PC")
			assert.Equal(t, errors.New("data kosong"), err)
			assert.Equal(t, &PenulisBukDummy, data)
		})

		t.Run("Success Single data by penulis buku", func(t *testing.T) {
			PenulisBukDummy := []models.Penulis_Buku{
				{
					IDPenulis:     "01H019WHW4A45KH65AH9XW89PC",
					PenulisBuku:   "ini penulis",
					AlamatPenulis: "ini alamat penulis",
					EmailPenulis:  "a@a.com",
					Deskripsi:     "ini deskripsi buku",
				},
			}
			AuthorbukRepoInterfaceMock.On("CariPenulisBuku", PenulisBukDummy[0].PenulisBuku).Return(PenulisBukDummy, nil).Once()
			data, err := AuthorbukServices.FindPenulisBuku(PenulisBukDummy[0].PenulisBuku)
			assert.Equal(t, nil, err)
			assert.Equal(t, PenulisBukDummy, data)
		})

		t.Run("Failed Single no data by find data", func(t *testing.T) {
			PenulisBukDummy := []models.Penulis_Buku{}
			PenbukRepoInterfaceMock.On("CariPenulisBuku", "cari apa").Return(PenulisBukDummy, nil).Once()

			data, err := PenbukServices.FindPenulisBuku("cari apa")
			assert.Equal(t, errors.New("data kosong"), err)
			assert.Equal(t, PenulisBukDummy, data)
		})
	})

}
func TestInsertAuthorbook(t *testing.T) {
	t.Run("Success insert Penulis buku", func(t *testing.T) {
		PenulisBukDummy := models.Penulis_Buku{
			IDPenulis:     "01H019WHW4A45KH65AH9XW89PC",
			PenulisBuku:   "ini penulis",
			AlamatPenulis: "ini alamat penulis",
			EmailPenulis:  "a@a.com",
			Deskripsi:     "ini deskripsi buku",
		}
		AuthorbukRepoInterfaceMock.On("CreatePenulisBuku", mock.AnythingOfType("models.Penulis_Buku")).Return(PenulisBukDummy, nil).Once()

		penbukInsert := models.Penulis_Buku{
			IDPenulis:     "01H019WHW4A45KH65AH9XW89PC",
			PenulisBuku:   "ini penulis",
			AlamatPenulis: "ini alamat penulis",
			EmailPenulis:  "a@a.com",
			Deskripsi:     "ini deskripsi buku",
		}
		data, err := AuthorbukServices.CreatePenulisBuku(penbukInsert)
		assert.Equal(t, nil, err)
		assert.Equal(t, penbukInsert, data)
	})

	t.Run("Failed insert nama penulis buku", func(t *testing.T) {
		PenulisBukDummy := models.Penulis_Buku{
			IDPenulis:     "01H019WHW4A45KH65AH9XW89PC",
			PenulisBuku:   "ini",
			AlamatPenulis: "ini alamat penulis",
			EmailPenulis:  "a@a.com",
			Deskripsi:     "ini deskripsi buku",
		}
		AuthorbukRepoInterfaceMock.On("CreatePenulisBuku", PenulisBukDummy).Return(PenulisBukDummy, errors.New("inputan nama penulis buku minimal 5 karakter dan maksimal 255")).Once()

		penbukInsert := models.Penulis_Buku{
			IDPenulis:     "01H019WHW4A45KH65AH9XW89PC",
			PenulisBuku:   "ini",
			AlamatPenulis: "ini alamat penulis",
			EmailPenulis:  "a@a.com",
			Deskripsi:     "ini deskripsi buku",
		}
		_, err := AuthorbukServices.CreatePenulisBuku(penbukInsert)
		assert.Equal(t, errors.New("inputan nama penulis buku minimal 5 karakter dan maksimal 255"), err)
	})

	t.Run("Failed insert alamat penulis buku", func(t *testing.T) {
		PenulisBukDummy := models.Penulis_Buku{
			IDPenulis:     "01H019WHW4A45KH65AH9XW89PC",
			PenulisBuku:   "ini penulis",
			AlamatPenulis: "ini",
			EmailPenulis:  "a@a.com",
			Deskripsi:     "ini deskripsi buku",
		}
		AuthorbukRepoInterfaceMock.On("CreatePenulisBuku", PenulisBukDummy).Return(PenulisBukDummy, errors.New("inputan alamat penulis buku minimal 5 karakter dan maksimal 255")).Once()

		penbukInsert := models.Penulis_Buku{
			IDPenulis:     "01H019WHW4A45KH65AH9XW89PC",
			PenulisBuku:   "ini penulis",
			AlamatPenulis: "ini",
			EmailPenulis:  "a@a.com",
			Deskripsi:     "ini deskripsi buku",
		}
		_, err := AuthorbukServices.CreatePenulisBuku(penbukInsert)
		assert.Equal(t, errors.New("inputan alamat penulis buku minimal 5 karakter dan maksimal 255"), err)
	})

	t.Run("Failed insert email penulis", func(t *testing.T) {
		PenulisBukDummy := models.Penulis_Buku{
			IDPenulis:     "01H019WHW4A45KH65AH9XW89PC",
			PenulisBuku:   "ini penulis",
			AlamatPenulis: "ini alamat penulis",
			EmailPenulis:  "aa.com",
			Deskripsi:     "ini deskripsi buku",
		}
		AuthorbukRepoInterfaceMock.On("CreatePenulisBuku", PenulisBukDummy).Return(PenulisBukDummy, errors.New("bukan format email")).Once()

		penbukInsert := models.Penulis_Buku{
			IDPenulis:     "01H019WHW4A45KH65AH9XW89PC",
			PenulisBuku:   "ini penulis",
			AlamatPenulis: "ini alamat penulis",
			EmailPenulis:  "aa.com",
			Deskripsi:     "ini deskripsi buku",
		}
		_, err := AuthorbukServices.CreatePenulisBuku(penbukInsert)
		assert.Equal(t, errors.New("bukan format email"), err)
	})

	t.Run("Failed insert deskripsi penulis", func(t *testing.T) {
		PenulisBukDummy := models.Penulis_Buku{
			IDPenulis:     "01H019WHW4A45KH65AH9XW89PC",
			PenulisBuku:   "ini penulis",
			AlamatPenulis: "ini alamat penulis",
			EmailPenulis:  "a@a.com",
			Deskripsi:     "hdeZv8VETZOu2TaSwJ3slWzhGmTNlPwqBKfjAvFPemYvAHaAebx9VrFANWkszRSgQ40tLvowywQPqLWjmIvOrOsh1E9NCnop9zX2pq0SegbRDgzIRQCHjXvWt2MggQhlM8eOIyeQmjTiglC7DtNY1QAG8wf4BGIhMfTgv91Je3mwkAFMNh7TkafeoNbdyeqjKthI97cQcYvKAecW4udq817nooUMr0HFLmCpykJx4LqhYAlW1jBVuHKSQAUWoRed",
		}
		AuthorbukRepoInterfaceMock.On("CreatePenulisBuku", PenulisBukDummy).Return(PenulisBukDummy, errors.New("inputan deskripsi penulis buku maksimal 255")).Once()

		penbukInsert := models.Penulis_Buku{
			IDPenulis:     "01H019WHW4A45KH65AH9XW89PC",
			PenulisBuku:   "ini penulis",
			AlamatPenulis: "ini alamat penulis",
			EmailPenulis:  "a@a.com",
			Deskripsi:     "hdeZv8VETZOu2TaSwJ3slWzhGmTNlPwqBKfjAvFPemYvAHaAebx9VrFANWkszRSgQ40tLvowywQPqLWjmIvOrOsh1E9NCnop9zX2pq0SegbRDgzIRQCHjXvWt2MggQhlM8eOIyeQmjTiglC7DtNY1QAG8wf4BGIhMfTgv91Je3mwkAFMNh7TkafeoNbdyeqjKthI97cQcYvKAecW4udq817nooUMr0HFLmCpykJx4LqhYAlW1jBVuHKSQAUWoRed",
		}
		_, err := AuthorbukServices.CreatePenulisBuku(penbukInsert)
		assert.Equal(t, errors.New("inputan deskripsi penulis buku maksimal 255"), err)
	})

}

func TestUpdateAuthorbook(t *testing.T) {
	t.Run("Success update penulis buku", func(t *testing.T) {
		PenulisBukDummy := models.Penulis_Buku{
			IDPenulis:     "01H019WHW4A45KH65AH9XW89PC",
			PenulisBuku:   "ini penulis",
			AlamatPenulis: "ini alamat penulis",
			EmailPenulis:  "a@a.com",
			Deskripsi:     "ini deskripsi buku",
		}
		AuthorbukRepoInterfaceMock.On("HitungDataPenulisBuku", mock.Anything).Return(int64(1)).Once()
		AuthorbukRepoInterfaceMock.On("UpdatePenulisBuku", PenulisBukDummy).Return(PenulisBukDummy, nil).Once()

		penbukUpdate := models.Penulis_Buku{
			IDPenulis:     "01H019WHW4A45KH65AH9XW89PC",
			PenulisBuku:   "ini penulis",
			AlamatPenulis: "ini alamat penulis",
			EmailPenulis:  "a@a.com",
			Deskripsi:     "ini deskripsi buku",
		}
		data, err := AuthorbukServices.UpdatePenulisBuku(penbukUpdate)
		assert.Equal(t, nil, err)
		assert.Equal(t, penbukUpdate, data)
	})

	t.Run("Failed update penulis empty id", func(t *testing.T) {
		PenulisBukDummy := models.Penulis_Buku{
			IDPenulis:     "01H019WHW4A45KH65AH9XW89PC",
			PenulisBuku:   "ini penulis",
			AlamatPenulis: "ini alamat penulis",
			EmailPenulis:  "a@a.com",
			Deskripsi:     "ini deskripsi buku",
		}
		AuthorbukRepoInterfaceMock.On("HitungDataPenulisBuku", mock.Anything).Return(int64(0)).Once()
		AuthorbukRepoInterfaceMock.On("UpdatePenulisBuku", PenulisBukDummy).Return(PenulisBukDummy, errors.New("data kosong")).Once()

		penbukInsert := models.Penulis_Buku{
			IDPenulis:     "01H019WHW4A45KH65AH9XW89PC",
			PenulisBuku:   "ini penulis",
			AlamatPenulis: "ini alamat penulis",
			EmailPenulis:  "a@a.com",
			Deskripsi:     "ini deskripsi buku",
		}
		_, err := AuthorbukServices.UpdatePenulisBuku(penbukInsert)
		assert.Equal(t, errors.New("data kosong"), err)
	})
}

func TestHapusDataAuthorbuk(t *testing.T) {
	t.Run("Success hapus data", func(t *testing.T) {
		penbukDummy := models.Penulis_Buku{
			IDPenulis:     "01H019WHW4A45KH65AH9XW89PC",
			PenulisBuku:   "ini penulis",
			AlamatPenulis: "ini alamat penulis",
			EmailPenulis:  "a@a.com",
			Deskripsi:     "ini deskripsi buku",
		}
		AuthorbukRepoInterfaceMock.On("HitungDataPenulisBuku", mock.Anything).Return(int64(1)).Once()
		AuthorbukRepoInterfaceMock.On("DeletePenulisBuku", mock.Anything).Return(nil).Once()
		err := AuthorbukServices.HapusPenulisBuku(penbukDummy.IDPenulis)
		assert.Equal(t, nil, err)
	})

	t.Run("failed hapus data empty id", func(t *testing.T) {
		penbukDummy := models.Penulis_Buku{
			IDPenulis:     "01H019WHW4A45KH65AH9XW89PC",
			PenulisBuku:   "ini penulis",
			AlamatPenulis: "ini alamat penulis",
			EmailPenulis:  "a@a.com",
			Deskripsi:     "ini deskripsi buku",
		}
		AuthorbukRepoInterfaceMock.On("HitungDataPenulisBuku", mock.Anything).Return(int64(0)).Once()
		AuthorbukRepoInterfaceMock.On("DeletePenulisBuku", mock.Anything).Return(nil).Once()
		err := AuthorbukServices.HapusPenulisBuku(penbukDummy.IDPenulis)
		assert.Equal(t, errors.New("data kosong"), err)
	})
}
