package utils

import (
	"bytes"
	"log"
	"net/http"
	"time"
)

//Ayuda a interceptar las respuestas de tus handlers
type responseRecorder struct {
	http.ResponseWriter
	statusCode     int
	body           *bytes.Buffer // Aquí guardaremos una copia del mensaje
}

//Se intercepta el momento en que se escribe el status (200, 400, etc)
func (rec *responseRecorder) WriteHeader(statusCode int) {
	rec.statusCode = statusCode
	rec.ResponseWriter.WriteHeader(statusCode)
}

//Se intercepta el momento en que se escribe el cuerpo de la respuesta (el mensaje de error o éxito)
func (rec *responseRecorder) Write(b []byte) (int, error) {
	rec.body.Write(b)
	return rec.ResponseWriter.Write(b) // Lo enviamos al frontend
}

func Logger(next http.HandlerFunc) http.HandlerFunc{
return func(w http.ResponseWriter, r *http.Request) {
		inicio := time.Now()

		//Se instancia el recorder para interceptar la respuesta del handler
		recorder := &responseRecorder{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
			body:           &bytes.Buffer{},
		}

		//Esto ejecuta el handler real, pero con nuestro recorder para interceptar la respuesta
		next(recorder, r)

		duracion := time.Since(inicio)

		//Se evalua el status code para determinar si es un error o no. Si es un error, se imprime el mensaje de error en el log.
		if recorder.statusCode >= 400 {
			// Si el status es 400 o más, es error
			mensajeError := bytes.TrimSpace(recorder.body.Bytes()) // Quitamos saltos de línea
			log.Printf("[ERROR] %s %s | Status: %d | Tiempo: %v | Razón: %s", 
				r.Method, r.URL.RequestURI(), recorder.statusCode, duracion, mensajeError)
		} else {
			// Si todo salió bien (200, 201) se hace un log normal sin el mensaje de error
			log.Printf("[INFO] %s %s | Status: %d | Tiempo: %v", 
				r.Method, r.URL.RequestURI(), recorder.statusCode, duracion)
		}
	}
}