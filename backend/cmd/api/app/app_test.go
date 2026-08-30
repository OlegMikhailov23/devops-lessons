package app

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"

	"gitlab.praktikum-services.ru/Stasyan/momo-store/cmd/api/dependencies"
)

func TestFakeAppIntegrational(t *testing.T) {
	store, err := dependencies.NewFakeDumplingsStore()
	assert.NoError(t, err)
	app, err := NewInstance(store)
	assert.NoError(t, err)

	t.Run("create_order", func(t *testing.T) {
		for i := 1; i <= 10; i++ {
			t.Run("id"+strconv.Itoa(i), func(t *testing.T) {
				r := httptest.NewRequest("POST", "/orders", nil)
				w := httptest.NewRecorder()
				app.CreateOrderController(w, r)

				assert.Equal(t, http.StatusOK, w.Code)
				assert.Equal(t, "application/json", w.Header().Get("Content-Type"))
				fmt.Fprintln(os.Stdout, "_____")
				fmt.Fprintln(os.Stdout, w.Body.String())
				fmt.Fprintln(os.Stdout, "_____")

				expectedJSON, err := json.Marshal(map[string]interface{}{"id": i})
				assert.NoError(t, err)
				assert.JSONEq(t, string(expectedJSON), w.Body.String())
			})
		}
	})

	t.Run("list_dumplings", func(t *testing.T) {
		r := httptest.NewRequest("GET", "/packs", nil)
		w := httptest.NewRecorder()
		app.ListDumplingsController(w, r)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, "application/json", w.Header().Get("Content-Type"))

		fmt.Fprintln(os.Stdout, "_____")
		fmt.Fprintln(os.Stdout, w.Body.String())
		fmt.Fprintln(os.Stdout, "_____")

		expectedJSON := "{\"results\":[{\"id\":1,\"name\":\"Пельмени\",\"price\":5,\"description\":\"С говядиной\",\"image\":\"https://momo-images-040-14.storage.yandexcloud.net/ziyazi.jpg\"},{\"id\":2,\"name\":\"Хинкали\",\"price\":3.5,\"description\":\"Со свининой\",\"image\":\"https://momo-images-040-14.storage.yandexcloud.net/hinkali.jpg\"},{\"id\":3,\"name\":\"Манты\",\"price\":2.75,\"description\":\"С мясом молодых бычков\",\"image\":\"https://momo-images-040-14.storage.yandexcloud.net/manti.jpg\"},{\"id\":4,\"name\":\"Буузы\",\"price\":4,\"description\":\"С телятиной и луком\",\"image\":\"https://momo-images-040-14.storage.yandexcloud.net/buzi.jpg\"},{\"id\":5,\"name\":\"Цзяоцзы\",\"price\":7.25,\"description\":\"С говядиной и свининой\",\"image\":\"https://momo-images-040-14.storage.yandexcloud.net/ziyazi.jpg\"},{\"id\":6,\"name\":\"Гедза\",\"price\":3.5,\"description\":\"С соевым мясом\",\"image\":\"https://momo-images-040-14.storage.yandexcloud.net/gedza.jpg\"},{\"id\":7,\"name\":\"Дим-самы\",\"price\":2.65,\"description\":\"С уткой\",\"image\":\"https://momo-images-040-14.storage.yandexcloud.net/dimsam.jpg\"},{\"id\":8,\"name\":\"Момо\",\"price\":5,\"description\":\"С бараниной\",\"image\":\"https://momo-images-040-14.storage.yandexcloud.net/momo.jpg\"},{\"id\":9,\"name\":\"Вонтоны\",\"price\":4.1,\"description\":\"С креветками\",\"image\":\"https://momo-images-040-14.storage.yandexcloud.net/vonton.jpg\"},{\"id\":10,\"name\":\"Баоцзы\",\"price\":4.2,\"description\":\"С капустой\",\"image\":\"https://momo-images-040-14.storage.yandexcloud.net/baozi.jpg\"},{\"id\":11,\"name\":\"Кундюмы\",\"price\":5.45,\"description\":\"С грибами\",\"image\":\"https://momo-images-040-14.storage.yandexcloud.net/kondumi.jpg\"},{\"id\":12,\"name\":\"Курзе\",\"price\":3.25,\"description\":\"С крабом\",\"image\":\"https://momo-images-040-14.storage.yandexcloud.net/kurze.jpg\"},{\"id\":13,\"name\":\"Бораки\",\"price\":4,\"description\":\"С говядиной и бараниной\",\"image\":\"https://momo-images-040-14.storage.yandexcloud.net/baraki.jpg\"},{\"id\":14,\"name\":\"Равиоли\",\"price\":2.9,\"description\":\"С рикоттой\",\"image\":\"https://momo-images-040-14.storage.yandexcloud.net/ravioly.jpg\"}]}\n"

		assert.NoError(t, err)
		assert.JSONEq(t, string(expectedJSON), w.Body.String())
	})

	t.Run("healthcheck", func(t *testing.T) {
		r := httptest.NewRequest("GET", "/health", nil)
		w := httptest.NewRecorder()
		app.HealthcheckController(w, r)

		assert.Equal(t, http.StatusOK, w.Code)
	})
}
