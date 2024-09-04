package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

// Variable para almacenar la clave que se pasará con la opción --hide
var hideKey string

// secretsCmd representa el comando `secrets`
var secretsCmd = &cobra.Command{
	Use:   "secrets",
	Short: "Comando para gestionar secretos en archivos YAML",
	Run: func(cmd *cobra.Command, args []string) {
		// Verificar si se pasó la opción --hide
		if hideKey != "" {
			err := hideSecrets(hideKey)
			if err != nil {
				fmt.Printf("Error al ocultar secretos: %v\n", err)
				os.Exit(1)
			}
			fmt.Println("Se ocultaron los secretos exitosamente.")
		} else {
			fmt.Println("Debe proporcionar la opción --hide con la clave que desea ocultar.")
		}
	},
}

func init() {
	// Agregar el comando al root
	rootCmd.AddCommand(secretsCmd)

	// Definir la opción --hide con la opción corta -k para evitar conflictos con -h (help)
	secretsCmd.Flags().StringVarP(&hideKey, "hide", "k", "", "Clave cuyo valor será reemplazado por 'secret' en los archivos YAML")
}

func hideSecrets(clave string) error {
	// Obtener todos los archivos .yaml y .yml en el directorio actual
	files, err := filepath.Glob("*.y*ml")
	if err != nil {
		return fmt.Errorf("no se pudieron encontrar archivos YAML: %v", err)
	}

	// Iterar sobre cada archivo encontrado
	for _, file := range files {
		err := processFile(file, clave)
		if err != nil {
			return err
		}
	}

	return nil
}

func processFile(file, clave string) error {
	// Leer el archivo
	data, err := os.ReadFile(file)
	if err != nil {
		return fmt.Errorf("no se pudo leer el archivo %s: %v", file, err)
	}

	// Unmarshal del contenido YAML en un mapa genérico
	var content map[interface{}]interface{}
	err = yaml.Unmarshal(data, &content)
	if err != nil {
		return fmt.Errorf("error al parsear el archivo YAML %s: %v", file, err)
	}

	// Buscar y reemplazar la clave en el mapa
	changed := replaceValue(content, clave)

	// Si se realizó un cambio, guardar el archivo
	if changed {
		newData, err := yaml.Marshal(content)
		if err != nil {
			return fmt.Errorf("error al generar YAML para %s: %v", file, err)
		}
		err = os.WriteFile(file, newData, 0644)
		if err != nil {
			return fmt.Errorf("error al escribir el archivo %s: %v", file, err)
		}
		fmt.Printf("Actualizado: %s\n", file)
	}

	return nil
}

// replaceValue reemplaza el valor de la clave especificada en el mapa y sus submapas
func replaceValue(content map[interface{}]interface{}, clave string) bool {
	changed := false
	for k, v := range content {
		// Verifica si la clave es un string y coincide con la clave que buscamos
		if key, ok := k.(string); ok && key == clave {
			content[k] = "secret"
			changed = true
		} else if nestedMap, ok := v.(map[interface{}]interface{}); ok {
			// Si es un mapa anidado, llamar recursivamente
			if replaceValue(nestedMap, clave) {
				changed = true
			}
		}
	}
	return changed
}
