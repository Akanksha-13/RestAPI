package apis

//validate inputs
//call the services asynchrnously
//return the response with proper status code
import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"

	"github.com/A292460/CRUDAPI/models"
	"github.com/xeipuuv/gojsonschema"
)

func CreateValkeyInstance(w http.ResponseWriter, r *http.Request) {
	schemaFilePath := "models/valkey_instance.json"
	schemaBytes, err := os.ReadFile(schemaFilePath)
	if err != nil {
		slog.Error("Failed to read schema file", "error", err.Error(), "status", http.StatusInternalServerError)
		http.Error(w, err.Error()+"\n Unable to read schema file", http.StatusInternalServerError)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		slog.Error("Failed to read request body", "error", err.Error(), "status", http.StatusInternalServerError)
		http.Error(w, err.Error()+"\n Unable to read request body", http.StatusInternalServerError)
		return
	}

	schemaLoader := gojsonschema.NewStringLoader(string(schemaBytes))
	// fmt.Println(schemaLoader)
	documentLoader := gojsonschema.NewStringLoader(string(body))
	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		slog.Error("Failed to validate schema", "error", err.Error(), "status", http.StatusInternalServerError)
		http.Error(w, err.Error()+"\n Unable to validate schema", http.StatusInternalServerError)
		return
	}
	if result.Valid() {
		fmt.Println("The Json Payload is valid")
		slog.Info("Validation successful", "status", http.StatusOK)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("The JSON Payload is valid"))
	} else {
		fmt.Println("The Json Payload is not valid. see errors :")
		for _, desc := range result.Errors() {
			fmt.Printf("- %s: %s\n", desc.Field(), desc.Description())
			slog.Error("Validation error", "field", desc.Field(), "description", desc.Description(), "status", http.StatusBadRequest)
		}
		http.Error(w, "The JSON Payload is not valid. See server logs for details.", http.StatusBadRequest)
	}
	var valkeyInstance models.ValkeyRequest
	err = json.Unmarshal(body, &valkeyInstance) //
	if err != nil {
		slog.Error("Failed to unmarshal JSON", "error", err.Error(), "status", http.StatusBadRequest)
		http.Error(w, err.Error()+"\n Unable to unmarshal JSON", http.StatusBadRequest)
		return
	}
	slog.Info("Unmarshalling successful", "valkeyInstance", valkeyInstance)

}
