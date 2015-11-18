package g

const (
	blogPrefix               = "b_"
	catalogPrefix            = "c_"
	newemployeeblogPrefix    = "nb_"
	newemployeecatalogPrefix = "nc_"
)

func BlogCachePut(key string, val interface{}) error {
	return Cache.Put(blogPrefix+key, val, blogCacheExpire)
}

func CatalogCachePut(key string, val interface{}) error {
	return Cache.Put(catalogPrefix+key, val, catalogCacheExpire)
}

func BlogCacheGet(key string) interface{} {
	return Cache.Get(blogPrefix + key)
}

func CatalogCacheGet(key string) interface{} {
	return Cache.Get(catalogPrefix + key)
}

func CatalogCacheDel(key string) error {
	return Cache.Delete(catalogPrefix + key)
}

func BlogCacheDel(key string) error {
	return Cache.Delete(blogPrefix + key)
}

//************************************************
func NBlogCachePut(key string, val interface{}) error {
	return Cache.Put(newemployeeblogPrefix+key, val, blogCacheExpire)
}

func NCatalogCachePut(key string, val interface{}) error {
	return Cache.Put(newemployeecatalogPrefix+key, val, catalogCacheExpire)
}

func NBlogCacheGet(key string) interface{} {
	return Cache.Get(newemployeeblogPrefix + key)
}

func NCatalogCacheGet(key string) interface{} {
	return Cache.Get(newemployeecatalogPrefix + key)
}

func NCatalogCacheDel(key string) error {
	return Cache.Delete(newemployeecatalogPrefix + key)
}

func NBlogCacheDel(key string) error {
	return Cache.Delete(newemployeeblogPrefix + key)
}
