# Traefik middleware cookie filter plugin

Simple plugin that removes all except specific cookies from a request.

If you do not provide any cookies that should be kept or you provide 
an empty list then all cookies will be removed.

Based on: https://github.com/traefik/plugindemo

## Usage

Add the plugin in your static configuration

```yaml
# Static configuration
experimental:
  plugins:
    cookieFilterPlugin:
      moduleName: github.com/SwissDataScienceCenter/cookiefilter
      version: "0.0.1"
```

Use the plugin in your dynamic configuration like this

```yaml
# Dynamic configuration

http:
  routers:
    my-router:
      rule: host(`demo.localhost`)
      service: service-foo
      entryPoints:
        - web
      middlewares:
        - cookiefilter

  services:
   service-foo:
      loadBalancer:
        servers:
          - url: http://127.0.0.1:5000
  
  middlewares:
    cookiefilter:
      plugin:
        cookieFilterPlugin:
          keepCookies:
            - cookieName1
            - cookieName2
```

The middleware defined above would make it so that requests to `service-foo` 
are stripped of all cookies except the ones with names `cookieName1` and `cookieName2`.
