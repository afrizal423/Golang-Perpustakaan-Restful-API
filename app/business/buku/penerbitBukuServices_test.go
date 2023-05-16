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

var PenbukRepoInterfaceMock = &mocks.IBukuRepository{Mock: mock.Mock{}}
var PenbukServices = buku.BukuService{Repository: PenbukRepoInterfaceMock}

func TestInsertPenbuk(t *testing.T) {
	t.Run("Success insert penerbit buku", func(t *testing.T) {
		penbukDummy := models.Penerbit_Buku{
			IDPenerbit:     "01H019WHW4A45KH65AH9XW89PC",
			PenerbitBuku:   "penerbit",
			AlamatPenerbit: "alamat penerbit",
			TelpPenerbit:   "031-031031",
			EmailPenerbit:  "penerbit@buku.com",
			Deskripsi:      "ini deskripsi penerbit buku",
		}
		PenbukRepoInterfaceMock.On("CreatePenerbitBuku", mock.AnythingOfType("models.Penerbit_Buku")).Return(penbukDummy, nil).Once()

		penbukInsert := models.Penerbit_Buku{
			IDPenerbit:     "01H019WHW4A45KH65AH9XW89PC",
			PenerbitBuku:   "penerbit",
			AlamatPenerbit: "alamat penerbit",
			TelpPenerbit:   "031-031031",
			EmailPenerbit:  "penerbit@buku.com",
			Deskripsi:      "ini deskripsi penerbit buku",
		}

		data, err := PenbukServices.CreatePenerbitBuku(penbukInsert)
		assert.Equal(t, nil, err)
		assert.Equal(t, penbukInsert, data)
	})

	t.Run("Failed insert data penerbit buku", func(t *testing.T) {
		t.Run("Nama penerbit buku", func(t *testing.T) {
			penbukDummy := models.Penerbit_Buku{
				IDPenerbit:     "01H019WHW4A45KH65AH9XW89PC",
				PenerbitBuku:   "pene",
				AlamatPenerbit: "alamat penerbit",
				TelpPenerbit:   "031-031031",
				EmailPenerbit:  "penerbit@buku.com",
				Deskripsi:      "ini deskripsi penerbit buku",
			}
			PenbukRepoInterfaceMock.On("CreatePenerbitBuku", penbukDummy).Return(penbukDummy, errors.New("inputan minimal 5 karakter dan maksimal 255")).Once()

			penbukInsert := models.Penerbit_Buku{
				IDPenerbit:     "01H019WHW4A45KH65AH9XW89PC",
				PenerbitBuku:   "pene",
				AlamatPenerbit: "alamat penerbit",
				TelpPenerbit:   "031-031031",
				EmailPenerbit:  "penerbit@buku.com",
				Deskripsi:      "ini deskripsi penerbit buku",
			}

			_, err := PenbukServices.CreatePenerbitBuku(penbukInsert)
			assert.Equal(t, errors.New("inputan nama penerbit buku minimal 5 karakter dan maksimal 255"), err)
		})

		t.Run("Alamat penerbit buku", func(t *testing.T) {
			penbukDummy := models.Penerbit_Buku{
				IDPenerbit:     "01H019WHW4A45KH65AH9XW89PC",
				PenerbitBuku:   "penerbit",
				AlamatPenerbit: "ala",
				TelpPenerbit:   "031-031031",
				EmailPenerbit:  "penerbit@buku.com",
				Deskripsi:      "ini deskripsi penerbit buku",
			}
			PenbukRepoInterfaceMock.On("CreatePenerbitBuku", penbukDummy).Return(penbukDummy, errors.New("inputan minimal 5 karakter dan maksimal 255")).Once()

			penbukInsert := models.Penerbit_Buku{
				IDPenerbit:     "01H019WHW4A45KH65AH9XW89PC",
				PenerbitBuku:   "penerbit",
				AlamatPenerbit: "ala",
				TelpPenerbit:   "031-031031",
				EmailPenerbit:  "penerbit@buku.com",
				Deskripsi:      "ini deskripsi penerbit buku",
			}

			_, err := PenbukServices.CreatePenerbitBuku(penbukInsert)
			assert.Equal(t, errors.New("inputan alamat penerbit minimal 5 karakter dan maksimal 255"), err)
		})

		t.Run("Telepon penerbit buku", func(t *testing.T) {
			penbukDummy := models.Penerbit_Buku{
				IDPenerbit:     "01H019WHW4A45KH65AH9XW89PC",
				PenerbitBuku:   "penerbit",
				AlamatPenerbit: "alamat penerbit",
				TelpPenerbit:   "031",
				EmailPenerbit:  "penerbit@buku.com",
				Deskripsi:      "ini deskripsi penerbit buku",
			}
			PenbukRepoInterfaceMock.On("CreatePenerbitBuku", penbukDummy).Return(penbukDummy, errors.New("inputan telp penerbit minimal 5 karakter nomor telp")).Once()

			penbukInsert := models.Penerbit_Buku{
				IDPenerbit:     "01H019WHW4A45KH65AH9XW89PC",
				PenerbitBuku:   "penerbit",
				AlamatPenerbit: "alamat penerbit",
				TelpPenerbit:   "031",
				EmailPenerbit:  "penerbit@buku.com",
				Deskripsi:      "ini deskripsi penerbit buku",
			}

			_, err := PenbukServices.CreatePenerbitBuku(penbukInsert)
			assert.Equal(t, errors.New("inputan telp penerbit minimal 5 karakter nomor telp"), err)
		})

		t.Run("Email penerbit buku", func(t *testing.T) {
			penbukDummy := models.Penerbit_Buku{
				IDPenerbit:     "01H019WHW4A45KH65AH9XW89PC",
				PenerbitBuku:   "penerbit",
				AlamatPenerbit: "alamat penerbit",
				TelpPenerbit:   "031-321215",
				EmailPenerbit:  "penerbitbuku.com",
				Deskripsi:      "ini deskripsi penerbit buku",
			}
			PenbukRepoInterfaceMock.On("CreatePenerbitBuku", penbukDummy).Return(penbukDummy, errors.New("bukan format email")).Once()

			penbukInsert := models.Penerbit_Buku{
				IDPenerbit:     "01H019WHW4A45KH65AH9XW89PC",
				PenerbitBuku:   "penerbit",
				AlamatPenerbit: "alamat penerbit",
				TelpPenerbit:   "031-2541321",
				EmailPenerbit:  "penerbitbuku.com",
				Deskripsi:      "ini deskripsi penerbit buku",
			}

			_, err := PenbukServices.CreatePenerbitBuku(penbukInsert)
			assert.Equal(t, errors.New("bukan format email"), err)
		})

		t.Run("Deskripsi penerbit buku", func(t *testing.T) {
			penbukDummy := models.Penerbit_Buku{
				IDPenerbit:     "01H019WHW4A45KH65AH9XW89PC",
				PenerbitBuku:   "penerbit",
				AlamatPenerbit: "alamat penerbit",
				TelpPenerbit:   "031-123456",
				EmailPenerbit:  "penerbit@buku.com",
				Deskripsi:      "ini",
			}
			PenbukRepoInterfaceMock.On("CreatePenerbitBuku", penbukDummy).Return(penbukDummy, errors.New("inputan minimal 5 karakter dan maksimal 255")).Once()

			penbukInsert := models.Penerbit_Buku{
				IDPenerbit:     "01H019WHW4A45KH65AH9XW89PC",
				PenerbitBuku:   "penerbit",
				AlamatPenerbit: "alamat penerbit",
				TelpPenerbit:   "031-123456",
				EmailPenerbit:  "penerbit@buku.com",
				Deskripsi:      "ini",
			}

			_, err := PenbukServices.CreatePenerbitBuku(penbukInsert)
			assert.Equal(t, errors.New("inputan Deskripsi penerbit minimal 5 karakter dan maksimal 255"), err)
		})
	})
}

