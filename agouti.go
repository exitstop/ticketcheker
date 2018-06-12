package main

import (
  "log"
	"fmt"
  "github.com/sclevine/agouti"
	//"strconv"
  "regexp"
  //"./sendemail"
	"time"
  "./lCommon"
  //"github.com/sqweek/dialog"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
   driver := agouti.PhantomJS()
  //driver := agouti.ChromeDriver()
  // driver := agouti.ChromeDriver(
  //   agouti.ChromeOptions("args", []string{"--headless", "--disable-gpu", "--no-sandbox"}),
  // )

  if err := driver.Start(); err != nil {
    log.Fatal("Failed to start driver:", err)
  }

	//capabilities := agouti.NewCapabilities().With("javascriptEnabled")

  //page, err := driver.NewPage(agouti.Desired(capabilities))
  page, err := driver.NewPage()
  if err != nil {
    log.Fatal("Failed to open page:", err)
  }


  if err := page.Navigate("https://tickets.fifa.com/Services/ADService.html?lang=ru"); err != nil {
    log.Fatal("Failed to navigate:", err)
  }


  type sf struct {
    name string
    cube []int
  }
  var filter []sf
  filter = append(filter, sf{name:"Матч16-Колумбия:Япония-Саранск", cube : []int{1,1,1,1,1}})

  lastCount := 0

  for true {
    duration := time.Second
    time.Sleep(duration)

    //var value string
    var count int
    var count1 int
    var count2 int
    var count3 int
    page.RunScript("return document.getElementsByClassName('categoryBox zeroAvailability').length", nil, &count)
    page.RunScript("return document.getElementsByClassName('categoryBox lowAvailability').length", nil, &count1)
    page.RunScript("return document.getElementsByClassName('categoryBox yellowAvailability').length", nil, &count2)
    page.RunScript("return document.getElementsByClassName('categoryBox greenAvailability').length", nil, &count3)
    fmt.Println(count);

    m := make( map[string]interface{} )

    var temp1 []string
    page.RunScript(`var fruits = [];
      var tm = document.getElementsByClassName('categoryBox zeroAvailability');
      for(var i = 0; i < tm.length; i ++){
        var tmName = tm[i].parentElement.parentElement.getElementsByClassName('pull-left');
        Array.prototype.forEach.call(tmName, function(el) { var t = el.childNodes[1].className; if(t) fruits.push(t) });
      }
      return fruits;`,
       nil, &temp1)
    var temp1Int []int
    for index, element := range temp1{
      var re = regexp.MustCompile(`[[:space:]]`)
      temp1[index] = re.ReplaceAllString(element, "")
      if( temp1[index] == "categoryBoxzeroAvailability" ){
        temp1Int = append(temp1Int, 0)
      }
      if( temp1[index] == "categoryBoxlowAvailability" ){
        temp1Int = append(temp1Int, 1)
      }
      if( temp1[index] == "categoryBoxyellowAvailability" ){
        temp1Int = append(temp1Int, 2)
      }
      if( temp1[index] == "categoryBoxgreenAvailability" ){
        temp1Int = append(temp1Int, 3)
      }
      //fmt.Println(strconv.Itoa(index) + ": "+ temp1[index])
      if(index !=0 && (index+1) % 5 == 0){
        //fmt.Println(" ")
      }
    }


    var temp []string
    page.RunScript("var fruits = [];"+
      `var tm = document.getElementsByClassName('categoryBox zeroAvailability');
      //Array.prototype.forEach.call(tm, function(el) {fruits.push(el.textContent);});
      //Array.prototype.forEach.call(tm, function(el) {fruits.push(el.parentElement.parentElement.previousSibling.previousSibling.firstChild.innerHTML);});
      //Array.prototype.forEach.call(tm, function(el) {fruits.push(el.parentElement.parentElement.previousSibling.previousSibling.childNodes[1].tagName);});
      Array.prototype.forEach.call(tm, function(el) {fruits.push(el.parentElement.parentElement.previousSibling.previousSibling.childNodes[1].textContent);});
      return fruits;`,
       nil, &temp)

    for index, element := range temp{
      var re = regexp.MustCompile(`[[:space:]]`)
      temp[index] = re.ReplaceAllString(element, "")
      //fmt.Println(strconv.Itoa(index) + ": "+ temp[index])
      m[temp[index]] = []int{ temp1Int[index*5],temp1Int[index*5 + 1],temp1Int[index*5 + 2],temp1Int[index*5 + 3],temp1Int[index*5 + 4]}
      if(index !=0 && (index+1) % 5 == 0){
        //fmt.Println(" ")
      }
    }

    soundPlay := true
    for index, element := range m{
      v := element.([]int)
      //fmt.Println(index + "            : "+ strconv.Itoa(v[0]) + " "+ strconv.Itoa(v[1]) + " " + strconv.Itoa(v[2]) + " " +strconv.Itoa(v[3]) + " " + strconv.Itoa(v[4]))
      for indFilt, elemFilt := range filter{
              fmStr := fmt.Sprintf("%50.50s %d %d %d %d %d", index, v[0], v[1], v[2], v[3], v[4])
              fmt.Println(fmStr)
        if( elemFilt.name == index){
          for ii := 0; ii < 5; ii++{
            if(elemFilt.cube[ii] > 0 && v[ii] > 0 && v[ii] != (elemFilt.cube[ii]-1)){
              filter[indFilt].cube[ii] = 1 + v[ii]
              fmStr := fmt.Sprintf("%50.50s %d %d %d %d %d", index, v[0], v[1], v[2], v[3], v[4])
              fmt.Println(fmStr)
              //dialog.Message("%s", "Please select a file").Title("Hello world!").Info()
              if soundPlay == true{
                if err := lCommon.PlayMusic("./sound/nemeckaja-rech-i-signal-trevogi.mp3", 5 ) ; err != nil {
                  fmt.Println("no sound")
                }
              }
              soundPlay = false
              //break;
            }

          }
        }
      }
    }

    if count - lastCount > 80 && lastCount != 0{
      if err := lCommon.PlayMusic("./sound/nemeckaja-rech-i-signal-trevogi.mp3", 5 ) ; err != nil {
          fmt.Println("no sound")
      }
    }
    if count != 0{
      lastCount = count
    }

    duration = time.Second * 10 
    time.Sleep(duration)
    page.Refresh()
  }
  if err := driver.Stop(); err != nil {
    log.Fatal("Failed to close pages and stop WebDriver:", err)
  }
}
