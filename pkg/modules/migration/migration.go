package migration

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
)

// MigrateFiles maneja la migración de archivos del sistema de archivos
func MigrateFiles(ctx context.Context, sourcePath string, destPath string) error {
	if _, err := os.Stat(sourcePath); os.IsNotExist(err) {
		return fmt.Errorf("ruta de origen no existe: %v", err)
	}

	if err := os.MkdirAll(destPath, 0755); err != nil {
		return fmt.Errorf("no se pudo crear directorio destino: %v", err)
	}

	return filepath.Walk(sourcePath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Obtener ruta relativa
		relPath, err := filepath.Rel(sourcePath, path)
		if err != nil {
			return fmt.Errorf("error al obtener ruta relativa: %v", err)
		}

		destFilePath := filepath.Join(destPath, relPath)

		if info.IsDir() {
			return os.MkdirAll(destFilePath, 0755)
		}

		// Copiar archivo
		content, err := os.ReadFile(path)
		if err != nil {
			return fmt.Errorf("error al leer archivo: %v", err)
		}

		if err := os.WriteFile(destFilePath, content, 0644); err != nil {
			return fmt.Errorf("error al escribir archivo: %v", err)
		}

		return nil
	})
}
