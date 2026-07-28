package rest

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"dbms-project/internal/db"
)

type TransportRoutePayload struct {
	UID            int32  `json:"u_id"`
	TransportRoute string `json:"transport_route"`
	EstTravelTime  string `json:"est_travel_time"`
}

func (h *Handler) HandleGetUniversityTransportRoutes(w http.ResponseWriter, r *http.Request) {
	uIDStr := r.URL.Query().Get("u_id")
	uID, err := strconv.Atoi(uIDStr)
	if err != nil || uID <= 0 {
		http.Error(w, "Invalid or missing u_id", http.StatusBadRequest)
		return
	}

	routes, err := h.Queries.GetUniversityTransportRoutes(r.Context(), int32(uID))
	if err != nil {
		http.Error(w, "Failed to retrieve transport routes: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if routes == nil {
		routes = []db.UniversityTransport{}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(routes)
}

func (h *Handler) HandleCreateUniversityTransportRoute(w http.ResponseWriter, r *http.Request) {
	var req TransportRoutePayload
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid payload", http.StatusBadRequest)
		return
	}

	route := strings.TrimSpace(req.TransportRoute)
	estTime := strings.TrimSpace(req.EstTravelTime)

	if req.UID <= 0 || route == "" || estTime == "" {
		http.Error(w, "u_id, transport_route, and est_travel_time are required", http.StatusBadRequest)
		return
	}

	_, err := h.Queries.InsertUniversityTransportRoute(r.Context(), db.InsertUniversityTransportRouteParams{
		UID:            req.UID,
		TransportRoute: route,
		EstTravelTime:  estTime,
	})
	if err != nil {
		http.Error(w, "Failed to create transport route: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Transport route created successfully",
	})
}

func (h *Handler) HandleUpdateUniversityTransportRoute(w http.ResponseWriter, r *http.Request) {
	var req TransportRoutePayload
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid payload", http.StatusBadRequest)
		return
	}

	route := strings.TrimSpace(req.TransportRoute)
	estTime := strings.TrimSpace(req.EstTravelTime)

	if req.UID <= 0 || route == "" || estTime == "" {
		http.Error(w, "u_id, transport_route, and est_travel_time are required", http.StatusBadRequest)
		return
	}

	err := h.Queries.UpdateUniversityTransportRoute(r.Context(), db.UpdateUniversityTransportRouteParams{
		UID:            req.UID,
		TransportRoute: route,
		EstTravelTime:  estTime,
	})
	if err != nil {
		http.Error(w, "Failed to update transport route: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Transport route updated successfully",
	})
}

func (h *Handler) HandleDeleteUniversityTransportRoute(w http.ResponseWriter, r *http.Request) {
	uIDStr := r.URL.Query().Get("u_id")
	uID, err := strconv.Atoi(uIDStr)
	if err != nil || uID <= 0 {
		http.Error(w, "Invalid or missing u_id", http.StatusBadRequest)
		return
	}

	route := strings.TrimSpace(r.URL.Query().Get("transport_route"))
	if route == "" {
		http.Error(w, "transport_route query parameter is required", http.StatusBadRequest)
		return
	}

	err = h.Queries.DeleteUniversityTransportRoute(r.Context(), db.DeleteUniversityTransportRouteParams{
		UID:            int32(uID),
		TransportRoute: route,
	})
	if err != nil {
		http.Error(w, "Failed to delete transport route: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Transport route deleted successfully",
	})
}
