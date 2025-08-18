package migration

import (
	"context"
	"os"
	"path/filepath"
	"testing"
)

func TestMigrateFiles(t *testing.T) {
	// Crear directorios temporales para pruebas
	sourceDir, err := os.MkdirTemp("", "source")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(sourceDir)

	destDir, err := os.MkdirTemp("", "dest")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(destDir)

	// Crear estructura de prueba
	testFiles := map[string][]byte{
		"file1.txt":      []byte("contenido1"),
		"dir/file2.txt":  []byte("contenido2"),
		"dir/dir2/file3": []byte("contenido3"),
	}

	for path, content := range testFiles {
		fullPath := filepath.Join(sourceDir, path)
		dir := filepath.Dir(fullPath)
		if err := os.MkdirAll(dir, 0755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(fullPath, content, 0644); err != nil {
			t.Fatal(err)
		}
	}

	// Ejecutar migración
	ctx := context.Background()
	if err := MigrateFiles(ctx, sourceDir, destDir); err != nil {
		t.Errorf("MigrateFiles() error = %v", err)
	}

	// Verificar resultados
	for path, expectedContent := range testFiles {
		destPath := filepath.Join(destDir, path)
		content, err := os.ReadFile(destPath)
		if err != nil {
			t.Errorf("No se pudo leer archivo migrado %s: %v", path, err)
			continue
		}
		if string(content) != string(expectedContent) {
			t.Errorf("Contenido incorrecto para %s. Obtenido = %s, Esperado = %s",
				path, string(content), string(expectedContent))
		}
	}
}
