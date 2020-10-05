// main.go
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Article - Our struct for all articles
type Article struct {
	Id      string `json:"Id"`
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

var Articles []Article

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage with the Dockerile!")
	fmt.Println("Endpoint Hit: homePage")
}

func simpleSearch(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: SimpleSearch")
	fmt.Fprintf(w, "<?xml version=\"1.0\"?><response method=\"simplesearch\" sessionkey=\"96C94952-1618q4B9B-85B1-57C4EC616395\" success=\"Y\"><errors></errors><request><auth password=\"Hotelbeds\" username=\"Hotelbedscruises\" /> <method action=\"simplesearch\" currency=\"USD\" sessionkey=\"96C94952-1618q4B9B-85B1-57C4EC616395\" sitename=\"hotelbeds.site.traveltek.net\" status=\"Live\" teamid=\"0\" type=\"cruise\" userid=\"3148318\"><searchdetail adults=\"2\" children=\"0\" enddate=\"2021-07-30\" resultkey=\"default\" sid=\"42360\" startdate=\"2021-07-01\" type=\"cruise\" /></method></request><results><cruise airbalcony=\"0\" airbalconypricecode=\"\" airinside=\"0\" airinsidepricecode=\"\" airoutside=\"0\" airoutsidepricecode=\"\" airport=\"\" airportname=\"\" airsuite=\"0\" airsuitepricecode=\"\" altvoyagecode=\"\" codetocruiseid=\"1582325\" copyandmedia=\"\" cruisebalcony=\"0\" cruisebalconypricecode=\"\" cruisebalconypricecodedescription=\"\" cruisebalconypricecodeiconurl=\"\" cruiseid=\"255940\" cruiseinside=\"0\" cruiseinsidepricecode=\"\" cruiseinsidepricecodedescription=\"\" cruiseinsidepricecodeiconurl=\"\" cruiseoutside=\"0\" cruiseoutsidepricecode=\"\" cruiseoutsidepricecodedescription=\"\" cruiseoutsidepricecodeiconurl=\"\" cruisesuite=\"60308.00\" cruisesuitepricecode=\"N\" cruisesuitepricecodedescription=\"\" cruisesuitepricecodeiconurl=\"\" currency=\"USD\" departuk=\"N\" description=\"\" displaycruiseonly=\"1\" domesticdeparture=\"0\" enddate=\"2021-08-29\" engine=\"Seabourn\" enginesource=\"live\" groupids=\"\" hascruiseonly=\"Y\" hasflights=\"N\" hidden=\"N\" iata=\"\" localpricing=\"Y\" marketid=\"9\" name=\"Lands Of The Sagas\" ncf=\"5759.00\" nettprice=\"60308.00\" nights=\"53\" nofly=\"N\" nofuelsupplement=\"\" ownerid=\"system\" price=\"60308.00\" priority=\"999\" privatenotes=\"\" ratecode=\"N\" ratecodedescription=\"\" ratecodeiconurl=\"\" resultiscruiseonly=\"Y\" resultkey=\"default\" resultno=\"302_313.0\" resultweight=\"0\" returndate=\"2021-08-29\" returnuk=\"N\" roundtrip=\"N\" saildate=\"2021-07-07\" sailnights=\"53\" scurrency=\"USD\" seadays=\"14\" searchdetail=\"\" searchno=\"0\" senior=\"0\" soldout=\"0\" special=\"no\" specialsgroup=\"\" sprice=\"60308\" startdate=\"2021-07-07\" startprice=\"30154.00\" stoplive=\"0\" systemgroup=\"\" taxes=\"0.00\" type=\"cruise\" voyagecode=\"6134B\" whatsincluded=\"\" zoneid=\"\"><cruisesuiteprices><price appliesto=\"booking\" marker=\"nett\" name=\"Supplier Price\" runorder=\"0\" type=\"set\" value=\"60308\" /><price appliesto=\"booking\" marker=\"supplier\" name=\"Supplier Price\" runorder=\"5\" type=\"set\" value=\"60308\" /></cruisesuiteprices><grade><cabin></cabin>       <rate></rate></grade><heroimg></heroimg><line code=\"SBN\" engine=\"Seabourn\" id=\"24\" logourl=\"http://static.traveltek.net/cruisepics/logos/seabourn.gif\" name=\"Seabourn\" niceurl=\"seabourn\" /><ports><port id=\"654\" name=\"Reykjavik\" /><port id=\"292\" name=\"Skjolden\" /><port id=\"5621\" name=\"Tasiilaq, Greenland\" /><port id=\"5443\" name=\"Lerwick, Scotland\" /><port id=\"617\" name=\"Kirkwall, Scotland\" /><port id=\"659\" name=\"Scrabster\" /><port id=\"3132\" name=\"Aberdeen\" /><port id=\"523\" name=\"Newcastle, UK\" /><port id=\"469\" name=\"Dover, UK\" /><port id=\"865\" name=\"Portland, Dorset\" /><port id=\"850\" name=\"Bantry, County Cork\" /><port id=\"3417\" name=\"Foynes\" /><port id=\"1489\" name=\"Galway\" /><port id=\"2210\" name=\"Killybegs\" /><port id=\"782\" name=\"Londonderry\" /><port id=\"454\" name=\"Belfast\" /><port id=\"86\" name=\"Dublin\" /></ports><prices><price appliesto=\"booking\" marker=\"nett\" name=\"Supplier Price\" runorder=\"0\" type=\"set\" value=\"60308\" /><price appliesto=\"booking\" marker=\"supplier\" name=\"Supplier Price\" runorder=\"5\" type=\"set\" value=\"60308\" /></prices><regions><region name=\"United Kingdom\" regionid=\"16\" /><region name=\"Polar Regions\" regionid=\"18\" /></regions><sessiondetail></sessiondetail><ship code=\"SQ\" id=\"564\" imagecaption=\"Seabourn Quest\" imageurl=\"http://static.traveltek.net/cruisepics/local_shipimages/1595326033.png\" name=\"Seabourn Quest\" niceurl=\"seabourn/seabourn-quest\" rating=\"6\" smallimageurl=\"http://static.traveltek.net/cruisepics/local_shipimages_small/1595326033.png\" /><uniqueportids>782</uniqueportids><uniqueportids>1298</uniqueportids><uniqueportids>1234</uniqueportids><uniqueportids>86</uniqueportids>           <uniqueportnames>Londonderry</uniqueportnames><upropdata></upropdata></cruise></results></response>")
	//  json.NewEncoder(w).Encode(Articles)
}

