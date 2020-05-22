package main

func main() {
	CacheRef := GlobalCacheRef{}
	Router := RouterConfig{}
	CacheRef.intializeGlobalCache()
	Router.initializeRouter(CacheRef)
	Router.registerRoutes()
}
