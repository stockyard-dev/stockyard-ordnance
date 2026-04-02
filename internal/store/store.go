package store
import ("database/sql";"fmt";"os";"path/filepath";"time";_ "modernc.org/sqlite")
type DB struct{db *sql.DB}
type Release struct{
	ID string `json:"id"`
	Name string `json:"name"`
	Version string `json:"version"`
	Platform string `json:"platform"`
	Artifact string `json:"artifact_url"`
	Checksum string `json:"checksum"`
	Downloads int `json:"downloads"`
	CreatedAt string `json:"created_at"`
}
func Open(d string)(*DB,error){if err:=os.MkdirAll(d,0755);err!=nil{return nil,err};db,err:=sql.Open("sqlite",filepath.Join(d,"ordnance.db")+"?_journal_mode=WAL&_busy_timeout=5000");if err!=nil{return nil,err}
db.Exec(`CREATE TABLE IF NOT EXISTS releases(id TEXT PRIMARY KEY,name TEXT NOT NULL,version TEXT DEFAULT '',platform TEXT DEFAULT '',artifact_url TEXT DEFAULT '',checksum TEXT DEFAULT '',downloads INTEGER DEFAULT 0,created_at TEXT DEFAULT(datetime('now')))`)
return &DB{db:db},nil}
func(d *DB)Close()error{return d.db.Close()}
func genID()string{return fmt.Sprintf("%d",time.Now().UnixNano())}
func now()string{return time.Now().UTC().Format(time.RFC3339)}
func(d *DB)Create(e *Release)error{e.ID=genID();e.CreatedAt=now();_,err:=d.db.Exec(`INSERT INTO releases(id,name,version,platform,artifact_url,checksum,downloads,created_at)VALUES(?,?,?,?,?,?,?,?)`,e.ID,e.Name,e.Version,e.Platform,e.Artifact,e.Checksum,e.Downloads,e.CreatedAt);return err}
func(d *DB)Get(id string)*Release{var e Release;if d.db.QueryRow(`SELECT id,name,version,platform,artifact_url,checksum,downloads,created_at FROM releases WHERE id=?`,id).Scan(&e.ID,&e.Name,&e.Version,&e.Platform,&e.Artifact,&e.Checksum,&e.Downloads,&e.CreatedAt)!=nil{return nil};return &e}
func(d *DB)List()[]Release{rows,_:=d.db.Query(`SELECT id,name,version,platform,artifact_url,checksum,downloads,created_at FROM releases ORDER BY created_at DESC`);if rows==nil{return nil};defer rows.Close();var o []Release;for rows.Next(){var e Release;rows.Scan(&e.ID,&e.Name,&e.Version,&e.Platform,&e.Artifact,&e.Checksum,&e.Downloads,&e.CreatedAt);o=append(o,e)};return o}
func(d *DB)Delete(id string)error{_,err:=d.db.Exec(`DELETE FROM releases WHERE id=?`,id);return err}
func(d *DB)Count()int{var n int;d.db.QueryRow(`SELECT COUNT(*) FROM releases`).Scan(&n);return n}