func TestGetDataPenBuk(t *testing.T) {
	t.Run("Get All data", func(t *testing.T) {
		t.Run("Success Get All data", func(t *testing.T) {
			penbukDummy := []models.Penerbit_Buku{
				{
					IDPenerbit:     "01H019WHW4A45KH65AH9XW89PC",
					PenerbitBuku:   "penerbit",
					AlamatPenerbit: "alamat penerbit",
					TelpPenerbit:   "031",
					EmailPenerbit:  "penerbit@buku.com",
					Deskripsi:      "ini deskripsi penerbit buku",
				},
			}
			PenbukRepoInterfaceMock.On("GetAllPenerbitBuku").Return(penbukDummy, nil).Once()

			data, err := PenbukServices.GetAllPenerbitBuku()
			assert.Equal(t, nil, err)
			assert.Equal(t, penbukDummy, data)
		})
		t.Run("Failed get all no data", func(t *testing.T) {
			penbukDummy := []models.Penerbit_Buku{}
			PenbukRepoInterfaceMock.On("GetAllPenerbitBuku").Return(penbukDummy, nil).Once()

			data, err := PenbukServices.GetAllPenerbitBuku()
			assert.Equal(t, errors.New("data kosong"), err)
			assert.Equal(t, penbukDummy, data)
		})
	})

	t.Run("Single data", func(t *testing.T) {
		t.Run("Success Single data", func(t *testing.T) {
			penbukDummy := models.Penerbit_Buku{
				IDPenerbit:     "01H019WHW4A45KH65AH9XW89PC",
				PenerbitBuku:   "penerbit",
				AlamatPenerbit: "alamat penerbit",
				TelpPenerbit:   "031",
				EmailPenerbit:  "penerbit@buku.com",
				Deskripsi:      "ini deskripsi penerbit buku",
			}
			PenbukRepoInterfaceMock.On("GetPenerbitBukuById", penbukDummy.IDPenerbit).Return(&penbukDummy, nil).Once()

			data, err := PenbukServices.GetPenerbitBukuById(penbukDummy.IDPenerbit)
			assert.Equal(t, nil, err)
			assert.Equal(t, &penbukDummy, data)
		})
		t.Run("Failed Single no data", func(t *testing.T) {
			penbukDummy := models.Penerbit_Buku{}
			PenbukRepoInterfaceMock.On("GetPenerbitBukuById", "01H019WHW4A45KH65AH9XW89PC").Return(&penbukDummy, nil).Once()

			data, err := PenbukServices.GetPenerbitBukuById("01H019WHW4A45KH65AH9XW89PC")
			assert.Equal(t, errors.New("data kosong"), err)
			assert.Equal(t, &penbukDummy, data)
		})
	})

}

