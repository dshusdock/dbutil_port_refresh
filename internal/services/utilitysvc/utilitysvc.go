package utilitysvc

import (
	"bufio"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"dshusdock/go_project/internal/constants"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net"
	"os"
	"strings"
)

type UtilitySvc struct {
}

type DBLoadParams struct {
	CmdType 	string
	DBLdrInfo 	constants.DBLoaderInfo
}

type Config struct {
    Host_ip  	string `json:"host_ip"`
	App_port   	int    `json:"app_port"`
	Log_level  	string `json:"log_level"`
}

func NewUtilitySvc() *UtilitySvc {
	return &UtilitySvc{}
}

func (u *UtilitySvc) CmdBuilder(p DBLoadParams) string {

	switch p.CmdType {
	case "find":
		return "find " + GroomAllicatStr(p.DBLdrInfo.FileDir) + " -name " + p.DBLdrInfo.SQLFileName
	case "copy":
		return "cp " + p.DBLdrInfo.FileDir + "/" + p.DBLdrInfo.SQLFileName + " /tmp"
	case "unzip":
		return "gzip -d /tmp/" + p.DBLdrInfo.SQLFileName
	case "createdb":
		return "mysql -udunkin -pdunkin123 -e 'create database " + p.DBLdrInfo.DBName + "';"
	case "loaddb":
		sql := strings.Split(p.DBLdrInfo.SQLFileName, ".")
		return "mysql -udunkin -pdunkin123 " + p.DBLdrInfo.DBName + " < /tmp/" + sql[0] + ".sql"
	}
	return ""
}

func GroomAllicatStr(str string) string {
	str = strings.Replace(str, "\\", "/", -1)
	str = strings.Replace(str, "$", "", -1)
	return str
}

func GenerateRandomString(length int) (string, error) {
	// Create a byte slice to store random bytes
	randomBytes := make([]byte, length)

	// Read random bytes from the crypto/rand package
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}

	// Encode the random bytes into a base64 string
	randomString := base64.URLEncoding.EncodeToString(randomBytes)

	// Return the random string
	return randomString[:length], nil
}

func EncryptValue(text string) ([]byte, error) {
	// key := []byte("passphrasewhichneedstobe32bytes!")
	key := []byte("keepstrackofthepreviousvalueok!!")
	c, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	return gcm.Seal(nonce, nonce, []byte(text), nil), nil
}	

func DecryptValue(ciphertext []byte) (string, error) {
	// key := []byte("passphrasewhichneedstobe32bytes!")
	key := []byte("keepstrackofthepreviousvalueok!!")
	c, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return "", err
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return "", err
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}

func GetHostIP() string {
	addrs, err := net.InterfaceAddrs()
    if err != nil {
        slog.Info(err.Error())
        return ""
    }

    for _, addr := range addrs {
        if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
            if ipNet.IP.To4() != nil {
				slog.Info("IP address", "IP", ipNet.IP.String())
            }
        }
    }
	return ""
}

func ReadTxtConfigFile() {
	// Open the config file
    file, err := os.Open("config.txt")
    if err != nil {
        slog.Info("Error opening file:", err)
        return
    }
    defer file.Close()

    // Create a map to store the config values
    config := make(map[string]string)

    // Read the file line by line
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        // Split the line into name and value
        parts := strings.SplitN(line, "=", 2)
        if len(parts) == 2 {
            name := strings.TrimSpace(parts[0])
            value := strings.TrimSpace(parts[1])
            config[name] = value
        }
    }

    if err := scanner.Err(); err != nil {
        slog.Info("Error reading file:", err)
        return
    }

    // Print the config values
    for name, value := range config {
    	slog.Info("%s: %s\n", name, value)
    }
}

func ReadJsonConfigFile() (ip string, port string, log_level string) {
	// Open the config file
	file, err := os.Open("app_config.json")
	if err != nil {
		slog.Info("Error opening file:", err)
		return
	}
	defer file.Close()

	// Read the file contents
	byteValue, err := io.ReadAll(file)
	if err != nil {
		slog.Info("Error reading file:", err)
		return
	}

	// Unmarshal the JSON data into a Config struct
	var config Config
	err = json.Unmarshal(byteValue, &config)
	if err != nil {
		slog.Info("Error parsing JSON:", "error", err)
		return
	}

	// Print the config values
	slog.Info("Host", "IP", config.Host_ip)
	slog.Info("Host", "Port", config.App_port)
	slog.Info("Log", "Level", config.Log_level)
	return config.Host_ip, fmt.Sprintf("%d", config.App_port), config.Log_level
  
}


		
