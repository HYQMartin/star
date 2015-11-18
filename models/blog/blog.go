package blog

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"star/g"
	. "star/models"
	"time"
)

func OneById(id int64) *Blog {
	if id <= 0 {
		return nil
	}

	key := fmt.Sprintf("%d", id)
	val := g.BlogCacheGet(key)
	if val == nil {
		if p := OneByIdInDB(id); p != nil {
			g.BlogCachePut(key, *p)
			return p
		}
		return nil
	}
	ret := val.(Blog)
	return &ret
}

func OneByIdInDB(id int64) *Blog {
	if id <= 0 {
		return nil
	}

	o := Blog{Id: id}
	err := orm.NewOrm().Read(&o, "Id")
	if err != nil {
		return nil
	}
	return &o
}

func IdByIdent(ident string) int64 {
	if ident == "" {
		return 0
	}

	val := g.BlogCacheGet(ident)
	if val == nil {
		if p := OneByIdentInDB(ident); p != nil {
			g.BlogCachePut(ident, p.Id)
			return p.Id
		} else {
			return 0
		}
	}

	return val.(int64)
}

func IdentExists(ident string) bool {
	id := IdByIdent(ident)
	return id > 0
}

func OneByIdent(ident string) *Blog {
	id := IdByIdent(ident)
	return OneById(id)
}

func OneByIdentInDB(ident string) *Blog {
	if ident == "" {
		return nil
	}

	c := Blog{Ident: ident}
	err := orm.NewOrm().Read(&c, "Ident")
	if err != nil {
		return nil
	}

	return &c
}

func IdsInDB(catalog_id int64) []int64 {
	if catalog_id <= 0 {
		return []int64{}
	}

	var blogs []Blog
	Blogs().Filter("CatalogId", catalog_id).Filter("Status", 1).OrderBy("-Created").All(&blogs, "Id")
	size := len(blogs)
	if size == 0 {
		return []int64{}
	}

	ret := make([]int64, size)
	for i := 0; i < size; i++ {
		ret[i] = blogs[i].Id
	}

	return ret
}

func ReadBlogContent(b *Blog) *BlogContent {
	if b.Id <= 0 || b.BlogContentId <= 0 {
		return nil
	}

	key := fmt.Sprintf("content_of_%d_%d", b.Id, b.BlogContentLastUpdate)
	val := g.BlogCacheGet(key)
	if val == nil {
		if p := readBlogContentInDB(b); p != nil {
			g.BlogCachePut(key, *p)
			return p
		}
		return nil
	}
	ret := val.(BlogContent)
	return &ret
}

func readBlogContentInDB(b *Blog) *BlogContent {
	o := BlogContent{Id: b.BlogContentId}
	err := orm.NewOrm().Read(&o, "Id")
	if err != nil {
		return nil
	}
	return &o
}

func Ids(catalog_id int64) []int64 {
	if catalog_id <= 0 {
		return []int64{}
	}

	key := fmt.Sprintf("article_ids_of_%d", catalog_id)
	val := g.BlogCacheGet(key)
	if val == nil {
		if ids := IdsInDB(catalog_id); len(ids) != 0 {
			g.BlogCachePut(key, ids)
			return ids
		} else {
			return []int64{}
		}
	}

	return val.([]int64)
}

func ByCatalog(catalog_id int64, offset, limit int) []*Blog {
	ids := Ids(catalog_id)
	size := len(ids)
	if size == 0 {
		return []*Blog{}
	}

	if size > limit {
		end := offset + limit
		if end > size {
			end = size
		}

		ids = ids[offset:end]
	}

	size = len(ids)
	ret := make([]*Blog, size)
	for i := 0; i < size; i++ {
		ret[i] = OneById(ids[i])
		ret[i].Content = ReadBlogContent(ret[i])
	}
	return ret
}

func Save(this *Blog, blogContent string) (int64, error) {
	if IdentExists(this.Ident) {
		return 0, fmt.Errorf("blog english identity exists")
	}

	bc := &BlogContent{Content: blogContent}
	or := orm.NewOrm()
	blogContentId, e := or.Insert(bc)
	if e != nil {
		return 0, e
	}

	this.BlogContentId = blogContentId
	this.BlogContentLastUpdate = time.Now().Unix()

	id, err := or.Insert(this)
	if err == nil {
		g.BlogCacheDel(fmt.Sprintf("article_ids_of_%d", this.CatalogId))
	}

	return id, err
}

func Del(b *Blog) error {
	num, err := Blogs().Filter("Id", b.Id).Delete()
	if err != nil {
		return err
	}

	if num > 0 {
		g.BlogCacheDel(fmt.Sprintf("article_ids_of_%d", b.CatalogId))
		BlogContents().Filter("Id", b.BlogContentId).Delete()
	}

	return nil
}

func Update(b *Blog, content string) error {
	if b.Id == 0 {
		return fmt.Errorf("primary key:id not set")
	}

	bc := ReadBlogContent(b)
	if content != "" && bc.Content != content {
		bc.Content = content
		_, e := orm.NewOrm().Update(bc)
		if e != nil {
			return e
		}
		b.BlogContentLastUpdate = time.Now().Unix()
	}

	_, err := orm.NewOrm().Update(b)
	if err == nil {
		g.BlogCacheDel(fmt.Sprintf("%d", b.Id))
	}
	return err
}

func Blogs() orm.QuerySeter {
	return orm.NewOrm().QueryTable(new(Blog))
}

func BlogContents() orm.QuerySeter {
	return orm.NewOrm().QueryTable(new(BlogContent))
}

//********************************************************

