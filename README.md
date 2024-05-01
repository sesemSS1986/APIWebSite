<b>Начало работы</b>

Загрузите пакет GO

	go get github.com/sesemSS1986/APIWebSite


<b>Быстрый старт</b>

Для непосредственного взаимодействия с REST API нужно создать клиента. Примеры использования находятся в директории examples или ниже в тексте.



<b>Example</b>

	func main() {
	
	  c := Client{
	
	  Url: "http://localhost:1323",
	
	  Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiQWRtaW4iLCJhZG1pbiI6dHJ1ZSwiZXhwIjoxNzEzNDI3MDg1fQ.ecbIV7kyj6l31HnFLDfSX_Zr1ybMSmfnP4YlDkKcNQY", }
	
	
	  //Метод получения информации по пользователям.
	
	  resdataU := c.GetUser()
	
	  //Метод получения информации по товарам.
	
	  resdataS := c.GetShop()
	
	   //Метод получения информации по товарам в корзинах у пользователей.
	
	  resdataB := c.GetBasket()
	
	  //Метод получения информации по сформированным заказам.
	
	  resdataO := c.GetOrders()
	
	  
		
	 //Получение всей информации
	
	  log.Println(resdataU)
	
	  log.Println(resdataS)
	
	  log.Println(resdataB)
	
	  log.Println(resdataO)
	
	  
		
	 //Получение более конкретизированной информации
	
	  for i, v := range resdataU {
	
	  //получение именён пользователей
	
	  log.Println(i, v.FirstName)
	
	  }
	
	}
