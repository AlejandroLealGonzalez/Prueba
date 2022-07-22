package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./public"))) //Creamos el servidor para archivos estaticos
	http.HandleFunc("/upload", upload)                      //llamamos controlador de la ruta /upload
	http.ListenAndServe(":8080", nil)                       //Servidor escuchando por el puerto 8080
}

func upload(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		chn := r.FormValue("channels")                //captura el valor del canal seleccionado
		file, handle, err := r.FormFile("fileClient") //obtenemos el archivo del formulario
		if err != nil {
			fmt.Fprint(w, "Error al cargar el archivo %v", err) //mostramos el error de carga al cliente
			return
		}
		defer file.Close() //Cerramos el archivo

		data, err := ioutil.ReadAll(file) //guardamos el archivo en data
		if err != nil {
			fmt.Fprint(w, "Error al leer el archivo %v", err) //mostramos error de lectura al cliente
			return
		}
		if chn == "1" {
			err = ioutil.WriteFile("./Channels/FilesCh1/"+handle.Filename, data, 0666) //guardamos el archivo en Files, "0666" -> permisos Unix
			if err != nil {
				fmt.Fprint(w, "Error al guardar el archivo %v", err)
			}
		} else {
			err = ioutil.WriteFile("./Channels/FilesCh1/"+handle.Filename, data, 0666) //guardamos el archivo en Files, "0666" -> permisos Unix
			if err != nil {
				fmt.Fprint(w, "Error al guardar el archivo %v", err)
			}
		}
		fmt.Fprint(w, "Cargado exitosamente!!!")

	}
}
