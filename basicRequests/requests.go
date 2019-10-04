package main
import (
	"fmt"
	"net/http"
	"io/ioutil"
	//"os"
	"regexp"
    "strings"
)


func main(){
	req, err := http.Get("https://www.amazon.com.br/gp/browse.html?node=6740748011&ref_=nav_em_T1_0_4_6_1__books_all")
   
    if err != nil{
    	fmt.Println(err)	
    }

    defer req.Body.Close()

    body, err := ioutil.ReadAll(req.Body)
    if err != nil {
    	fmt.Println(err)	
    }
	get_product_grid_item(string(body))
}

func get_title(html string){
	re := regexp.MustCompile("<a class=\"a-link-normal(.|\n)*?>(.|\n)*?</a>")
	titulos := re.FindAllString(string(html), -1)
	if titulos == nil{
		fmt.Println("Nenhum titulo encontrado")
	}else{
		//fmt.Println(titulos);
		for _, t := range titulos{
			s := strings.Split(t,"title=\"")
			s = strings.Split(s[1],"\"")
			fmt.Println( s[0] ) //s[0]
		}
	}
}
//acs_product-price__list
func get_price(html string){
	re := regexp.MustCompile(`<span class="a-size-mini a-color-secondary acs_product-price__list a-text-strike">(.|\n)*?</span>`)
	precos_s_desc := re.FindAllString(string(html), -1)

	if precos_s_desc == nil{
		fmt.Println("Preço não encontrado")
	}else{
		for _, t := range precos_s_desc{
			s := strings.Split(t,">")
			s = strings.Split(s[1],"<")
			fmt.Println( s[0] )
		}
	}

}

func get_product_grid_item(html string){
	re := regexp.MustCompile("<li class=\"a-carousel-card acswidget-carousel__card(.|\n)*?>(.|\n)*?</li>")
	div_produtos := re.FindAllString(string(html), -1)
	if div_produtos == nil{
		fmt.Println("Nenhum produtos encontrado")
	}else{
		for _, div := range div_produtos{
			get_title(div)
			get_price(div)
		}
	}
}












//func getLinks( html string){
//	re := regexp.MustCompile("<a href=(.|\n)*?>(.|\n)*?</a>")
//    links := re.FindAllString(string(html), -1)
//
//    if links == nil{
//    	fmt.Println("Nenhum link encontrado")
//    }else{
//
//    	for _, link := range links{
//    		//fmt.Println( link )
//            //fmt.Println("\n")
//            getHREF(link)
//            //fmt.Println("\n")
//    	}
//    }
//}
//
//
//func getP( html string) {
//	re := regexp.MustCompile("<span (.|\n)*?>(.|\n)*?</span>")
//    paragrafos := re.FindAllString(string(html), -1)
//
//    if paragrafos == nil{
//    	fmt.Println("Nenhum paragrafo encontrado")
//    }else{
//    	for _, p := range paragrafos{
//    		fmt.Println(p)
//    	}
//    }
//}
//
//func getHREF(html string){
//    re := regexp.MustCompile("href=\"(.|\n)*?\"")
//    url := re.FindAllString(string(html), -1)
//    if url == nil{
//        fmt.Println("Nenhum paragrafo encontrado")
//    }else{
//        for _, u := range url{
//            s := strings.Split(u,"\"")
//            if(strings.Contains(s[1],"2019/09/27")){
//				fmt.Println(s[1])
//			}
//        }
//    }
//}


