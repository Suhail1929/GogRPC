package data_test

import (
	"Bakri-Souhail/GoGrpcClient/data"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func Test_Data(t *testing.T) {
	/* ReadFiles */

	// Création d'un fichier temporaire
	tmpfile, err := os.CreateTemp("", "test*.json")
	if err != nil {
		t.Fatal(err)
	}
	
	// Écrire des données JSON dans le fichier temporaire
	_, err = tmpfile.WriteString(`{"key": "value"}`)
	if err != nil {
		t.Fatal(err)
	}
	// Lire les données du fichier temporaire
	datas, err := data.ReadFiles(tmpfile)
	if err != nil {
		t.Fatalf("Erreur inattendue : %v", err)
	}
	
	// Fermer le fichier pour s'assurer que les données sont écrites sur le disque
	if err := tmpfile.Close(); err != nil {
		t.Fatal(err)
	}

	
	// Vérifier si les données lues sont correctes
	expectedData := data.Data{Json: `{"key": "value"}`}
	assert.Equal(t, expectedData, datas, "Les données lues ne correspondent pas")
	
	defer os.Remove(tmpfile.Name()) // Supprimer le fichier temporaire une fois le test terminé
	
	/* OpenJsonFile */

	// Testez l'ouverture d'un fichier JSON existant
	file, err := data.OpenJsonFile("1" , "")
	assert.NoError(t, err)
	assert.NotNil(t, file)
	file.Close()

	// Testez l'ouverture d'un fichier JSON inexistant
	_, err = data.OpenJsonFile("6" , "")
	assert.Error(t, err)
}