func TestUpdateDataPenBuk(t *testing.T) {
	t.Run("Success update data", func(t *testing.T) {
		var DummyEdit models.Penerbit_Buku
		PenbukRepoInterfaceMock.On("HitungDataPenerbitBuku", mock.Anything).Return(int64(1)).Once()
		PenbukRepoInterfaceMock.On("UpdatePenerbitBuku", mock.AnythingOfType("models.Penerbit_Buku")).Return(models.Penerbit_Buku{}, nil).Once()
		dataUpdate := models.Penerbit_Buku{
			IDPenerbit:     "01H019WHW4A45KH65AH9XW89PC",
			PenerbitBuku:   "penerbit",
			AlamatPenerbit: "alamat penerbit",
			TelpPenerbit:   "0310341",
			EmailPenerbit:  "penerbit@buku.com",
			Deskripsi:      "ini deskripsi penerbit buku",
		}
		data, err := PenbukServices.UpdatePenerbitBuku(dataUpdate)
		assert.Equal(t, nil, err)
		assert.Equal(t, DummyEdit, data)
	})

	t.Run("failed update data empty id", func(t *testing.T) {
		var DummyEdit models.Penerbit_Buku
		PenbukRepoInterfaceMock.On("HitungDataPenerbitBuku", mock.Anything).Return(int64(0)).Once()
		PenbukRepoInterfaceMock.On("UpdatePenerbitBuku", mock.AnythingOfType("models.Penerbit_Buku")).Return(models.Penerbit_Buku{}, nil).Once()
		dataUpdate := models.Penerbit_Buku{
			IDPenerbit:     "01H019WHW4A45KH65AH9XW89PC",
			PenerbitBuku:   "penerbit",
			AlamatPenerbit: "alamat penerbit",
			TelpPenerbit:   "031",
			EmailPenerbit:  "penerbit@buku.com",
			Deskripsi:      "ini deskripsi penerbit buku",
		}
		data, err := PenbukServices.UpdatePenerbitBuku(dataUpdate)
		assert.Equal(t, errors.New("data kosong"), err)
		assert.Equal(t, DummyEdit, data)
	})
}

func TestHapusDataPenbuk(t *testing.T) {
	t.Run("Success hapus data", func(t *testing.T) {
		penbukDummy := models.Penerbit_Buku{
			IDPenerbit:     "01H019WHW4A45KH65AH9XW89PC",
			PenerbitBuku:   "penerbit",
			AlamatPenerbit: "alamat penerbit",
			TelpPenerbit:   "031",
			EmailPenerbit:  "penerbit@buku.com",
			Deskripsi:      "ini deskripsi penerbit buku",
		}
		PenbukRepoInterfaceMock.On("HitungDataPenerbitBuku", mock.Anything).Return(int64(1)).Once()
		PenbukRepoInterfaceMock.On("DeletePenerbitBuku", mock.Anything).Return(nil).Once()
		err := PenbukServices.HapusPenerbitBuku(penbukDummy.IDPenerbit)
		assert.Equal(t, nil, err)
	})

	t.Run("failed hapus data empty id", func(t *testing.T) {
		penbukDummy := models.Penerbit_Buku{
			IDPenerbit:     "01H019WHW4A45KH65AH9XW89PC",
			PenerbitBuku:   "penerbit",
			AlamatPenerbit: "alamat penerbit",
			TelpPenerbit:   "031",
			EmailPenerbit:  "penerbit@buku.com",
			Deskripsi:      "ini deskripsi penerbit buku",
		}
		PenbukRepoInterfaceMock.On("HitungDataPenerbitBuku", mock.Anything).Return(int64(0)).Once()
		PenbukRepoInterfaceMock.On("DeletePenerbitBuku", mock.Anything).Return(nil).Once()
		err := PenbukServices.HapusPenerbitBuku(penbukDummy.IDPenerbit)
		assert.Equal(t, errors.New("data kosong"), err)
	})
}
