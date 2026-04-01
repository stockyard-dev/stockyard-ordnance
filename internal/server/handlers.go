package server
import("encoding/json";"net/http";"strconv";"github.com/stockyard-dev/stockyard-ordnance/internal/store")
func(s *Server)handleList(w http.ResponseWriter,r *http.Request){list,_:=s.db.List();if list==nil{list=[]store.Scan{}};writeJSON(w,200,list)}
func(s *Server)handleCreate(w http.ResponseWriter,r *http.Request){var req struct{Target string `json:"target"`};json.NewDecoder(r.Body).Decode(&req);if req.Target==""{writeError(w,400,"target required");return};sc,_:=s.db.CreateScan(req.Target);writeJSON(w,202,sc)}
func(s *Server)handleAddVuln(w http.ResponseWriter,r *http.Request){id,_:=strconv.ParseInt(r.PathValue("id"),10,64);var v store.Vulnerability;json.NewDecoder(r.Body).Decode(&v);v.ScanID=id;if v.Package==""{writeError(w,400,"package required");return};s.db.AddVuln(&v);writeJSON(w,201,v)}
func(s *Server)handleGetVulns(w http.ResponseWriter,r *http.Request){id,_:=strconv.ParseInt(r.PathValue("id"),10,64);list,_:=s.db.GetVulns(id);if list==nil{list=[]store.Vulnerability{}};writeJSON(w,200,list)}
func(s *Server)handleDelete(w http.ResponseWriter,r *http.Request){id,_:=strconv.ParseInt(r.PathValue("id"),10,64);s.db.Delete(id);writeJSON(w,200,map[string]string{"status":"deleted"})}
func(s *Server)handleOverview(w http.ResponseWriter,r *http.Request){m,_:=s.db.Stats();writeJSON(w,200,m)}
