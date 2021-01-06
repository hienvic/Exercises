package main 
  
import ( 
    "fmt"
    "strings"
    "io/ioutil"
    "encoding/json"
    "log"
) 

type Instance struct {
    VCPU float64 `json:"vCPU"`
    VRam float64 `json:"vRam"`
    Counts float64 `json:"counts"`
}

func checkFile(filePath string) bool{
    if strings.HasSuffix(filePath,".json") {
        return true
    } else {
        return false
    }
}

func inputFilePath() string{
    var filePath string
    fmt.Print("Enter config path: ")
    fmt.Scanln(&filePath)
    check := checkFile(filePath)
    if !check {
        fmt.Print("Invalid file, enter path: ")
        fmt.Scanln(&filePath)
    } 
    return filePath
}

func readFile() ([]byte, error) {
    data, err := ioutil.ReadFile(inputFilePath())
    if err != nil {
        fmt.Println(err)
    } 
    return data, err
}

func main() {
	oldMap := make(map[string]Instance)
	for {
		data, _ := readFile()
		var result map[string][]interface{}
		newMap := make(map[string]Instance)
		if err := json.Unmarshal(data, &result); err != nil {
			log.Fatalf("JSON unmarshaling failed: %s", err)
		}
		for _, v := range result["Instances"] {
			instance := v.(map[string]interface{})
			newMap[instance["type"].(string)] = Instance{
				instance["vCPU"].(float64),
				instance["vRam"].(float64),
				instance["counts"].(float64),
			}
		}
		for key, _ := range newMap {
			_, ok := oldMap[key]
			if ok == false {
				fmt.Println(key, "Provision", newMap[key].Counts)
			} else {
				if newMap[key].Counts > oldMap[key].Counts {
					fmt.Println(key, "Provision", newMap[key].Counts - oldMap[key].Counts)
				} else if newMap[key].Counts < oldMap[key].Counts {
					fmt.Println(key, "Delete", oldMap[key].Counts - newMap[key].Counts)
				} else {
                   
				}
			}
		}
		oldMap = newMap
	}
}