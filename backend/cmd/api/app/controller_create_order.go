package app

import (
    "encoding/json"
    "net/http"

    "go.uber.org/zap"

    "gitlab.praktikum-services.ru/Stasyan/momo-store/internal/logger"
    "gitlab.praktikum-services.ru/Stasyan/momo-store/internal/store/dumplings"
)

func (i *Instance) CreateOrderController(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

    // accept optional items payload
    type reqItem struct {
        ProductID int64  `json:"product_id"`
        Count     uint32 `json:"count"`
    }
    var req struct {
        Items []reqItem `json:"items"`
    }
    _ = json.NewDecoder(r.Body).Decode(&req)

    var items []dumplings.OrderItem
    for _, it := range req.Items {
        items = append(items, dumplings.OrderItem{
            Pack:  dumplings.Product{ID: it.ProductID},
            Count: it.Count,
        })
    }
	id, err := i.store.CreateOrder(ctx, items...)
	if err != nil {
		logger.Log.Error("cannot create order", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).
		Encode(map[string]interface{}{
			"id": id,
		})

	i.ordersCounter.Inc()
}