func NOneById(id int64) *NewemployeeBlog {
	if id <= 0 {
		return nil
	}

	key := fmt.Sprintf("%d", id)
	val := g.NBlogCacheGet(key)
	if val == nil {
		if p := NOneByIdInDB(id); p != nil {
			g.NBlogCachePut(key, *p)
			return p
		}
		return nil
	}
	ret := val.(NewemployeeBlog)
	return &ret
}

func NOneByIdInDB(id int64) *NewemployeeBlog {
	if id <= 0 {
		return nil
	}

	o := NewemployeeBlog{Id: id}
	err := orm.NewOrm().Read(&o, "Id")
	if err != nil {
		return nil
	}
	return &o
}

func NIdByIdent(ident string) int64 {
	if ident == "" {
		return 0
	}

	val := g.NBlogCacheGet(ident)
	if val == nil {
		if p := NOneByIdentInDB(ident); p != nil {
			g.NBlogCachePut(ident, p.Id)
			return p.Id
		} else {
			return 0
		}
	}

	return val.(int64)
}

func NIdentExists(ident string) bool {
	id := NIdByIdent(ident)
	return id > 0
}

func NOneByIdent(ident string) *NewemployeeBlog {
	id := NIdByIdent(ident)
	return NOneById(id)
}

func NOneByIdentInDB(ident string) *NewemployeeBlog {
	if ident == "" {
		return nil
	}

	c := NewemployeeBlog{Ident: ident}
	err := orm.NewOrm().Read(&c, "Ident")
	if err != nil {
		return nil
	}

	return &c
}

func NIdsInDB(catalog_id int64) []int64 {
	if catalog_id <= 0 {
		return []int64{}
	}

	var blogs []NewemployeeBlog
	NBlogs().Filter("CatalogId", catalog_id).Filter("Status", 1).OrderBy("-Created").All(&blogs, "Id")
	size := len(blogs)
	if size == 0 {
		return []int64{}
	}

	ret := make([]int64, size)
	for i := 0; i < size; i++ {
		ret[i] = blogs[i].Id
	}

	return ret
}

func NReadBlogContent(b *NewemployeeBlog) *NewemployeeBlogContent {
	if b.Id <= 0 || b.BlogContentId <= 0 {
		return nil
	}

	key := fmt.Sprintf("content_of_%d_%d", b.Id, b.BlogContentLastUpdate)
	val := g.NBlogCacheGet(key)
	if val == nil {
		if p := NreadBlogContentInDB(b); p != nil {
			g.NBlogCachePut(key, *p)
			return p
		}
		return nil
	}
	ret := val.(NewemployeeBlogContent)
	return &ret
}

func NreadBlogContentInDB(b *NewemployeeBlog) *NewemployeeBlogContent {
	o := NewemployeeBlogContent{Id: b.BlogContentId}
	err := orm.NewOrm().Read(&o, "Id")
	if err != nil {
		return nil
	}
	return &o
}

func NIds(catalog_id int64) []int64 {
	if catalog_id <= 0 {
		return []int64{}
	}

	key := fmt.Sprintf("article_ids_of_%d", catalog_id)
	val := g.NBlogCacheGet(key)
	if val == nil {
		if ids := NIdsInDB(catalog_id); len(ids) != 0 {
			g.NBlogCachePut(key, ids)
			return ids
		} else {
			return []int64{}
		}
	}

	return val.([]int64)
}

func NByCatalog(catalog_id int64, offset, limit int) []*NewemployeeBlog {
	ids := NIds(catalog_id)
	size := len(ids)
	if size == 0 {
		return []*NewemployeeBlog{}
	}

	if size > limit {
		end := offset + limit
		if end > size {
			end = size
		}

		ids = ids[offset:end]
	}

	size = len(ids)
	ret := make([]*NewemployeeBlog, size)
	for i := 0; i < size; i++ {
		ret[i] = NOneById(ids[i])
		ret[i].Content = NReadBlogContent(ret[i])
	}
	return ret
}

func NSave(this *NewemployeeBlog, blogContent string) (int64, error) {
	if NIdentExists(this.Ident) {
		return 0, fmt.Errorf("blog english identity exists")
	}

	bc := &NewemployeeBlogContent{Content: blogContent}
	or := orm.NewOrm()
	blogContentId, e := or.Insert(bc)
	if e != nil {
		return 0, e
	}

	this.BlogContentId = blogContentId
	this.BlogContentLastUpdate = time.Now().Unix()

	id, err := or.Insert(this)
	if err == nil {
		g.NBlogCacheDel(fmt.Sprintf("article_ids_of_%d", this.CatalogId))
	}

	return id, err
}

func NDel(b *NewemployeeBlog) error {
	num, err := Blogs().Filter("Id", b.Id).Delete()
	if err != nil {
		return err
	}

	if num > 0 {
		g.NBlogCacheDel(fmt.Sprintf("article_ids_of_%d", b.CatalogId))
		NBlogContents().Filter("Id", b.BlogContentId).Delete()
	}

	return nil
}

func NUpdate(b *NewemployeeBlog, content string) error {
	if b.Id == 0 {
		return fmt.Errorf("primary key:id not set")
	}

	bc := NReadBlogContent(b)
	if content != "" && bc.Content != content {
		bc.Content = content
		_, e := orm.NewOrm().Update(bc)
		if e != nil {
			return e
		}
		b.BlogContentLastUpdate = time.Now().Unix()
	}

	_, err := orm.NewOrm().Update(b)
	if err == nil {
		g.NBlogCacheDel(fmt.Sprintf("%d", b.Id))
	}
	return err
}

func NBlogs() orm.QuerySeter {
	return orm.NewOrm().QueryTable(new(NewemployeeBlog))
}

func NBlogContents() orm.QuerySeter {
	return orm.NewOrm().QueryTable(new(NewemployeeBlogContent))
}
