package buku_test

import (
	"testing"

	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/app/business/buku"
	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/app/business/buku/mocks"
	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/app/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var JenbukRepoInterfaceMock = &mocks.IBukuRepository{Mock: mock.Mock{}}
var JenbukServices = buku.BukuService{Repository: JenbukRepoInterfaceMock}

func TestInsertJenbuk(t *testing.T) {
	JenbukDummy := models.Jenis_Buku{
		IDJenis:   "01H019WHW4A45KH65AH9XW89PC",
		JenisBuku: "ini jenis buku",
		Deskripsi: "ini deskripsi",
	}
	JenbukRepoInterfaceMock.On("CreateJenisBuku", mock.AnythingOfType("models.Jenis_Buku")).Return(JenbukDummy, nil).Once()
	JenbukInsert := models.Jenis_Buku{
		IDJenis:   "01H019WHW4A45KH65AH9XW89PC",
		JenisBuku: "ini jenis buku",
		Deskripsi: "ini deskripsi",
	}
	data, err := JenbukServices.CreateJenisBuku(JenbukInsert)
	assert.Equal(t, nil, err)
	assert.Equal(t, JenbukInsert, data)
}

func TestGetDataJenisBuku(t *testing.T) {
	t.Run("Get All data", func(t *testing.T) {
		JenbukDummy := []models.Jenis_Buku{
			{
				IDJenis:   "01H019WHW4A45KH65AH9XW89PC",
				JenisBuku: "ini jenis buku",
				Deskripsi: "ini deskripsi",
			},
		}
		JenbukRepoInterfaceMock.On("GetAllJenisBuku").Return(JenbukDummy, nil).Once()

		data, err := JenbukServices.GetAllJenisBuku()
		assert.Equal(t, nil, err)
		assert.Equal(t, JenbukDummy, data)
	})

	t.Run("Get Single data", func(t *testing.T) {
		JenbukDummy := models.Jenis_Buku{
			IDJenis:   "01H019WHW4A45KH65AH9XW89PC",
			JenisBuku: "ini jenis buku",
			Deskripsi: "ini deskripsi",
		}
		JenbukRepoInterfaceMock.On("GetJenisBukuById", JenbukDummy.IDJenis).Return(&JenbukDummy, nil).Once()

		data, err := JenbukServices.GetJenisBukuById(JenbukDummy.IDJenis)
		assert.Equal(t, nil, err)
		assert.Equal(t, &JenbukDummy, data)
	})

	t.Run("Get Single empty data", func(t *testing.T) {
		JenbukDummy := models.Jenis_Buku{}
		JenbukRepoInterfaceMock.On("GetJenisBukuById", "01H019WHW4A45KH65AH9XW89PC").Return(&JenbukDummy, nil).Once()

		data, err := JenbukServices.GetJenisBukuById("01H019WHW4A45KH65AH9XW89PC")
		assert.Equal(t, nil, err)
		assert.Equal(t, &JenbukDummy, data)
	})

	t.Run("Get All find data", func(t *testing.T) {
		JenbukDummy := []models.Jenis_Buku{
			{
				IDJenis:   "01H019WHW4A45KH65AH9XW89PC",
				JenisBuku: "ini jenis buku",
				Deskripsi: "ini deskripsi",
			},
		}
		JenbukRepoInterfaceMock.On("CariJenisBuku", "ini jenis buku").Return(JenbukDummy, nil).Once()

		data, err := JenbukServices.FindJenisBuku("ini jenis buku")
		assert.Equal(t, nil, err)
		assert.Equal(t, JenbukDummy, data)
	})

	t.Run("Get All find data empty", func(t *testing.T) {
		JenbukDummy := []models.Jenis_Buku{}
		JenbukRepoInterfaceMock.On("CariJenisBuku", "ini jenis buku").Return(JenbukDummy, nil).Once()

		data, err := JenbukServices.FindJenisBuku("ini jenis buku")
		assert.Equal(t, nil, err)
		assert.Equal(t, JenbukDummy, data)
	})
}

func TestUpdateDataJenbuk(t *testing.T) {
	var DummyEdit models.Jenis_Buku
	JenbukRepoInterfaceMock.On("HitungDataJenisBuku", mock.Anything).Return(int64(1)).Once()
	JenbukRepoInterfaceMock.On("UpdateJenisBuku", mock.AnythingOfType("models.Jenis_Buku")).Return(models.Jenis_Buku{}, nil).Once()
	dataUpdate := models.Jenis_Buku{
		IDJenis:   "01H019WHW4A45KH65AH9XW89PC",
		JenisBuku: "ini jenis buku",
		Deskripsi: "ini deskripsi",
	}
	data, err := JenbukServices.UpdateJenisBuku(dataUpdate)
	assert.Equal(t, nil, err)
	assert.Equal(t, DummyEdit, data)
}

func TestHapusDataJenbuk(t *testing.T) {
	JenbukDummy := models.Jenis_Buku{
		IDJenis:   "01H019WHW4A45KH65AH9XW89PC",
		JenisBuku: "ini jenis buku",
		Deskripsi: "ini deskripsi",
	}
	JenbukRepoInterfaceMock.On("HitungDataJenisBuku", mock.Anything).Return(int64(1)).Once()
	JenbukRepoInterfaceMock.On("DeleteJenisBuku", mock.Anything).Return(nil).Once()
	err := JenbukServices.HapusJenisBuku(JenbukDummy.IDJenis)
	assert.Equal(t, nil, err)
}
