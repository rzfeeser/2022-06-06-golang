package main
 
import (
    "encoding/xml"
    "io/ioutil"
)
 
type notes struct {
    To      string `xml:"to"`
    From    string `xml:"from"`
    Subject string `xml:"subject"`
    Body    string `xml:"body"`
}
 
func main() {
    note := notes{To: "Kevin McCallister",
        From:    "Mom",
        Subject: "Be home soon!",
        Body:    "Getting the first flight home! Take care of the house until we return.",
    }
 
    file, _ := xml.MarshalIndent(note, "", " ")
 
    _ = ioutil.WriteFile("HomeAlone.xml", file, 0644)
 
}

