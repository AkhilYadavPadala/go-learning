package main
import "fmt"
func main(){
	var cpu float64;
	var memory float64;
	var latency float64;
	fmt.Println("Enter values sepearated by space")
	fmt.Scan(&cpu,&memory,&latency)
	answer:=getSystemStatus(cpu,memory,latency)
	fmt.Println(answer)
}

func getSystemStatus(cpu,memory,latency float64) string {
	if cpu<80 && memory<90 && latency<100 {
		return "System Status: Healthy"
	}else{
		return "System Status: Critical"
	}
}