package main
import "fmt"

func main()  {
	languages := make(map[string]string)

	languages["JS"] = "javascript";
	languages["RB"] = "ruby";
	languages["PY"] = "python";
	
	fmt.Println(languages);

	fmt.Println(languages["JS"]);
	

	delete(languages,"JS");

	for key, value := range languages{
		fmt.Println(key+" "+value)
	}
}