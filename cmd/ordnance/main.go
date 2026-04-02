package main
import ("fmt";"log";"net/http";"os";"github.com/stockyard-dev/stockyard-ordnance/internal/server";"github.com/stockyard-dev/stockyard-ordnance/internal/store")
func main(){port:=os.Getenv("PORT");if port==""{port="9720"};dataDir:=os.Getenv("DATA_DIR");if dataDir==""{dataDir="./ordnance-data"}
db,err:=store.Open(dataDir);if err!=nil{log.Fatalf("ordnance: %v",err)};defer db.Close();srv:=server.New(db)
fmt.Printf("\n  Ordnance — vulnerability scanner\n  Dashboard:  http://localhost:%s/ui\n  API:        http://localhost:%s/api\n\n",port,port)
log.Printf("ordnance: listening on :%s",port);log.Fatal(http.ListenAndServe(":"+port,srv))}
