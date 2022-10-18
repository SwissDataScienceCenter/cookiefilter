# Traefik middleware cookie filter plugin

Simple plugin that removes all except specific cookies from a request.

If you do not provide any cookies that should be kept or you provide 
an empty list then all cookies will be removed.

Note that this plugin does not in any way modify cookies or anything
else in responses, only in requests.

Based on: https://github.com/traefik/plugindemo

## Demo

Navigate to the [demo](https://github.com/SwissDataScienceCenter/cookiefilter/tree/main/demo) 
folder in the repo to run a quick docker-compose
demonstration of this plugin. It includes additional information on
how to start and use the demo.

The demo also illustrates how the plugin can be loaded in a traefik
docker image and used without relying on the traefik Pilot. For more
information about packaging plugins in an image see 
[here](https://traefik.io/blog/using-private-plugins-in-traefik-proxy-2-5/).

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
        cookiefilter:
          keepCookies:
            - cookieName1
            - cookieName2
```

The middleware defined above would make it so that requests to `service-foo` 
are stripped of all cookies except the ones with names `cookieName1` and `cookieName2`.
