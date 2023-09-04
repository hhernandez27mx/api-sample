package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"azt.com/api-sample/model/entity"
)

func HomeContact(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Página de contacto")
}

func AdddContactHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Se esperaba una solicitud POST", http.StatusMethodNotAllowed)
		return
	}

	var data entity.Contact
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Error al decodificar JSON", http.StatusBadRequest)
		return
	}

	// Procesar los datos recibidos
	result := fmt.Sprintf("Hola, %s. Tienes %d años y tu correo es %s.", data.Name, data.Age, data.Email)

	// Enviar una respuesta al cliente
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := ResponseApi{0, result}
	json.NewEncoder(w).Encode(response)
}
