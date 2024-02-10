package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func main() {

	// Directorio del cual queremos listar los archivos
	directorio := "./files"

	// Llama a la función que lista los archivos
	archivos, err := listarArchivos(directorio)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("por favor selecciona un numero de archivo")
	for i := 0; i < len(archivos); i++ {
		fmt.Printf("- %d -> %s \n", i, archivos[i])
	}
	var indice int
	fmt.Scanln(&indice)
	// Abre el archivo
	var ruta string = "./files/" + archivos[indice]
	file, err := os.Open(ruta)
	if err != nil {
		fmt.Println("Error al abrir el archivo:", err)
		return
	}
	defer file.Close()

	// Crea un nuevo lector bufio
	scanner := bufio.NewScanner(file)
	var contador int = 0
	// Itera sobre cada línea del archivo
	for scanner.Scan() {
		// Lee la línea actual y la imprime
		fmt.Println(scanner.Text(), isAnagram(scanner.Text()))
		if isAnagram(scanner.Text()) {
			contador++
		}
	}
	fmt.Println(contador)
	// Verifica errores de escaneo
	if err := scanner.Err(); err != nil {
		fmt.Println("Error al escanear el archivo:", err)
	}
}

func listarArchivos(directorio string) ([]string, error) {
	var archivos []string

	// Lee el directorio
	err := filepath.Walk(directorio, func(path string, info os.FileInfo, err error) error {
		// Ignora el error de directorio inaccesible
		if err != nil {
			return nil
		}

		// Ignora los directorios y agrega los archivos a la lista
		if !info.IsDir() {
			archivos = append(archivos, info.Name())
		}
		return nil
	})

	// Maneja los errores
	if err != nil {
		return nil, err
	}

	return archivos, nil
}
func isAnagram(addr string) bool {
	a := strings.Split(addr, ":")

	if len(a[0]) == len(a[1]) {
		if orderAlfabetic(strings.ToLower(a[0])) == orderAlfabetic(strings.ToLower(a[1])) {
			return true
		}
	}
	return false
}

func orderAlfabetic(addr string) string {
	letras := []rune(addr)
	sort.Slice(letras, func(i, j int) bool { return letras[i] < letras[j] })
	palabraOrdenada := string(letras)
	return palabraOrdenada
}
