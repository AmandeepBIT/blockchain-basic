package blockchain

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/google/uuid"
)

func ReadDir(dirname string) []os.FileInfo {

	files, err := ioutil.ReadDir(dirname)
	if err != nil {
		log.Fatal(err)
	}
	return files
}

func ReadFile(file string) []byte {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	return data
}

func ConvertDecryptFiles() {

	files := ReadDir(EncryptedLoc)

	filename := DecryptedLoc + "final.txt"

	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {

		fmt.Println(f.Name())
		databyte := ReadFile(EncryptedLoc + f.Name())

		data := DecryptFile(string(databyte))
		length, err := io.WriteString(file, data)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("file length is ", length, data)
	}
	defer file.Close()
	databyte := ReadFile(filename)
	fmt.Println("Actual data of saved file is ", string(databyte))

}

func CreateChunksAndEncrypt() {

	split := 4
	file, err := os.Open(fileNameC)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	texts := make([]string, 0)
	for scanner.Scan() {
		text := scanner.Text()
		texts = append(texts, text)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	lengthPerSplit := len(texts) / split
	for i := 0; i < split; i++ {
		if i+1 == split {
			chunkTexts := texts[i*lengthPerSplit:]
			fmt.Println(chunkTexts)
			writefile(EncryptFile(strings.Join(chunkTexts, "\n")))
		} else {
			chunkTexts := texts[i*lengthPerSplit : (i+1)*lengthPerSplit]
			fmt.Println(chunkTexts)
			writefile(EncryptFile(strings.Join(chunkTexts, "\n")))
		}
	}
}

func writefile(data string) {
	file, err := os.Create("./chunks/encrypted/chunks-" + uuid.New().String() + ".txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	file.WriteString(data)
}
