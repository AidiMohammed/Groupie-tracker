package tools	

import (
	"bufio"
	"os"
	"strings"
)

// loadEnvFile lit un ficher .env et charger les varibales d'enviranment 
func LoadEnvFile(filePath string) error {
	//ouvrire le ficher .env
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Lires les lignes de fichier 
	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		line := scanner.Text()
		
		//ignore les lignes vides ou les commentaires
		if strings.TrimSpace(line) == "" || strings.HasPrefix(line,"#"){
			continue
		}

		parts := strings.SplitN(line, "=",2)
		if len(parts) != 2 {
			continue 
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		value = strings.Trim(value, `"'`)
		os.Setenv(key,value)
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

// Function to check if a file exists
func FileExists(filepath string) bool {

	file, err := os.Stat(filepath)
	return !os.IsNotExist(err) && !file.IsDir()
}