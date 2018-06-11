package main

import (
  "log"
	"fmt"
  "github.com/sclevine/agouti"
	"strconv"
  "regexp"
  "./sendemail"
	"time"
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

    var temp1 []string
    page.RunScript(`var fruits = [];
      var tm = document.getElementsByClassName('categoryBox zeroAvailability');
      for(var i = 0; i < tm.length; i ++){
        var tmName = tm[i].parentElement.parentElement.getElementsByClassName('pull-left');
        Array.prototype.forEach.call(tmName, function(el) { var t = el.childNodes[1].className; if(t) fruits.push(t) });
      }
      return fruits;`,
       nil, &temp1)
    for index, element := range temp1{
      var re = regexp.MustCompile(`[[:space:]]`)
      temp1[index] = re.ReplaceAllString(element, "")
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
      if(index !=0 && (index+1) % 5 == 0){
        //fmt.Println(" ")
      }
    }

    //fmt.Println(temp)

    //for i := 0; i < count; i++{
      //sendFmt := "return document.getElementsByClassName('categoryBox zeroAvailability')["+strconv.Itoa(i)+"].textContent"
      //page.RunScript(sendFmt, nil, &value)
      //fmt.Println(value);
    //}
    //page.RunScript("return window.someObject;", nil, &value)
    //dat, err := ioutil.ReadFile("bundle.js")
    //check(err)
    //var number int
    //page.RunScript(string(dat), nil, &value)
    //fmt.Println(number)

    //log.Println(page.HTML());
    //sectionTitle, err := page.FindByID("stylesheet").Text()
    //sectionTitle, err := page.FindByClass("header").Text()
    //log.Println(sectionTitle)

    //sendemail.Send(temp[0])
    sendemail.Send("zeroAvailability = " + strconv.Itoa(count) +
    " lowAvailability = " + strconv.Itoa(count1) +
    " yellowAvailability = " + strconv.Itoa(count2) +
    " grenAvailability = " + strconv.Itoa(count3), 
    "exitstop@list.ru", "gbe643412@gmail.com", "fgjkriJDdjrjhfhIF73hfd" )
    duration = time.Second * 60
    time.Sleep(duration)
    page.Refresh()
  }
  if err := driver.Stop(); err != nil {
    log.Fatal("Failed to close pages and stop WebDriver:", err)
  }
}