/*func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    key := vars["id"]

    for _, article := range Articles {
        if article.Id == key {
            json.NewEncoder(w).Encode(article)
        }
    }
}*/
func cabinGrades(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<?xml version=\"1.0\"?><response method=\"getcabingrades\" success=\"N\"><errors><error code=\"420\" text=\"Error getting rate codes. HASH(0x7f245d604460)\" /></errors><request><auth password=\"Hotelbeds\" username=\"Hotelbedscruises\" /><method action=\"getcabingrades\" language=\"en\" ratecode=\"\" resultno=\"302_0.0\" sessionkey=\"107E5256_72F8u4347-803A-5CFD384926B3\" /></request></response>")
}
func rateCodes(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<?xml version=\"1.0\"?><response method=\"getratecodes\" sessionkey=\"107E5256_72F8u4347-803A-5CFD384926B3\" success=\"Y\"><errors></errors><request><auth password=\"Hotelbeds\" username=\"Hotelbedscruises\" /><method action=\"getratecodes\" language=\"en\" resultno=\"302_0.0\" sessionkey=\"107E5256_72F8u4347-803A-5CFD384926B3\" /></request><results><farecodes Code=\"5244\" ConditionCode=\"0067/00000000/RSO51SAV\" Text=\"Invalid Sail Date\" /></results></response>")
}

func createNewArticle(w http.ResponseWriter, r *http.Request) {
	// get the body of our POST request
	// unmarshal this into a new Article struct
	// append this to our Articles array.
	reqBody, _ := ioutil.ReadAll(r.Body)
	var article Article
	json.Unmarshal(reqBody, &article)
	// update our global Articles array to include
	// our new Article
	Articles = append(Articles, article)

	json.NewEncoder(w).Encode(article)
}

func deleteArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for index, article := range Articles {
		if article.Id == id {
			Articles = append(Articles[:index], Articles[index+1:]...)
		}
	}

}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/simpleSearch", simpleSearch).Methods("POST")
	myRouter.HandleFunc("/cabinGrades", cabinGrades)
	myRouter.HandleFunc("/article/{id}", deleteArticle).Methods("DELETE")
	myRouter.HandleFunc("/rateCodes", rateCodes)
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	Articles = []Article{
		Article{Id: "1", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		Article{Id: "2", Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}
	handleRequests()
}
