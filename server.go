package main

import (
    "github.com/gin-gonic/gin"
    "errors"
    "log"
    "regexp"
    "net/http"
    "io/ioutil"
    )

type Request struct {
    Site []string 
    SearchText string
}

type Response struct {
    FoundAtSite string
}



func main() {
    r := gin.Default()
    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200,gin.H{
            "main":"default",
        })    
    })
    
    r.POST("/checkText", checkTextHandler);
    
    
    
    r.Run("0.0.0.0:8080") // listen and server on 0.0.0.0:8080
}




func checkTextHandler(c *gin.Context){
    var jsonreq Request;
       

    
    if c.BindJSON(&jsonreq) == nil && len(jsonreq.Site) > 0 && len(jsonreq.SearchText) > 0   {
        var resjson Response;
        
        
        
        i:=searchOnEachSiteOfArr(jsonreq.Site,jsonreq.SearchText)
        
        if i >=0 {
            resjson.FoundAtSite = jsonreq.Site[i]
        
            log.Print(`>SearchText: "`,jsonreq.SearchText,`"   found on Site [` ,i,`] : `, jsonreq.Site[i] )
        
            c.JSON(200,resjson)
        }else{
            c.AbortWithError(204,errors.New(">word not found"))
        }
        
    }else{
      c.AbortWithError(204,errors.New(">invalid request"))
    }
    
    
}




func getSiteContent( url string ) string {
    res,_ := http.Get(url);
    
    txt,_ := ioutil.ReadAll(res.Body)
    res.Body.Close()
    

    return string(txt)
}



func pureHTMLTextContent(src string) (string){

    tagexp := regexp.MustCompile(`<.*?/?.*?>`)
    scriptexp := compileRegExpMatchingTagAndAllInsideIt("script")
    styleexp := compileRegExpMatchingTagAndAllInsideIt("style")
    

    return tagexp.ReplaceAllLiteralString( scriptexp.ReplaceAllLiteralString(styleexp.ReplaceAllLiteralString(src," ")," ")," ")// match: re(helloworld) on: hell<tag>wo<tag>rld should be correct

}


func compileRegExpMatchingTagAndAllInsideIt(tagname string) *regexp.Regexp {
    return regexp.MustCompile(`(?s:<\s*` + tagname +`.*?>.*?<\s*/\s*`+ tagname +`\s*>)`)
}








func searchOnEachSiteOfArr(sites []string, search string) int {
    resrch := regexp.MustCompile(search)
    
    for k,val := range sites {
        if( resrch.MatchString( pureHTMLTextContent(getSiteContent(val) ) )   ){
          return k  
        }
    }
    
    return -1
}





