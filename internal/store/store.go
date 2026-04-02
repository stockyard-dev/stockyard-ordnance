package store
import ("database/sql";"fmt";"os";"path/filepath";"time";_ "modernc.org/sqlite")
type DB struct{db *sql.DB}
type Release struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Version string `json:"version"`
	Changelog string `json:"changelog"`
	AssetCount int `json:"asset_count"`
	Downloads int `json:"downloads"`
	Status string `json:"status"`
	PublishedAt string `json:"published_at"`
	CreatedAt string `json:"created_at"`
}
func Open(d string)(*DB,error){if err:=os.MkdirAll(d,0755);err!=nil{return nil,err};db,err:=sql.Open("sqlite",filepath.Join(d,"ordnance.db")+"?_journal_mode=WAL&_busy_timeout=5000");if err!=nil{return nil,err}
db.Exec(`CREATE TABLE IF NOT EXISTS releases(id TEXT PRIMARY KEY,name TEXT NOT NULL,version TEXT DEFAULT '',changelog TEXT DEFAULT '',asset_count INTEGER DEFAULT 0,downloads INTEGER DEFAULT 0,status TEXT DEFAULT 'draft',published_at TEXT DEFAULT '',created_at TEXT DEFAULT(datetime('now')))`)
return &DB{db:db},nil}
func(d *DB)Close()error{return d.db.Close()}
func genID()string{return fmt.Sprintf("%d",time.Now().UnixNano())}
func now()string{return time.Now().UTC().Format(time.RFC3339)}
func(d *DB)Create(e *Release)error{e.ID=genID();e.CreatedAt=now();_,err:=d.db.Exec(`INSERT INTO releases(id,name,version,changelog,asset_count,downloads,status,published_at,created_at)VALUES(?,?,?,?,?,?,?,?,?)`,e.ID,e.Name,e.Version,e.Changelog,e.AssetCount,e.Downloads,e.Status,e.PublishedAt,e.CreatedAt);return err}
func(d *DB)Get(id string)*Release{var e Release;if d.db.QueryRow(`SELECT id,name,version,changelog,asset_count,downloads,status,published_at,created_at FROM releases WHERE id=?`,id).Scan(&e.ID,&e.Name,&e.Version,&e.Changelog,&e.AssetCount,&e.Downloads,&e.Status,&e.PublishedAt,&e.CreatedAt)!=nil{return nil};return &e}
func(d *DB)List()[]Release{rows,_:=d.db.Query(`SELECT id,name,version,changelog,asset_count,downloads,status,published_at,created_at FROM releases ORDER BY created_at DESC`);if rows==nil{return nil};defer rows.Close();var o []Release;for rows.Next(){var e Release;rows.Scan(&e.ID,&e.Name,&e.Version,&e.Changelog,&e.AssetCount,&e.Downloads,&e.Status,&e.PublishedAt,&e.CreatedAt);o=append(o,e)};return o}
func(d *DB)Update(e *Release)error{_,err:=d.db.Exec(`UPDATE releases SET name=?,version=?,changelog=?,asset_count=?,downloads=?,status=?,published_at=? WHERE id=?`,e.Name,e.Version,e.Changelog,e.AssetCount,e.Downloads,e.Status,e.PublishedAt,e.ID);return err}
func(d *DB)Delete(id string)error{_,err:=d.db.Exec(`DELETE FROM releases WHERE id=?`,id);return err}
func(d *DB)Count()int{var n int;d.db.QueryRow(`SELECT COUNT(*) FROM releases`).Scan(&n);return n}

func(d *DB)Search(q string, filters map[string]string)[]Release{
    where:="1=1"
    args:=[]any{}
    if q!=""{
        where+=" AND (name LIKE ?)"
        args=append(args,"%"+q+"%");
    }
    if v,ok:=filters["status"];ok&&v!=""{where+=" AND status=?";args=append(args,v)}
    rows,_:=d.db.Query(`SELECT id,name,version,changelog,asset_count,downloads,status,published_at,created_at FROM releases WHERE `+where+` ORDER BY created_at DESC`,args...)
    if rows==nil{return nil};defer rows.Close()
    var o []Release;for rows.Next(){var e Release;rows.Scan(&e.ID,&e.Name,&e.Version,&e.Changelog,&e.AssetCount,&e.Downloads,&e.Status,&e.PublishedAt,&e.CreatedAt);o=append(o,e)};return o
}

func(d *DB)Stats()map[string]any{
    m:=map[string]any{"total":d.Count()}
    rows,_:=d.db.Query(`SELECT status,COUNT(*) FROM releases GROUP BY status`)
    if rows!=nil{defer rows.Close();by:=map[string]int{};for rows.Next(){var s string;var c int;rows.Scan(&s,&c);by[s]=c};m["by_status"]=by}
    return m
}
