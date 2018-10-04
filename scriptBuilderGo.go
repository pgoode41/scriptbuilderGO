package main

import (
  "fmt"
  "net/http"
	"io/ioutil"
  "encoding/json"
  "os"
  "log"
  "strings"
)

func main()  {
  //Pull JSON And Store It In A Variable
  url := "https://swapi.co/api/people"
  req, _ := http.NewRequest("GET", url, nil)
  res, _ := http.DefaultClient.Do(req)
  defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

  //Parse Through JSON
  var character Character
  json.Unmarshal(body, &character)
  people := character.Results

  // Create Headers Directory If It Doesn't Exist.
  headerPath := "/opt/goJson/Headers"
  if _, err := os.Stat(headerPath); os.IsNotExist(err) {
    os.Mkdir(headerPath, 0764)
    os.Chmod(headerPath, 0764)
  }

  //Change To Header Directory.
  if err := os.Chdir("/opt/goJson/Headers"); err != nil {
    panic(err)
  }

  //List Of Illegal Characters For Variables.(Allows Spaces)
  //Use formatedJsonNameNOSpaces On Variable Directly
  badCharacterList := []string {"(", ")", `\`, `\`, ",", ".", "!", "@", "#", "$", "%", "^", "*", "+", "=", "[", "]", ";",
     `:`, "<", ">", "{", "}", "\\", "?", "/", "|", "`", "~", "_"}

  //Loop Through JSON
  for i := 0; i < len(people); i++ {
      //Create Variables From JSON
      // Using Name Method From Custom Type
      jsonName := people[i].Name
      //Removing Illegal Characters From Variable
      for _, badCharacter := range badCharacterList {
        jsonName = strings.Replace(jsonName, badCharacter, "", -1)
      }
      //Stores Data That Doesn't Contain Illegal Characters, But Allows Spaces.
      formatedJsonNameSpaces := strings.Title(jsonName)
      //Stores Data That Doesn't Contain Illegal Characters, But DOES NOT Allow Spaces.
      formatedJsonNameNOSpaces := strings.Replace(formatedJsonNameSpaces, " ", "", -1)
      //##################################################################################################################
      // VARIABLES BEING GENERATED START
      //##################################################################################################################
      namesync := formatedJsonNameNOSpaces
      //##################################################################################################################
      // VARIABLES BEING GENERATED END
      //##################################################################################################################
      //Create Header File.
      headerFile, err := os.Create(namesync+".sh")
      os.Chmod(namesync+".sh", 0764)
      if err != nil {
        log.Fatal("Cannot create header file", err)
      }
      defer headerFile.Close()
      //Creating Bash Header Files With Veriables From JSON
      fmt.Fprintln(headerFile, "#!/bin/bash")
		  fmt.Fprintln(headerFile, "#####################################################################################################################")
		  fmt.Fprintln(headerFile, "#This script will change the ComputerName and LocalHostName/ bonjour names to the name of the current logged in user.")
		  fmt.Fprintln(headerFile, "#####################################################################################################################")
		  fmt.Fprintln(headerFile, "\n")
		  fmt.Fprintln(headerFile, "#####################################################################################################################")
		  fmt.Fprintln(headerFile, "#####################################################################################################################")
		  fmt.Fprintln(headerFile, "#This pulls data gatherd from python variable generator script.")
		  fmt.Fprintln(headerFile, "#After this section is configured, the script can use those values as variables in the script.")
		  fmt.Fprintln(headerFile, "\n")
		  fmt.Fprintln(headerFile, "\n")
	    fmt.Fprintln(headerFile, "namesync="+"'"+namesync+"'")
	    fmt.Fprintln(headerFile, "\n")
		  fmt.Fprintln(headerFile, "#####################################################################################################################")
		  fmt.Fprintln(headerFile, "#####################################################################################################################")
		  fmt.Fprintln(headerFile, "\n")
		  fmt.Fprintln(headerFile, "\n")
      //
      //Grab Contents From Bash Template
      input, err := ioutil.ReadFile("/opt/goJson/BashTemplates/testMerge.sh")
      //Append Bash Template Contents To JSON Generated Header File
      file, err := os.OpenFile(namesync+".sh", os.O_WRONLY|os.O_APPEND, 0764)
        if err != nil {
                fmt.Println(err)
        }
        fmt.Fprint(file, string(input))
        //Close Merged Finished Merged File
        defer file.Close()

    }
}


//JSON Struct For Unmarshal
type Characters struct {
	Characters []Character `json:"users"`
}

type Character struct {
	Count    int         `json:"count"`
	Next     string      `json:"next"`
	Previous interface{} `json:"previous"`
	Results  []struct {
		Name      string    `json:"name"`
		Height    string    `json:"height"`
		Mass      string    `json:"mass"`
		HairColor string    `json:"hair_color"`
		SkinColor string    `json:"skin_color"`
		EyeColor  string    `json:"eye_color"`
		BirthYear string    `json:"birth_year"`
		Gender    string    `json:"gender"`
		Homeworld string    `json:"homeworld"`
		Films     []string  `json:"films"`
		Species   []string  `json:"species"`
		Vehicles  []string  `json:"vehicles"`
		Starships []string  `json:"starships"`
		//Created   time.Time `json:"created"`
	//Edited    time.Time `json:"edited"`
		URL       string    `json:"url"`
	} `json:"results"`

}
